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

package helper

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CopyFields copy the source of a struct to destination of struct with terraform types.
// Unsigned integers are not properly handled.
func CopyFields(ctx context.Context, source, destination interface{}) error {
	tflog.Debug(ctx, "Copy fields", map[string]interface{}{
		"source":      source,
		"destination": destination,
	})
	sourceValue := reflect.ValueOf(source)
	destinationValue := reflect.ValueOf(destination)

	// Check if destination is a pointer to a struct
	if destinationValue.Kind() != reflect.Ptr || destinationValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("destination is not a pointer to a struct")
	}

	// if source is a pointer, use the Elem() method to get the value that the pointer points to
	if sourceValue.Kind() == reflect.Ptr {
		sourceValue = sourceValue.Elem()
	}

	if sourceValue.Kind() != reflect.Struct {
		return fmt.Errorf("source is not a struct")
	}

	// Get the type of the destination struct
	// destinationType := destinationValue.Elem().Type()
	for i := 0; i < sourceValue.NumField(); i++ {
		sourceFieldTag := getFieldJSONTag(sourceValue, i)

		tflog.Debug(ctx, "Converting source field", map[string]interface{}{
			"sourceFieldTag":  sourceFieldTag,
			"sourceFieldKind": sourceValue.Field(i).Kind().String(),
		})

		sourceField := sourceValue.Field(i)
		if sourceField.Kind() == reflect.Ptr {
			sourceField = sourceField.Elem()
		}
		if !sourceField.IsValid() {
			tflog.Error(ctx, "source field is not valid", map[string]interface{}{
				"sourceFieldTag": sourceFieldTag,
				"sourceField":    sourceField,
			})
			continue
		}

		destinationField := getFieldByTfTag(destinationValue.Elem(), sourceFieldTag)
		if destinationField.IsValid() && destinationField.CanSet() {

			tflog.Debug(ctx, "debugging source field", map[string]interface{}{
				"sourceField Interface": sourceField.Interface(),
			})
			// Convert the source value to the type of the destination field dynamically
			var destinationFieldValue attr.Value

			switch sourceField.Kind() {
			case reflect.String:
				destinationFieldValue = types.StringValue(sourceField.String())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				destinationFieldValue = types.Int64Value(sourceField.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if sourceField.Uint() > math.MaxInt64 {
					return fmt.Errorf("source field value is too large for int64")
				}
				destinationFieldValue = types.Int64Value(int64(sourceField.Uint())) // #nosec G115 --- validated, Error returned if value is too large for int64
			case reflect.Float32, reflect.Float64:
				// destinationFieldValue = types.Float64Value(sourceField.Float())
				destinationFieldValue = types.NumberValue(big.NewFloat(sourceField.Float()))
			case reflect.Bool:
				destinationFieldValue = types.BoolValue(sourceField.Bool())
			case reflect.Array, reflect.Slice:
				if destinationField.Type().Kind() == reflect.Slice {
					arr := reflect.ValueOf(sourceField.Interface())
					slice := reflect.MakeSlice(destinationField.Type(), arr.Len(), arr.Cap())
					for index := 0; index < arr.Len(); index++ {
						value := arr.Index(index)
						v := slice.Index(index)
						switch v.Kind() {
						case reflect.Ptr:
							newDes := reflect.New(v.Type().Elem()).Interface()
							err := CopyFields(ctx, value.Interface(), newDes)
							if err != nil {
								return err
							}
							slice.Index(index).Set(reflect.ValueOf(newDes))
						case reflect.Struct:
							newDes := reflect.New(v.Type()).Interface()
							err := CopyFields(ctx, value.Interface(), newDes)
							if err != nil {
								return err
							}
							slice.Index(index).Set(reflect.ValueOf(newDes).Elem())
						}
					}
					destinationField.Set(slice)
				} else if /* check if destination is types.Set */ _, ok := destinationField.Interface().(types.Set); ok {
					destinationFieldValue = copySliceToSetType(ctx, sourceField.Interface())

				} else {
					destinationFieldValue = copySliceToTargetField(ctx, sourceField.Interface())
				}
			case reflect.Struct:
				// placeholder for improvement, need to consider both go struct and types.Object
				switch destinationField.Kind() {
				case reflect.Ptr:
					newDes := reflect.New(destinationField.Type().Elem()).Interface()
					err := CopyFields(ctx, sourceField.Interface(), newDes)
					if err != nil {
						return err
					}
					destinationField.Set(reflect.ValueOf(newDes))
				case reflect.Struct:
					newDes := reflect.New(destinationField.Type()).Interface()
					err := CopyFields(ctx, sourceField.Interface(), newDes)
					if err != nil {
						return err
					}
					destinationField.Set(reflect.ValueOf(newDes).Elem())
				}
				continue

			default:
				tflog.Error(ctx, "unsupported source field type", map[string]interface{}{
					"sourceField": sourceField,
				})
				continue
			}

			if destinationField.Type() == reflect.TypeOf(destinationFieldValue) {
				destinationField.Set(reflect.ValueOf(destinationFieldValue))
			}
		}
	}

	return nil
}

