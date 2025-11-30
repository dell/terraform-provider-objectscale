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

import "net/http"

// Helper functions for recursively fetching paginated data from ObjectScale

// PaginatedResp is an interface that represents a paginated response.
// It provides methods for getting the next marker and the paginated response data.
type PaginatedResp[T any] interface {
	// GetNextMarker returns the next marker for pagination.
	GetNextMarker() *string
	// GetPaginatedResp returns the paginated response data.
	GetPaginatedResp() []T
}

// PaginatedReq is an interface that represents a paginated request.
// It provides methods for executing the request and setting the marker for pagination.
type PaginatedReq[self any, resp any] interface {
	// Execute executes the request and returns the response, HTTP response, and any error.
	Execute() (resp, *http.Response, error)
	// Marker sets the marker for pagination.
	Marker(marker string) self
}

// GetAllInstances is a function that recursively fetches all paginated data from ObjectScale.
// It takes a paginated request as input and returns the combined paginated data and any error.
func GetAllInstances[T any, rs PaginatedResp[T], rq PaginatedReq[rq, rs]](in rq) ([]T, error) {
	// Initialize an empty slice to store the combined paginated data.
	var ret []T
	// Loop until there is no more paginated data.
	for {
		// Execute the request and get the response, HTTP response, and any error.
		resp, _, err := in.Execute()
		// If there is an error, return the error.
		if err != nil {
			return nil, err
		}
		// Append the paginated response data to the combined data.
		ret = append(ret, resp.GetPaginatedResp()...)
		// If there is no next marker, break the loop.
		if nm := resp.GetNextMarker(); nm == nil || *nm == "" {
			break
		} else {
			in = in.Marker(*nm)
		}
	}
	// Return the combined paginated data and no error.
	return ret, nil
}
