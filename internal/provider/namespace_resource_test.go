package provider

import (
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var locals = `
locals {
  rgs = {
    "rg1": "urn:storageos:ReplicationGroupInfo:55ca12b2-e908-4bac-a5fe-3fdaa975e3eb:global",
    "rg2": "urn:storageos:ReplicationGroupInfo:cd8bffcb-7a99-4023-82a8-982054fd73c2:global",
    "rg3": "urn:storageos:ReplicationGroupInfo:e0b539a3-6ddd-4412-b4d0-ce08049f64cd:global",
  }
}
`

// Test to Fetch Namespaces
func TestAccNSRs(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// create
				Config: ProviderConfigForTesting + locals + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg1"]
					allowed_vpools_list         = [local.rgs["rg2"]]
					disallowed_vpools_list      = [local.rgs["rg3"]]
					is_compliance_enabled = true
					# is_encryption_enabled = true
					default_audit_delete_expiration = 3600
				}
				`,
			},
			{
				// import
				Config: ProviderConfigForTesting + locals + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg1"]
					allowed_vpools_list         = [local.rgs["rg2"]]
					disallowed_vpools_list      = [local.rgs["rg3"]]
					is_compliance_enabled = true
					# is_encryption_enabled = true
					default_audit_delete_expiration = 3600
				}
				`,
				ResourceName: "objectscale_namespace.all",
				ImportState:  true,
			},
			{
				// import error
				Config: ProviderConfigForTesting + locals + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg1"]
					allowed_vpools_list         = [local.rgs["rg2"]]
					disallowed_vpools_list      = [local.rgs["rg3"]]
					is_compliance_enabled = true
					# is_encryption_enabled = true
					default_audit_delete_expiration = 3600
				}
				`,
				ResourceName:  "objectscale_namespace.all",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(".*Error reading namespace.*"),
				ImportStateId: "invalid-id",
			},

			// TODO: something wrong with default_audit_delete_expiration update
			// {
			// 	// update default audit delete expiration
			// 	Config: ProviderConfigForTesting + locals + `
			// 	resource"objectscale_namespace" "all" {
			// 		name                        = "testacc_namespace"
			// 		default_data_services_vpool = local.rgs["rg1"]
			// 		default_audit_delete_expiration = 1200
			// 	}
			// 	`,
			// },
			{
				// update vpools
				Config: ProviderConfigForTesting + locals + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg3"]
					allowed_vpools_list         = [local.rgs["rg3"]]
					disallowed_vpools_list      = [local.rgs["rg2"]]
					# TODO: something wrong with default_audit_delete_expiration update
					# default_audit_delete_expiration = 0
				}
				`,
			},
		},
	})
}

// Test to Fetch Namespaces
func TestAccNSRsAll(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// create
				Config: ProviderConfigForTesting + locals + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg1"]
					allowed_vpools_list         = [local.rgs["rg2"]]
					disallowed_vpools_list      = [local.rgs["rg3"]]
					namespace_admins             = "admin1,admin2"
					user_mapping = [{
						domain = "domain"
						groups = ["group1", "group2"]
						attributes = [
							{
								key = "key1"
								value = ["value1", "value2"]
							},
							{
								key = "key2"
								value = ["value3", "value4"]
							}
						]
					}]
					external_group_admins = "admin1@foo,admin2@bar"
				}
				`,
			},
			{
				// update vpools
				Config: ProviderConfigForTesting + locals + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg3"]
					allowed_vpools_list         = [local.rgs["rg3"]]
					disallowed_vpools_list      = [local.rgs["rg2"]]
					namespace_admins             = "admin2,admin3"
					user_mapping = [{
						domain = "domain2"
						groups = ["group3", "group4"]
						attributes = [
							{
								key = "key3"
								value = ["value5", "value6"]
							},
							{
								key = "key4"
								value = ["value7", "value8"]
							}
						]
					}]
					default_bucket_block_size = 1024
					external_group_admins = "admin3@foo,admin4@bar"
					is_stale_allowed = true
					is_object_lock_with_ado_allowed = true
					# default_audit_delete_expiration = 3600
				}
				`,
			},
			{
				// remove from lists
				Config: ProviderConfigForTesting + locals + `
				resource"objectscale_namespace" "all" {
					name = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg3"]
					allowed_vpools_list = []
					disallowed_vpools_list = []
					namespace_admins = ""
					user_mapping = [{
						domain = "domain2"
						groups = []
						attributes = []
					}]
					default_bucket_block_size = -1
					external_group_admins = ""
					is_stale_allowed = false
					is_object_lock_with_ado_allowed = false
					# default_audit_delete_expiration = 0
				}
				`,
			},
			{
				// remove all from lists 2
				Config: ProviderConfigForTesting + locals + `
				resource"objectscale_namespace" "all" {
					name = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg3"]
					user_mapping = []
				}
				`,
			},
		},
	})
}

// Tests for vPoolDiff
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
