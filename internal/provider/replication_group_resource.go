/*
Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"context"
	"fmt"
	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ReplicationGroupResource{}
var _ resource.ResourceWithImportState = &ReplicationGroupResource{}

func NewReplicationGroupResource() resource.Resource {
	return &ReplicationGroupResource{}
}

// ReplicationGroupResource defines the resource implementation.
type ReplicationGroupResource struct {
	resourceProviderConfig
}

func (r *ReplicationGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_replication_group"
}

func (r *ReplicationGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This resource allows end user to Provision and manage Dell ObjectScale Replication Groups.",
		MarkdownDescription: "This resource allows end user to Provision and manage Dell ObjectScale Replication Groups.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Identifier of the Replication Group.",
				MarkdownDescription: "Identifier of the Replication Group.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Description:         "Name of the Replication Group.",
				MarkdownDescription: "Name of the Replication Group.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"zone_mappings": schema.SetNestedAttribute{
				Description:         "List of zones (VDC + Storage Pool) which will be used for replication.",
				MarkdownDescription: "List of zones (VDC + Storage Pool) which will be used for replication.",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vdc": schema.StringAttribute{
							Description:         "Virtual Data Center ID.",
							MarkdownDescription: "Virtual Data Center ID.",
							Required:            true,
						},
						"storage_pool": schema.StringAttribute{
							Description:         "Storage Pool ID.",
							MarkdownDescription: "Storage Pool ID.",
							Required:            true,
						},
						"is_replication_target": schema.BoolAttribute{
							Description:         "In passive replication groups, one zone acts as the target. This attribute must be set to \"true\" for the zone which will act as the replication target.",
							MarkdownDescription: "In passive replication groups, one zone acts as the target. This attribute must be set to `true` for the zone which will act as the replication target.",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
					},
				},
				Validators: []validator.Set{
					setvalidator.SizeAtLeast(1),
				},
			},
			"type": schema.StringAttribute{
				Description:         "Type of the Replication Group (Active/Passive). Cannot be updated.",
				MarkdownDescription: "Type of the Replication Group (Active/Passive). Cannot be updated.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				Description:         "Description of the Replication Group.",
				MarkdownDescription: "Description of the Replication Group.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"enable_rebalancing": schema.BoolAttribute{
				Description:         "Enable Rebalancing.",
				MarkdownDescription: "Enable Rebalancing.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"allow_all_namespaces": schema.BoolAttribute{
				Description:         "Whether to allow all namespaces.",
				MarkdownDescription: "Whether to allow all namespaces.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"replicate_to_all_sites": schema.BoolAttribute{
				Description:         "Whether to replicate to all sites (for Active configuration). Cannot be updated.",
				MarkdownDescription: "Whether to replicate to all sites (for Active configuration). Cannot be updated.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
}

// helper function to marshal GET response to tfsdk model.
func (r *ReplicationGroupResource) respToModel(rg *clientgen.DataServiceVpoolServiceGetDataServiceStoreResponse) models.ReplicationGroupResourceModel {
	return models.ReplicationGroupResourceModel{
		ID:                 helper.TfStringNN(rg.Id),
		Name:               helper.TfStringNN(rg.Name),
		Description:        helper.TfStringNN(rg.Description),
		EnableRebalancing:  helper.TfBoolNN(rg.EnableRebalancing),
		AllowAllNamespaces: helper.TfBoolNN(rg.IsAllowAllNamespaces),
		FullRep:            helper.TfBoolNN(rg.IsFullRep),
		ZoneMappings: helper.SetNotNull(rg.VarrayMappings,
			func(vmr clientgen.DataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInnerVarrayMappingsInner) types.Object {
				return helper.Object(models.ReplicationGroupResourceZoneMapping{
					Name:                helper.TfString(vmr.Name),
					Value:               helper.TfString(vmr.Value),
					IsReplicationTarget: helper.TfBool(vmr.IsReplicationTarget),
				})
			}),
		Type: types.StringValue(map[bool]string{
			true:  "Passive",
			false: "Active",
		}[rg.UseReplicationTarget != nil && *rg.UseReplicationTarget]),
	}
}

// helper func that converts known false values to nil.
func rgRsIsTrue[T interface{ types.Bool | bool }](in T) *bool {
	var ret *bool
	switch inTyped := any(in).(type) {
	case types.Bool:
		ret = helper.ValueToPointer[bool](inTyped)
	case bool:
		ret = &inTyped
	}
	// if known true, return it
	if ret != nil && *ret {
		return ret
	}
	// everything else is nil
	return nil
}

// helper function to unmarshal zone mappings from tfsdk to json.
func (r *ReplicationGroupResource) getZoneList(in models.ReplicationGroupResourceModel) []clientgen.DataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInnerVarrayMappingsInner {
	return helper.ValueListTransform(in.ZoneMappings,
		func(in models.ReplicationGroupResourceZoneMapping) clientgen.DataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInnerVarrayMappingsInner {
			return clientgen.DataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInnerVarrayMappingsInner{
				Name:                helper.ValueToPointer[string](in.Name),
				Value:               helper.ValueToPointer[string](in.Value),
				IsReplicationTarget: rgRsIsTrue(in.IsReplicationTarget),
			}
		})
}

// helper function to check if asymmetric replication.
func (r *ReplicationGroupResource) isAsymmetricReplication(in []clientgen.DataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInnerVarrayMappingsInner) bool {
	// if any zone is configured as a replication target, we are in asymmetric replication
	for _, zm := range in {
		if zm.IsReplicationTarget != nil && *zm.IsReplicationTarget {
			return true
		}
	}
	return false
}

// helper function to check if all zones are known.
func (r *ReplicationGroupResource) checkAllZonesKnown(inM models.ReplicationGroupResourceModel) bool {
	in := inM.ZoneMappings
	if in.IsUnknown() {
		return false
	}
	// unmarshal zone mappings to tfsdk model
	zoneMappingsTf := helper.ValueListTransform(in,
		func(in models.ReplicationGroupResourceZoneMapping) models.ReplicationGroupResourceZoneMapping {
			return in
		})
	// check if all zones are known
	for _, zm := range zoneMappingsTf {
		if zm.IsReplicationTarget.IsUnknown() ||
			zm.Name.IsUnknown() ||
			zm.Value.IsUnknown() {
			return false
		}
	}
	return true
}

// Config Validation.
func (r *ReplicationGroupResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {

	var conf models.ReplicationGroupResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &conf)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// zone mapping validations
	if r.checkAllZonesKnown(conf) {
		// unmarshal zone mappings
		zoneMappings := r.getZoneList(conf)

		// Validation 1: check for duplicate zones, identified by vdc + storage pool
		zmap := make(map[string]struct{}, len(zoneMappings))
		for _, zm := range zoneMappings {
			// id should be unique in set
			id := *zm.Name + ":" + *zm.Value
			if _, ok := zmap[id]; ok {
				resp.Diagnostics.AddError(
					"Duplicate zone mapping",
					fmt.Sprintf("A zone mapping with vdc %s and storage_pool %s appears more than once", *zm.Name, *zm.Value),
				)
				return
			} else {
				// add to map so we can check for duplicates in next set of items
				zmap[id] = struct{}{}
			}
		}

		// Validation 2: A Passive Replication Group must have exactly three zones, 2 source and 1 target
		asymmetricReplicationPlan := r.isAsymmetricReplication(zoneMappings)
		if asymmetricReplicationPlan {
			if len(zoneMappings) != 3 {
				resp.Diagnostics.AddError(
					"Invalid number of zones for Passive Replication Group",
					"A Passive Replication Group must have exactly three zones, 2 source and 1 target",
				)
			} else {
				targetZones := 0
				for _, zm := range zoneMappings {
					if zm.IsReplicationTarget != nil && *zm.IsReplicationTarget {
						targetZones++
					}
				}
				if targetZones != 1 {
					resp.Diagnostics.AddError(
						"Invalid number of source and target zones for Passive Replication Group",
						"A Passive Replication Group must have exactly two source zones and exactly one target zone.",
					)
				}
			}
		}
	}
}

// Plan Modify.
func (r *ReplicationGroupResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// If the entire plan is null, the resource is planned for destruction.
	if req.Plan.Raw.IsNull() {
		resp.Diagnostics.AddWarning("Deletion of Replication Group is not supported.",
			"If this plan is applied, this resource will be removed from the state, but will not be destroyed on ObjectScale.")
		return
	}

	var plan models.ReplicationGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If the entire state is null, the resource is planned for creation.
	// create plan validations
	if req.State.Raw.IsNull() {
		if r.checkAllZonesKnown(plan) {
			// Action 1: Compute Active/Passive
			asymmetricReplicationPlan := r.isAsymmetricReplication(r.getZoneList(plan))
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("type"), types.StringValue(map[bool]string{
				true:  "Passive",
				false: "Active",
			}[asymmetricReplicationPlan]))...)

			// Validation 1: IsFullRep is applicable only for symmetric replication
			if asymmetricReplicationPlan && helper.IsKnown(plan.FullRep) && plan.FullRep.ValueBool() {
				resp.Diagnostics.AddError("replicate_to_all_sites can be set to true only for Active replication", "")
			}
		}
		return
	}

	// update plan validations
	var state models.ReplicationGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validation 1: replicate_to_all_sites cannot be modified
	if helper.IsChangedNN(plan.FullRep, state.FullRep) {
		resp.Diagnostics.AddError("Error updating Replication Group", "replicate_to_all_sites cannot be modified")
	}

	// if plan zone mappings is known, run zone mapping validations
	if r.checkAllZonesKnown(plan) {
		// Validation 1: A replication group cannot be converted from Active to Passive configuration
		// No use checking for passive to active, since its possible to remove the passive zone (making its Terraform configuration indistinguishable from active).
		asymmetricReplicationPlan := r.isAsymmetricReplication(r.getZoneList(plan))
		if state.Type.ValueString() == "Active" && asymmetricReplicationPlan {
			resp.Diagnostics.AddError("Cannot update replication group type from Active to Passive",
				"Please contact Dell Technologies customer support if want to update this replication group's type.")
		}

		// Validation 2: Cannot add and remove zones at the same time
		addzones, removezones := r.zoneDiff(plan, state)
		if len(addzones) > 0 && len(removezones) > 0 {
			resp.Diagnostics.AddError("Cannot add and remove zones in one operation",
				"Please create a plan to only add or remove zones at a time."+
					" Please also note that, after one zone update operation, the replication group may not accept any more zone updates for some time.")
		}

		// Validation 3: If removing zones, throw warning
		if len(removezones) > 0 {
			resp.Diagnostics.AddWarning("This plan will remove zones from the replication group.",
				"Are you sure you want to remove VDC from this Replication Group?\n"+
					"Removing zones from replication group may result in data loss.\n"+
					"We recommend contacting customer support before performing this operation.\n"+
					"Data loss may occur if prerequisite procedures are not properly followed.\n"+
					"Verify the following conditions:\n"+
					"- Ensure that Geo replication is up-to-date.\n"+
					"- Replication to/from VDC for the Replication Group will be disabled.\n"+
					"- Recovery will be initiated. Data may not be available until recovery is complete.\n"+
					"- Removal is permanent; the site cannot be added back to this replication group.\n"+
					"- Data associated with this replication group will be permanently deleted from this VDC.\n"+
					"- In cases where XOR encoding is utilized and the RG falls below 3 VDCs, the XOR encoded data"+
					" will have to be replaced with fully replicated copies, which could significantly increase storage required to fully protect the data.")
		}
	}
}

// Create.
func (r *ReplicationGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.ReplicationGroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// determine if we are dealing with asymmetric replication (ie. Passive Config)
	zoneMappings := r.getZoneList(plan)
	// run the API call
	createdRG, _, err := r.client.GenClient.DataVpoolApi.DataServiceVpoolServiceCreateDataServiceVpool(ctx).
		DataServiceVpoolServiceCreateDataServiceVpoolRequest(clientgen.DataServiceVpoolServiceCreateDataServiceVpoolRequest{
			Name:                 plan.Name.ValueString(),
			ZoneMappings:         zoneMappings,
			Description:          helper.ValueToPointer[string](plan.Description),
			EnableRebalancing:    helper.ValueToPointer[bool](plan.EnableRebalancing),
			IsAllowAllNamespaces: helper.ValueToPointer[bool](plan.AllowAllNamespaces),
			IsFullRep:            rgRsIsTrue(plan.FullRep),
			UseReplicationTarget: rgRsIsTrue(plan.Type.ValueString() == "Passive"),
		}).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating Replication Group", err.Error())
		return
	}

	// Save data into Terraform state
	state := r.respToModel(&clientgen.DataServiceVpoolServiceGetDataServiceStoreResponse{
		Id:                   createdRG.Id,
		Name:                 createdRG.Name,
		VarrayMappings:       createdRG.VarrayMappings,
		Description:          createdRG.Description,
		EnableRebalancing:    createdRG.EnableRebalancing,
		IsAllowAllNamespaces: createdRG.IsAllowAllNamespaces,
		IsFullRep:            createdRG.IsFullRep,
		UseReplicationTarget: createdRG.UseReplicationTarget,
	})
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Read.
func (r *ReplicationGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.ReplicationGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rg, _, err := r.client.GenClient.DataVpoolApi.
		DataServiceVpoolServiceGetDataServiceStore(ctx, state.ID.ValueString()).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Error reading Replication Group state", err.Error())
		return
	}

	if rg.Id == nil {
		// if wrong ID is passed to API, this happens
		resp.Diagnostics.AddError("Error reading Replication Group state", "Replication Group not found with ID "+state.ID.ValueString())
	}

	state2 := r.respToModel(rg)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state2)...)
}

// helper function to create diff lists of zone mappings, used in update.
func (r *ReplicationGroupResource) zoneDiff(plan, state models.ReplicationGroupResourceModel) (add, remove []clientgen.DataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInnerVarrayMappingsInner) {
	plist, slist := r.getZoneList(plan), r.getZoneList(state)
	// each zone has a unique Name + Value.
	// create maps by these unique signature and create diff lists
	pmap, smap := make(map[string]*clientgen.DataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInnerVarrayMappingsInner, len(plist)),
		make(map[string]*clientgen.DataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInnerVarrayMappingsInner, len(slist))
	for _, p := range plist {
		pmap[*p.Name+*p.Value] = &p
	}
	for _, s := range slist {
		smap[*s.Name+*s.Value] = &s
	}
	for id, p := range pmap {
		// is in plan
		if _, ok := smap[id]; !ok {
			// not in state; so add
			add = append(add, *p)
		}
	}
	for id, s := range smap {
		// is in state
		if _, ok := pmap[id]; !ok {
			// not in plan; so remove
			remove = append(remove, *s)
		}
	}
	return add, remove
}

// Update.
func (r *ReplicationGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan models.ReplicationGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// basic attributes update
	basicreq := clientgen.DataServiceVpoolServicePutDataServiceVpoolRequest{
		Description:        helper.ValueToPointer[string](plan.Description),
		EnableRebalancing:  helper.ValueToPointer[bool](plan.EnableRebalancing),
		AllowAllNamespaces: helper.ValueToPointer[bool](plan.AllowAllNamespaces),
	}
	if helper.IsChangedNN(plan.Name, state.Name) {
		basicreq.Name = helper.ValueToPointer[string](plan.Name)
	}
	_, _, err := r.client.GenClient.DataVpoolApi.
		DataServiceVpoolServicePutDataServiceVpool(ctx, state.ID.ValueString()).
		DataServiceVpoolServicePutDataServiceVpoolRequest(basicreq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error updating Replication Group attributes", err.Error())
		return
	}

	// zone mappings update
	add, remove := r.zoneDiff(plan, state)
	// zone mappings addition
	if len(add) > 0 {
		_, _, err := r.client.GenClient.DataVpoolApi.
			DataServiceVpoolServiceAddToVpool(ctx, state.ID.ValueString()).
			DataServiceVpoolServiceAddToVpoolRequest(clientgen.DataServiceVpoolServiceAddToVpoolRequest{
				Mappings: add,
			}).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error adding new zones to Replication Group", err.Error())
			return
		}
	}

	// zone mappings removal
	if len(remove) > 0 {
		_, _, err := r.client.GenClient.DataVpoolApi.
			DataServiceVpoolServiceRemoveFromVpool(ctx, state.ID.ValueString()).
			DataServiceVpoolServiceRemoveFromVpoolRequest(clientgen.DataServiceVpoolServiceRemoveFromVpoolRequest{
				Mappings: remove,
			}).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error removing zones from Replication Group", err.Error())
			return
		}
	}

	// wait 30 seconds for replication group to be updated
	time.Sleep(30 * time.Second)
	// Read updated data
	rg, _, err := r.client.GenClient.DataVpoolApi.
		DataServiceVpoolServiceGetDataServiceStore(ctx, state.ID.ValueString()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading Replication Group state after update", err.Error())
		return
	}
	state2 := r.respToModel(rg)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state2)...)
}

// Delete.
func (r *ReplicationGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// just remove from state
	resp.State.RemoveResource(ctx)
}

// ImportState.
func (r *ReplicationGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	allRGResp, _, err := r.client.GenClient.DataVpoolApi.DataServiceVpoolServiceGetDataServiceVpools(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting the list of replication groups",
			err.Error(),
		)
		return
	}
	// Find the resource with the matching name
	for _, rg := range allRGResp.DataServiceVpool {
		if *rg.Name == req.ID {
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), rg.Id)...)
			return
		}
	}
	// return error if not found
	resp.Diagnostics.AddError(fmt.Sprintf("Could not find replication group with name %s", req.ID), "")
}
