package provider

import (
	"context"
	"fmt"
	"strings"

	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &ObjectUserDataSource{}
	_ datasource.DataSourceWithConfigure = &ObjectUserDataSource{}
)

type ObjectUserDataSource struct {
	datasourceProviderConfig
}

func NewObjectUserDataSource() datasource.DataSource {
	return &ObjectUserDataSource{}
}

func (d *ObjectUserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_object_user"
}

func (d *ObjectUserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Fetch list of Object users.",
		MarkdownDescription: "Fetch list of Object users.",

		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Internal ID for this data source.",
				MarkdownDescription: "Internal ID for this data source.",
			},

			"namespace": schema.StringAttribute{
				Optional:            true,
				Description:         "Namespace containing object users.",
				MarkdownDescription: "Namespace containing object users.",
			},

			"name": schema.StringAttribute{
				Optional:            true,
				Description:         "Filter object users by username.",
				MarkdownDescription: "Filter object users by username.",
			},
			"tag": schema.StringAttribute{
				Optional:            true,
				Description:         "Filter object users by tag. 'tag' and 'value' are required together.",
				MarkdownDescription: "Filter object users by tag. 'tag' and 'value' are required together.",
			},
			"value": schema.StringAttribute{
				Optional:            true,
				Description:         "Filter object users by tag value. 'tag' and 'value' are required together.",
				MarkdownDescription: "Filter object users by tag value. 'tag' and 'value' are required together.",
			},
			"users": schema.ListNestedAttribute{
				Computed:            true,
				Description:         "List of object users matching the filters.",
				MarkdownDescription: "List of object users matching the filters.",

				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description:         "Name of the user.",
							MarkdownDescription: "Name of the user.",
							Computed:            true,
						},
						"id": schema.StringAttribute{
							Description:         "ID of the user.",
							MarkdownDescription: "ID of the user.",
							Computed:            true,
						},
						"namespace": schema.StringAttribute{
							Description:         "Namespace to which the user belongs to.",
							MarkdownDescription: "Namespace to which the user belongs to.",
							Computed:            true,
						},
						"created": schema.StringAttribute{
							Description:         "Timestamp of the creation of the object user.",
							MarkdownDescription: "Timestamp of the creation of the object user.",
							Computed:            true,
						},
						"locked": schema.BoolAttribute{
							Description:         "Lock status of the object user.",
							MarkdownDescription: "Lock status of the object user.",
							Computed:            true,
						},
						"secret_keys": schema.SingleNestedAttribute{
							Computed:            true,
							Description:         "List of secret keys for the object user.",
							MarkdownDescription: "List of secret keys for the object user.",
							Attributes: map[string]schema.Attribute{

								"secret_key_1_id": schema.StringAttribute{
									Computed:            true,
									Description:         "ID of the first secret key.",
									MarkdownDescription: "ID of the first secret key.",
								},

								"secret_key_1": schema.StringAttribute{
									Computed:            true,
									Description:         "First secret key for the object user.",
									MarkdownDescription: "First secret key for the object user.",
								},

								"secret_key_1_exist": schema.BoolAttribute{
									Computed:            true,
									Description:         "If the first secret key exists.",
									MarkdownDescription: "If the first secret key exists.",
								},
								"key_timestamp_1": schema.StringAttribute{
									Computed:            true,
									Description:         "Timestamp when the first secret key was created.",
									MarkdownDescription: "Timestamp when the first secret key was created.",
								},
								"key_expiry_timestamp_1": schema.StringAttribute{
									Computed:            true,
									Description:         "Timestamp when the first secret key expires.",
									MarkdownDescription: "Timestamp when the first secret key expires.",
								},

								"secret_key_2_id": schema.StringAttribute{
									Computed:            true,
									Description:         "ID of the second secret key.",
									MarkdownDescription: "ID of the second secret key.",
								},

								"secret_key_2": schema.StringAttribute{
									Computed:            true,
									Description:         "Second secret key for the object user.",
									MarkdownDescription: "Second secret key for the object user.",
								},

								"secret_key_2_exist": schema.BoolAttribute{
									Computed:            true,
									Description:         "If the second secret key exists.",
									MarkdownDescription: "If the second secret key exists.",
								},
								"key_timestamp_2": schema.StringAttribute{
									Computed:            true,
									Description:         "Timestamp when the second secret key was created.",
									MarkdownDescription: "Timestamp when the second secret key was created.",
								},
								"key_expiry_timestamp_2": schema.StringAttribute{
									Computed:            true,
									Description:         "Timestamp when the second secret key expires.",
									MarkdownDescription: "Timestamp when the second secret key expires.",
								},
							},
						},
						"tags": schema.ListNestedAttribute{
							Description:         "Tags associated to the user.",
							MarkdownDescription: "Tags associated to the user.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Key of the tag associated to the user.",
										MarkdownDescription: "Key of the tag associated to the user.",
										Computed:            true,
									},
									"value": schema.StringAttribute{
										Description:         "Value of the tag associated to the user.",
										MarkdownDescription: "Value of the tag associated to the user.",
										Computed:            true,
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

func (d *ObjectUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.ObjectUserDatasourceModel
	// Load inputs
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Tag.ValueString() != "" && data.Value.ValueString() == "" || data.Tag.ValueString() == "" && data.Value.ValueString() != "" {
		resp.Diagnostics.AddError(
			"Error retrieving object user",
			"value and tag are required together",
		)
		return
	}

	var finalUsers []models.ObjectUser
	if !data.Name.IsNull() {
		// CASE 1 — USERNAME PROVIDED
		username := data.Name.ValueString()

		users, err := d.listUsersByName(ctx, username)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error retrieving object user by name",
				fmt.Sprintf("Failed retrieving user %s: %s", username, err),
			)
			return
		}
		finalUsers = users
	} else if !data.Tag.IsNull() {
		// CASE 2 — TAG FILTER
		Namespace := data.Namespace.ValueString()
		Tag := data.Tag.ValueString()
		Value := data.Value.ValueString()

		tagUsers, err := d.listUsersByTag(ctx, Namespace, Tag, Value)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error retrieving users by tag",
				fmt.Sprintf("Unable to retrieve object users with tag filter: %s", err),
			)
			return
		}

		finalUsers = append(finalUsers, tagUsers...)
	} else if !data.Namespace.IsNull() {
		// CASE 3 — NAMESPACE FILTER
		Namespace := data.Namespace.ValueString()

		namespaceUsers, err := d.listUsersByNamespace(ctx, Namespace)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error retrieving namespace object users",
				fmt.Sprintf("Unable to retrieve namespace object users %s: %s", Namespace, err),
			)
			return
		}

		finalUsers = append(finalUsers, namespaceUsers...)
	} else {
		// CASE 4 — LIST ALL USERS
		allUsers, err := d.listAllUsers(ctx)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error listing Object users",
				fmt.Sprintf("Error listing Object users: %s", err),
			)
			return
		}
		finalUsers = append(finalUsers, allUsers...)
	}

	// save state
	data.Id = types.StringValue("object_user_datasource")
	data.Users = finalUsers
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *ObjectUserDataSource) listUsersByNamespace(ctx context.Context, namespace string) ([]models.ObjectUser, error) {

	listResp, _, err := d.client.GenClient.UserManagementApi.
		UserManagementServiceGetUsersForNamespace(ctx, namespace).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("listing object users for namespace %q: %w", namespace, err)
	}

	items := listResp.Blobuser

	var users []models.ObjectUser
	var user_list models.ObjectUser
	for _, u := range items {
		user_list, err = d.getUser(ctx, u.Userid)
		if err != nil {
			return nil, fmt.Errorf("listing object users for namespace %q: %w", namespace, err)
		}
		users = append(users, user_list)
	}

	return users, nil
}

