/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ObjectUserResource{}
var _ resource.ResourceWithImportState = &ObjectUserResource{}

func NewObjectUserResource() resource.Resource {
	return &ObjectUserResource{}
}

// ObjectUserResource defines the resource implementation.
type ObjectUserResource struct {
	resourceProviderConfig
}

func (r *ObjectUserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_object_user"
}

func (r *ObjectUserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "ObjectScale supports creation and management of IAM Users within a namespace.",
		Description:         "ObjectScale supports creation and management of IAM Users within a namespace.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description:         "Name of the user. Required.",
				MarkdownDescription: "Name of the user. Required.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Description:         "ID of the user. Required.",
				MarkdownDescription: "ID of the user. Required.",
				Computed:            true,
			},
			"namespace": schema.StringAttribute{
				Description:         "Namespace to which the user belongs to.",
				MarkdownDescription: "Namespace to which the user belongs to.",
				Required:            true,
			},
			"created": schema.StringAttribute{
				Description:         "Secret key of the object user",
				MarkdownDescription: "Secret key of the object user",
				Computed:            true,
			},
			"locked": schema.BoolAttribute{
				Description:         "Timestamp of the secret key for the object user.",
				MarkdownDescription: "Timestamp of the secret key for the object user.",
				Computed:            true,
				Optional:            true,
			},
			"tags": schema.ListNestedAttribute{
				Description:         "Tags associated to the user. Default: []. Updatable.",
				MarkdownDescription: "Tags associated to the user. Default: []. Updatable.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description:         "Key of the tag associated to the user.",
							MarkdownDescription: "Key of the tag associated to the user.",
							Optional:            true,
							Computed:            true,
						},
						"value": schema.StringAttribute{
							Description:         "Value of the tag associated to the user.",
							MarkdownDescription: "Value of the tag associated to the user.",
							Optional:            true,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (r *ObjectUserResource) tagJson(a models.ObjectUserTags) clientgen.UserManagementServiceAddUserRequestTagsInner {
	return clientgen.UserManagementServiceAddUserRequestTagsInner{
		Name:  helper.ValueToPointer[string](a.Name),
		Value: helper.ValueToPointer[string](a.Value),
	}
}
func (r *ObjectUserResource) modelToJson(plan models.ObjectUserResourceModel) clientgen.UserManagementServiceAddUserRequest {
	return clientgen.UserManagementServiceAddUserRequest{
		Namespace: plan.Namespace.ValueString(),
		User:      plan.Name.ValueString(),
		Tags:       helper.ValueListTransform(plan.Tags, r.tagJson),
	}
}

func (r *ObjectUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Creating object user")
	var plan models.ObjectUserResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}
	planJson := r.modelToJson(plan)
	object_user_req := r.client.GenClient.UserManagementApi.UserManagementServiceAddUser(ctx)
	_, _, err := object_user_req.UserManagementServiceAddUserRequest(
		clientgen.UserManagementServiceAddUserRequest{
			Namespace: planJson.Namespace,
			User:      planJson.User,
			Tags:       planJson.Tags,
		}).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating user", err.Error())
		return
	}

	object_user, _, err := r.client.GenClient.UserManagementApi.UserManagementServiceGetUserInfo(ctx, plan.Name.ValueString()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading user after creation", err.Error())
		return
	}
	data := r.getModel(&clientgen.UserManagementServiceGetUserInfoResponse{
		Tag:       object_user.Tag,
		Locked:    object_user.Locked,
		Created:   object_user.Created,
		Namespace: object_user.Namespace,
		Name:      object_user.Name,
	}, helper.TfStringNN(&object_user.Name))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ObjectUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Info(ctx, "reading user")
	var state models.ObjectUserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	object_user, _, err := r.client.GenClient.UserManagementApi.UserManagementServiceGetUserInfo(ctx, state.Name.ValueString()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading user", err.Error())
		return
	}
	data := r.getModel(&clientgen.UserManagementServiceGetUserInfoResponse{
		Tag:       object_user.Tag,
		Locked:    object_user.Locked,
		Created:   object_user.Created,
		Namespace: object_user.Namespace,
		Name:      object_user.Name,
	}, helper.TfStringNN(&object_user.Name))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ObjectUserResource) getModel(
	object_user *clientgen.UserManagementServiceGetUserInfoResponse,
	user_id types.String,
) models.ObjectUserResourceModel {
	return models.ObjectUserResourceModel{

		Locked:    helper.TfBoolNN(&object_user.Locked),
		Created:   helper.TfStringNN(&object_user.Created),
		Namespace: helper.TfStringNN(&object_user.Namespace),
		Name:      helper.TfStringNN(&object_user.Name),
		Id:        user_id,
		Tags: helper.ListNotNull(object_user.Tag,
			func(v clientgen.UserManagementServiceAddUserRequestTagsInner) types.Object {
				return helper.Object(models.ObjectUserTags{
					Name:  helper.TfStringNN(v.Name),
					Value: helper.TfStringNN(v.Value),
				})
			}),
	}
}

// computes the difference between two Iam Tag sets (lists).
func TagsDiff(first, second []clientgen.UserManagementServiceAddUserRequestTagsInner) []clientgen.UserManagementServiceAddUserRequestTagsInner {
	var diff []clientgen.UserManagementServiceAddUserRequestTagsInner
	smap := make(map[string]struct{}, len(second))
	for _, v := range second {
		smap[*v.Name] = struct{}{}
	}
	for _, v := range first {
		if _, ok := smap[*v.Name]; !ok {
			diff = append(diff, v)
		}
	}
	return diff
}

