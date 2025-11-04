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
