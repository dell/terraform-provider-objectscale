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

def NormalizeObjectScaleModels(json_obj: dict) -> dict:
    """
    Normalize ObjectScale specific models.
    """
    ret = _normalizeObjectScaleIamResponseMetadata(json_obj)
    ret = _normalizeObjectScaleBasicResponseMetadata(ret)
    return ret
