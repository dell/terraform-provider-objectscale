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

# Available actions: Create, Update, Delete and Import
# After `terraform apply` of this example file it will create a new user access key with the name set in `name` attribute on the ObjectScale

resource "objectscale_iam_user_access_key" "test_iam_user_access_key" {
  username  = "sample_user_1"
  namespace = "ns1"
  status    = "Active"
  id        = "AKIA80817B9F1F4C72CB"


}

# After the execution of above resource block, access key would have been created on the user of the ObjectScale array. For more information, Please check the terraform state file. 
