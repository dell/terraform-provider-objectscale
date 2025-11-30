# Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://mozilla.org/MPL/2.0/


# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# The command is
# terraform import objectscale_iam_inline_policy.example <namespace>:<entity_type>:<entity_name>
# <namespace>: Namespace to which the IAM entity belongs.
# <entity_type>: One of user, group, or role.
# <entity_name>: The name of the user, group, or role.
# Examples:
# Import inline policies for a user
# terraform import objectscale_iam_inline_policy.example ns1:user:userTest1
# Import inline policies for a group
# terraform import objectscale_iam_inline_policy.example ns1:group:groupTest1
# Import inline policies for a role
# terraform import objectscale_iam_inline_policy.example ns1:role:roleTest1

# after running this command, populate the policies field and other required parameters in the config file to start managing this resource.
# Note: running "terraform show" after importing shows the current config/state of the resource. You can copy/paste that config to make it easier to manage the resource.
