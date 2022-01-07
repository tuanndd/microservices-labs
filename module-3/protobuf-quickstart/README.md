# Mục đích
Thử nghiệm protobuf

# Hướng dẫn 
```bash
# install protoc
https://github.com/protocolbuffers/protobuf/releases/tag/v3.19.1

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go mod tidy

# compile proto
protoc --go_out=. --go_opt=paths=source_relative proto/addressbook.proto

# test tạo address book
go run cmd/add_person.go file.dat

# test xem address book
go run cmd/list_people.go file.dat

less file.dat
```

# Tài liệu tham khảo:
- https://developers.google.com/protocol-buffers/docs/gotutorial