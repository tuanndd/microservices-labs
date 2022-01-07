# Mục đích
Thử nghiệm RabbitMQ

# Hướng dẫn
```bash
# chuẩn bị
docker run --rm -it -p 5672:5672 -p 15672:15672 rabbitmq:3-management

go mod tidy
```

### Tutorial 1: Hello World
![](figure-1.png)
- Gởi message có routing_key = queue_name
```bash
cd tutorial-1

go run send.go
go run receive.go
```

### Tutorial 2: Work queues
![](figure-2.png)
- Exchange = ""
- Gởi message có routing_key = queue_name
- Message được roudrobin đến các consumer

```bash
cd tutorial-2

go run worker.go
go run worker.go

go run new_task.go a..........
go run new_task.go b..........
go run new_task.go c..........
go run new_task.go d..........
```

### Tutorial 3: Publish/Subscribe
![](figure-3.png)
- tạo exchange_name = "logs", exchange_type = "fanout"
- producer send message đến exchange với routing_key=""
- consumer bind queue đến exchange, routing_Key=""
- mỗi consumer nhận được 1 bản copy của message

```bash
cd tutorial-3

go run receive_logs.go
go run receive_logs.go

go run emit_log.go
```

### Tutorial 4: Routing
![](figure-4.png)
- tao exchange có type=direct
- producer gởi message đến exhange có routing_key=A
- consumer binding đến exchange với routing_key=A
- mỗi consumer nhận 1 bản copy của message

```bash
cd tutorial-4
go run receive_logs_direct.go warning error
go run receive_logs_direct.go info warning error

go run emit_log_direct.go error "Run. Run. Or it will explode."
go run emit_log_direct.go info "App running."
```
### Tutorial 5: Topics
![](figure-5.png)
- tạo exchange có type=topic
- producer gởi message đến exchange có routing_key=a.1 hoặc a.2, ... , b.1, b.2
- consumer lần lượt bind queue đến exchange với exchange routing_key=a.* và routing_key=b.*  (routing_key="#" nghĩa là nhận tất cả message)

```bash
cd tutorial-5

go run receive_logs_topic.go "#"
go run receive_logs_topic.go "kern.*"
go run receive_logs_topic.go "*.critical"

go run emit_log_topic.go "kern.critical" "A critical kernel error"
go run emit_log_topic.go "app.critical" "A critical app error"
```

### Tutorial 6: RPC
![](figure-6.png)
- producer gởi message đến request_queue gồm content, correlation_id=rand()=x, reply_queue=A
- consumer nhận message trên request_queue, xử lý và gởi message chứa kết quả đến reply_queue voi correlation_id=x

```bash
cd tutorial-6

go run rpc_server.go
# => [x] Awaiting RPC requests

go run rpc_client.go 30
# => [x] Requesting fib(30)
```
# Tài liệu tham khảo
- https://www.rabbitmq.com/getstarted.html