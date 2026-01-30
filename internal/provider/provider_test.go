/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"terraform-provider-objectscale/internal/client"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// var testProvider provider.Provider

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"objectscale": providerserver.NewProtocol6WithError(New("test")()),
}

// var FunctionMocker *mockey.Mocker

var ProviderConfigForTesting = ``
var username, password, endpoint, insecure, rgs string
var logoutUser, logoutPassword string

func init() {
	_, err := loadEnvFile("objectscale.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	endpoint = setDefault(os.Getenv("OBJECTSCALE_ENDPOINT"), "http://localhost:3007")
	username = setDefault(os.Getenv("OBJECTSCALE_USERNAME"), "test")
	password = setDefault(os.Getenv("OBJECTSCALE_PASSWORD"), "test")
	logoutUser = setDefault(os.Getenv("OBJECTSCALE_LOGOUT_USERNAME"), "logouttest")
	logoutPassword = setDefault(os.Getenv("OBJECTSCALE_LOGOUT_PASSWORD"), "logouttest")
	insecure = setDefault(os.Getenv("OBJECTSCALE_INSECURE"), "true")
	rgs = `
		data "objectscale_replication_group" "preq_replication_group" {
		}

		locals {
			rgs = {
				"rg1": data.objectscale_replication_group.preq_replication_group.replication_groups.0.id,
				"rg2": data.objectscale_replication_group.preq_replication_group.replication_groups.1.id,
				"rg3": data.objectscale_replication_group.preq_replication_group.replication_groups.2.id,
			}
		}
		`

	ProviderConfigForTesting = fmt.Sprintf(`
		provider "objectscale" {
			username = "%s"
			password = "%s"
			endpoint = "%s"
			insecure = "%s"
			timeout = 120
		}
	`, username, password, endpoint, insecure)
}

func testAccPreCheck(t *testing.T) {
	if v := username; v == "" {
		t.Fatal("OBJECTSCALE_USERNAME must be set for acceptance tests")
	}

	if v := password; v == "" {
		t.Fatal("OBJECTSCALE_PASSWORD must be set for acceptance tests")
	}

	if v := endpoint; v == "" {
		t.Fatal("OBJECTSCALE_ENDPOINT must be set for acceptance tests")
	}

	// // Before each test clear out the mocker
	// if FunctionMocker != nil {
	// 	FunctionMocker.UnPatch()
	// }
}

func testUserTokenCleanup(t *testing.T) {
	client, err := client.NewClient(
		endpoint,
		logoutUser,
		logoutPassword,
		insecure == "true",
		120,
	)

	if err != nil {
		t.Fatal("Could not login with logout user", err.Error())
	}

	// logout the username
	_, _, err = client.GenClient.AuthenticationApi.AuthenticationResourceLogout(context.TODO()).Username(username).Execute()
	if err != nil {
		t.Fatal("Could not logout test user", err.Error())
	}

	// logout the logout user
	_, _, err = client.GenClient.AuthenticationApi.AuthenticationResourceLogout(context.TODO()).Execute()
	if err != nil {
		t.Fatal("Could not logout the logout user", err.Error())
	}
}

func setDefault(osInput string, defaultStr string) string {
	if osInput == "" {
		return defaultStr
	}
	return osInput
}

// loadEnvFile used to read env file and set params.
func loadEnvFile(path string) (map[string]string, error) {
	envMap := make(map[string]string)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		envMap[key] = value
		// Set the environment variable for system access
		if err := os.Setenv(key, value); err != nil {
			return nil, fmt.Errorf("error setting environment variable %s: %w", key, err)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return envMap, nil
}
