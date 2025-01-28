wget https://golang.org/dl/go1.22.11.linux-amd64.tar.gz
sudo tar -xf go1.22.11.linux-amd64.tar.gz -C /usr/local

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v29.3/protoc-29.3-linux-x86_64.zip
sudo unzip -o protoc-29.3-linux-x86_64.zip -d /usr/local bin/protoc
sudo unzip -o protoc-29.3-linux-x86_64.zip -d /usr/local 'include/*'
rm -f protoc-29.3-linux-x86_64.zip

export PATH=$PATH:$HOME/go/bin

protoc --go_out=./pkg/a2l/ --go_opt=paths=source_relative --go-grpc_out=./pkg/a2l/ --go-grpc_opt=paths=source_relative ./protobuf/*.proto
