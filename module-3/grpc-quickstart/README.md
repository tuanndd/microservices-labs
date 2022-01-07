# Mục đích
Thử nghiệm GRPC

# Hướng dẫn
```bash
# install modules
go install google.golang.org/protobuf/cmd/protoc-gen-go
#go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

go mod tidy
go get google.golang.org/protobuf

# compile proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/helloworld.proto

# run test
go run server.go
go run client.go
```
# Tài liệu tham khảo
- https://grpc.io/docs/languages/go/quickstart/