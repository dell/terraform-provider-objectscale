package provider

import (
	"context"
	"fmt"
	"terraform-provider-objectscale/internal/client"
	"terraform-provider-objectscale/internal/clientgen"
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
		tflog.Error(ctx, "IAMGroupsDataSource.Read: config load failed")
		return
	}

	namespace := data.Namespace.ValueString()
	ctx = tflog.SetField(ctx, "namespace", namespace)

	var filterGroupName, filterUserName *string
	if !data.GroupName.IsNull() && !data.GroupName.IsUnknown() {
		val := data.GroupName.ValueString()
		filterGroupName = &val
		ctx = tflog.SetField(ctx, "group_name", val)
	}

	if !data.UserName.IsNull() && !data.UserName.IsUnknown() {
		val := data.UserName.ValueString()
		filterUserName = &val
		ctx = tflog.SetField(ctx, "user_name", val)
	}

	tflog.Debug(ctx, "IAMGroupsDataSource.Read: config loaded")

	var allGroups []clientgen.IamServiceListGroupsResponseListGroupsResultGroupsInner
	var marker *string

	// Fetch groups with proper API
	for {
		tflog.Trace(ctx, "IAMGroupsDataSource.Read: API call", map[string]interface{}{
			"marker": marker,
		})

		if filterUserName != nil {
			// ListGroupsForUser API
			var listRespForUser *clientgen.IamServiceListGroupsForUserResponse
			apiReq := d.client.GenClient.IamApi.IamServiceListGroupsForUser(ctx).
				XEmcNamespace(namespace).
				UserName(*filterUserName)

			if marker != nil {
				apiReq = apiReq.Marker(*marker)
			}

			listRespForUser, _, err := apiReq.Execute()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error Listing IAM Groups for User",
					fmt.Sprintf("Could not list groups for user %s: %s", *filterUserName, err),
				)
				return
			}

			if listRespForUser.ListGroupsForUserResult != nil {
				for _, g := range listRespForUser.ListGroupsForUserResult.Groups {
					allGroups = append(allGroups, clientgen.IamServiceListGroupsResponseListGroupsResultGroupsInner{
						GroupName: stringValue(g.GroupName),
						GroupId:   stringValue(g.GroupId),
						Arn:       stringValue(g.Arn),
						Path:      stringValue(g.Path),
					})
				}

				if listRespForUser.ListGroupsForUserResult.IsTruncated != nil &&
					*listRespForUser.ListGroupsForUserResult.IsTruncated &&
					listRespForUser.ListGroupsForUserResult.Marker != nil {
					marker = listRespForUser.ListGroupsForUserResult.Marker
					tflog.Debug(ctx, "IAMGroupsDataSource.Read: continuing pagination for user", map[string]interface{}{
						"next_marker": *marker,
					})
					continue
				}
			}
		} else {
			// ListGroups API
			var listResp *clientgen.IamServiceListGroupsResponse
			apiReq := d.client.GenClient.IamApi.IamServiceListGroups(ctx).XEmcNamespace(namespace)
			if marker != nil {
				apiReq = apiReq.Marker(*marker)
			}

			listResp, _, err := apiReq.Execute()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error Listing IAM Groups",
					fmt.Sprintf("Could not list groups: %s", err),
				)
				return
			}

			if listResp.ListGroupsResult != nil {
				allGroups = append(allGroups, listResp.ListGroupsResult.Groups...)
				if listResp.ListGroupsResult.IsTruncated != nil &&
					*listResp.ListGroupsResult.IsTruncated &&
					listResp.ListGroupsResult.Marker != nil {
					marker = listResp.ListGroupsResult.Marker
					tflog.Debug(ctx, "IAMGroupsDataSource.Read: continuing pagination", map[string]interface{}{
						"next_marker": *marker,
					})
					continue
				}
			}
		}

		break
	}

	tflog.Debug(ctx, "IAMGroupsDataSource.Read: total groups fetched", map[string]interface{}{
		"total": len(allGroups),
	})

	// Filter by group_name if provided
	var filtered []clientgen.IamServiceListGroupsResponseListGroupsResultGroupsInner
	if filterGroupName != nil {
		for _, g := range allGroups {
			if g.GroupName == *filterGroupName {
				filtered = append(filtered, g)
			}
		}
	} else {
		filtered = allGroups
	}

	// Build final state
	var groups []models.IAMGroupModel
	for _, g := range filtered {
		groupCtx := tflog.SetField(ctx, "current_group", g.GroupName)

		// Always initialize users slice
		users := []types.String{}

		// Fetch users if group_name filter is provided
		if filterGroupName != nil {
			getResp, _, err := d.client.GenClient.IamApi.IamServiceGetGroup(ctx).
				XEmcNamespace(namespace).
				GroupName(g.GroupName).
				Execute()
			if err == nil && getResp.GetGroupResult != nil {
				for _, u := range getResp.GetGroupResult.Users {
					users = append(users, types.StringValue(u))
				}
			} else if err != nil {
				tflog.Warn(groupCtx, "IAMGroupsDataSource.Read: GetGroup API failed", map[string]interface{}{
					"error": err.Error(),
				})
			}
		}

		groups = append(groups, models.IAMGroupModel{
			GroupName:  types.StringValue(g.GroupName),
			GroupId:    types.StringValue(g.GroupId),
			Arn:        types.StringValue(g.Arn),
			Path:       types.StringValue(g.Path),
			CreateDate: types.StringValue(g.CreateDate.Format(time.RFC3339)),
			Users:      users,
		})
	}

	data.ID = types.StringValue("iam_groups_" + namespace)
	data.Groups = groups

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	tflog.Trace(ctx, "IAMGroupsDataSource.Read: completed")
}

// Helper function to safely convert *string to string
func stringValue(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

