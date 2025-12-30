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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccManagementUserResourceForLocalUserCRUD(t *testing.T) {
	defer testUserTokenCleanup(t)

	resourceName := "objectscale_management_user.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// create Local User
			{
				Config: ProviderConfigForTesting + testAccManagementUserResourceLocalUserConfig1(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "localuser1"),
					resource.TestCheckResourceAttr(resourceName, "type", "LOCAL_USER"),
					resource.TestCheckResourceAttr(resourceName, "name", "localuser1"),
					resource.TestCheckResourceAttr(resourceName, "password", "pass123"),
					resource.TestCheckResourceAttr(resourceName, "system_administrator", "false"),
					resource.TestCheckResourceAttr(resourceName, "system_monitor", "false"),
					resource.TestCheckResourceAttr(resourceName, "security_administrator", "false"),
				),
			},
			// update error on non-updatable fields
			{
				Config:      ProviderConfigForTesting + testAccManagementUserResourceADLDAPUserConfig1(),
				ExpectError: regexp.MustCompile("Error updating management user"),
			},
			// update Local User
			{
				Config: ProviderConfigForTesting + testAccManagementUserResourceLocalUserConfig2(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "localuser1"),
					resource.TestCheckResourceAttr(resourceName, "type", "LOCAL_USER"),
					resource.TestCheckResourceAttr(resourceName, "name", "localuser1"),
					resource.TestCheckResourceAttr(resourceName, "password", "pass1234"),
					resource.TestCheckResourceAttr(resourceName, "system_administrator", "true"),
					resource.TestCheckResourceAttr(resourceName, "system_monitor", "true"),
					resource.TestCheckResourceAttr(resourceName, "security_administrator", "true"),
				),
			},
			// delete Local User
		},
	})
}

func TestAccManagementUserResourceForADLDAPUserCRUD(t *testing.T) {
	defer testUserTokenCleanup(t)

	resourceName := "objectscale_management_user.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// create AD/LDAP User
			{
				Config: ProviderConfigForTesting + testAccManagementUserResourceADLDAPUserConfig1(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "user1@domain"),
					resource.TestCheckResourceAttr(resourceName, "type", "AD_LDAP_USER"),
					resource.TestCheckResourceAttr(resourceName, "name", "user1@domain"),
					resource.TestCheckResourceAttr(resourceName, "system_administrator", "false"),
					resource.TestCheckResourceAttr(resourceName, "system_monitor", "false"),
					resource.TestCheckResourceAttr(resourceName, "security_administrator", "false"),
				),
			},
			// update error on non-updatable fields
			{
				Config:      ProviderConfigForTesting + testAccManagementUserResourceLocalUserConfig1(),
				ExpectError: regexp.MustCompile("Error updating management user"),
			},
			// update AD/LDAP User
			{
				Config: ProviderConfigForTesting + testAccManagementUserResourceADLDAPUserConfig2(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "user1@domain"),
					resource.TestCheckResourceAttr(resourceName, "type", "AD_LDAP_USER"),
					resource.TestCheckResourceAttr(resourceName, "name", "user1@domain"),
					resource.TestCheckResourceAttr(resourceName, "system_administrator", "true"),
					resource.TestCheckResourceAttr(resourceName, "system_monitor", "true"),
					resource.TestCheckResourceAttr(resourceName, "security_administrator", "true"),
				),
			},
			// delete AD/LDAP User
		},
	})
}

func TestAccManagementUserResourceForADLDAPGroupCRUD(t *testing.T) {
	defer testUserTokenCleanup(t)

	resourceName := "objectscale_management_user.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// create AD/LDAP Group
			{
				Config: ProviderConfigForTesting + testAccManagementUserResourceADLDAPGroupConfig1(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "group1@domain"),
					resource.TestCheckResourceAttr(resourceName, "type", "AD_LDAP_GROUP"),
					resource.TestCheckResourceAttr(resourceName, "name", "group1@domain"),
					resource.TestCheckResourceAttr(resourceName, "system_administrator", "false"),
					resource.TestCheckResourceAttr(resourceName, "system_monitor", "false"),
					resource.TestCheckResourceAttr(resourceName, "security_administrator", "false"),
				),
			},
			// update error on non-updatable fields
			{
				Config:      ProviderConfigForTesting + testAccManagementUserResourceLocalUserConfig1(),
				ExpectError: regexp.MustCompile("Error updating management user"),
			},
			// update AD/LDAP Group
			{
				Config: ProviderConfigForTesting + testAccManagementUserResourceADLDAPGroupConfig2(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "group1@domain"),
					resource.TestCheckResourceAttr(resourceName, "type", "AD_LDAP_GROUP"),
					resource.TestCheckResourceAttr(resourceName, "name", "group1@domain"),
					resource.TestCheckResourceAttr(resourceName, "system_administrator", "true"),
					resource.TestCheckResourceAttr(resourceName, "system_monitor", "true"),
					resource.TestCheckResourceAttr(resourceName, "security_administrator", "true"),
				),
			},
			// delete AD/LDAP Group
		},
	})
}

