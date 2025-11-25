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
	"strings"
	"terraform-provider-objectscale/internal/client"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ApplyPolicies(client *client.Client, ctx context.Context, plan models.IAMInlinePolicyResourceModel) (models.IAMInlinePolicyResourceModel, error) {
	// Determine namespace
	namespace := plan.Namespace.ValueString()

	// Determine entity type and name
	var entityType, entityName string
	if !plan.Username.IsNull() && !plan.Username.IsUnknown() {
		entityType = "User"
		entityName = plan.Username.ValueString()
	} else if !plan.Groupname.IsNull() && !plan.Groupname.IsUnknown() {
		entityType = "Group"
		entityName = plan.Groupname.ValueString()
	} else if !plan.Rolename.IsNull() && !plan.Rolename.IsUnknown() {
		entityType = "Role"
		entityName = plan.Rolename.ValueString()
	}

	// Step 1: Get current policies from ObjectScale
	var currentPolicies []string
	switch entityType {
	case "User":
		listResp, _, err := client.GenClient.IamApi.IamServiceListUserPolicies(ctx).
			XEmcNamespace(namespace).
			UserName(entityName).
			Execute()
		if err != nil {
			return plan, fmt.Errorf("failed to list policies: %w", err)
		}
		currentPolicies = listResp.ListUserPoliciesResult.PolicyNames

	case "Group":
		listResp, _, err := client.GenClient.IamApi.IamServiceListGroupPolicies(ctx).
			XEmcNamespace(namespace).
			GroupName(entityName).
			Execute()
		if err != nil {
			return plan, fmt.Errorf("failed to list policies: %w", err)
		}
		currentPolicies = listResp.ListGroupPoliciesResult.PolicyNames

	case "Role":
		listResp, _, err := client.GenClient.IamApi.IamServiceListRolePolicies(ctx).
			XEmcNamespace(namespace).
			RoleName(entityName).
			Execute()
		if err != nil {
			return plan, fmt.Errorf("failed to list policies: %w", err)
		}
		currentPolicies = listResp.ListRolePoliciesResult.PolicyNames
	}

	// Convert desired policies to map for quick lookup
	desiredMap := make(map[string]string)
	for _, p := range plan.Policies {
		desiredMap[p.Name.ValueString()] = p.Document.ValueString()
	}

	// Step 2: Delete policies not in desired config
	for _, existing := range currentPolicies {
		if _, found := desiredMap[existing]; !found {
			switch entityType {
			case "User":
				_, _, err := client.GenClient.IamApi.IamServiceDeleteUserPolicy(ctx).
					XEmcNamespace(namespace).
					UserName(entityName).
					PolicyName(existing).
					Execute()
				if err != nil {
					return plan, fmt.Errorf("failed to delete policy %s: %w", existing, err)
				}

			case "Group":
				_, _, err := client.GenClient.IamApi.IamServiceDeleteGroupPolicy(ctx).
					XEmcNamespace(namespace).
					GroupName(entityName).
					PolicyName(existing).
					Execute()
				if err != nil {
					return plan, fmt.Errorf("failed to delete policy %s: %w", existing, err)
				}

			case "Role":
				_, _, err := client.GenClient.IamApi.IamServiceDeleteRolePolicy(ctx).
					XEmcNamespace(namespace).
					RoleName(entityName).
					PolicyName(existing).
					Execute()
				if err != nil {
					return plan, fmt.Errorf("failed to delete policy %s: %w", existing, err)
				}
			}
		}
	}

	// Step 3: Create or Update desired policies
	for name, doc := range desiredMap {
		switch entityType {
		case "User":
			_, _, err := client.GenClient.IamApi.IamServicePutUserPolicy(ctx).
				XEmcNamespace(namespace).
				UserName(entityName).
				PolicyName(name).
				PolicyDocument(doc).
				Execute()
			if err != nil {
				return plan, fmt.Errorf("failed to apply policy %s: %w", name, err)
			}

		case "Group":
			_, _, err := client.GenClient.IamApi.IamServicePutGroupPolicy(ctx).
				XEmcNamespace(namespace).
				GroupName(entityName).
				PolicyName(name).
				PolicyDocument(doc).
				Execute()
			if err != nil {
				return plan, fmt.Errorf("failed to apply policy %s: %w", name, err)
			}

		case "Role":
			_, _, err := client.GenClient.IamApi.IamServicePutRolePolicy(ctx).
				XEmcNamespace(namespace).
				RoleName(entityName).
				PolicyName(name).
				PolicyDocument(doc).
				Execute()
			if err != nil {
				return plan, fmt.Errorf("failed to apply policy %s: %w", name, err)
			}
		}
	}

	// Set ID - format: <namespace>:<entity_type>:<entity_name>
	plan.ID = types.StringValue(fmt.Sprintf("%s:%s:%s", namespace, strings.ToLower(entityType), entityName))

	return plan, nil
}
