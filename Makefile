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

clean:
	sudo rm -f ${BINARY}
	sudo rm -f terraform-provider-${NAME}_*
	sudo rm -rf ./objectscale-client

docker-build-linux: client-checkout
	docker run --rm -it -v ./objectscale-client:/io -w /io/c ghcr.io/rust-cross/rust-musl-cross:x86_64-musl cargo rustc --crate-type=staticlib --release
	docker run --rm -it -v .:/src -w /src -e CC="gcc" -e CGO_LDFLAGS="-L/src/objectscale-client/target/x86_64-unknown-linux-musl/release/" golang:1.23-alpine sh -c "apk add --no-cache musl-dev build-base && go build -ldflags=\"-linkmode external -extldflags '-static'\" -o ${BINARY}"

release: clean docker-build-linux
	cp terraform-provider-objectscale terraform-provider-${NAME}_v${VERSION}
	zip -j terraform-provider-${NAME}_${VERSION}_${OS_ARCH}.zip terraform-provider-${NAME}_v${VERSION}
	cp terraform-registry-manifest.json terraform-provider-${NAME}_${VERSION}_manifest.json
	shasum -a 256 *.zip terraform-provider-${NAME}_${VERSION}_manifest.json > terraform-provider-${NAME}_${VERSION}_SHA256SUMS
