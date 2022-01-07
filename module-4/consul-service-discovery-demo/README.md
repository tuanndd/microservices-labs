# Mục đích
Thử nghiệm dùng consul làm service discovery và centralize config

# Hướng dẫn
```bash
# chuẩn bị
# run consul agent
consul agent -dev -client 0.0.0.0

# tạo KV để lưu config cho app
http://vagrant-ip:8500/ui/dc1/kv/create

key=product-configuration

value=
{
    "categories": [
        "one",
        "two",
        "three",
        "four"
    ]
}

go mod tidy

# run services
go run product-service/main.go
go run user-service/main.go

# xem thông tin services
http://vagrant-ip:8500/ui/dc1/services

# test services
http://vagrant-ip:8100/product-configuration
http://vagrant-ip:8080/user-products (sẽ tìm địa chỉ của product-service từ consul và gởi request đến http://vagrant-ip:8100/products)
```

# Tài liệu tham khảo
 - https://medium.com/amartha-engineering/leveraging-consul-as-service-discovery-health-checking-and-key-value-kv-store-in-go-e475cf63ab6a