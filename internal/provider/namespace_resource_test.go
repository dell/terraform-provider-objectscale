package provider

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"terraform-provider-objectscale/internal/client"
	"terraform-provider-objectscale/internal/clientgen"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
)

// Test to Fetch Namespaces.
func TestAccNSRs(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	defer testUserTokenCleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// create
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg1"]
					allowed_vpools_list         = [local.rgs["rg2"]]
					disallowed_vpools_list      = [local.rgs["rg3"]]
					is_compliance_enabled = true
					default_audit_delete_expiration = 3600
				}
				`,
			},
			{
				// import
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg1"]
					allowed_vpools_list         = [local.rgs["rg2"]]
					disallowed_vpools_list      = [local.rgs["rg3"]]
					is_compliance_enabled = true
					default_audit_delete_expiration = 3600
				}
				`,
				ResourceName: "objectscale_namespace.all",
				ImportState:  true,
			},
			{
				// import error
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg1"]
					allowed_vpools_list         = [local.rgs["rg2"]]
					disallowed_vpools_list      = [local.rgs["rg3"]]
					is_compliance_enabled = true
					default_audit_delete_expiration = 3600
				}
				`,
				ResourceName:  "objectscale_namespace.all",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(".*Error reading namespace.*"),
				ImportStateId: "invalid-id",
			},
			{
				// update name error
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace2"
					default_data_services_vpool = local.rgs["rg3"]
				}
				`,
				ExpectError: regexp.MustCompile(".*not[[:space:]]updatable.*"),
			},
			{
				// update vpools
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg3"]
					allowed_vpools_list         = [local.rgs["rg3"]]
					disallowed_vpools_list      = [local.rgs["rg2"]]
				}
				`,
			},
		},
	})
}

// Test to Fetch Namespaces.
func TestAccNSRsAll(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	defer testUserTokenCleanup(t)
	var upM, rcUpdateM *mockey.Mocker

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// create
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace2"
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
					retention_classes = [
						{
							name = "testacc1"
							period = 1000
						},
						{
							name = "testacc2"
							period = 2000
						}
					]
					quota = {
						notification_size = 100
						block_size = 124
					}
				}
				`,
			},
			{
				// update vpools
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace2"
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
					retention_classes = [
						{
							name = "testacc1"
							period = 500
						},
						{
							name = "testacc3"
							period = 3000
						},
						{
							name = "testacc2"
							period = 2000
						}
					]
					quota = {
						notification_size = 90
						block_size = 124
					}
					root_user_password = "password1"
				}
				`,
			},
			{
				// remove from lists
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name = "testacc_namespace2"
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
					quota = {
						notification_size = 100
						block_size = 224
					}
					root_user_password = "password2"
					current_root_user_password = "password1"
				}
				`,
			},
			// remove retention classes error
			{
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name = "testacc_namespace2"
					default_data_services_vpool = local.rgs["rg3"]
					retention_classes = []
				}
				`,
				ExpectError: regexp.MustCompile(".*removal of retention classes.*"),
			},
			{
				// remove all from lists 2
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name = "testacc_namespace2"
					default_data_services_vpool = local.rgs["rg3"]
					user_mapping = []
					quota = {
						notification_size = -1
						block_size = -1
					}
				}
				`,
			},
			{
				// update ns error mock
				PreConfig: func() {
					upM = mockey.Mock((*clientgen.NamespaceApiService).NamespaceServiceUpdateNamespaceExecute).
						Return(nil, nil, fmt.Errorf("error")).Build()
				},
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace2"
					default_data_services_vpool = "dummy"
				}
				`,
				ExpectError: regexp.MustCompile("Error updating namespace"),
			},
			{
				// update retention class error mock
				PreConfig: func() {
					upM.UnPatch()
					upM = mockey.Mock((*clientgen.NamespaceApiService).NamespaceServiceUpdateNamespaceExecute).
						Return(nil, nil, nil).Build()
					rcUpdateM = mockey.Mock((*clientgen.NamespaceApiService).NamespaceServiceCreateRetentionClassExecute).
						Return(nil, nil, fmt.Errorf("{}")).Build()
				},
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace2"
					default_data_services_vpool = local.rgs["rg3"]
					retention_classes = [
						{
							name = "testacc1"
							period = 500
						},
						{
							name = "testacc2"
							period = 2000
						},
						{
							name = "testacc3"
							period = 3000
						},
						{
							name = "testacc4"
							period = -4000
						}
					]
				}
				`,
				ExpectError: regexp.MustCompile("Error adding retention classes"),
			},
		},
	})
	upM.UnPatch()
	rcUpdateM.UnPatch()
}

