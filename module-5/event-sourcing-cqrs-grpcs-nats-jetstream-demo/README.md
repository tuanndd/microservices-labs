# Mục đích
Thử nghiệm event sourcing, cqrs

![](figure-1.png)

### Code-flow để tạo order mới:
| Step | Service | Subscribe vào subject | Xử lý nội bộ | Push message đến subject |
| - | - | - | - | - |
| 1 | ordersvc | | nhận http request từ client | order.created |
| 2 | stream-processor-* | order.created | transaction{sql_insert_order()} | |
| 2 | paymentsvc | order.created | | order.payment.debited |
| 3 | reviewsvc | order.payment.debited | sql_update_order(status="Approved") | order.approved  

Hàm CreateEvent() của eventstore grpc server sẽ insert_db(event) -> push_msg(subject=event.eventType)  
# Hướng dẫn
### Chuẩn bị
```bash
## start server nats jetstream
# edit js.conf
nats-server -c jetstream.conf -V

## start CockroachDB cluster
cockroach start --insecure --store=ordersdb1 --listen-addr=localhost:26257 --http-addr=localhost:8080 --join=localhost:26257,localhost:26258,localhost:26259

cockroach start --insecure --store=ordersdb2 --listen-addr=localhost:26258 --http-addr=localhost:8081 --join=localhost:26257,localhost:26258,localhost:26259

cockroach start --insecure --store=ordersdb3 --listen-addr=localhost:26259 --http-addr=localhost:8082 --join=localhost:26257,localhost:26258,localhost:26259

cockroach init --insecure --host=localhost:26257

# tạo db
cockroach sql --insecure --host=localhost:26257
$> CREATE DATABASE ordersdb;

## compile proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/*.proto

go mod tidy
```

### Run app
```bash
# run services
go run eventstore\eventstore.go
go run ordersvc\main.go

go run stream-processor-1\main.go
go run stream-processor-2\main.go

go run reviewsvc\main.go

# test
curl -X POST -H "Content-Type: application/json" -d @order.json http://localhost:3000/api/orders
```

### Monitor
#### CockroachDB
http://localhost:8080

#### NATS
```bash
# stream info
nats stream list
nats stream info order

# stream message
nats stream view order

# consumer info
nats consumer list order
nats consumer info order paymentsvc
```

# Tài liệu tham khảo
- https://shijuvar.medium.com/building-microservices-with-event-sourcing-cqrs-in-go-using-grpc-nats-streaming-and-cockroachdb-983f650452aa

