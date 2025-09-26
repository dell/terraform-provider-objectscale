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

package helper

import (
	"context"
	"fmt"
	"terraform-provider-objectscale/internal/models"

	objectscale "github.com/dell/objectscale-client/golang/pkg"
)

func BuildNamespaceFromPlan(ctx context.Context, plan *models.NamespaceResourceModel) (*objectscale.Namespace, error) {
	retentionClasses := &objectscale.RetentionClasses{
		RetentionClass: []objectscale.RetentionClass{},
	}
	if !plan.RetentionClasses.IsUnknown() {
		if err := assignObjectToField(ctx, plan.RetentionClasses, retentionClasses); err != nil {
			return nil, fmt.Errorf("error parsing retention classes: %v", err)
		}
	}

	userMapping := []objectscale.UserMapping{}

	if !plan.UserMapping.IsNull() && !plan.UserMapping.IsUnknown() {
		var userMappingList []models.UserMappingResource
		diags := plan.UserMapping.ElementsAs(ctx, &userMappingList, false)
		if diags.HasError() {
			return nil, fmt.Errorf("error parsing user mapping list")
		}

		for _, userMappingItem := range userMappingList {
			item := &objectscale.UserMapping{}

			if err := readFromState(ctx, userMappingItem, item); err != nil {
				return nil, fmt.Errorf("error parsing user mapping: %v", err)
			}

			var attributeList []models.AttributeResource
			diags := userMappingItem.Attributes.ElementsAs(ctx, &attributeList, false)
			if diags.HasError() {
				return nil, fmt.Errorf("error parsing attribute list")
			}

			attributes := []objectscale.Attribute{}
			for _, attributeItem := range attributeList {
				item := &objectscale.Attribute{}

				if err := readFromState(ctx, attributeItem, item); err != nil {
					return nil, fmt.Errorf("error parsing attribute: %v", err)
				}
				attributes = append(attributes, *item)
			}
			item.Attributes = attributes

			userMapping = append(userMapping, *item)
		}
	}

	namespace := &objectscale.Namespace{
		Name:                     plan.Name.ValueString(),
		DefaultDataServicesVpool: plan.DefaultDataServicesVpool.ValueString(),

		NamespaceAdmins:              plan.NamespaceAdmins.ValueString(),
		IsEncryptionEnabled:          plan.IsEncryptionEnabled.ValueBool(),
		DefaultBucketBlockSize:       plan.DefaultBucketBlockSize.ValueInt64(),
		ExternalGroupAdmins:          plan.ExternalGroupAdmins.ValueString(),
		IsStaleAllowed:               plan.IsStaleAllowed.ValueBool(),
		IsObjectLockWithAdoAllowed:   plan.IsObjectLockWithAdoAllowed.ValueBool(),
		IsComplianceEnabled:          plan.IsComplianceEnabled.ValueBool(),
		NotificationSize:             plan.NotificationSize.ValueInt64(),
		BlockSize:                    plan.BlockSize.ValueInt64(),
		NotificationSizeInCount:      plan.NotificationSizeInCount.ValueInt64(),
		BlockSizeInCount:             plan.BlockSizeInCount.ValueInt64(),
		DefaultAuditDeleteExpiration: plan.DefaultAuditDeleteExpiration.ValueInt64(),

		RetentionClasses: *retentionClasses,
		UserMapping:      userMapping,
		// For the list type, need to pass a new empty list
		AllowedVpoolsList:    []string{},
		DisallowedVpoolsList: []string{},
	}

	return namespace, nil
}
