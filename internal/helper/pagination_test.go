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
	"fmt"
	"net/http"
	"terraform-provider-objectscale/internal/clientgen"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

// Test Mock Fetch Paginated Namespaces
func TestMockPagination(t *testing.T) {
	times := 0
	values := []clientgen.NamespaceServiceGetNamespacesResponseNamespaceInner{
		{
			Name: SetDefault(nil, "a"),
		},
		{
			Name: SetDefault(nil, "b"),
		},
		{
			Name: SetDefault(nil, "c"),
		},
	}
	FunctionMocker := mockey.Mock(clientgen.ApiNamespaceServiceGetNamespacesRequest.Execute).
		To(func() (*clientgen.NamespaceServiceGetNamespacesResponse, *http.Response, error) {
			ret := &clientgen.NamespaceServiceGetNamespacesResponse{
				Namespace: []clientgen.NamespaceServiceGetNamespacesResponseNamespaceInner{
					values[times],
				},
			}
			if times < 2 {
				ret.NextMarker = SetDefault(nil, "dummy")
			}
			times++
			return ret, nil, nil
		}).Build()

	dsreq := clientgen.ApiNamespaceServiceGetNamespacesRequest{}
	ret, err := GetAllInstances(dsreq)
	assert.Nil(t, err)
	assert.NotNil(t, ret)
	assert.Equal(t, values, ret)
	FunctionMocker.UnPatch()
}

// Test Mock Fetch Paginated Namespaces Error
func TestMockPaginationError(t *testing.T) {
	FunctionMocker := mockey.Mock(clientgen.ApiNamespaceServiceGetNamespacesRequest.Execute).
		Return(nil, nil, fmt.Errorf("mock error")).Build()

	dsreq := clientgen.ApiNamespaceServiceGetNamespacesRequest{}
	ret, err := GetAllInstances(dsreq)
	assert.NotNil(t, err)
	assert.Nil(t, ret)

	FunctionMocker.UnPatch()
}
