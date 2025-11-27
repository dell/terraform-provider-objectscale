package provider

import (
	"context"
	"fmt"
	"terraform-provider-objectscale/internal/client"
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"
	"time"

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
	client *client.Client
}

func (d *IAMGroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_groups"
}

// Schema describes the data source arguments and results.
func (d *IAMGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieve IAM Groups from ObjectScale IAM.",
		Description:         "Retrieve IAM Groups from ObjectScale IAM.",
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

// Configure loads the API client.
func (d *IAMGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T.", req.ProviderData),
		)
		return
	}

	d.client = client
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
				"Error listing groups for user",
				fmt.Sprintf("Unable to list groups for user %s: %s", userName, err.Error()),
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
				"Error listing IAM groups",
				fmt.Sprintf("Failed to list groups: %s", err.Error()),
			)
			return
		}

		finalGroups = groups
	}

	// Save state
	data.ID = types.StringValue("iam_groups_" + ns)
	data.Groups = finalGroups

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "IAMGroupsDataSource.Read: completed")
}

func (d *IAMGroupsDataSource) listGroupsForUser(ctx context.Context, namespace, userName string) ([]models.IAMGroupModel, error) {
	var out []models.IAMGroupModel
	var marker *string

	for {
		req := d.client.GenClient.IamApi.IamServiceListGroupsForUser(ctx).
			UserName(userName).
			XEmcNamespace(namespace)

		if marker != nil {
			req = req.Marker(*marker)
		}

		resp, _, err := req.Execute()
		if err != nil {
			return nil, err
		}

		if resp.ListGroupsForUserResult != nil {
			for _, g := range resp.ListGroupsForUserResult.Groups {
				out = append(out, models.IAMGroupModel{
					GroupName: helper.TfString(g.GroupName),
					GroupId:   helper.TfString(g.GroupId),
					Arn:       helper.TfString(g.Arn),
					Path:      helper.TfString(g.Path),
				})
			}

			if resp.ListGroupsForUserResult.IsTruncated != nil &&
				*resp.ListGroupsForUserResult.IsTruncated &&
				resp.ListGroupsForUserResult.Marker != nil {
				marker = resp.ListGroupsForUserResult.Marker
				continue
			}
		}

		break
	}

	return out, nil
}

func (d *IAMGroupsDataSource) getGroupByName(ctx context.Context, namespace, groupName string) ([]models.IAMGroupModel, error) {
	var out []models.IAMGroupModel
	var marker *string

	for {
		req := d.client.GenClient.IamApi.IamServiceGetGroup(ctx).
			GroupName(groupName).
			XEmcNamespace(namespace)

		if marker != nil {
			req = req.Marker(*marker)
		}

		resp, _, err := req.Execute()
		if err != nil {
			return nil, fmt.Errorf("GetGroup API failed for %s: %w", groupName, err)
		}

		if resp.GetGroupResult != nil && resp.GetGroupResult.Group != nil {
			group := models.IAMGroupModel{
				GroupName:  types.StringValue(*resp.GetGroupResult.Group.GroupName),
				GroupId:    types.StringValue(*resp.GetGroupResult.Group.GroupId),
				Arn:        types.StringValue(*resp.GetGroupResult.Group.Arn),
				Path:       types.StringValue(*resp.GetGroupResult.Group.Path),
				CreateDate: types.StringValue(*resp.GetGroupResult.Group.CreateDate),
			}

			out = append(out, group)

			if resp.GetGroupResult.IsTruncated != nil && *resp.GetGroupResult.IsTruncated && resp.GetGroupResult.Marker != nil {
				marker = resp.GetGroupResult.Marker
				continue
			}
		}

		break
	}

	return out, nil
}

func (d *IAMGroupsDataSource) listAllGroups(ctx context.Context, namespace string) ([]models.IAMGroupModel, error) {
	var out []models.IAMGroupModel
	var marker *string

	for {
		req := d.client.GenClient.IamApi.IamServiceListGroups(ctx).
			XEmcNamespace(namespace)

		if marker != nil {
			req = req.Marker(*marker)
		}

		resp, _, err := req.Execute()
		if err != nil {
			return nil, err
		}

		if resp.ListGroupsResult != nil {
			for _, g := range resp.ListGroupsResult.Groups {
				out = append(out, models.IAMGroupModel{
					GroupName:  types.StringValue(g.GroupName),
					GroupId:    types.StringValue(g.GroupId),
					Arn:        types.StringValue(g.Arn),
					Path:       types.StringValue(g.Path),
					CreateDate: types.StringValue(g.CreateDate.Format(time.RFC3339)),
				})
			}

			if resp.ListGroupsResult.IsTruncated != nil &&
				*resp.ListGroupsResult.IsTruncated &&
				resp.ListGroupsResult.Marker != nil {
				marker = resp.ListGroupsResult.Marker
				continue
			}
		}

		break
	}

	return out, nil
}
