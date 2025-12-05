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
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// TfString - Converts *string to types.String, returns types.StringNull if input is nil.
func TfString[T ~string](in *T) types.String {
	if in == nil {
		return types.StringNull()
	}
	return types.StringValue(string(*in))
}

func TfStringNN[T ~string](in *T) types.String {
	if in == nil {
		return types.StringValue("")
	}
	return types.StringValue(string(*in))
}

// TfStringFromPTime - Converts *time.Time to types.String, returns types.StringNull if input is nil.
func TfStringFromPTime(in *time.Time) types.String {
	if in == nil {
		return types.StringNull()
	}
	return types.StringValue((*in).String())
}

// TfBool - Converts *bool to types.Bool, returns types.BoolNull if input is nil.
func TfBool(in *bool) types.Bool {
	if in == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*in)
}

func TfBoolNN(in *bool) types.Bool {
	if in == nil {
		return types.BoolValue(false)
	}
	return types.BoolValue(*in)
}

// TfInt64 - Converts *int64 to types.Int64, returns types.Int64Null if input is nil.
func TfInt64(in *int64) types.Int64 {
	if in == nil {
		return types.Int64Null()
	}
	return types.Int64Value(*in)
}

func TfInt64NN(in *int64) types.Int64 {
	if in == nil {
		return types.Int64Value(0)
	}
	return types.Int64Value(*in)
}

// TfObject - Converts input using the transform transform function, returns empty output if input is nil.
func TfObject[tfT any, jT any](in *jT, transform func(jT) tfT) tfT {
	if in == nil {
		var ret tfT
		return ret
	}
	return transform(*in)
}

// ValueToPointer - Extracts Go value pointer from attr.Value
// Returns nil if input is not known
// Supported types: types.String, types.Bool
// We can add more types in the future when required

type GoTypes interface {
	~bool | ~string | ~int64
}

func ValueToPointer[T GoTypes, VT attr.Value](in VT) *T {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	var ret any
	switch inv := any(in).(type) {
	case types.String:
		ret = inv.ValueString()
	case types.Bool:
		ret = inv.ValueBool()
	case types.Int64:
		ret = inv.ValueInt64()
	}

	switch retv := ret.(type) {
	case T:
		return &retv
	}
	return nil
}

func ValueToList[T GoTypes](in types.List) []T {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	var ret []T
	for _, i := range in.Elements() {
		x := ValueToPointer[T](i)
		if x != nil {
			ret = append(ret, *x)
		}
	}
	return ret
}

func ValueObjectTransform[T any, Tf any](in types.Object, transform func(Tf) T) T {
	var ret Tf
	in.As(context.Background(), &ret, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	return transform(ret)
}

type TfCollection interface {
	attr.Value
	ElementsAs(context.Context, interface{}, bool) diag.Diagnostics
}

func ValueListTransform[T any, Tf any](in TfCollection, transform func(Tf) T) []T {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	var unm []Tf
	var ret []T
	in.ElementsAs(context.Background(), &unm, true)
	for _, i := range unm {
		ret = append(ret, transform(i))
	}
	return ret
}

// SliceTransform - Applies the transform function to each element in a slice.
func SliceTransform[tfT any, jT any](in []jT, transform func(jT) tfT) []tfT {
	ret := make([]tfT, len(in))
	for i, v := range in {
		ret[i] = transform(v)
	}
	return ret
}

func ListTransform[T any, V attr.Value](in []T, transform func(T) V) types.List {
	var dummy T
	if in == nil {
		return types.ListNull(transform(dummy).Type(context.Background()))
	}
	return types.ListValueMust(
		transform(dummy).Type(context.Background()),
		SliceTransform(in, func(x T) attr.Value {
			return transform(x)
		}),
	)
}

func SetNotNull[T any, V attr.Value](in []T, transform func(T) V) types.Set {
	var dummy T
	if in == nil {
		in = make([]T, 0)
	}
	return types.SetValueMust(
		transform(dummy).Type(context.Background()),
		SliceTransform(in, func(x T) attr.Value {
			return transform(x)
		}),
	)
}

func ListNotNull[T any, V attr.Value](in []T, transform func(T) V) types.List {
	if in == nil {
		in = make([]T, 0)
	}
	return ListTransform(in, transform)
}

// SetDefault - Returns pointer to default value if input is nil, otherwise returns input.
func SetDefault[T any](in *T, defaultVal T) *T {
	if in != nil {
		return in
	}
	return &defaultVal
}

func Object(in any) types.Object {
	intypes := make(map[string]attr.Type)
	invals := StructToValuesReflection(in)
	for k, v := range invals {
		intypes[k] = v.Type(context.Background())
	}
	return types.ObjectValueMust(intypes, invals)
}

// TfInt64From32 safely converts a *int32 to types.Int64 for Terraform models.
func TfInt64From32(v *int32) types.Int64 {
	if v == nil {
		return types.Int64Null()
	}
	return types.Int64Value(int64(*v))
}
