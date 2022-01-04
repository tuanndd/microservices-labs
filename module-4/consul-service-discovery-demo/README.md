# Mục đích
Thử nghiệm dùng consul làm service discovery và centralize config

# Hướng dẫn
```bash
# chuẩn bị
# run consul agent
consul agent -dev

# tạo KV để lưu config cho app
http://localhost:8500/ui/dc1/kv/create
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

# run services
go run product-service\main.go
go run user-service\main.go

# xem thông tin services
http://localhost:8500/ui/dc1/services

# test services
http://localhost:8100/product-configuration
http://localhost:8080/user-products (sẽ tìm địa chỉ của product-service từ consul và gởi request đến http://localhost:8100/products)
```

# Tài liệu tham khảo
 - https://medium.com/amartha-engineering/leveraging-consul-as-service-discovery-health-checking-and-key-value-kv-store-in-go-e475cf63ab6a