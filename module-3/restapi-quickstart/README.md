# Mục đích
Thử nghiệm build REST API

# Hướng dẫn
```bash
go get ...

# start server
go run main.go

# test
curl http://localhost:8080/books
curl -X POST -d @create.json -H "Content-Type: application/json" http://localhost:8080/books
curl http://localhost:8080/books/1
curl -X PATCH -d @update.json -H "Content-Type: application/json" http://localhost:8080/books/1
curl -X DELETE http://localhost:8080/books/1
curl http://localhost:8080/books/1
```

# Tài liệu tham khảo
- https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
