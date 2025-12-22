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

package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ObjectUserResourceModel struct {
	Tags      types.List   `tfsdk:"tags"`
	Name      types.String `tfsdk:"name"`
	Namespace types.String `tfsdk:"namespace"`
	Locked    types.Bool   `tfsdk:"locked"`
	Created   types.String `tfsdk:"created"`
	Id        types.String `tfsdk:"id"`
}

type ObjectUserTags struct {
	// A single-valued attribute indicating the user's IDP domain
	Name types.String `tfsdk:"name"`
	// Attributes
	Value types.String `tfsdk:"value"`
}
