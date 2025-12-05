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
    """In GetRoleResponse, Result property should be GetRoleResult.
       Inner property Role should be normalised to IamRole.
       Also for UpdateRoleResponse inner property Role should be normalised to IamRole.
       Also for CreateRoleResponse inner property Role should be normalised to IamRole.
       Also for ListRolesResponse inner property Roles should be normalised to IamRole.
    """

    # Check GetRoleResponse exists
    if (
        "components" in json_obj and
        "schemas" in json_obj["components"] and
        "IamService_GetRoleResponse" in json_obj["components"]["schemas"]
    ):
        get_role_schema = json_obj["components"]["schemas"]["IamService_GetRoleResponse"]

        if (
            "properties" in get_role_schema and
            "Result" in get_role_schema["properties"] and
            "properties" in get_role_schema["properties"]["Result"] and
            "Role" in get_role_schema["properties"]["Result"]["properties"]
        ):
            common_role = get_role_schema["properties"]["Result"]["properties"]["Role"]

            if (
                "properties" in common_role and
                "Tags" in common_role["properties"] and
                "items" in common_role["properties"]["Tags"]
            ):
                common_role["properties"]["Tags"]["items"] = {
                    "$ref": "#/components/schemas/IamTagKeyValue"
                }

            json_obj['components']['schemas']['IamRole'] = common_role
            json_obj['components']['schemas']['IamRoleResult'] = {
                "type": "object",
                "properties": {
                    "Role": {
                        "$ref": "#/components/schemas/IamRole"
                    }
                }
            }

    # GetRole Response normalization
    if "IamService_GetRoleResponse" in json_obj['components']['schemas']:
        props = json_obj['components']['schemas']['IamService_GetRoleResponse']['properties']
        if "Result" in props:
            props['GetRoleResult'] = props['Result']
            del props['Result']
            props['GetRoleResult'] = {
                "$ref": "#/components/schemas/IamRoleResult"
            }

    # UpdateRole Response normalization
    if "IamService_UpdateRoleResponse" in json_obj['components']['schemas']:
        props = json_obj['components']['schemas']['IamService_UpdateRoleResponse']['properties']
        if "Result" in props:
            props['UpdateRoleResult'] = props['Result']
            del props['Result']
            props['UpdateRoleResult'] = {
                "$ref": "#/components/schemas/IamRoleResult"
            }

    # CreateRole Response normalization
    if "IamService_CreateRoleResponse" in json_obj['components']['schemas']:
        props = json_obj['components']['schemas']['IamService_CreateRoleResponse']['properties']
        if "Result" in props:
            props['CreateRoleResult'] = props['Result']
            del props['Result']
            props['CreateRoleResult'] = {
                "$ref": "#/components/schemas/IamRoleResult"
            }

    # ListRoles Response normalization
    if "IamService_ListRolesResponse" in json_obj['components']['schemas']:
        props = json_obj['components']['schemas']['IamService_ListRolesResponse']['properties']
        if "Result" in props:
            props['ListRolesResult'] = props['Result']
            del props['Result']

            # Ensure sub-properties exist
            if (
                "properties" in props['ListRolesResult'] and
                "member" in props['ListRolesResult']["properties"]
            ):
                props['ListRolesResult']["properties"]["Roles"] = (
                    props['ListRolesResult']["properties"]["member"]
                )
                del props['ListRolesResult']["properties"]["member"]

                if "items" in props['ListRolesResult']["properties"]["Roles"]:
                    props['ListRolesResult']["properties"]["Roles"]['items'] = {
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