func (d *ObjectUserDataSource) listUsersByName(ctx context.Context, name string) ([]models.ObjectUser, error) {

	var users []models.ObjectUser
	var user_list models.ObjectUser
	user_list, err := d.getUser(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("listing users for user %q: %w", name, err)
	}
	users = append(users, user_list)

	return users, nil
}

func (d *ObjectUserDataSource) listUsersByTag(ctx context.Context, namespace, tag, value string) ([]models.ObjectUser, error) {

	req := d.client.GenClient.UserManagementApi.
		UserManagementServiceQueryUsers(ctx)

	if ns := strings.TrimSpace(namespace); ns != "" {
		req = req.Namespace(ns)
	}

	if tg := strings.TrimSpace(tag); tg != "" {
		req = req.Tag(tg)
	}

	if val := strings.TrimSpace(value); val != "" {
		req = req.Value(val)
	}

	// Execute the request after conditionally setting fields.
	listResp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("listing users for tag %q and value %q: %w", tag, value, err)
	}

	items := listResp.Blobuser

	var users []models.ObjectUser
	var user_list models.ObjectUser
	for _, u := range items {
		user_list, err = d.getUser(ctx, u.Userid)
		if err != nil {
			return nil, fmt.Errorf("listing users for tag %q and value %q: %w", tag, value, err)
		}
		users = append(users, user_list)
	}

	return users, nil
}
func (d *ObjectUserDataSource) listAllUsers(ctx context.Context) ([]models.ObjectUser, error) {

	listResp, _, err := d.client.GenClient.UserManagementApi.
		UserManagementServiceGetAllUsers(ctx).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("listing all users: %w", err)
	}

	items := listResp.Blobuser

	var users []models.ObjectUser
	var user_list models.ObjectUser
	for _, u := range items {
		user_list, err = d.getUser(ctx, u.Userid)
		if err != nil {
			return nil, fmt.Errorf("listing all users: %w", err)
		}
		users = append(users, user_list)
	}

	return users, nil
}

