package provider

import (
	"context"
	"fmt"

	"terraform-provider-objectscale/internal/client"
	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/helper"
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
		Description:         "Fetch IAM user information for a specific ObjectScale namespace.",
		MarkdownDescription: "Fetch IAM user information for a specific ObjectScale namespace.",

		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Internal ID for this data source.",
				MarkdownDescription: "Internal ID for this data source.",
			},

			"namespace": schema.StringAttribute{
				Required:            true,
				Description:         "Namespace containing IAM users.",
				MarkdownDescription: "Namespace containing IAM users.",
			},

			"username": schema.StringAttribute{
				Optional:            true,
				Description:         "Filter users by username.",
				MarkdownDescription: "Filter users by username.",
			},

			"groupname": schema.StringAttribute{
				Optional:            true,
				Description:         "Filter users who belong to the given group name.",
				MarkdownDescription: "Filter users who belong to the given group name.",
			},

			"users": schema.ListNestedAttribute{
				Computed:            true,
				Description:         "List of IAM users matching the filters.",
				MarkdownDescription: "List of IAM users matching the filters.",

				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"username": schema.StringAttribute{
							Computed:            true,
							Description:         "IAM username.",
							MarkdownDescription: "IAM username.",
						},

						"id": schema.StringAttribute{
							Computed:            true,
							Description:         "Unique ObjectScale IAM user ID (maps to UserId).",
							MarkdownDescription: "Unique ObjectScale IAM user ID (maps to UserId).",
						},

						"arn": schema.StringAttribute{
							Computed:            true,
							Description:         "IAM user ARN.",
							MarkdownDescription: "IAM user ARN.",
						},

						"create_date": schema.StringAttribute{
							Computed:            true,
							Description:         "The timestamp when the user was created.",
							MarkdownDescription: "The timestamp when the user was created.",
						},

						"path": schema.StringAttribute{
							Computed:            true,
							Description:         "IAM user path.",
							MarkdownDescription: "IAM user path.",
						},

						"permissions_boundary": schema.StringAttribute{
							Computed:            true,
							Description:         "Policy ARN used as permissions boundary. Can be empty if not set.",
							MarkdownDescription: "Policy ARN used as permissions boundary. Can be empty if not set.",
						},

						"tags": schema.ListNestedAttribute{
							Computed:            true,
							Description:         "List of tags assigned to the user (sorted by key in Read function).",
							MarkdownDescription: "List of tags assigned to the user (sorted by key in Read function).",

							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Computed:            true,
										Description:         "Tag key.",
										MarkdownDescription: "Tag key.",
									},
									"value": schema.StringAttribute{
										Computed:            true,
										Description:         "Tag value.",
										MarkdownDescription: "Tag value.",
									},
								},
							},
						},

						"access_keys": schema.ListNestedAttribute{
							Computed:            true,
							Description:         "List of access keys for the IAM user.",
							MarkdownDescription: "List of access keys for the IAM user.",

							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"access_key_id": schema.StringAttribute{
										Computed:            true,
										Description:         "Access key ID.",
										MarkdownDescription: "Access key ID.",
									},

									"create_date": schema.StringAttribute{
										Computed:            true,
										Description:         "Timestamp when the access key was created.",
										MarkdownDescription: "Timestamp when the access key was created.",
									},

									"status": schema.StringAttribute{
										Computed:            true,
										Description:         "Status of the access key (Active/Inactive).",
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

	var finalUsers []models.IAMUser
	if !data.Username.IsNull() {
		// CASE 1 — USERNAME PROVIDED → DIRECTLY CALL GetUser
		username := data.Username.ValueString()

		users, err := d.getUser(ctx, ns, username)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error retrieving IAM user by name",
				fmt.Sprintf("Failed retrieving user %s: %v", username, err),
			)
			return
		}
		finalUsers = users
	} else if !data.Groupname.IsNull() {
		// CASE 2 — GROUP FILTER (preload usernames)
		groupName := data.Groupname.ValueString()

		groupUsers, err := d.listUsersByGroup(ctx, ns, groupName)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error retrieving IAM group",
				fmt.Sprintf("Unable to retrieve IAM group %q: %v", groupName, err),
			)
			return
		}

		finalUsers = append(finalUsers, groupUsers...)
	} else {
		// CASE 3 — LIST ALL USERS (when username not provided)
		allUsers, err := d.listAllUsers(ctx, ns)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error listing IAM users",
				fmt.Sprintf("Error listing IAM users: %v", err),
			)
			return
		}
		finalUsers = append(finalUsers, allUsers...)
	}

	// ---- fetch tags and access keys for each user ----
	for i, u := range finalUsers {
		username := u.UserName.ValueString()

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
					Key:   helper.TfString(t.Key),
					Value: helper.TfString(t.Value),
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
					AccessKeyId: helper.TfString(k.AccessKeyId),
					CreateDate:  helper.TfString(k.CreateDate),
					Status:      helper.TfString(k.Status),
				})
			}
		}

		finalUsers[i].Tags = tags
		finalUsers[i].AccessKeys = accessKeys
	}

	// save state
	data.ID = types.StringValue("iam_user_datasource")
	data.Users = finalUsers
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *IAMUserDataSource) getUser(ctx context.Context, namespace, username string) ([]models.IAMUser, error) {
	// CASE 1 — USERNAME PROVIDED → DIRECTLY CALL GetUser
	getResp, _, err := d.client.GenClient.IamApi.
		IamServiceGetUser(ctx).
		UserName(username).
		XEmcNamespace(namespace).
		Execute()

	if err != nil {
		return nil, err
	}

	if getResp == nil || getResp.GetUserResult == nil || getResp.GetUserResult.User == nil {
		return nil, fmt.Errorf("no result in response")
	}

	u := getResp.GetUserResult.User

	// ---- final user object ----
	userObj := models.IAMUser{
		ID:                  helper.TfString(u.UserId),
		UserName:            helper.TfString(u.UserName),
		Arn:                 helper.TfString(u.Arn),
		Path:                helper.TfString(u.Path),
		CreateDate:          helper.TfString(u.CreateDate),
		PermissionsBoundary: d.getPermissionBoundary(u.PermissionsBoundary),
	}
	return []models.IAMUser{userObj}, nil
}

