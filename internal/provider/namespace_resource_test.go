package provider

import (
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test to Fetch Namespaces.
func TestAccNSRs(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				resource"objectscale_namespace" "all" {
					name = "rounak_namespace"
					default_data_services_vpool = "urn:storageos:ReplicationGroupInfo:c6b3238b-0c46-4973-a1df-9ea83d9ba1d1:global"
					allowed_vpools_list = ["urn:storageos:ReplicationGroupInfo:9411d78e-5afb-42da-9d1f-c6c77948082b:global"]
					# root_user_password = "ChangeMe"
				}
				`,
			},
			{
				Config: ProviderConfigForTesting + `
				resource"objectscale_namespace" "all" {
					name = "rounak_namespace"
					default_data_services_vpool = "urn:storageos:ReplicationGroupInfo:c6b3238b-0c46-4973-a1df-9ea83d9ba1d1:global"
					disallowed_vpools_list = ["urn:storageos:ReplicationGroupInfo:9411d78e-5afb-42da-9d1f-c6c77948082b:global"]
					# root_user_password = "ChangeMe"
				}
				`,
			},
		},
	})
}

// Tests for vPoolDiff.
func TestVPoolDiff(t *testing.T) {
	r := &NamespaceResource{}

	tests := []struct {
		name     string
		first    []string
		second   []string
		expected []string
	}{
		{
			name:     "No difference",
			first:    []string{"a", "b", "c"},
			second:   []string{"a", "b", "c"},
			expected: []string{},
		},
		{
			name:     "All elements different",
			first:    []string{"c", "d"},
			second:   []string{"a", "b"},
			expected: []string{"c", "d"},
		},
		{
			name:     "Some elements different",
			first:    []string{"b", "c"},
			second:   []string{"a", "b"},
			expected: []string{"c"},
		}, {
			name:     "Empty second list",
			first:    []string{"x", "y"},
			second:   []string{},
			expected: []string{"x", "y"},
		},
		{
			name:     "Empty first list",
			first:    []string{},
			second:   []string{"x", "y"},
			expected: []string{},
		},
		{
			name:     "Duplicates in first list",
			first:    []string{"a", "b", "b", "c"},
			second:   []string{"a"},
			expected: []string{"b", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := r.vpoolDiff(tt.first, tt.second)
			if len(result) == 0 && len(tt.expected) == 0 {
				return // treat both as equal
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("vpoolDiff(%v, %v) = %v; expected %v", tt.first, tt.second, result, tt.expected)
			}
		})
	}
}