func getFieldJSONTag(sourceValue reflect.Value, i int) string {
	sourceFieldTag := sourceValue.Type().Field(i).Tag.Get("tf")
	//sourceFieldTag = strings.TrimSuffix(sourceFieldTag, ",omitempty")
	return sourceFieldTag
}

func getFieldByTfTag(destinationValue reflect.Value, tagValue string) reflect.Value {
	for j := 0; j < destinationValue.NumField(); j++ {
		field := destinationValue.Type().Field(j)
		if field.Tag.Get("tfsdk") == tagValue {
			return destinationValue.Field(j)
		}
	}
	return reflect.Value{}
}

func copySliceToSetType(ctx context.Context, fields any) types.Set {
	listVal := copySliceToTargetField(ctx, fields)
	if listVal.IsUnknown() {
		return types.SetUnknown(listVal.ElementType(ctx))
	}
	setValue, _ := types.SetValue(listVal.ElementType(ctx), listVal.Elements())
	return setValue
}

func copySliceToTargetField(ctx context.Context, fields interface{}) types.List {
	var objects []attr.Value
	attrTypeMap := make(map[string]attr.Type)

	// get the attrType for Object
	structElem := reflect.ValueOf(fields).Type().Elem()
	switch structElem.Kind() {
	case reflect.String:
		listValue, _ := types.ListValueFrom(ctx, types.StringType, fields)
		return listValue
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		listValue, _ := types.ListValueFrom(ctx, types.Int64Type, fields)
		return listValue
	case reflect.Float32, reflect.Float64:
		listValue, _ := types.ListValueFrom(ctx, types.Float64Type, fields)
		return listValue
	case reflect.Bool:
		listValue, _ := types.ListValueFrom(ctx, types.BoolType, fields)
		return listValue
	case reflect.Struct:
		for fieldIndex := 0; fieldIndex < structElem.NumField(); fieldIndex++ {
			field := structElem.Field(fieldIndex)
			tag := field.Tag.Get("json")
			tag = strings.TrimSuffix(tag, ",omitempty")
			fieldType := field.Type
			if fieldType.Kind() == reflect.Ptr {
				fieldType = fieldType.Elem()
			}

			switch fieldType.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				attrTypeMap[tag] = types.Int64Type
			case reflect.String:
				attrTypeMap[tag] = types.StringType
			case reflect.Float32, reflect.Float64:
				attrTypeMap[tag] = types.NumberType
			}
		}
		// iterate the slice
		arr := reflect.ValueOf(fields)
		for index := 0; index < arr.Len(); index++ {
			valueMap := make(map[string]attr.Value)
			// iterate the fields
			elem := arr.Index(index)
			for fieldIndex := 0; fieldIndex < elem.NumField(); fieldIndex++ {
				tag := elem.Type().Field(fieldIndex).Tag.Get("json")
				tag = strings.TrimSuffix(tag, ",omitempty")
				eleField := elem.Field(fieldIndex)
				eleFieldType := eleField.Type()
				if eleFieldType.Kind() == reflect.Ptr {
					eleFieldType = eleFieldType.Elem()
					eleField = eleField.Elem()
				}
				switch eleFieldType.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					valueMap[tag] = types.Int64Value(eleField.Int())
				case reflect.String:
					valueMap[tag] = types.StringValue(eleField.String())
				case reflect.Float32, reflect.Float64:
					valueMap[tag] = types.NumberValue(big.NewFloat(eleField.Float()))
				}
			}
			object, _ := types.ObjectValue(attrTypeMap, valueMap)
			objects = append(objects, object)
		}
		listValue, _ := types.ListValue(types.ObjectType{AttrTypes: attrTypeMap}, objects)
		return listValue
	}
	return types.ListUnknown(types.StringType)
}

