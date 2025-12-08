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

func ApplyPolicyARNs(client *client.Client, ctx context.Context, plan models.IAMManagedPolicyResourceModel, currentState *models.IAMManagedPolicyResourceModel) (models.IAMManagedPolicyResourceModel, error) {
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

	// Step 1: Get the current policy ARNs
	var currentPolicyARNs []string
	if currentState != nil {
		// Use state for Update functionality
		diags := currentState.PolicyARNs.ElementsAs(ctx, &currentPolicyARNs, false)
		if diags.HasError() {
			return plan, fmt.Errorf("failed to read policy_arns from state")
		}
	} else {
		currentPolicyARNs = []string{}

		// Call ListAttached<entity>Policies API
		var policyARNs []string
		var marker string

		switch entityType {
		case "User":
			for {
				listReq := client.GenClient.IamApi.IamServiceListAttachedUserPolicies(ctx).
					XEmcNamespace(namespace).
					UserName(entityName)

				if marker != "" {
					listReq = listReq.Marker(marker)
				}

				listResp, _, err := listReq.Execute()
				if err != nil {
					return plan, fmt.Errorf("failed to list policy arns: %w", err)
				}

				for _, p := range listResp.ListAttachedUserPoliciesResult.AttachedPolicies {
					policyARNs = append(policyARNs, *p.PolicyArn)
				}

				markerPtr := listResp.ListAttachedUserPoliciesResult.Marker
				if markerPtr == nil || *markerPtr == "" {
					break
				}
				marker = *markerPtr
			}
		case "Group":
			for {
				listReq := client.GenClient.IamApi.IamServiceListAttachedGroupPolicies(ctx).
					XEmcNamespace(namespace).
					GroupName(entityName)

				if marker != "" {
					listReq = listReq.Marker(marker)
				}

				listResp, _, err := listReq.Execute()
				if err != nil {
					return plan, fmt.Errorf("failed to list policy arns: %w", err)
				}

				for _, p := range listResp.ListAttachedGroupPoliciesResult.AttachedPolicies {
					policyARNs = append(policyARNs, *p.PolicyArn)
				}

				markerPtr := listResp.ListAttachedGroupPoliciesResult.Marker
				if markerPtr == nil || *markerPtr == "" {
					break
				}
				marker = *markerPtr
			}
		case "Role":
			for {
				listReq := client.GenClient.IamApi.IamServiceListAttachedRolePolicies(ctx).
					XEmcNamespace(namespace).
					RoleName(entityName)

				if marker != "" {
					listReq = listReq.Marker(marker)
				}

				listResp, _, err := listReq.Execute()
				if err != nil {
					return plan, fmt.Errorf("failed to list policy arns: %w", err)
				}

				for _, p := range listResp.ListAttachedRolePoliciesResult.AttachedPolicies {
					policyARNs = append(policyARNs, *p.PolicyArn)
				}

				markerPtr := listResp.ListAttachedRolePoliciesResult.Marker
				if markerPtr == nil || *markerPtr == "" {
					break
				}
				marker = *markerPtr
			}
		}

		currentPolicyARNs = append(currentPolicyARNs, policyARNs...)
	}

	// Step 2: Get the desired policy arns from the plan
	var desiredPolicyARNs []string
	diags := plan.PolicyARNs.ElementsAs(ctx, &desiredPolicyARNs, false)
	if diags.HasError() {
		return plan, fmt.Errorf("failed to read policy_arns from plan")
	}

	// Step 3: Detach policy arns in current but not in desired
	detachPolicyARNs := computeDiff(currentPolicyARNs, desiredPolicyARNs)
	for _, arn := range detachPolicyARNs {
		// call detach API
		switch entityType {
		case "User":
			_, _, err := client.GenClient.IamApi.IamServiceDetachUserPolicy(ctx).
				XEmcNamespace(namespace).
				UserName(entityName).
				PolicyArn(arn).
				Execute()
			if err != nil {
				return plan, fmt.Errorf("failed to detach policy arn %s: %w", arn, err)
			}

		case "Group":
			_, _, err := client.GenClient.IamApi.IamServiceDetachGroupPolicy(ctx).
				XEmcNamespace(namespace).
				GroupName(entityName).
				PolicyArn(arn).
				Execute()
			if err != nil {
				return plan, fmt.Errorf("failed to detach policy arn %s: %w", arn, err)
			}

		case "Role":
			_, _, err := client.GenClient.IamApi.IamServiceDetachRolePolicy(ctx).
				XEmcNamespace(namespace).
				RoleName(entityName).
				PolicyArn(arn).
				Execute()
			if err != nil {
				return plan, fmt.Errorf("failed to detach policy arn %s: %w", arn, err)
			}
		}
	}

	// Step 4: Attach policy arns in desired but not in current
	attachPolicyARNs := computeDiff(desiredPolicyARNs, currentPolicyARNs)
	for _, arn := range attachPolicyARNs {
		// call attach API
		switch entityType {
		case "User":
			_, _, err := client.GenClient.IamApi.IamServiceAttachUserPolicy(ctx).
				XEmcNamespace(namespace).
				UserName(entityName).
				PolicyArn(arn).
				Execute()
			if err != nil {
				return plan, fmt.Errorf("failed to attach policy arn %s: %w", arn, err)
			}

		case "Group":
			_, _, err := client.GenClient.IamApi.IamServiceAttachGroupPolicy(ctx).
				XEmcNamespace(namespace).
				GroupName(entityName).
				PolicyArn(arn).
				Execute()
			if err != nil {
				return plan, fmt.Errorf("failed to attach policy arn %s: %w", arn, err)
			}

		case "Role":
			_, _, err := client.GenClient.IamApi.IamServiceAttachRolePolicy(ctx).
				XEmcNamespace(namespace).
				RoleName(entityName).
				PolicyArn(arn).
				Execute()
			if err != nil {
				return plan, fmt.Errorf("failed to attach policy arn %s: %w", arn, err)
			}
		}
	}

	// Set ID - format: <namespace>:<entity_type>:<entity_name>
	plan.ID = types.StringValue(fmt.Sprintf("%s:%s:%s", namespace, strings.ToLower(entityType), entityName))

	return plan, nil
}

// computeDiff returns items present in 'a' but not in 'b'.
// Order of 'a' is preserved.
func computeDiff(a, b []string) []string {
	bSet := make(map[string]struct{}, len(b))
	for _, v := range b {
		bSet[v] = struct{}{}
	}

	out := make([]string, 0, len(a))
	for _, v := range a {
		if _, found := bSet[v]; !found {
			out = append(out, v)
		}
	}
	return out
}