func TestAccNsRsUpdateCommon(t *testing.T) {
	r := &NamespaceResource{
		resourceProviderConfig: resourceProviderConfig{
			client: &client.Client{
				GenClient: &clientgen.APIClient{
					NamespaceApi: &clientgen.NamespaceApiService{},
				},
			},
		},
	}
	tests := []struct {
		name  string
		state clientgen.NamespaceServiceGetNamespaceResponse
		plan  clientgen.NamespaceServiceGetNamespaceResponse
		mock  func() *mockey.Mocker
	}{
		{
			name: "manageRetentionClasses update error",
			state: clientgen.NamespaceServiceGetNamespaceResponse{
				RetentionClasses: &clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerRetentionClasses{
					RetentionClass: []clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerRetentionClassesRetentionClassInner{
						{
							Name:   getpointer("c1"),
							Period: getpointer[int64](500),
						},
					},
				},
			},
			plan: clientgen.NamespaceServiceGetNamespaceResponse{
				RetentionClasses: &clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerRetentionClasses{
					RetentionClass: []clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerRetentionClassesRetentionClassInner{
						{
							Name:   getpointer("c1"),
							Period: getpointer[int64](300),
						},
					},
				},
			},
			mock: func() *mockey.Mocker {
				return mockey.Mock((*clientgen.NamespaceApiService).NamespaceServiceUpdateRetentionClassExecute).
					Return(nil, nil, fmt.Errorf("{}")).Build()
			},
		},
		{
			name: "manageRetentionClasses create error",
			state: clientgen.NamespaceServiceGetNamespaceResponse{
				RetentionClasses: &clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerRetentionClasses{},
			},
			plan: clientgen.NamespaceServiceGetNamespaceResponse{
				RetentionClasses: &clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerRetentionClasses{
					RetentionClass: []clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerRetentionClassesRetentionClassInner{
						{
							Name:   getpointer("c1"),
							Period: getpointer[int64](500),
						},
					},
				},
			},
			mock: func() *mockey.Mocker {
				return mockey.Mock((*clientgen.NamespaceApiService).NamespaceServiceCreateRetentionClassExecute).
					Return(nil, nil, fmt.Errorf("{}")).Build()
			},
		},
		{
			name: "manage quotas error",
			state: clientgen.NamespaceServiceGetNamespaceResponse{
				NotificationSize: getpointer[int64](-1),
				BlockSize:        getpointer[int64](-1),
				RetentionClasses: &clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerRetentionClasses{},
			},
			plan: clientgen.NamespaceServiceGetNamespaceResponse{
				NotificationSize: getpointer[int64](500),
				BlockSize:        getpointer[int64](500),
				RetentionClasses: &clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerRetentionClasses{},
			},
			mock: func() *mockey.Mocker {
				return mockey.Mock((*clientgen.NamespaceApiService).NamespaceServiceUpdateNamespaceQuotaExecute).
					Return(nil, nil, fmt.Errorf("{}")).Build()
			},
		},
		{
			name: "read error",
			state: clientgen.NamespaceServiceGetNamespaceResponse{
				NotificationSize: getpointer[int64](-1),
				BlockSize:        getpointer[int64](-1),
				RetentionClasses: &clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerRetentionClasses{},
			},
			plan: clientgen.NamespaceServiceGetNamespaceResponse{
				NotificationSize: getpointer[int64](-1),
				BlockSize:        getpointer[int64](-1),
				RetentionClasses: &clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerRetentionClasses{},
			},
			mock: func() *mockey.Mocker {
				return mockey.Mock((*clientgen.NamespaceApiService).NamespaceServiceGetNamespaceExecute).
					Return(nil, nil, fmt.Errorf("{}")).Build()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttmock := tt.mock()
			dummys, dummyp := tt.state, tt.plan
			dummys.Id = getpointer("todo")
			_, err := r.updateCommon(context.TODO(), &dummyp, &dummys)
			ttmock.UnPatch()
			assert.NotNil(t, err)
		})
	}
}