func assignObjectToField(ctx context.Context, source basetypes.ObjectValue, destination interface{}) error {
	destElemType := reflect.TypeOf(destination).Elem()
	isPtr := false
	if destElemType.Kind() == reflect.Ptr {
		isPtr = true
		destElemType = destElemType.Elem()
	}
	targetObject := reflect.New(destElemType).Elem()
	attrMap := source.Attributes()
	for key, val := range attrMap {
		destinationField, err := getFieldByJSONTag(targetObject.Addr().Interface(), key)
		if err != nil {
			// skip current field
			continue
		}
		if destinationField.IsValid() && destinationField.CanSet() {
			switch val.Type(ctx) {
			case basetypes.StringType{}:
				stringVal, ok := val.(basetypes.StringValue)
				if !ok || stringVal.IsNull() || stringVal.IsUnknown() {
					continue
				}
				targetValue := stringVal.ValueString()
				if destinationField.Kind() == reflect.Ptr && destinationField.Type().Elem().Kind() == reflect.String {
					destinationField.Set(reflect.ValueOf(&targetValue))
				}
				if destinationField.Type().Kind() == reflect.String {
					destinationField.Set(reflect.ValueOf(targetValue))
				}
			case basetypes.Int64Type{}:
				intVal, ok := val.(basetypes.Int64Value)
				if !ok || intVal.IsNull() || intVal.IsUnknown() {
					continue
				}
				if destinationField.Kind() == reflect.Int64 {
					destinationField.Set(reflect.ValueOf(intVal.ValueInt64()))
				}
				if destinationField.Kind() == reflect.Ptr && destinationField.Type().Elem().Kind() == reflect.Int64 {
					destinationField.Set(reflect.ValueOf(intVal.ValueInt64Pointer()))
				}
				if destinationField.Kind() == reflect.Int32 {
					destinationField.Set(reflect.ValueOf(int32(intVal.ValueInt64())))
				}
				if destinationField.Kind() == reflect.Ptr && destinationField.Type().Elem().Kind() == reflect.Int32 {
					val := int32(intVal.ValueInt64())
					destinationField.Set(reflect.ValueOf(&val))
				}
			case basetypes.BoolType{}:
				boolVal, ok := val.(basetypes.BoolValue)
				if !ok || boolVal.IsNull() || boolVal.IsUnknown() {
					continue
				}
				if destinationField.Kind() == reflect.Ptr {
					destinationField.Set(reflect.ValueOf(boolVal.ValueBoolPointer()))
				} else {
					destinationField.Set(reflect.ValueOf(boolVal.ValueBool()))
				}
			case basetypes.NumberType{}:
				floatVal, ok := val.(basetypes.NumberValue)
				if !ok || floatVal.IsNull() || floatVal.IsUnknown() {
					continue
				}
				bigFloat := floatVal.ValueBigFloat()
				if destinationField.Kind() == reflect.Ptr {
					if destinationField.Type().Elem().Kind() == reflect.Float64 {
						bigFloatVal, _ := bigFloat.Float64()
						floatVal, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", bigFloatVal), 64)
						destinationField.Set(reflect.ValueOf(&floatVal))
					}
					if destinationField.Type().Elem().Kind() == reflect.Float32 {
						bigFloatVal, _ := bigFloat.Float32()
						floatVal, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", bigFloatVal), 32)
						float32Val := float32(floatVal)
						destinationField.Set(reflect.ValueOf(&float32Val))
					}
				} else {
					if destinationField.Kind() == reflect.Float64 {
						bigFloatVal, _ := bigFloat.Float64()
						floatVal, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", bigFloatVal), 64)
						destinationField.Set(reflect.ValueOf(floatVal))
					}
					if destinationField.Kind() == reflect.Float32 {
						bigFloatVal, _ := bigFloat.Float32()
						floatVal, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", bigFloatVal), 32)
						destinationField.Set(reflect.ValueOf(float32(floatVal)))
					}
				}
			default:
				typeString := val.Type(ctx).String()
				if strings.HasPrefix(typeString, "types.ObjectType") {
					objVal, ok := val.(basetypes.ObjectValue)
					if !ok || objVal.IsNull() || objVal.IsUnknown() {
						continue
					}
					err := assignObjectToField(ctx, objVal, destinationField.Addr().Interface())
					if err != nil {
						return err
					}
				} else if strings.HasPrefix(typeString, "types.ListType") {
					listVal, ok := val.(basetypes.ListValue)
					if !ok || listVal.IsNull() || listVal.IsUnknown() {
						continue
					}
					list, err := getFieldListVal(ctx, listVal, destinationField.Interface())
					if err != nil {
						return err
					}
					if reflect.TypeOf(destinationField.Interface()).Kind() == reflect.Ptr {
						destinationField.Set(reflect.New(destinationField.Type().Elem()))
						destinationField.Elem().Set(list)
					} else {
						destinationField.Set(list)
					}
				}
			}
		}
	}
	if isPtr {
		reflect.ValueOf(destination).Elem().Set(targetObject.Addr())
	} else {
		reflect.ValueOf(destination).Elem().Set(targetObject)
	}
	return nil
}

