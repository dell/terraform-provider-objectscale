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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// I-19 — Read SP metadata datasource after SP create populates parsed fields.
func TestAcc_I19_DataSource_SPMetadata(t *testing.T) {
	defer testUserTokenCleanup(t)
	cfg := spHCL("objectscale.example.com") + `
data "objectscale_iam_service_provider_metadata" "md" {
  depends_on = [objectscale_iam_service_provider.sp]
}
`
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider_metadata.md", "metadata_xml"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider_metadata.md", "entity_id"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider_metadata.md", "acs_url"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider_metadata.md", "authn_requests_signed"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider_metadata.md", "want_assertions_signed"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider_metadata.md", "signing_certificate"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider_metadata.md", "name_id_formats.#"),
				),
			},
		},
	})
}
