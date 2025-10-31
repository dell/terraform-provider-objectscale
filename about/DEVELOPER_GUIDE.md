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
