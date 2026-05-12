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
	"fmt"
	"regexp"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccVDCCertificateDataSource_Read(t *testing.T) {
	loginM := loginMocker()
	defer loginM.UnPatch()
	mocker := mockey.Mock(GetVDCKeystore).Return("-----BEGIN CERTIFICATE-----\ntest\n-----END CERTIFICATE-----", nil).Build()
	defer mocker.UnPatch()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `data "objectscale_vdc_certificate" "test" {}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.objectscale_vdc_certificate.test", "id", "vdc_certificate_datasource"),
					resource.TestCheckResourceAttrSet("data.objectscale_vdc_certificate.test", "certificate_chain"),
				),
			},
		},
	})
}

func TestAccVDCCertificateDataSource_ReadError(t *testing.T) {
	loginM := loginMocker()
	defer loginM.UnPatch()
	mocker := mockey.Mock(GetVDCKeystore).Return("", fmt.Errorf("permission denied")).Build()
	defer mocker.UnPatch()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + `data "objectscale_vdc_certificate" "test" {}`,
				ExpectError: regexp.MustCompile(`Error reading VDC certificate`),
			},
		},
	})
}
