# Mục đích
Thử nghiệm event sourcing, cqrs

![](figure-1.png)

### Code-flow để tạo order mới:
| Step | Service | Subscribe vào subject | Xử lý nội bộ | Push message đến subject |
| - | - | - | - | - |
| 1 | ordersvc | | nhận request từ client | order.created |
| 2 | stream-processor-* | order.created | transaction{sql_insert_order()} | |
| 2 | paymentsvc | order.created | | order.payment.debited |
| 3 | reviewsvc | order.payment.debited | sql_update_order(status="Approved") | order.approved  

Hàm CreateEvent() của eventstore grpc server sẽ insert_db(event) -> push_msg(subject=event.eventType)  
# Hướng dẫn
### Chuẩn bị
```bash
## start NATS JetStream
docker run --rm -it -p 4222:4222 nats -js

# start CockroachDB cluster
docker network create -d bridge roachnet

docker run -d \
--name=roach1 \
--hostname=roach1 \
--net=roachnet \
-p 26257:26257 -p 8080:8080  \
-v "${PWD}/cockroach-data/roach1:/cockroach/cockroach-data"  \
cockroachdb/cockroach:v21.2.3 start \
--insecure \
--join=roach1,roach2,roach3

docker run -d \
--name=roach2 \
--hostname=roach2 \
--net=roachnet \
-v "${PWD}/cockroach-data/roach2:/cockroach/cockroach-data" \
cockroachdb/cockroach:v21.2.3 start \
--insecure \
--join=roach1,roach2,roach3

docker run -d \
--name=roach3 \
--hostname=roach3 \
--net=roachnet \
-v "${PWD}/cockroach-data/roach3:/cockroach/cockroach-data" \
cockroachdb/cockroach:v21.2.3 start \
--insecure \
--join=roach1,roach2,roach3

docker exec -it roach1 ./cockroach init --insecure

# tạo database
docker exec -it roach1 ./cockroach sql --insecure
> CREATE DATABASE ordersdb;


## compile proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/*.proto

go mod tidy
```

### Run app
```bash
# run services
go run eventstore/main.go

go run ordersvc/main.go

go run stream-processor-1/main.go
go run stream-processor-2/main.go

go run paymentsvc/main.go

go run reviewsvc/main.go

# test
curl -X POST -H "Content-Type: application/json" -d @order.json http://localhost:3000/api/orders

# stop CockroachDB cluster 
docker stop roach1 roach2 roach3
docker rm roach1 roach2 roach3
sudo rm -rf cockroach-data

```

### Monitor
#### CockroachDB
http://vagrant-ip:8080

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

