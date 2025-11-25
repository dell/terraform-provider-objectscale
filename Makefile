# Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

# http://mozilla.org/MPL/2.0/

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Copyright (c) 2023-2024 Dell Inc., or its subsidiaries. All Rights Reserved.

# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

# http://mozilla.org/MPL/2.0/

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=registry.terraform.io
NAMESPACE=dell
NAME=objectscale
BINARY=terraform-provider-${NAME}
VERSION?=1.0.0
OS_ARCH=linux_amd64

OPENAPI_CMD?=java -Xmx16G -jar openapi-generator-cli-6.6.0.jar
OPENAPI_GEN_DIR=internal/clientgen
OPENAPI_SOURCE_DIR=clientgen_utils/openapi_specs
OPENAPI_FULL_PATH=${OPENAPI_SOURCE_DIR}/ecs_metadata_openapi_4.1.json
OPENAPI_FILTERED_PATH=${OPENAPI_SOURCE_DIR}/ecs_metadata_openapi_4.1_filtered.json

default: install

build:
	go build -o ${BINARY}

install: uninstall build
	rm -rfv ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	find examples -type d -name ".terraform" -exec rm -rfv "{}" +;
	find examples -type f -name "trace.*" -delete
	find examples -type f -name "*.tfstate" -delete
	find examples -type f -name "*.hcl" -delete
	find examples -type f -name "*.backup" -delete
	rm -rf trace.*
	
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

uninstall:
	rm -rfv ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	find examples -type d -name ".terraform" -exec rm -rfv "{}" +;
	find examples -type f -name "trace.*" -delete
	find examples -type f -name "*.tfstate" -delete
	find examples -type f -name "*.hcl" -delete
	find examples -type f -name "*.backup" -delete
	rm -rf trace.*

build_spec:
	python3 clientgen_utils/main.py --input ${OPENAPI_FULL_PATH} --output ${OPENAPI_FILTERED_PATH}

build_client: build_spec
	rm -rf ${OPENAPI_GEN_DIR}
	mkdir -p ${OPENAPI_GEN_DIR}
	${OPENAPI_CMD} generate -i ${OPENAPI_FILTERED_PATH} \
		-g go --type-mappings integer+unsigned64=uint64  -o ${OPENAPI_GEN_DIR} \
		--global-property apis,models,supportingFiles=client.go:README.md:configuration.go:response.go:utils.go,modelTests=false,apiTests=false,modelDocs=false \
		-c clientgen_utils/config.yaml
		
	cd ${OPENAPI_GEN_DIR} && goimports -w .

extract_template:
	${OPENAPI_CMD} author template -g go

check:
	terraform fmt -recursive examples/
	gofmt -s -w .
	golangci-lint run --fix --timeout 5m
	go vet

gosec:
	gosec -quiet  ./...

generate:
	terraform fmt -recursive examples/
	go generate ./...

cover:
	rm -f coverage.*
	go test -coverprofile=coverage.out ./...
	go tool cover -html coverage.out -o coverage.html

clean:
	rm -f ${BINARY}
	rm -f terraform-provider-${NAME}_*
	rm -rf ./objectscale-client

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64
