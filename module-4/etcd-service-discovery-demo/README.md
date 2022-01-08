# Mục đích
Thử nghiệm dùng etcd làm service discovery

| Role | Logic code |
| - | - |
| register | tạo K="/web/node1", V="localhost:8000" kết hợp với extend lease định kỳ để giữ KV tồn tại |
| discovery | watch KV /web/* , /gRPC/* để phát hiện service up/down |
# Hướng dẫn
```bash
# chuẩn bị
# run etcd server
etcd


go mod tidy

# test
go run discovery.go
go run register.go
```

# Tài liệu tham khảo
- https://www.pixelstech.net/article/1615108646-Service-discovery-with-etcd