func getFieldByJSONTag(destination interface{}, tag string) (reflect.Value, error) {
	destElemVal := reflect.ValueOf(destination).Elem()
	destElemType := destElemVal.Type()

	for i := 0; i < destElemType.NumField(); i++ {
		field := destElemType.Field(i)
		jsonTag := field.Tag.Get("tf")
		// if strings.Contains(jsonTag, ",") {
		// 	jsonTag = strings.TrimSuffix(jsonTag, ",omitempty")
		// }
		if jsonTag == tag {
			return destElemVal.Field(i), nil
		}
	}

	return reflect.Value{}, fmt.Errorf("field with tag %s not found in destination", tag)
}

func getFieldListVal(ctx context.Context, source basetypes.ListValue, destination interface{}) (reflect.Value, error) {
	destType := reflect.TypeOf(destination)
	if destType.Kind() == reflect.Ptr {
		destType = destType.Elem()
	}
	listLen := len(source.Elements())
	targetList := reflect.MakeSlice(destType, listLen, listLen)
	listElemType := source.ElementType(ctx)
	for i, listElem := range source.Elements() {
		switch listElemType {
		case basetypes.StringType{}:
			strVal, ok := listElem.(basetypes.StringValue)
			if !ok || strVal.IsNull() || strVal.IsUnknown() {
				continue
			}
			if destType.Elem().Kind() == reflect.Ptr {
				targetList.Index(i).Elem().Set(reflect.ValueOf(strVal.ValueStringPointer()))
			} else {
				targetList.Index(i).Set(reflect.ValueOf(strVal.ValueString()))
			}
		case basetypes.Int64Type{}:
			strVal, ok := listElem.(basetypes.Int64Value)
			if !ok || strVal.IsNull() || strVal.IsUnknown() {
				continue
			}
			if destType.Elem().Kind() == reflect.Ptr {
				targetList.Index(i).Elem().Set(reflect.ValueOf(strVal.ValueInt64Pointer()))
			} else {
				targetList.Index(i).Set(reflect.ValueOf(strVal.ValueInt64()))
			}
		case basetypes.BoolType{}:
			strVal, ok := listElem.(basetypes.BoolValue)
			if !ok || strVal.IsNull() || strVal.IsUnknown() {
				continue
			}
			if destType.Elem().Kind() == reflect.Ptr {
				targetList.Index(i).Elem().Set(reflect.ValueOf(strVal.ValueBoolPointer()))
			} else {
				targetList.Index(i).Set(reflect.ValueOf(strVal.ValueBool()))
			}
		default:
			typeString := listElemType.String()
			if strings.HasPrefix(typeString, "types.ListType") {
				listVal, ok := listElem.(basetypes.ListValue)
				if !ok || listVal.IsNull() || listVal.IsUnknown() {
					continue
				}
				val, err := getFieldListVal(ctx, listVal, targetList.Index(i).Interface())
				if err != nil {
					return targetList, err
				}
				if reflect.TypeOf(targetList.Index(i).Interface()).Kind() == reflect.Ptr {
					targetList.Index(i).Set(reflect.New(targetList.Index(i).Type().Elem()))
					targetList.Index(i).Elem().Set(val)
				} else {
					targetList.Index(i).Set(val)
				}
			} else if strings.HasPrefix(typeString, "types.ObjectType") {
				objVal, ok := listElem.(basetypes.ObjectValue)
				if !ok || objVal.IsNull() || objVal.IsUnknown() {
					continue
				}
				err := assignObjectToField(ctx, objVal, targetList.Index(i).Addr().Interface())
				if err != nil {
					return targetList, err
				}
			}
		}
	}
	return targetList, nil
}

