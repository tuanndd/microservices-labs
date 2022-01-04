# Mục đích
Thử nghiệm dùng etcd làm service discovery

# Hướng dẫn
```bash
# chuẩn bị
# run etcd server

go clean -modcache
go get go.etcd.io/etcd/clientv3
go get google.golang.org/grpc@v1.26.0

# test
go run discovery.go
go run register.go
```

# Tài liệu tham khảo
- https://www.pixelstech.net/article/1615108646-Service-discovery-with-etcd
