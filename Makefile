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

TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=registry.terraform.io
NAMESPACE=dell
NAME=objectscale
VERSION=1.0.0

ifeq ($(OS),Windows_NT)
    MIRROR_DIR_PREFIX = $(APPDATA)/terraform.d
    LIB_NAME = objectscale_client.dll
    BINARY = terraform-provider-${NAME}.exe
    OS_ARCH = windows_amd64
else
    OS_NAME := $(shell uname -s)
    ifeq ($(OS_NAME),Darwin)
        MIRROR_DIR_PREFIX = ~/.terraform.d
        LIB_NAME = libobjectscale_client.dylib
        BINARY = terraform-provider-${NAME}
        OS_ARCH = darwin_amd64
	else
        MIRROR_DIR_PREFIX = ~/.terraform.d
        LIB_NAME = libobjectscale_client.so
        BINARY = terraform-provider-${NAME}
        OS_ARCH = linux_amd64
    endif
endif

default: install

build:
	go mod download
	CGO_ENABLED=1 go build -o ${BINARY}

install: uninstall build
	mkdir -p ${MIRROR_DIR_PREFIX}/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ${MIRROR_DIR_PREFIX}/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp ./objectscale-client/target/release/${LIB_NAME} ${MIRROR_DIR_PREFIX}/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

uninstall:
	rm -rfv ${MIRROR_DIR_PREFIX}/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	find examples -type d -name ".terraform" -exec rm -rfv "{}" +;
	find examples -type f -name "trace.*" -delete
	find examples -type f -name "*.tfstate" -delete
	find examples -type f -name "*.hcl" -delete
	find examples -type f -name "*.backup" -delete
	rm -rf trace.*

client-checkout:
	git clone -b main https://github.com/dell/objectscale-client.git

client-build: clean client-checkout
	cd ./objectscale-client/c && cargo build --release

test: check
	go test -i $(TEST) || exit 1       
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

check:
	terraform fmt -recursive examples/
	gofmt -s -w .
	golangci-lint run --fix --timeout 5m
	go vet

gosec:
	gosec -quiet -log gosec.log -out=gosecresults.csv -fmt=csv -exclude=G104 ./...

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

generate:
	go generate ./...

cover:
	rm -f coverage.*
	go test -coverprofile=coverage.out ./...
	go tool cover -html coverage.out -o coverage.html

clean:
	rm -f ${BINARY}
	rm -f terraform-provider-${NAME}_*
	rm -rf ./objectscale-client

build-linux-amd64:
	docker run --rm -it -v .:/io -w /io ghcr.io/rust-cross/rust-musl-cross:x86_64-musl /bin/bash -c "cd objectscale-client/c && cargo rustc --crate-type=staticlib --release; cd /tmp && curl -LO https://go.dev/dl/go1.24.7.linux-amd64.tar.gz && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.7.linux-amd64.tar.gz; cd /io && CC='x86_64-unknown-linux-musl-gcc' GOFLAGS='-buildvcs=false' CGO_LDFLAGS='-L/io/objectscale-client/target/x86_64-unknown-linux-musl/release/' /usr/local/go/bin/go build -ldflags=\"-linkmode external -extldflags '-static'\" -o ${BINARY}"

build-linux-arm64:
	docker run --rm -it -v .:/io -w /io ghcr.io/rust-cross/rust-musl-cross:aarch64-musl /bin/bash -c "cd objectscale-client/c && cargo rustc --crate-type=staticlib --release; cd /tmp && curl -LO https://go.dev/dl/go1.24.7.linux-amd64.tar.gz && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.7.linux-amd64.tar.gz; cd /io && CC='aarch64-unknown-linux-musl-gcc' GOFLAGS='-buildvcs=false' GOARCH=arm64 CGO_ENABLED=1 CGO_LDFLAGS='-L/io/objectscale-client/target/aarch64-unknown-linux-musl/release/' /usr/local/go/bin/go build -ldflags=\"-linkmode external -extldflags '-static'\" -o ${BINARY}"

build-linux-arm:
	docker run --rm -it -v .:/io -w /io ghcr.io/rust-cross/rust-musl-cross:arm-musleabi /bin/bash -c "cd objectscale-client/c && cargo rustc --crate-type=staticlib --release; cd /tmp && curl -LO https://go.dev/dl/go1.24.7.linux-amd64.tar.gz && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.7.linux-amd64.tar.gz; cd /io && CC='arm-unknown-linux-musleabi-gcc' GOFLAGS='-buildvcs=false' GOARCH=arm CGO_ENABLED=1 CGO_LDFLAGS='-L/io/objectscale-client/target/arm-unknown-linux-musleabi/release/' /usr/local/go/bin/go build -ldflags=\"-linkmode external -extldflags '-static'\" -o ${BINARY}"

release:
	cp terraform-provider-objectscale terraform-provider-${NAME}_v${VERSION}
	zip -j terraform-provider-${NAME}_${VERSION}_${OS_ARCH}.zip terraform-provider-${NAME}_v${VERSION}
	cp terraform-registry-manifest.json terraform-provider-${NAME}_${VERSION}_manifest.json
	shasum -a 256 *.zip terraform-provider-${NAME}_${VERSION}_manifest.json > terraform-provider-${NAME}_${VERSION}_SHA256SUMS