func (d *ObjectUserDataSource) getUser(ctx context.Context, username string) (models.ObjectUser, error) {
	objectUser, _, err_user := d.client.GenClient.UserManagementApi.
		UserManagementServiceGetUserInfo(ctx, username).
		Execute()
	if err_user != nil {
		return models.ObjectUser{}, fmt.Errorf("reading user %q: %w", username, err_user)
	}

	obj_access_key, _, err_access_key := d.client.GenClient.UserSecretKeyApi.
		UserSecretKeyServiceGetKeysForUser(ctx, username).
		Execute()

	if err_access_key != nil {
		return models.ObjectUser{}, fmt.Errorf("reading user secret keys %q: %w", username, err_access_key)
	}

	var obj_user models.ObjectUser
	obj_user = models.ObjectUser{
		Id:        helper.TfString(&objectUser.Name),
		Name:      helper.TfString(&objectUser.Name),
		Namespace: helper.TfString(&objectUser.Namespace),
		Locked:    helper.TfBool(&objectUser.Locked),
		Created:   helper.TfString(&objectUser.Created),
		Tags: helper.ListNotNull(objectUser.Tag,
			func(v clientgen.UserManagementServiceAddUserRequestTagsInner) types.Object {
				return helper.Object(models.ObjectUserTags{
					Name:  helper.TfStringNN(v.Name),
					Value: helper.TfStringNN(v.Value),
				})
			}),
		SecretKey: models.ObjectUserAccessKey{
			SecretKey1Id:        helper.TfString(obj_access_key.SecretKey1Id),
			SecretKey1:          helper.TfString(obj_access_key.SecretKey1),
			SecretKey1Exist:     helper.TfBool(obj_access_key.SecretKey1Exist),
			KeyTimestamp1:       helper.TfString(obj_access_key.KeyTimestamp1),
			KeyExpiryTimestamp1: helper.TfString(obj_access_key.KeyExpiryTimestamp1),
			SecretKey2Id:        helper.TfString(obj_access_key.SecretKey2Id),
			SecretKey2:          helper.TfString(obj_access_key.SecretKey2),
			SecretKey2Exist:     helper.TfBool(obj_access_key.SecretKey2Exist),
			KeyTimestamp2:       helper.TfString(obj_access_key.KeyTimestamp2),
			KeyExpiryTimestamp2: helper.TfString(obj_access_key.KeyExpiryTimestamp2),
		},
	}

	return obj_user, nil
}
