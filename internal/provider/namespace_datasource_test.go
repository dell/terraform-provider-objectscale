package provider

import (
	"os"
	"regexp"
	"terraform-provider-objectscale/internal/clientgen"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test to Fetch Namespaces.
func TestAccNSDs(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_namespace" "all" {
				}
				`,
			},
			{
				// fetch invalid
				Config: ProviderConfigForTesting + rgs + `
				data "objectscale_namespace" "all" {
					name = "invalid-id"
				}
				`,
				ExpectError: regexp.MustCompile(`Error getting namespaces`),
			},
			{
				// fetch one
				Config: ProviderConfigForTesting + rgs + `
				resource "objectscale_namespace" "preq" {
					name                        = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg1"]
				}
				data "objectscale_namespace" "one" {
					name = objectscale_namespace.preq.id
				}
				`,
			},
			{
				// fetch all by prefix
				Config: ProviderConfigForTesting + rgs + `
				resource "objectscale_namespace" "preq" {
					name                        = "testacc_namespace"
					default_data_services_vpool = local.rgs["rg1"]
				}
				data "objectscale_namespace" "prefix" {
					name = "testacc*"
				}
				`,
			},
		},
	})
}

func TestNamespaceDataSource_updateNamespaceState(t *testing.T) {
	type args struct {
		namespaces []clientgen.NamespaceServiceGetNamespacesResponseNamespaceInner
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test updateNamespaceState with empty namespaces",
			args: args{
				namespaces: []clientgen.NamespaceServiceGetNamespacesResponseNamespaceInner{},
			},
		},
		{
			name: "test updateNamespaceState with one namespace",
			args: args{
				namespaces: []clientgen.NamespaceServiceGetNamespacesResponseNamespaceInner{
					{
						Name: getpointer("namespace1"),
						Id:   getpointer("id1"),
						UserMapping: []clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerUserMappingInner{
							{
								Domain: "todo",
								Groups: []string{"todo"},
								Attributes: []clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerUserMappingInnerAttributesInner{
									{
										Key:   "todokey",
										Value: []string{"todovalue"},
									},
								},
							},
						},
						RetentionClasses: &clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerRetentionClasses{
							RetentionClass: []clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerRetentionClassesRetentionClassInner{
								{
									Name:   getpointer("class1"),
									Period: getpointer[int64](500),
								},
							},
						},
						Link: &clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerLink{
							Rel:  getpointer(""),
							Href: getpointer(""),
						},
					},
				},
			},
		},
		{
			name: "test updateNamespaceState with multiple namespaces",
			args: args{
				namespaces: []clientgen.NamespaceServiceGetNamespacesResponseNamespaceInner{
					{
						Name: getpointer("namespace1"),
						Id:   getpointer("id1"),
						Link: &clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerLink{
							Rel:  getpointer(""),
							Href: getpointer(""),
						},
					},
					{
						Name: getpointer("namespace2"),
						Id:   getpointer("id2"),
						Link: &clientgen.NamespaceServiceGetNamespacesResponseNamespaceInnerLink{
							Rel:  getpointer(""),
							Href: getpointer(""),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NamespaceDataSource{}.updateNamespaceState(tt.args.namespaces)
		})
	}
}
