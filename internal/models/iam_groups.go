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

import "github.com/hashicorp/terraform-plugin-framework/types"

type IAMGroupsDatasourceModel struct {
    ID        types.String   `tfsdk:"id"`
    Namespace types.String   `tfsdk:"namespace"`
    GroupName types.String   `tfsdk:"group_name"`
    UserName  types.String   `tfsdk:"user_name"`  // <-- add this
    Groups    []IAMGroupModel `tfsdk:"groups"`
}

type IAMGroupModel struct {
    GroupName  types.String   `tfsdk:"group_name"`
    GroupId    types.String   `tfsdk:"group_id"`
    Arn        types.String   `tfsdk:"arn"`
    Path       types.String   `tfsdk:"path"`
    CreateDate types.String   `tfsdk:"create_date"`
    Users      []types.String `tfsdk:"users"`
}
