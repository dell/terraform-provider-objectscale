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

def _normalizeObjectScaleIamRoleResponse(json_obj: dict) -> dict:
    """Normalize IamRole-related responses safely and modularly."""

    def get_path(obj, path: list):
        """Safely get a nested value. Returns None if any key is missing."""
        for p in path:
            if not isinstance(obj, dict) or p not in obj:
                return None
            obj = obj[p]
        return obj

    def set_path(obj, path: list, value):
        """Safely set a nested value only if all parents exist."""
        cur = obj
        for p in path[:-1]:
            if p not in cur or not isinstance(cur[p], dict):
                return False
            cur = cur[p]
        cur[path[-1]] = value
        return True

    def delete_path(obj, path: list):
        """Safely delete a path only if all keys exist."""
        cur = obj
        for p in path[:-1]:
            if p not in cur or not isinstance(cur[p], dict):
                return False
            cur = cur[p]
        return cur.pop(path[-1], None) is not None


    # Build Common IamRole Schema
    common_role_path = [
        "components", "schemas",
        "IamService_GetRoleResponse", "properties",
        "Result", "properties", "Role"
    ]
    common_role = get_path(json_obj, common_role_path)

    if common_role:
        # Fix Tags schema
        tags_path = common_role["properties"].get("Tags", {})
        if "properties" in common_role and "Tags" in common_role["properties"]:
            common_role["properties"]["Tags"]["items"] = {
                "$ref": "#/components/schemas/IamTagKeyValue"
            }

        # Create IamRole schema
        json_obj["components"]["schemas"]["IamRole"] = common_role

        # Create IamRoleResult schema
        json_obj["components"]["schemas"]["IamRoleResult"] = {
            "type": "object",
            "properties": {
                "Role": {"$ref": "#/components/schemas/IamRole"}
            }
        }

    # Helper to normalize Role responses
    def normalize_response(root_schema: str, result_key: str):
        base_path = ["components", "schemas", root_schema, "properties"]

        result_obj = get_path(json_obj, base_path + ["Result"])
        if result_obj is None:
            return

        delete_path(json_obj, base_path + ["Result"])

        set_path(json_obj, base_path + [result_key], {
            "$ref": "#/components/schemas/IamRoleResult"
        })

    normalize_response("IamService_GetRoleResponse", "GetRoleResult")
    normalize_response("IamService_UpdateRoleResponse", "UpdateRoleResult")
    normalize_response("IamService_CreateRoleResponse", "CreateRoleResult")


    # Normalize ListRolesResponse

    list_roles_path = [
        "components", "schemas",
        "IamService_ListRolesResponse", "properties"
    ]
    list_roles_props = get_path(json_obj, list_roles_path)

    if list_roles_props and "Result" in list_roles_props:
        list_roles_props["ListRolesResult"] = list_roles_props.pop("Result")

        lr = list_roles_props["ListRolesResult"]["properties"]
        if "member" in lr:
            lr["Roles"] = lr.pop("member")
            lr["Roles"]["items"] = {
                "$ref": "#/components/schemas/IamRole"
            }

    return json_obj

def _normalizeObjectScaleIamResponseMetadata(json_obj: dict) -> dict:
    """
    Normalize ObjectScale specific response metadata.
    """
    anyitems = False
    commonType = {
        "type": "object",
        "properties": {
            "RequestId": {
                "type": "string"
            }
        }
    }
    for _, obj in json_obj['components']['schemas'].items():
        if 'ResponseMetadata' in obj['properties']:
            if obj['properties']['ResponseMetadata'] == commonType:
                obj['properties']['ResponseMetadata'] = {
						"$ref": "#/components/schemas/IamResponseMetadata"
					}
                anyitems = True
    if anyitems:
        json_obj['components']['schemas']['IamResponseMetadata'] = commonType

    return json_obj

def _normalizeObjectScaleBasicResponseMetadata(json_obj: dict) -> dict:
    """
    Some responses are basically just response metadata.
    Normalize all such models.
    """

    basicResp = {
        "type": "object",
        "properties": {
            "ResponseMetadata": {
                "$ref": "#/components/schemas/IamResponseMetadata"
            }
        }
    }
    basicRespRef = "#/components/schemas/BasicResponse"

    for _, path in json_obj['paths'].items():
        for _, method in path.items():
            for _, resp in method.get('responses', {}).items():
                refstrf = resp.get('content', {}).get('application/json', {}).get('schema', {}).get('$ref')
                if refstrf:
                    refstr = refstrf.split('/')[-1]
                    ref = json_obj['components']['schemas'][refstr]
                    if ref == basicResp:
                        resp['content']['application/json']['schema']["$ref"] = basicRespRef
                        json_obj['components']['schemas']['BasicResponse'] = basicResp
                        del json_obj['components']['schemas'][refstr]

    return json_obj

def _normalizeObjectScaleIamTags(json_obj: dict) -> dict:
    """
    Add ObjectScale specific marker to the OpenAPI spec.
    Fixes specs of IAM tagging untagging APIs query parameters.
    Adds 'x-indexed-kv' to any api with Tags.member.N query parameter.
    Adds 'x-indexed-key-only' to any api with TagKeys parameters that have only keys.
    """
    # this is ok in spec, but we shall simplify it for our usecase
    commonTagKeyType = {
        "type": "object",
        "properties": {
            "key": {
                "type": "string",
                "description": "The name of the tag."
            }
        }
    }
    # this type is wrong in the spec
    commonTagKeyValueType = {
        "type": "object",
        "properties": {
            "key": {
                "type": "string",
                "description": "The name of the tag."
            },
            "value": {
                "type": "string",
                "description": "The value of the tag."
            }
        }
    }
    for path, obj in json_obj['paths'].items():
        if '/iam?' not in path:
            continue
        for param in obj.get('post', {}).get('parameters', []):
            if param['name'] == 'TagKeys':
                param['x-indexed-kv'] = 'true'
                param['x-indexed-key-only'] = 'true'
                param["schema"] = {
                    "type": "array",
                    "items": {
                        "$ref": "#/components/schemas/IamTagKey"
                    }
                }
                json_obj['components']['schemas']['IamTagKey'] = commonTagKeyType

            if param['name'] == 'Tags.member.N':
                param['x-indexed-kv'] = 'true'
                param["schema"] = {
                        "type": "array",
                        "items": {
                            "$ref": "#/components/schemas/IamTagKeyValue"
                        }
                    }
                json_obj['components']['schemas']['IamTagKeyValue'] = commonTagKeyValueType
    return json_obj

def NormalizeObjectScaleModels(json_obj: dict) -> dict:
    """
    Normalize ObjectScale specific models.
    """
    ret = _normalizeObjectScaleIamResponseMetadata(json_obj)
    ret = _normalizeObjectScaleBasicResponseMetadata(ret)
    ret = _normalizeObjectScaleIamTags(ret)
    ret = _normalizeObjectScaleIamRoleResponse(ret)
    return ret
