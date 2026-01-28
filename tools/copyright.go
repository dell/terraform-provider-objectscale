//go:build tools

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

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	factTypeResource   = "resource"
	factTypeDatasource = "data"
)

type Fact struct {
	Note string
}

var facts = map[string]map[string]map[string]Fact{
	"Identity & Access Management (IAM)": {
		"iam_group":             {factTypeResource: {}, factTypeDatasource: {}},
		"iam_policy":            {factTypeResource: {}, factTypeDatasource: {}},
		"iam_inline_policy":     {factTypeResource: {}, factTypeDatasource: {}},
		"iam_policy_attachment": {factTypeResource: {}}, // no datasource
		"iam_role":              {factTypeResource: {}, factTypeDatasource: {}},
		"iam_user":              {factTypeResource: {}, factTypeDatasource: {}},
		"iam_user_access_key":   {factTypeResource: {}},
		"iam_management_user":   {factTypeResource: {}, factTypeDatasource: {}},
		"iam_group_membership":  {factTypeResource: {}},
	},
	"Object User": {
		"object_user":            {factTypeResource: {}, factTypeDatasource: {}},
		"object_user_secret_key": {factTypeResource: {}}, // no datasource
	},
	"Management User": {
		"management_user": {factTypeResource: {}, factTypeDatasource: {}},
	},
	"Namespacing / Tenancy": {
		"namespace": {factTypeResource: {}, factTypeDatasource: {}},
	},
	"Object Storage Containers": {
		"bucket": {factTypeResource: {
			Note: "> **Warning:** Deleting a bucket using this resource will also delete all data contained within the bucket. Ensure you have backed up any important data before performing a destroy operation.",
		}, factTypeDatasource: {}},
	},
	"Data Protection": {
		"replication_group": {
			factTypeResource: {
				Note: "~> **Note:** Deletion of Replication Group is not supported." +
					" If this resource gets planned for deletion, it will simply be removed from the state." +
					" But the Replication Group will not be destroyed on the ObjectScale array." +
					"\n\n!> **Caution:** This resource does support removal of zones from Replication Group." +
					" But be cautious that removing zones from replication group may result in data loss.\n" +
					"We recommend contacting customer support before performing this operation.\n" +
					"Data loss may occur if prerequisite procedures are not properly followed.\n" +
					"Verify the following conditions:<br/>" +
					"- Ensure that Geo replication is up-to-date.<br/>" +
					"- Replication to/from VDC for the Replication Group will be disabled.<br/>" +
					"- Recovery will be initiated. Data may not be available until recovery is complete.<br/>" +
					"- Removal is permanent; the site cannot be added back to this replication group.<br/>" +
					"- Data associated with this replication group will be permanently deleted from this VDC.<br/>" +
					"- In cases where XOR encoding is utilized and the RG falls below 3 VDCs, the XOR encoded data" +
					" will have to be replaced with fully replicated copies, which could significantly increase storage required to fully protect the data.",
			},
			factTypeDatasource: {},
		},
	},
	"Storage Topology & Capacity Domains": {
		"storage_pool": {factTypeDatasource: {}}, // resource not developed yet
		"vdc":          {factTypeDatasource: {}}, // resource not developed yet
	},
}

// getCopyrightYear get copyright year from git log and file modify time.
func getCopyrightYear(filePath string) (string, error) {
	// Sanitize the filePath
	filePath = filepath.Clean(filePath)
	currYear := fmt.Sprintf("%d", time.Now().Year())
	cmd := exec.Command("bash", "-c", "git log --follow --format=%cd --date=format:%Y "+filePath+" | sort -u") // #nosec G204
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	fmt.Println(filePath, "git-log: (", string(output), ") ")
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	// if newly created
	if lines[0] == "" {
		return currYear, nil
	}
	startYear := lines[0]
	if len(lines) == 1 {
		return startYear, nil
	}
	endYear := lines[len(lines)-1]
	return startYear + "-" + endYear, nil
}

type FactNormalized struct {
	Fact
	SubCategory string
}

func normalizeFacts(in map[string]map[string]map[string]Fact) (resources, datasources map[string]FactNormalized) {
	resources = make(map[string]FactNormalized)
	datasources = make(map[string]FactNormalized)
	for subCategory, citem := range in {
		for name, nitem := range citem {
			for factType, fact := range nitem {
				if factType == factTypeResource {
					resources[name] = FactNormalized{Fact: fact, SubCategory: subCategory}
				} else if factType == factTypeDatasource {
					datasources[name] = FactNormalized{Fact: fact, SubCategory: subCategory}
				}
			}
		}
	}
	return resources, datasources
}

// main function to traveser docs folder and update copyright year.
func main() {

	err := filepath.Walk("docs/", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		path = filepath.Clean(path)
		file, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		year, err := getCopyrightYear(path)
		if err != nil {
			return err
		}
		println("Copyright Years: " + year + " " + path)
		replacedFile := strings.ReplaceAll(string(file), "<copyright-year>", year)

		// get the file and directory name by splitting by '/' and getting last 2
		pathHierarchy := strings.Split(path, "/")
		fileName := strings.TrimSuffix(pathHierarchy[len(pathHierarchy)-1], ".md")
		dirName := pathHierarchy[len(pathHierarchy)-2]

		var fnote, subCategory string
		resourceFacts, datasourceFacts := normalizeFacts(facts)
		// if dir is datasource
		if dirName == "data-sources" {
			// if note exist
			if note, ok := datasourceFacts[fileName]; ok {
				// add note
				if note.Note != "" {
					fnote = "\n\n" + note.Note
				}
				// add subcategory
				subCategory = note.SubCategory
			}
		}

		// if dir is resource
		if dirName == "resources" {
			// if note exist
			if note, ok := resourceFacts[fileName]; ok {
				// add note
				if note.Note != "" {
					fnote = "\n\n" + note.Note
				}
				// add subcategory
				subCategory = note.SubCategory
			}
		}

		// replace <subcategory>
		replacedFile = strings.ReplaceAll(replacedFile, "<subcategory>", subCategory)
		// replace <note>
		replacedFile = strings.ReplaceAll(replacedFile, "<note>", fnote)

		err = os.WriteFile(path, []byte(replacedFile), 0600)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
}
