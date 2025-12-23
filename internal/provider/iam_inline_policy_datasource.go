package provider

import (
	"context"
	"fmt"

	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &IAMInlinePolicyDataSource{}

func NewIAMInlinePolicyDataSource() datasource.DataSource {
	return &IAMInlinePolicyDataSource{}
}

type IAMInlinePolicyDataSource struct {
	datasourceProviderConfig
}

func (d *IAMInlinePolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_inline_policy"
}

// Schema describes the data source arguments and results.
func (d *IAMInlinePolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Description:         "Retrieve IAM inline policies for a user, group, or role.",
		MarkdownDescription: "Retrieve IAM inline policies for a user, group, or role.",

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

			"username": schema.StringAttribute{
				Description:         "Name of the IAM User.",
				MarkdownDescription: "Name of the IAM User.",
				Optional:            true,
			},

			"groupname": schema.StringAttribute{
				Description:         "Name of the IAM Group.",
				MarkdownDescription: "Name of the IAM Group.",
				Optional:            true,
			},

			"rolename": schema.StringAttribute{
				Description:         "Name of the IAM Role.",
				MarkdownDescription: "Name of the IAM Role.",
				Optional:            true,
			},

			"policies": schema.ListNestedAttribute{
				Description:         "Inline policies attached to the user, group, or role.",
				MarkdownDescription: "Inline policies attached to the user, group, or role.",
				Computed:            true,

				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"name": schema.StringAttribute{
							Description:         "Name of the inline policy.",
							MarkdownDescription: "Name of the inline policy.",
							Computed:            true,
						},

						"document": schema.StringAttribute{
							Description:         "Policy document in JSON format.",
							MarkdownDescription: "Policy document in JSON format.",
							Required:            true,
							CustomType:          jsontypes.NormalizedType{},
						},
					},
				},
			},
		},
	}
}

func (d *IAMInlinePolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Trace(ctx, "IAMInlinePolicyDataSource.Read: start")

	var data models.IAMInlinePolicyResourceModel

	// Load config
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ns := data.Namespace.ValueString()
	var finalPolicies []models.IAMInlinePolicyModel

	// Validate exactly one of username / groupname / rolename
	count := 0
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		count++
	}
	if !data.Groupname.IsNull() && !data.Groupname.IsUnknown() {
		count++
	}
	if !data.Rolename.IsNull() && !data.Rolename.IsUnknown() {
		count++
	}

	if count != 1 {
		resp.Diagnostics.AddError(
			"Invalid configuration",
			"Exactly one of username, groupname, or rolename must be specified.",
		)
		return
	}

	var err error

	// CASE 1 — user inline policies
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		finalPolicies, err = d.listUserInlinePolicies(
			ctx,
			ns,
			data.Username.ValueString(),
		)

		// CASE 2 — group inline policies
	} else if !data.Groupname.IsNull() && !data.Groupname.IsUnknown() {
		finalPolicies, err = d.listGroupInlinePolicies(
			ctx,
			ns,
			data.Groupname.ValueString(),
		)

		// CASE 3 — role inline policies
	} else {
		finalPolicies, err = d.listRoleInlinePolicies(
			ctx,
			ns,
			data.Rolename.ValueString(),
		)
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Namespace does not exist or error reading IAM inline policies",
			err.Error(),
		)
		return
	}

	// Save state
	data.ID = types.StringValue("iam_inline_policy_" + ns)
	data.Policies = finalPolicies

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "IAMInlinePolicyDataSource.Read: completed")
}

func (d *IAMInlinePolicyDataSource) listUserInlinePolicies(ctx context.Context, namespace, username string) ([]models.IAMInlinePolicyModel, error) {

	req := d.client.GenClient.IamApi.IamServiceListUserPolicies(ctx).
		UserName(username).
		XEmcNamespace(namespace)

	items, err := helper.GetAllInstances(req)
	if err != nil {
		return nil, err
	}

	policies := make([]models.IAMInlinePolicyModel, 0, len(items))

	for _, name := range items {

		getResp, _, err := d.client.GenClient.IamApi.IamServiceGetUserPolicy(ctx).
			UserName(username).
			PolicyName(name).
			XEmcNamespace(namespace).
			Execute()
		if err != nil {
			return nil, fmt.Errorf(
				"error getting inline policy info %s for user %s: %w",
				name, username, err,
			)
		}

		var doc string
		if getResp.GetUserPolicyResult.PolicyDocument != nil {
			doc = *getResp.GetUserPolicyResult.PolicyDocument
		}

		policies = append(policies, models.IAMInlinePolicyModel{
			Name:     helper.TfString(&name),
			Document: jsontypes.NewNormalizedValue(doc),
		})
	}

	return policies, nil
}

func (d *IAMInlinePolicyDataSource) listGroupInlinePolicies(ctx context.Context, namespace, groupname string) ([]models.IAMInlinePolicyModel, error) {

	req := d.client.GenClient.IamApi.IamServiceListGroupPolicies(ctx).
		GroupName(groupname).
		XEmcNamespace(namespace)

	items, err := helper.GetAllInstances(req)
	if err != nil {
		return nil, err
	}

	policies := make([]models.IAMInlinePolicyModel, 0, len(items))

	for _, name := range items {

		getResp, _, err := d.client.GenClient.IamApi.IamServiceGetGroupPolicy(ctx).
			GroupName(groupname).
			PolicyName(name).
			XEmcNamespace(namespace).
			Execute()
		if err != nil {
			return nil, fmt.Errorf(
				"error getting inline policy info %s for group %s: %w",
				name, groupname, err,
			)
		}

		var doc string
		if getResp.GetGroupPolicyResult.PolicyDocument != nil {
			doc = *getResp.GetGroupPolicyResult.PolicyDocument
		}

		policies = append(policies, models.IAMInlinePolicyModel{
			Name:     helper.TfString(&name),
			Document: jsontypes.NewNormalizedValue(doc),
		})
	}

	return policies, nil
}

func (d *IAMInlinePolicyDataSource) listRoleInlinePolicies(ctx context.Context, namespace, rolename string) ([]models.IAMInlinePolicyModel, error) {

	req := d.client.GenClient.IamApi.IamServiceListRolePolicies(ctx).
		RoleName(rolename).
		XEmcNamespace(namespace)

	items, err := helper.GetAllInstances(req)
	if err != nil {
		return nil, fmt.Errorf(
			"error listing role policy names for role %s: %w",
			rolename, err,
		)
	}

	policies := make([]models.IAMInlinePolicyModel, 0, len(items))

	for _, name := range items {

		getResp, _, err := d.client.GenClient.IamApi.IamServiceGetRolePolicy(ctx).
			RoleName(rolename).
			PolicyName(name).
			XEmcNamespace(namespace).
			Execute()
		if err != nil {
			return nil, fmt.Errorf(
				"error getting inline policy info %s for role %s: %w",
				name, rolename, err,
			)
		}

		var doc string
		if getResp.GetRolePolicyResult.PolicyDocument != nil {
			doc = *getResp.GetRolePolicyResult.PolicyDocument
		}

		policies = append(policies, models.IAMInlinePolicyModel{
			Name:     helper.TfString(&name),
			Document: jsontypes.NewNormalizedValue(doc),
		})
	}

	return policies, nil
}
