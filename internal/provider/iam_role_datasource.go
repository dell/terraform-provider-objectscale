package provider

import (
	"context"
	"fmt"

	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &IAMRoleDataSource{}
	_ datasource.DataSourceWithConfigure = &IAMRoleDataSource{}
)

type IAMRoleDataSource struct {
	datasourceProviderConfig
}

func NewIAMRoleDataSource() datasource.DataSource {
	return &IAMRoleDataSource{}
}

func (d *IAMRoleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_role"
}

func (d *IAMRoleDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This data source retrieves an ObjectScale IAM role by name/ID and gives its key attributes.",
		MarkdownDescription: "This data source retrieves an ObjectScale IAM role by name/ID and gives its key attributes.",

		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Internal ID for this data source.",
				MarkdownDescription: "Internal ID for this data source.",
			},

			"namespace": schema.StringAttribute{
				Required:            true,
				Description:         "Namespace to query IAM roles from.",
				MarkdownDescription: "Namespace to query IAM roles from.",
			},

			"role_name": schema.StringAttribute{
				Optional:            true,
				Description:         "Filter roles by name.",
				MarkdownDescription: "Filter roles by name.",
			},

			"roles": schema.ListNestedAttribute{
				Computed:            true,
				Description:         "List of IAM roles matching the provided filters.",
				MarkdownDescription: "List of IAM roles matching the provided filters.",

				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"arn": schema.StringAttribute{
							Computed:            true,
							Description:         "IAM role ARN.",
							MarkdownDescription: "IAM role ARN.",
						},

						"role_name": schema.StringAttribute{
							Computed:            true,
							Description:         "IAM role name.",
							MarkdownDescription: "IAM role name.",
						},

						"role_id": schema.StringAttribute{
							Computed:            true,
							Description:         "Unique ObjectScale IAM role ID (RoleId).",
							MarkdownDescription: "Unique ObjectScale IAM role ID (RoleId).",
						},

						"path": schema.StringAttribute{
							Computed:            true,
							Description:         "IAM role path.",
							MarkdownDescription: "IAM role path.",
						},

						"description": schema.StringAttribute{
							Computed:            true,
							Description:         "Role description.",
							MarkdownDescription: "Role description.",
						},

						"create_date": schema.StringAttribute{
							Computed:            true,
							Description:         "Timestamp when the role was created.",
							MarkdownDescription: "Timestamp when the role was created.",
						},

						"max_session_duration": schema.Int64Attribute{
							Computed:            true,
							Description:         "Maximum session duration allowed for the IAM role.",
							MarkdownDescription: "Maximum session duration allowed for the IAM role.",
						},

						"assume_role_policy": schema.StringAttribute{
							Computed:            true,
							Description:         "The trust policy document defining who can assume the role.",
							MarkdownDescription: "The trust policy document defining who can assume the role.",
						},

						"permissions_boundary": schema.SingleNestedAttribute{
							Computed:            true,
							Description:         "Permissions boundary applied to this role.",
							MarkdownDescription: "Permissions boundary applied to this role.",

							Attributes: map[string]schema.Attribute{
								"permissions_boundary_arn": schema.StringAttribute{
									Computed:            true,
									Description:         "ARN of the permissions boundary policy.",
									MarkdownDescription: "ARN of the permissions boundary policy.",
								},
								"permissions_boundary_type": schema.StringAttribute{
									Computed:            true,
									Description:         "Type of permissions boundary (always 'Policy').",
									MarkdownDescription: "Type of permissions boundary (always 'Policy').",
								},
							},
						},

						"tags": schema.ListNestedAttribute{
							Computed:            true,
							Description:         "List of tags assigned to the role.",
							MarkdownDescription: "List of tags assigned to the role.",

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
					},
				},
			},
		},
	}
}

