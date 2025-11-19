package provider

import (
	"context"
	"fmt"
	"time"
	"terraform-provider-objectscale/internal/client"
	"terraform-provider-objectscale/internal/clientgen"
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

    // Attach namespace + group_name to logs
    ctx = tflog.SetField(ctx, "namespace", namespace)

    if !data.GroupName.IsNull() && !data.GroupName.IsUnknown() {
        ctx = tflog.SetField(ctx, "group_name", data.GroupName.ValueString())
    }

    tflog.Debug(ctx, "IAMGroupsDataSource.Read: config loaded")

    // Determine group filter
    var filterGroupName *string
    if !data.GroupName.IsNull() && !data.GroupName.IsUnknown() {
        val := data.GroupName.ValueString()
        filterGroupName = &val
        tflog.Debug(ctx, "IAMGroupsDataSource.Read: applying group_name filter")
    }

    // 1. CALL LIST GROUPS (ALL)
    tflog.Debug(ctx, "IAMGroupsDataSource.Read: listing IAM groups")

    var allGroups []clientgen.IamServiceListGroupsResponseListGroupsResultGroupsInner
    var marker *string

    for {
        tflog.Trace(ctx, "IAMGroupsDataSource.Read: ListGroups API call", map[string]interface{}{
            "marker": marker,
        })

        apiReq := d.client.GenClient.IamApi.IamServiceListGroups(ctx).
            XEmcNamespace(namespace)

        if marker != nil {
            apiReq = apiReq.Marker(*marker)
        }

        listResp, _, err := apiReq.Execute()
        if err != nil {
            resp.Diagnostics.AddError(
                "Error Listing IAM Groups",
                fmt.Sprintf("Could not list groups: %s", err),
            )
            tflog.Error(ctx, "IAMGroupsDataSource.Read: ListGroups API error", map[string]interface{}{
                "error": err.Error(),
            })
            return
        }

        if listResp.ListGroupsResult != nil {
            tflog.Trace(ctx, "IAMGroupsDataSource.Read: received ListGroups page", map[string]interface{}{
                "count": len(listResp.ListGroupsResult.Groups),
            })

            allGroups = append(allGroups, listResp.ListGroupsResult.Groups...)

            if listResp.ListGroupsResult.IsTruncated != nil &&
                *listResp.ListGroupsResult.IsTruncated &&
                listResp.ListGroupsResult.Marker != nil {

                marker = listResp.ListGroupsResult.Marker

                tflog.Debug(ctx, "IAMGroupsDataSource.Read: continuing pagination", map[string]interface{}{
                    "next_marker": *marker,
                })

            } else {
                tflog.Debug(ctx, "IAMGroupsDataSource.Read: pagination complete")
                break
            }
        } else {
            tflog.Warn(ctx, "IAMGroupsDataSource.Read: ListGroupsResult was nil")
            break
        }
    }

    tflog.Debug(ctx, "IAMGroupsDataSource.Read: total groups fetched", map[string]interface{}{
        "total": len(allGroups),
    })

    // 2. FILTER IF group_name GIVEN
    var filtered []clientgen.IamServiceListGroupsResponseListGroupsResultGroupsInner

    if filterGroupName != nil {
        tflog.Debug(ctx, "IAMGroupsDataSource.Read: filtering groups by name")

        for _, g := range allGroups {
            if g.GroupName == *filterGroupName {
                filtered = append(filtered, g)
            }
        }

        tflog.Debug(ctx, "IAMGroupsDataSource.Read: filtered groups result", map[string]interface{}{
            "filtered_count": len(filtered),
        })
    } else {
        filtered = allGroups
    }

    // 3. Build final groups list for Terraform
    tflog.Trace(ctx, "IAMGroupsDataSource.Read: building group state models", map[string]interface{}{
        "group_count": len(filtered),
    })

    var groups []models.IAMGroupModel

    for _, g := range filtered {
        // Add per-group context for nested logs
        groupCtx := tflog.SetField(ctx, "current_group", g.GroupName)

        var users []types.String

        // Fetch group users only when group_name is filtered
        if filterGroupName != nil {
            tflog.Trace(groupCtx, "IAMGroupsDataSource.Read: calling GetGroup API")

            getResp, _, err := d.client.GenClient.IamApi.
                IamServiceGetGroup(ctx).
                XEmcNamespace(namespace).
                GroupName(g.GroupName).
                Execute()

            if err != nil {
                tflog.Warn(groupCtx, "IAMGroupsDataSource.Read: GetGroup failed", map[string]interface{}{
                    "error": err.Error(),
                })
            } else if getResp.GetGroupResult != nil {
                for _, u := range getResp.GetGroupResult.Users {
                    users = append(users, types.StringValue(u))
                }

                tflog.Trace(groupCtx, "IAMGroupsDataSource.Read: group users loaded", map[string]interface{}{
                    "user_count": len(users),
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

    // Set ID
    data.ID = types.StringValue("iam_groups_" + namespace)
    data.Groups = groups

    tflog.Debug(ctx, "IAMGroupsDataSource.Read: final state ready", map[string]interface{}{
        "groups_in_state": len(groups),
        "state_id":        data.ID.ValueString(),
    })

    // Save state
    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

    tflog.Trace(ctx, "IAMGroupsDataSource.Read: completed")
}