func (d *IAMUserDataSource) listUsersByGroup(ctx context.Context, namespace, groupName string) ([]models.IAMUser, error) {
	items, err := helper.GetAllInstances(d.client.GenClient.IamApi.IamServiceGetGroup(ctx).
		GroupName(groupName).
		XEmcNamespace(namespace))

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	var users []models.IAMUser
	for _, u := range items {
		users = append(users, models.IAMUser{
			UserName:            helper.TfString(u.UserName),
			ID:                  helper.TfString(u.UserId),
			Arn:                 helper.TfString(u.Arn),
			Path:                helper.TfString(u.Path),
			CreateDate:          helper.TfString(u.CreateDate),
			PermissionsBoundary: types.StringNull(),
		})
	}
	return users, nil
}

func (d *IAMUserDataSource) listAllUsers(ctx context.Context, namespace string) ([]models.IAMUser, error) {

	items, err := helper.GetAllInstances(d.client.GenClient.IamApi.IamServiceListUsers(ctx).XEmcNamespace(namespace))
	if err != nil {
		return nil, err
	}

	var users []models.IAMUser
	for _, u := range items {
		users = append(users, models.IAMUser{
			UserName:            helper.TfString(u.UserName),
			ID:                  helper.TfString(u.UserId),
			Arn:                 helper.TfString(u.Arn),
			Path:                helper.TfString(u.Path),
			CreateDate:          helper.TfString(u.CreateDate),
			PermissionsBoundary: types.StringNull(),
		})
	}
	return users, nil
}

func (d *IAMUserDataSource) getPermissionBoundary(prmissionsBoundary *clientgen.IamServiceGetUserResponseGetUserResultUserPermissionsBoundary) types.String {
	if prmissionsBoundary != nil && prmissionsBoundary.PermissionsBoundaryArn != nil {
		return types.StringValue(*prmissionsBoundary.PermissionsBoundaryArn)
	}
	return types.StringValue("")
}