func (d *IAMRoleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.IAMRoleDatasourceModel

	// Load inputs
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ns := data.Namespace.ValueString()
	var finalRoles []models.IAMRole

	// --- CASE 1: Get a single role by RoleName ---
	if !data.RoleName.IsNull() {
		roleName := data.RoleName.ValueString()

		roles, err := d.getRole(ctx, ns, roleName)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error retrieving IAM role",
				fmt.Sprintf("Failed retrieving role %s: %s", roleName, err.Error()),
			)
			return
		}

		finalRoles = roles
	} else {
		// --- CASE 2: List all roles ---
		allRoles, err := d.listAllRoles(ctx, ns)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error listing IAM roles", err.Error(),
			)
			return
		}

		finalRoles = allRoles
	}

	// Save state
	data.ID = types.StringValue("iam_role_datasource")
	data.Roles = finalRoles
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *IAMRoleDataSource) getRole(ctx context.Context, namespace, roleName string) ([]models.IAMRole, error) {
	getResp, _, err := d.client.GenClient.IamApi.
		IamServiceGetRole(ctx).
		RoleName(roleName).
		XEmcNamespace(namespace).
		Execute()

	if err != nil {
		return nil, err
	}

	if getResp == nil || getResp.GetRoleResult == nil || getResp.GetRoleResult.Role == nil {
		return nil, err
	}

	r := getResp.GetRoleResult.Role

	// Convert MaxSessionDuration (*int32 → types.Int64)
	var maxSessionDuration *int64
	if r.MaxSessionDuration != nil {
		v := int64(*r.MaxSessionDuration)
		maxSessionDuration = &v
	}

	roleObj := models.IAMRole{
		RoleId:              helper.TfString(r.RoleId),
		RoleName:            helper.TfString(r.RoleName),
		Arn:                 helper.TfString(r.Arn),
		AssumeRolePolicy:    helper.TfString(r.AssumeRolePolicyDocument),
		Path:                helper.TfString(r.Path),
		Description:         helper.TfString(r.Description),
		CreateDate:          helper.TfString(r.CreateDate),
		MaxSessionDuration:  helper.TfInt64(maxSessionDuration),
		PermissionsBoundary: convertPermissionsBoundary(r.PermissionsBoundary),
		Tags:                convertTags(r.Tags),
	}

	return []models.IAMRole{roleObj}, nil
}

func (d *IAMRoleDataSource) listAllRoles(ctx context.Context, namespace string) ([]models.IAMRole, error) {
	items, err := helper.GetAllInstances(
		d.client.GenClient.IamApi.
			IamServiceListRoles(ctx).
			XEmcNamespace(namespace),
	)
	if err != nil {
		return nil, err
	}

	var roles []models.IAMRole

	for _, r := range items {

		// Convert MaxSessionDuration (*int32 → types.Int64)
		var maxSessionDuration types.Int64
		if r.MaxSessionDuration != nil {
			v := int64(*r.MaxSessionDuration)
			maxSessionDuration = helper.TfInt64(&v)
		} else {
			maxSessionDuration = types.Int64Null()
		}

		roles = append(roles, models.IAMRole{
			RoleId:              helper.TfString(r.RoleId),
			RoleName:            helper.TfString(r.RoleName),
			Arn:                 helper.TfString(r.Arn),
			AssumeRolePolicy:    helper.TfString(r.AssumeRolePolicyDocument),
			Path:                helper.TfString(r.Path),
			Description:         helper.TfString(r.Description),
			CreateDate:          helper.TfString(r.CreateDate),
			MaxSessionDuration:  maxSessionDuration,
			PermissionsBoundary: convertPermissionsBoundary(r.PermissionsBoundary),
			Tags:                convertTags(r.Tags),
		})
	}

	return roles, nil
}

func convertTags(tags []clientgen.IamTagKeyValue) []models.IAMRoleTag {
	var result []models.IAMRoleTag
	for _, t := range tags {
		result = append(result, models.IAMRoleTag{
			Key:   helper.TfString(t.Key),
			Value: helper.TfString(t.Value),
		})
	}
	return result
}

func convertPermissionsBoundary(pb *clientgen.IamRolePermissionsBoundary) *models.IAMRolePermissionsBoundary {
	if pb == nil {
		return nil
	}
	return &models.IAMRolePermissionsBoundary{
		PermissionsBoundaryArn:  helper.TfString(pb.PermissionsBoundaryArn),
		PermissionsBoundaryType: helper.TfString(pb.PermissionsBoundaryType),
	}
}
