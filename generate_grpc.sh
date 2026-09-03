#!/bin/sh
set -eu

# Generates the protobuf and gRPC sources into ./pkg/a2l. Every tool it uses is pinned, so that the
# generated sources do not change without a change in this repository. The Go toolchain is the one
# the environment provides; go.mod declares the version the module is built with.
#
# The plugin versions must stay in step with the runtimes go.mod depends on: the code protoc-gen-go
# emits requires google.golang.org/protobuf of at least the same version, and the same holds for
# protoc-gen-go-grpc and google.golang.org/grpc.
PROTOC_VERSION=36.1
PROTOC_SHA256=c4bc672d9d49214dc8cafdceadf4df92182d6ca8e3ec65a56b2d7de5602669b4
PROTOC_ZIP="protoc-${PROTOC_VERSION}-linux-x86_64.zip"

PROTOC_GEN_GO_VERSION=v1.36.12
PROTOC_GEN_GO_GRPC_VERSION=v1.6.2

go install "google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOC_GEN_GO_VERSION}"
go install "google.golang.org/grpc/cmd/protoc-gen-go-grpc@${PROTOC_GEN_GO_GRPC_VERSION}"

curl -fsSLO "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}"
echo "${PROTOC_SHA256}  ${PROTOC_ZIP}" | sha256sum -c -
sudo unzip -o "${PROTOC_ZIP}" -d /usr/local bin/protoc
sudo unzip -o "${PROTOC_ZIP}" -d /usr/local 'include/*'
rm -f "${PROTOC_ZIP}"

export PATH=$PATH:$(go env GOPATH)/bin

protoc --go_out=./pkg/a2l/ --go_opt=paths=source_relative --go-grpc_out=./pkg/a2l/ --go-grpc_opt=paths=source_relative ./protobuf/*.proto