// ReadFromState read from model to openapi struct, model should not contain nested struct.
func readFromState(ctx context.Context, source, destination interface{}) error {
	sourceValue := reflect.ValueOf(source)
	destinationValue := reflect.ValueOf(destination)
	if destinationValue.Kind() != reflect.Ptr || destinationValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("destination is not a pointer to a struct")
	}
	if sourceValue.Kind() == reflect.Ptr {
		sourceValue = sourceValue.Elem()
	}
	if sourceValue.Kind() != reflect.Struct {
		return fmt.Errorf("source is not a struct")
	}
	for i := 0; i < sourceValue.NumField(); i++ {
		sourceFieldTag := sourceValue.Type().Field(i).Tag.Get("tfsdk")
		destinationField, err := getFieldByJSONTag(destinationValue.Elem().Addr().Interface(), sourceFieldTag)
		if err != nil {
			// Not found, skip the field
			continue
		}
		if destinationField.IsValid() && destinationField.CanSet() {
			switch sourceValue.Field(i).Interface().(type) {
			case basetypes.StringValue:
				stringVal, ok := sourceValue.Field(i).Interface().(basetypes.StringValue)
				if !ok || stringVal.IsNull() || stringVal.IsUnknown() {
					continue
				}
				targetValue := stringVal.ValueString()
				if destinationField.Kind() == reflect.Ptr && destinationField.Type().Elem().Kind() == reflect.String {
					destinationField.Set(reflect.ValueOf(&targetValue))
				}
				if destinationField.Type().Kind() == reflect.String {
					destinationField.Set(reflect.ValueOf(targetValue))
				}
			case basetypes.Int64Value:
				intVal, ok := sourceValue.Field(i).Interface().(basetypes.Int64Value)
				if !ok || intVal.IsNull() || intVal.IsUnknown() {
					continue
				}
				if destinationField.Kind() == reflect.Int64 {
					destinationField.Set(reflect.ValueOf(intVal.ValueInt64()))
				}
				if destinationField.Kind() == reflect.Ptr && destinationField.Type().Elem().Kind() == reflect.Int64 {
					destinationField.Set(reflect.ValueOf(intVal.ValueInt64Pointer()))
				}
				if destinationField.Kind() == reflect.Int32 {
					destinationField.Set(reflect.ValueOf(int32(intVal.ValueInt64())))
				}
				if destinationField.Kind() == reflect.Ptr && destinationField.Type().Elem().Kind() == reflect.Int32 {
					val := int32(intVal.ValueInt64())
					destinationField.Set(reflect.ValueOf(&val))
				}
			case basetypes.BoolValue:
				boolVal, ok := sourceValue.Field(i).Interface().(basetypes.BoolValue)
				if !ok || boolVal.IsNull() || boolVal.IsUnknown() {
					continue
				}
				if destinationField.Kind() == reflect.Ptr {
					destinationField.Set(reflect.ValueOf(boolVal.ValueBoolPointer()))
				} else {
					destinationField.Set(reflect.ValueOf(boolVal.ValueBool()))
				}
			case basetypes.NumberValue:
				floatVal, ok := sourceValue.Field(i).Interface().(basetypes.NumberValue)
				if !ok || floatVal.IsNull() || floatVal.IsUnknown() {
					continue
				}
				bigFloat := floatVal.ValueBigFloat()
				if destinationField.Kind() == reflect.Ptr {
					if destinationField.Type().Elem().Kind() == reflect.Float64 {
						bigFloatVal, _ := bigFloat.Float64()
						floatVal, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", bigFloatVal), 64)
						destinationField.Set(reflect.ValueOf(&floatVal))
					}
					if destinationField.Type().Elem().Kind() == reflect.Float32 {
						bigFloatVal, _ := bigFloat.Float32()
						floatVal, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", bigFloatVal), 32)
						float32Val := float32(floatVal)
						destinationField.Set(reflect.ValueOf(&float32Val))
					}
				} else {
					if destinationField.Kind() == reflect.Float64 {
						bigFloatVal, _ := bigFloat.Float64()
						floatVal, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", bigFloatVal), 64)
						destinationField.Set(reflect.ValueOf(floatVal))
					}
					if destinationField.Kind() == reflect.Float32 {
						bigFloatVal, _ := bigFloat.Float32()
						floatVal, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", bigFloatVal), 32)
						destinationField.Set(reflect.ValueOf(float32(floatVal)))
					}
				}
			case basetypes.ObjectValue:
				objVal, ok := sourceValue.Field(i).Interface().(basetypes.ObjectValue)
				if !ok || objVal.IsNull() || objVal.IsUnknown() {
					continue
				}
				err := assignObjectToField(ctx, objVal, destinationField.Addr().Interface())
				if err != nil {
					return err
				}
			case basetypes.ListValue:
				listVal, ok := sourceValue.Field(i).Interface().(basetypes.ListValue)
				if !ok || listVal.IsNull() || listVal.IsUnknown() {
					continue
				}
				list, err := getFieldListVal(ctx, listVal, destinationField.Interface())
				if err != nil {
					return err
				}
				if reflect.TypeOf(destinationField.Interface()).Kind() == reflect.Ptr {
					destinationField.Set(reflect.New(destinationField.Type().Elem()))
					destinationField.Elem().Set(list)
				} else {
					destinationField.Set(list)
				}
			}
		}
	}
	return nil
}