func TagsChanged(
	plan, state []clientgen.UserManagementServiceAddUserRequestTagsInner,
) []clientgen.UserManagementServiceAddUserRequestTagsInner {

	// Build a map of name -> value from state for quick lookup
	sMap := make(map[string]*string, len(state))
	for _, t := range state {
		if t.Name == nil {
			// skip entries without a name
			continue
		}
		sMap[*t.Name] = t.Value
	}

	var changed []clientgen.UserManagementServiceAddUserRequestTagsInner
	for _, p := range plan {
		if p.Name == nil {
			// skip entries without a name
			continue
		}
		sVal, ok := sMap[*p.Name]
		if !ok {
			// name not present in state -> not a "present on both" case
			continue
		}
		if !equalPtrString(p.Value, sVal) {
			// present on both, but value differs
			changed = append(changed, p)
		}
	}

	return changed
}

func equalPtrString(a, b *string) bool {
	switch {
	case a == nil && b == nil:
		return true
	case a == nil || b == nil:
		return false
	default:
		return *a == *b
	}
}

func (r *ObjectUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Info(ctx, "updating user")

	var plan, state models.ObjectUserResourceModel

	// Read Terraform plan and state data
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !plan.Tags.IsNull() || !plan.Tags.IsUnknown() || len(plan.Tags.Elements()) != 0 {
		tags_plan, tags_state := helper.ValueListTransform(plan.Tags, r.tagJson),
			helper.ValueListTransform(state.Tags, r.tagJson)

		tags_to_add, tags_to_remove := TagsDiff(tags_plan, tags_state),
			TagsDiff(tags_state, tags_plan)
		tags_to_change := TagsChanged(tags_plan, tags_state)

		if len(tags_to_remove) != 0 {
			_, _, err_untag := r.client.GenClient.UserManagementApi.UserManagementServiceRemoveUserTags(ctx, state.Name.ValueString()).
				UserManagementServiceRemoveUserTagsRequest(clientgen.UserManagementServiceRemoveUserTagsRequest{
					Tags: tags_to_remove,
				}).
				Execute()
			if err_untag != nil {
				resp.Diagnostics.AddError("Error removing user tags", err_untag.Error())
				return
			}
		}

		if len(tags_to_change) != 0 {
			_, _, err_untag := r.client.GenClient.UserManagementApi.UserManagementServiceUpdateUserTag(ctx, state.Name.ValueString()).
				UserManagementServiceUpdateUserTagRequest(clientgen.UserManagementServiceUpdateUserTagRequest{
					Tags: tags_to_change,
				}).
				Execute()
			if err_untag != nil {
				resp.Diagnostics.AddError("Error updating user tags", err_untag.Error())
				return
			}
		}
		if len(tags_to_add) != 0 {
			_, _, err_tag := r.client.GenClient.UserManagementApi.UserManagementServiceAddUserTag(ctx, state.Name.ValueString()).
				UserManagementServiceAddUserTagRequest(clientgen.UserManagementServiceAddUserTagRequest{
					Tags: tags_to_add,
				}).
				Execute()
			if err_tag != nil {
				resp.Diagnostics.AddError("Error adding user tags", err_tag.Error())
				return
			}
		}
	}
	if helper.IsChangedNN(plan.Locked, state.Locked) {
		_, _, err := r.client.GenClient.UserManagementApi.UserManagementServiceSetUserLock(ctx).
			UserManagementServiceSetUserLockRequest(clientgen.UserManagementServiceSetUserLockRequest{
				Namespace: state.Namespace.ValueStringPointer(),
				User:      state.Name.ValueString(),
				IsLocked:  plan.Locked.ValueBool(),
			}).
			Execute()

		if err != nil {
			resp.Diagnostics.AddError("Error updating locked status", err.Error())
			return
		}
	}
	object_user, _, err := r.client.GenClient.UserManagementApi.UserManagementServiceGetUserInfo(ctx, state.Name.ValueString()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading user after modification", err.Error())
		return
	}
	data := r.getModel(&clientgen.UserManagementServiceGetUserInfoResponse{
		Tag:       object_user.Tag,
		Locked:    object_user.Locked,
		Created:   object_user.Created,
		Namespace: object_user.Namespace,
		Name:      object_user.Name,
	}, helper.TfStringNN(&object_user.Name))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ObjectUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "Deleting object user")
	var state models.ObjectUserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.GenClient.UserManagementApi.UserManagementServiceRemoveUser(ctx).
		UserManagementServiceRemoveUserRequest(clientgen.UserManagementServiceRemoveUserRequest{
			Namespace: state.Namespace.ValueStringPointer(),
			User:      state.Name.ValueString(),
		}).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Object user",
			err.Error(),
		)
	}
}

func (r *ObjectUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Info(ctx, "Importing object user")

	object_user, _, err := r.client.GenClient.UserManagementApi.UserManagementServiceGetUserInfo(ctx, req.ID).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading user after import", err.Error())
		return
	}
	data := r.getModel(&clientgen.UserManagementServiceGetUserInfoResponse{
		Tag:       object_user.Tag,
		Locked:    object_user.Locked,
		Created:   object_user.Created,
		Namespace: object_user.Namespace,
		Name:      object_user.Name,
	}, helper.TfStringNN(&object_user.Name))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}