func TestAccManagementUserResourceForImport(t *testing.T) {
	defer testUserTokenCleanup(t)

	resourceName := "objectscale_management_user.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + testAccManagementUserResourceADLDAPUserConfig1(),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateId:     "invalid@name",
				ImportStateVerify: true,
				ExpectError:       regexp.MustCompile(`Import Management User failed`),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateId:     "user1@domain",
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccManagementUserResourceForErrorScenarios(t *testing.T) {
	defer testUserTokenCleanup(t)
	
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// absent type
			{
				Config:      ProviderConfigForTesting + testAccManagementUserResourceErrorConfig1(),
				ExpectError: regexp.MustCompile("Missing required argument"),
			},
			// absent name
			{
				Config:      ProviderConfigForTesting + testAccManagementUserResourceErrorConfig2(),
				ExpectError: regexp.MustCompile("Missing required argument"),
			},
			// invalid type
			{
				Config:      ProviderConfigForTesting + testAccManagementUserResourceErrorConfig3(),
				ExpectError: regexp.MustCompile("Invalid Attribute Value Match"),
			},
			// invalid name format for Local User
			{
				Config:      ProviderConfigForTesting + testAccManagementUserResourceErrorConfig4(),
				ExpectError: regexp.MustCompile("Invalid Name Format for LOCAL_USER"),
			},
			// invalid name format for AD/LDAP User
			{
				Config:      ProviderConfigForTesting + testAccManagementUserResourceErrorConfig5(),
				ExpectError: regexp.MustCompile("Invalid Name Format for AD_LDAP_USER/AD_LDAP_GROUP"),
			},
			// invalid name format for AD/LDAP Group
			{
				Config:      ProviderConfigForTesting + testAccManagementUserResourceErrorConfig6(),
				ExpectError: regexp.MustCompile("Invalid Name Format for AD_LDAP_USER/AD_LDAP_GROUP"),
			},
			// absent password when creating Local User
			{
				Config:      ProviderConfigForTesting + testAccManagementUserResourceErrorConfig7(),
				ExpectError: regexp.MustCompile("Password is required for LOCAL_USER"),
			},
			// present password when creating AD/LDAP User
			{
				Config:      ProviderConfigForTesting + testAccManagementUserResourceErrorConfig8(),
				ExpectError: regexp.MustCompile("Password is not applicable for AD_LDAP_USER/AD_LDAP_GROUP"),
			},
			// present password when creating AD/LDAP Group
			{
				Config:      ProviderConfigForTesting + testAccManagementUserResourceErrorConfig9(),
				ExpectError: regexp.MustCompile("Password is not applicable for AD_LDAP_USER/AD_LDAP_GROUP"),
			},
		},
	})
}

func testAccManagementUserResourceLocalUserConfig1() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "LOCAL_USER"
        name = "localuser1"
        password = "pass123"
    }
    `
}

func testAccManagementUserResourceLocalUserConfig2() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "LOCAL_USER"
        name = "localuser1"
        password = "pass1234"
        system_administrator = true
        system_monitor = true
        security_administrator = true
    }
    `
}

func testAccManagementUserResourceADLDAPUserConfig1() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "AD_LDAP_USER"
        name = "user1@domain"
    }
    `
}

func testAccManagementUserResourceADLDAPUserConfig2() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "AD_LDAP_USER"
        name = "user1@domain"
        system_administrator = true
        system_monitor = true
        security_administrator = true
    }
    `
}

func testAccManagementUserResourceADLDAPGroupConfig1() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "AD_LDAP_GROUP"
        name = "group1@domain"
    }
    `
}

func testAccManagementUserResourceADLDAPGroupConfig2() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "AD_LDAP_GROUP"
        name = "group1@domain"
        system_administrator = true
        system_monitor = true
        security_administrator = true
    }
    `
}

func testAccManagementUserResourceErrorConfig1() string {
	return `
    resource "objectscale_management_user" "example" {
        name = "localuser1"
    }
    `
}

func testAccManagementUserResourceErrorConfig2() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "LOCAL_USER"
    }
    `
}

func testAccManagementUserResourceErrorConfig3() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "INVALID_TYPE"
        name = "localuser1"
    }
    `
}

func testAccManagementUserResourceErrorConfig4() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "LOCAL_USER"
        name = "invalid@name"
    }
    `
}

func testAccManagementUserResourceErrorConfig5() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "AD_LDAP_USER"
        name = "invalid_name"
    }
    `
}

func testAccManagementUserResourceErrorConfig6() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "AD_LDAP_GROUP"
        name = "invalid_name"
    }
    `
}

func testAccManagementUserResourceErrorConfig7() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "LOCAL_USER"
        name = "localuser1"
    }
    `
}

func testAccManagementUserResourceErrorConfig8() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "AD_LDAP_USER"
        name = "user1@domain"
		password = "pass123"
    }
    `
}

func testAccManagementUserResourceErrorConfig9() string {
	return `
    resource "objectscale_management_user" "example" {
        type = "AD_LDAP_GROUP"
        name = "group1@domain"
		password = "pass123"
    }
    `
}
