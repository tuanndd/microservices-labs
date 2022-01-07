# Mục đích
Thử nghiệm build REST API

# Hướng dẫn
```bash
go mod tidy

# start server
go run main.go

# test
curl http://vagrant-ip:8080/books
curl -X POST -d @create.json -H "Content-Type: application/json" http://vagrant-ip:8080/books
curl http://vagrant-ip:8080/books/1
curl -X PATCH -d @update.json -H "Content-Type: application/json" http://vagrant-ip:8080/books/1
curl -X DELETE http://vagrant-ip:8080/books/1
curl http://vagrant-ip:8080/books/1
```

# Tài liệu tham khảo
- https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
