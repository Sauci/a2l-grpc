java -jar antlr.jar -Dlanguage=Go -visitor -o pkg/a2l/parser grammar/A2L.g4
protoc --go_out=./pkg/a2l/ --go_opt=paths=source_relative --go-grpc_out=./pkg/a2l/ --go-grpc_opt=paths=source_relative ./protobufs/*.proto
python -m grpc_tools.protoc -I./ --python_out=./../lis/a2l-api/ --pyi_out=./../lis/a2l-api/ --grpc_python_out=./../lis/a2l-api/ ./protobufs/*.proto
python -m grpc_tools.protoc -I./ --python_out=./../lis/liba2l/ --pyi_out=./../lis/liba2l/ --grpc_python_out=./../lis/liba2l/ ./protobufs/*.proto
go build --buildmode=c-shared -o a2lparser.dll ./cmd/a2l/a2l.go