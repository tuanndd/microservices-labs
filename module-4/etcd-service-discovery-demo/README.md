# Mục đích
Thử nghiệm dùng etcd làm service discovery

# Hướng dẫn
```bash
# chuẩn bị
# run etcd server
minikube stop
etcd


go mod tidy

# test
go run discovery.go
go run register.go
```

# Tài liệu tham khảo
- https://www.pixelstech.net/article/1615108646-Service-discovery-with-etcd
