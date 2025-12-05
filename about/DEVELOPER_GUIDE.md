<!--
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
-->

# Developer Guide

This is a guide for new contributors.

## Generating client from OpenAPI specs

The generated client is present in the `internal/clientgen` folder and is used as a Golang submodule.
The `clientgen_utils` folder contains everything required to generate the client package.

The file `clientgen_utils/openapi_specs/ecs_metadata_openapi_4.1.json` is the ObjectScale 4.1 OpenAPI specification. This file is processed to generate the file `clientgen/openapi_specs/ecs_metadata_openapi_4.1_filtered.json`.
The following types of processing is carried out:

1. We filter the spec to include only the APIs and models that we need. This keeps our client code size small.

The required APIs are specified in the `clientgen_utils/requiredApis.py` file.
Processing of the OpenAPI spec is carried out by all the other python files in the `clientgen_utils` folder.
The Makefile target `build_spec` runs the python program for processing the OpenAPI spec.

Then the client code is generated using the `build_client` Makefile target. This invokes openapi-generator-cli-6.6.0 to generate the client code using `clientgen_utils/config.yaml` and the mustache templates in `clientgen_utils/templates`. The code generation also requires `goimports` to be installed.

Installing openapi-generator-cli-6.6.0 involves running at the root of this repo
```
https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.6.0/openapi-generator-cli-6.6.0.jar
```
The openapi generator cli requires atleast 16GB RAM, otherwise it may crash mid-generation.
Installing goimports
```
go install golang.org/x/tools/cmd/goimports@latest
```

Generation of client code is a fully automated process and is verified by Github Actions on every PR.

# Running Acceptance Tests

In order to run acceptance tests, you need to set the following environment variables

```
TF_ACC=1
OBJECTSCALE_USERNAME=<username>
OBJECTSCALE_PASSWORD=<password>
OBJECTSCALE_ENDPOINT=https://<ip>:4443
OBJECTSCALE_LOGOUT_USERNAME=<logoutusername>
OBJECTSCALE_LOGOUT_PASSWORD=<logoutpassword>
```

The above is not an exhaustive list. More environment variables may be required for certain tests.
The above set of environment variables are required for all tests.

## The logout username and password

`OBJECTSCALE_USERNAME` and `OBJECTSCALE_PASSWORD` are used to authenticate to run all tests.
We found quite early on that running acceptance tests quickly hits the maximum token limit.

But an actual Terraform user will likely not run into this issue, since they will not likely be invoking Terraform in such quick succession.
For them, it is better if they can use the same token across all resources and datasources.

So we just made changes in the test cleanup logic. We introduced the variables `OBJECTSCALE_LOGOUT_USERNAME` and `OBJECTSCALE_LOGOUT_PASSWORD`.
These credentials will be used to revoke all tokens generated for the user `OBJECTSCALE_USERNAME` at the end of each acceptance test.