func loginMocker() *mockey.Mocker {
	loginH := make(http.Header)
	loginH.Set("X-SDS-Auth-Token", "todo")
	return mockey.Mock((*clientgen.AuthenticationApiService).AuthenticationResourceGetLoginTokenExecute).
		Return(nil, &http.Response{
			Header: loginH,
		}, nil).Build()
}

func TestAccNsRsCreateError(t *testing.T) {
	defer testUserTokenCleanup(t)
	loginM := loginMocker()
	createM := mockey.Mock((*clientgen.NamespaceApiService).NamespaceServiceCreateNamespaceExecute).
		Return(nil, nil, fmt.Errorf("error")).Build()
	deleteM := mockey.Mock((*clientgen.NamespaceApiService).NamespaceServiceDeactivateNamespaceExecute).
		Return(nil, nil, nil).Build()
	rcM := mockey.Mock((*clientgen.NamespaceApiService).NamespaceServiceCreateRetentionClassExecute).
		Return(nil, nil, fmt.Errorf("{}")).Build()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// create error
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace2"
					default_data_services_vpool = local.rgs["rg1"]
				}
				`,
				ExpectError: regexp.MustCompile(`Error creating namespace`),
			},
			{
				// create error after update
				PreConfig: func() {
					createM.UnPatch()
					createM = mockey.Mock((*clientgen.NamespaceApiService).NamespaceServiceCreateNamespaceExecute).
						Return(&clientgen.NamespaceServiceCreateNamespaceResponse{
							Id: getpointer("todo"),
						}, nil, nil).Build()
				},
				Config: ProviderConfigForTesting + namespace_preq_rgs + `
				resource"objectscale_namespace" "all" {
					name                        = "testacc_namespace2"
					default_data_services_vpool = local.rgs["rg1"]
					retention_classes = [
						{
							name = "testacc1"
							period = 1000
						}
					]
				}
				`,
				ExpectError: regexp.MustCompile(`Error adding retention classes`),
			},
		},
	})
	loginM.UnPatch()
	createM.UnPatch()
	deleteM.UnPatch()
	rcM.UnPatch()
}

func getpointer[T any](in T) *T {
	return &in
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

func TestNamespaceResource_boolToString(t *testing.T) {
	tests := []struct {
		name string
		b    *bool
		want *string
	}{
		{
			name: "bool is true",
			b:    func() *bool { b := true; return &b }(),
			want: func() *string { s := "true"; return &s }(),
		},
		{
			name: "bool is false",
			b:    func() *bool { b := false; return &b }(),
			want: func() *string { s := "false"; return &s }(),
		},
		{
			name: "bool is nil",
			b:    nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (&NamespaceResource{}).boolToString(tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boolToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNamespaceResource_stringToBool(t *testing.T) {
	tests := []struct {
		name string
		s    *string
		want *bool
	}{
		{
			name: "string is true",
			s:    func() *string { s := "true"; return &s }(),
			want: func() *bool { b := true; return &b }(),
		},
		{
			name: "string is false",
			s:    func() *string { s := "false"; return &s }(),
			want: func() *bool { b := false; return &b }(),
		},
		{
			name: "string is nil",
			s:    nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (&NamespaceResource{}).stringToBool(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
