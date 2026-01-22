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
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &IAMGroupsDataSource{}

func NewIAMGroupsDataSource() datasource.DataSource {
	return &IAMGroupsDataSource{}
}

type IAMGroupsDataSource struct {
	datasourceProviderConfig
}

func (d *IAMGroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_groups"
}

// Schema describes the data source arguments and results.
func (d *IAMGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This data source retrieves key attributes of Dell ObjectScale IAM groups.",
		Description:         "This data source retrieves key attributes of Dell ObjectScale IAM groups.",
		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Description:         "Identifier",
				MarkdownDescription: "Identifier",
				Computed:            true,
			},

			"namespace": schema.StringAttribute{
				Description:         "Name of the namespace.",
				MarkdownDescription: "Name of the namespace.",
				Required:            true,
			},

			"group_name": schema.StringAttribute{
				Description:         "Name of the IAM Group.",
				MarkdownDescription: "Name of the IAM Group.",
				Optional:            true,
			},

			"user_name": schema.StringAttribute{
				Description:         "Name of the IAM User.",
				MarkdownDescription: "Name of the IAM User to filter groups for.",
				Optional:            true,
			},

			"groups": schema.ListNestedAttribute{
				Description:         "List of IAM Groups",
				MarkdownDescription: "List of IAM Groups",
				Computed:            true,

				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"group_name": schema.StringAttribute{
							Description:         "Name of the IAM Group.",
							MarkdownDescription: "Name of the IAM Group.",
							Computed:            true,
						},

						"group_id": schema.StringAttribute{
							Description:         "Unique identifier for the group.",
							MarkdownDescription: "Unique identifier for the group.",
							Computed:            true,
						},

						"arn": schema.StringAttribute{
							Description:         "ARN of the IAM Group.",
							MarkdownDescription: "ARN of the IAM Group.",
							Computed:            true,
						},

						"path": schema.StringAttribute{
							Description:         "Path associated with the IAM Group.",
							MarkdownDescription: "Path associated with the IAM Group.",
							Computed:            true,
						},

						"create_date": schema.StringAttribute{
							Description:         "ISO 8601 creation timestamp of the IAM Group.",
							MarkdownDescription: "ISO 8601 creation timestamp of the IAM Group.",
							Computed:            true,
						},

						"users": schema.ListAttribute{
							Description:         "List of users belonging to this IAM Group.",
							MarkdownDescription: "List of users belonging to this IAM Group.",
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *IAMGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Trace(ctx, "IAMGroupsDataSource.Read: start")

	var data models.IAMGroupsDatasourceModel

	// Load config
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ns := data.Namespace.ValueString()

	var finalGroups []models.IAMGroupModel

	// CASE 1 — filter by user_name
	if !data.UserName.IsNull() && !data.UserName.IsUnknown() {

		userName := data.UserName.ValueString()
		groups, err := d.listGroupsForUser(ctx, ns, userName)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error retrieving groups for user",
				fmt.Sprintf("Unable to retrieve groups for user %s: %s", userName, err.Error()),
			)
			return
		}

		finalGroups = groups

	} else if !data.GroupName.IsNull() && !data.GroupName.IsUnknown() {

		// CASE 2 — filter by group_name
		groupName := data.GroupName.ValueString()

		groups, err := d.getGroupByName(ctx, ns, groupName)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error retrieving IAM group",
				fmt.Sprintf("Failed to retrieve group %s: %s", groupName, err.Error()),
			)
			return
		}

		finalGroups = groups

	} else {

		// CASE 3 — list all groups
		groups, err := d.listAllGroups(ctx, ns)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error retrieving IAM groups",
				fmt.Sprintf("Failed to retrieve groups: %s", err.Error()),
			)
			return
		}

		finalGroups = groups
	}

	if len(finalGroups) == 0 {
		resp.Diagnostics.AddError(
			"Invalid namespace",
			"The namespace does not exist.",
		)
		return
	}

	// Save state
	data.ID = types.StringValue("iam_groups_" + ns)
	data.Groups = finalGroups

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "IAMGroupsDataSource.Read: completed")
}

func (d *IAMGroupsDataSource) listGroupsForUser(ctx context.Context, namespace, userName string) ([]models.IAMGroupModel, error) {
	req := d.client.GenClient.IamApi.IamServiceListGroupsForUser(ctx).
		UserName(userName).
		XEmcNamespace(namespace)

	items, err := helper.GetAllInstances(req)
	if err != nil {
		return nil, err
	}

	return helper.SliceTransform(items, func(v clientgen.IamServiceListGroupsForUserResponseListGroupsForUserResultGroupsInner) models.IAMGroupModel {
		return models.IAMGroupModel{
			GroupName: helper.TfString(v.GroupName),
			GroupId:   helper.TfString(v.GroupId),
			Arn:       helper.TfString(v.Arn),
			Path:      helper.TfString(v.Path),
		}
	}), nil
}

func (d *IAMGroupsDataSource) getGroupByName(ctx context.Context, namespace, groupName string) ([]models.IAMGroupModel, error) {
	v, _, err := d.client.GenClient.IamApi.IamServiceGetGroup(ctx).
		GroupName(groupName).
		XEmcNamespace(namespace).Execute()
	if err != nil {
		return nil, err
	}

	return []models.IAMGroupModel{
		{
			GroupName:  helper.TfString(v.GetGroupResult.Group.GroupName),
			GroupId:    helper.TfString(v.GetGroupResult.Group.GroupId),
			Arn:        helper.TfString(v.GetGroupResult.Group.Arn),
			Path:       helper.TfString(v.GetGroupResult.Group.Path),
			CreateDate: helper.TfString(v.GetGroupResult.Group.CreateDate),
		},
	}, nil
}

func (d *IAMGroupsDataSource) listAllGroups(ctx context.Context, namespace string) ([]models.IAMGroupModel, error) {
	req := d.client.GenClient.IamApi.IamServiceListGroups(ctx).
		XEmcNamespace(namespace)

	items, err := helper.GetAllInstances(req)
	if err != nil {
		return nil, err
	}

	return helper.SliceTransform(items, func(v clientgen.IamServiceListGroupsResponseListGroupsResultGroupsInner) models.IAMGroupModel {
		return models.IAMGroupModel{
			GroupName:  types.StringValue(v.GroupName),
			GroupId:    types.StringValue(v.GroupId),
			Arn:        types.StringValue(v.Arn),
			Path:       types.StringValue(v.Path),
			CreateDate: types.StringValue(v.CreateDate),
		}
	}), nil
}
