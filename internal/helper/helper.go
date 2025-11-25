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

package helper

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
)

// StructToValuesReflection converts any struct with `tfsdk` tags
// into a map[string]attr.Value. Fields must implement attr.Value.
func StructToValuesReflection(s interface{}) map[string]attr.Value {
	result := make(map[string]attr.Value)

	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	// If it's a pointer, dereference
	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("tfsdk")
		if tag == "" {
			continue
		}

		fieldVal := val.Field(i)
		if !fieldVal.IsValid() {
			continue
		}

		// Only include fields that implement attr.Value
		if v, ok := fieldVal.Interface().(attr.Value); ok {
			result[tag] = v
		}
	}

	return result
}
