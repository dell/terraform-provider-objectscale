package provider

import (
	"context"
	"fmt"

	"terraform-provider-objectscale/internal/client"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	// "github.com/hashicorp/terraform-plugin-log/tflog"
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
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
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

	// 0. Optional GROUP FILTER → preload usernames in group
	groupUserSet := map[string]bool{}
	if !data.Groupname.IsNull() {
		groupName := data.Groupname.ValueString()

		groupResp, httpResp, err := d.client.GenClient.IamApi.
			IamServiceGetGroup(ctx).
			GroupName(groupName).
			XEmcNamespace(ns).
			Execute()

		if err != nil {
			if httpResp != nil && httpResp.StatusCode == 404 {
				resp.Diagnostics.AddError(
					"IAM Group Not Found",
					fmt.Sprintf("IAM Group %q does not exist in namespace %q", groupName, ns),
				)
				return
			}
			resp.Diagnostics.AddError("Error calling GetGroup", err.Error())
			return
		}
		if httpResp != nil && httpResp.StatusCode >= 400 {
			resp.Diagnostics.AddError(
				"Error calling GetGroup",
				fmt.Sprintf("HTTP %d returned for group %q", httpResp.StatusCode, groupName),
			)
			return
		}

		if groupResp != nil && groupResp.GetGroupResult != nil {
			for _, userObj := range groupResp.GetGroupResult.Users {
				if userObj.UserName != nil {
					groupUserSet[*userObj.UserName] = true
				}
			}
		}
	}

	// 1. List all users in the namespace
	listResp, _, err := d.client.GenClient.IamApi.
		IamServiceListUsers(ctx).
		XEmcNamespace(ns).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Error listing IAM users", err.Error())
		return
	}
	if listResp == nil || listResp.ListUsersResult == nil {
		resp.Diagnostics.AddError("Invalid ListUsers response", "ListUsersResult is nil")
		return
	}

	var finalUsers []models.IAMUser

	// 2. Loop through each user and fetch tags & access keys
	for _, u := range listResp.ListUsersResult.Users {
		if u.UserName == nil {
			continue
		}
		un := *u.UserName

		// Optional username filter
		if !data.Username.IsNull() && data.Username.ValueString() != un {
			continue
		}

		// Optional group filter
		if len(groupUserSet) > 0 && !groupUserSet[un] {
			continue
		}

		// Fetch user tags
		var tags []models.IAMUserTag
		tagsResp, _, tErr := d.client.GenClient.IamApi.
			IamServiceListUserTags(ctx).
			UserName(un).
			XEmcNamespace(ns).
			Execute()
		if tErr == nil && tagsResp != nil && tagsResp.ListUserTagsResult != nil && tagsResp.ListUserTagsResult.Tags != nil {
			for _, t := range tagsResp.ListUserTagsResult.Tags {
				tags = append(tags, models.IAMUserTag{
					Key:   safe(t.Key),
					Value: safe(t.Value),
				})
			}
		}

		// Fetch access keys
		var accessKeys []models.IAMUserAccessKey
		keysResp, _, kErr := d.client.GenClient.IamApi.
			IamServiceListAccessKeys(ctx).
			UserName(un).
			XEmcNamespace(ns).
			Execute()
		if kErr == nil && keysResp != nil && keysResp.ListAccessKeysResult != nil && keysResp.ListAccessKeysResult.AccessKeyMetadata != nil {
			for _, key := range keysResp.ListAccessKeysResult.AccessKeyMetadata {
				accessKeys = append(accessKeys, models.IAMUserAccessKey{
					AccessKeyId: safe(key.AccessKeyId),
					CreateDate:  safe(key.CreateDate),
					Status:      safe(key.Status),
				})
			}
		}

		// Build final user object
		userObj := models.IAMUser{
			ID:                  safe(u.UserId),
			UserName:            safe(u.UserName),
			Arn:                 safe(u.Arn),
			Path:                safe(u.Path),
			CreateDate:          safe(u.CreateDate),
			PermissionsBoundary: types.StringValue(""),
			Tags:                tags,
			AccessKeys:          accessKeys,
		}

		finalUsers = append(finalUsers, userObj)
	}

	// 3. Save final state
	data.ID = types.StringValue("iam_user_datasource")
	data.Users = finalUsers

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
