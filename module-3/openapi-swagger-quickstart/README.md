# Mục đích
- Thử nghiệm dùng Swagger chuyển REST API thành OpenAPI


# Hướng dẫn

```bash
# install modules
go mod tidy
go get -v github.com/swaggo/swag/cmd/swag

# generate swagger's docs
swag init

# run server
go run main.go

# test api
http://vagrant-ip:8080/swagger/index.html

```
# Tài liệu tham khảo
- https://golangexample.com/automatically-generate-restful-api-documentation-with-swagger-2-0-for-go/