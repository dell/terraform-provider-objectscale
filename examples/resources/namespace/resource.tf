/*
Copyright (c) 2023-2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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

resource "objectscale_namespace" "namespace" {
  name                        = "test_namespace"
  default_data_services_vpool = "urn:storageos:ReplicationGroupInfo:0e953ad1-94a5-4eb1-825a-d58d29e85434:global"
  retention_classes = {
    retention_class = [{
      name   = "r1"
      period = 1
    }]
  }
  user_mapping = [{
    attributes = [{
      key   = "key1"
      value = ["value1"]
    }]
    domain = "domain"
    groups = ["group"]
  }]
}

# After the execution of above resource block, namespace would have been created on the ObjectScale array. For more information, Please check the terraform state file. 
