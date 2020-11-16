# Implementation

# Building
## Generate ANTLR4 wrapper
- `java -jar antlr.jar -Dlanguage=Go -visitor -o pkg/a2l/parser grammar/A2L.g4`

## Generate gRPC GO code
- `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
- `go get -u google.golang.org/protobuf`
- `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/*.proto`

## Build GO shared library
- `go build --buildmode=c-shared -o a2lparser.dll ./cmd/a2l/a2l.go`

## Generate gRPC Python code
- `python -m pip install grpcio-tools`
- `python -m grpc_tools.protoc -I./protobuf/ --python_out=./script --pyi_out=./script --grpc_python_out=./script ./protobuf/A2L.proto`
