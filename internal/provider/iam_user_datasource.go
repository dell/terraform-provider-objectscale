package provider

import (
	"context"
	"fmt"

	"terraform-provider-objectscale/internal/client"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &IAMUserDataSource{}
	_ datasource.DataSourceWithConfigure = &IAMUserDataSource{}
)

type IAMUserDataSource struct {
	client *client.Client
}

func NewIAMUserDataSource() datasource.DataSource {
	return &IAMUserDataSource{}
}

func (d *IAMUserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_user"
}

func (d *IAMUserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *IAMUserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Fetch IAM user information for a specific ObjectScale namespace.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Internal ID for this data source.",
			},

			"namespace": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Namespace containing IAM users.",
			},

			"username": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filter users by username.",
			},

			"groupname": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Filter users who belong to the given group name.",
			},

			"users": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: "List of IAM users matching the filters.",

				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"username": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: " IAM username.",
						},

						"id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Unique ObjectScale IAM user ID (maps to UserId).",
						},

						"arn": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "IAM user ARN.",
						},

						"create_date": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The timestamp when the user was created.",
						},

						"path": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "IAM user path.",
						},

						"permissions_boundary": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Policy ARN used as permissions boundary. Can be empty if not set.",
						},

						"tags": schema.ListNestedAttribute{
							Computed:            true,
							MarkdownDescription: "List of tags assigned to the user (sorted by key in Read function).",

							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Tag key.",
									},
									"value": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Tag value.",
									},
								},
							},
						},

						"access_keys": schema.ListNestedAttribute{
							Computed:            true,
							MarkdownDescription: "List of access keys for the IAM user.",

							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"access_key_id": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Access key ID.",
									},
									"create_date": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Timestamp when the access key was created.",
									},
									"status": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Status of the access key (Active/Inactive).",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *IAMUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.IAMUserDatasourceModel

	// Load inputs
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ns := data.Namespace.ValueString()

	// Helper: *string → types.String
	safe := func(s *string) types.String {
		if s != nil {
			return types.StringValue(*s)
		}
		return types.StringValue("")
	}

	// CASE 1 — USERNAME PROVIDED → DIRECTLY CALL GetUser
	if !data.Username.IsNull() {
		username := data.Username.ValueString()

		getResp, _, err := d.client.GenClient.IamApi.
			IamServiceGetUser(ctx).
			UserName(username).
			XEmcNamespace(ns).
			Execute()

		if err != nil {
			resp.Diagnostics.AddError(
				"Error calling GetUser",
				fmt.Sprintf("Failed retrieving user %q: %v", username, err),
			)
			return
		}

		if getResp == nil || getResp.GetUserResult == nil || getResp.GetUserResult.User == nil {
			resp.Diagnostics.AddError("Invalid GetUserResponse", "Missing GetUserResult.User")
			return
		}

		u := getResp.GetUserResult.User

		// ---- load tags ----
		var tags []models.IAMUserTag
		tResp, _, _ := d.client.GenClient.IamApi.
			IamServiceListUserTags(ctx).
			UserName(username).
			XEmcNamespace(ns).
			Execute()

		if tResp != nil && tResp.ListUserTagsResult != nil {
			for _, t := range tResp.ListUserTagsResult.Tags {
				tags = append(tags, models.IAMUserTag{
					Key:   safe(t.Key),
					Value: safe(t.Value),
				})
			}
		}

		// ---- load access keys ----
		var accessKeys []models.IAMUserAccessKey
		kResp, _, _ := d.client.GenClient.IamApi.
			IamServiceListAccessKeys(ctx).
			UserName(username).
			XEmcNamespace(ns).
			Execute()

		if kResp != nil && kResp.ListAccessKeysResult != nil {
			for _, k := range kResp.ListAccessKeysResult.AccessKeyMetadata {
				accessKeys = append(accessKeys, models.IAMUserAccessKey{
					AccessKeyId: safe(k.AccessKeyId),
					CreateDate:  safe(k.CreateDate),
					Status:      safe(k.Status),
				})
			}
		}

		// PermissionsBoundary safe handling
		var pb types.String
		if u.PermissionsBoundary != nil && u.PermissionsBoundary.PermissionsBoundaryArn != nil {
			pb = types.StringValue(*u.PermissionsBoundary.PermissionsBoundaryArn)
		} else {
			pb = types.StringValue("")
		}

		// ---- final user object ----
		userObj := models.IAMUser{
			ID:                  safe(u.UserId),
			UserName:            safe(u.UserName),
			Arn:                 safe(u.Arn),
			Path:                safe(u.Path),
			CreateDate:          safe(u.CreateDate),
			PermissionsBoundary: pb,
			Tags:                tags,
			AccessKeys:          accessKeys,
		}

		// ---- save state ----
		data.ID = types.StringValue("iam_user_datasource")
		data.Users = []models.IAMUser{userObj}

		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	// CASE 2 — GROUP FILTER (preload usernames)

	groupUserSet := map[string]bool{}

	if !data.Groupname.IsNull() {
		groupName := data.Groupname.ValueString()

		gResp, hResp, err := d.client.GenClient.IamApi.
			IamServiceGetGroup(ctx).
			GroupName(groupName).
			XEmcNamespace(ns).
			Execute()

		if err != nil || hResp.StatusCode >= 400 {
			resp.Diagnostics.AddError(
				"Error calling GetGroup",
				fmt.Sprintf("Unable to retrieve IAM group %q: %v", groupName, err),
			)
			return
		}

		if gResp.GetGroupResult != nil {
			for _, u := range gResp.GetGroupResult.Users {
				if u.UserName != nil {
					groupUserSet[*u.UserName] = true
				}
			}
		}
	}

	// CASE 3 — LIST ALL USERS (when username not provided)

	listResp, _, err := d.client.GenClient.IamApi.
		IamServiceListUsers(ctx).
		XEmcNamespace(ns).
		Execute()

	if err != nil || listResp == nil || listResp.ListUsersResult == nil {
		resp.Diagnostics.AddError("Error listing IAM users", err.Error())
		return
	}

	var finalUsers []models.IAMUser

	for _, u := range listResp.ListUsersResult.Users {
		if u.UserName == nil {
			continue
		}

		username := *u.UserName

		// apply group filter
		if len(groupUserSet) > 0 && !groupUserSet[username] {
			continue
		}

		// ---- fetch tags ----
		var tags []models.IAMUserTag
		tResp, _, _ := d.client.GenClient.IamApi.
			IamServiceListUserTags(ctx).
			UserName(username).
			XEmcNamespace(ns).
			Execute()

		if tResp != nil && tResp.ListUserTagsResult != nil {
			for _, t := range tResp.ListUserTagsResult.Tags {
				tags = append(tags, models.IAMUserTag{
					Key:   safe(t.Key),
					Value: safe(t.Value),
				})
			}
		}

		// ---- fetch access keys ----
		var accessKeys []models.IAMUserAccessKey
		kResp, _, _ := d.client.GenClient.IamApi.
			IamServiceListAccessKeys(ctx).
			UserName(username).
			XEmcNamespace(ns).
			Execute()

		if kResp != nil && kResp.ListAccessKeysResult != nil {
			for _, k := range kResp.ListAccessKeysResult.AccessKeyMetadata {
				accessKeys = append(accessKeys, models.IAMUserAccessKey{
					AccessKeyId: safe(k.AccessKeyId),
					CreateDate:  safe(k.CreateDate),
					Status:      safe(k.Status),
				})
			}
		}

		// listUsers does NOT return permissions boundary → empty
		pb := types.StringValue("")

		// ---- build final user struct ----
		finalUsers = append(finalUsers, models.IAMUser{
			ID:                  safe(u.UserId),
			UserName:            safe(u.UserName),
			Arn:                 safe(u.Arn),
			Path:                safe(u.Path),
			CreateDate:          safe(u.CreateDate),
			PermissionsBoundary: pb,
			Tags:                tags,
			AccessKeys:          accessKeys,
		})
	}

	// save state
	data.ID = types.StringValue("iam_user_datasource")
	data.Users = finalUsers
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
