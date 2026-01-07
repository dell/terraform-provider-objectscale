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
# After `terraform apply` of this example file it will create a new namespace with the name set in `name` attribute on the ObjectScale

resource "objectscale_namespace" "all" {
  name                        = "namespace1"
  default_data_services_vpool = "urn:storageos:ReplicationGroupInfo:55ca12b2-e908-4bac-a5fe-3fdaa975e3eb:global"
  allowed_vpools_list = [
    "urn:storageos:ReplicationGroupInfo:cd8bffcb-7a99-4023-82a8-982054fd73c2:global"
  ]
  disallowed_vpools_list = [
    "urn:storageos:ReplicationGroupInfo:e0b539a3-6ddd-4412-b4d0-ce08049f64cd:global"
  ]
  namespace_admins = "admin2,admin3"
  user_mapping = [{
    domain = "domain2"
    groups = ["group3", "group4"]
    attributes = [
      {
        key   = "key3"
        value = ["value5", "value6"]
      },
      {
        key   = "key4"
        value = ["value7", "value8"]
      }
    ]
  }]
  default_bucket_block_size       = 1024
  external_group_admins           = "admin3@foo,admin4@bar"
  is_stale_allowed                = true
  is_object_lock_with_ado_allowed = true
  retention_classes = [
    {
      name   = "class1"
      period = 500
    },
    {
      name   = "class2"
      period = 2000
    }
  ]
  quota = {
    notification_size = 90
    block_size        = 124
  }
  root_user_password = "password1"
}

# After the execution of above resource block, namespace would have been created on the ObjectScale array. For more information, Please check the terraform state file. 
