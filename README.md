# Giới thiệu
Khóa học **Building Distributed Applications with Microservices** được xây dựng nhằm mục đích giúp người học tiết kiệm thời gian tìm hiểu, xây dựng được ứng dụng phân tán nhanh chóng nhưng vẫn đảm bảo tính ***dễ mở rộng, khả năng chịu lỗi và chi phí tối ưu***.

Khóa học được xây dựng dựa trên series bài https://www.nginx.com/blog/introduction-to-microservices/ kèm theo các bài lab minh họa có thể tham khảo để viết ứng dụng thực tế.

## Điều kiện tham gia khóa học
- Bạn có thể lập trình trên máy tính( prefer biết dùng Visual Code và cơ bản về Python, Golang)
- Bạn biết về Linux (run commands in Terminal)

# Nội dung chi tiết
### Module 1: [Introduction to Microservices](module-1/Introduction-to-Microservices.pptx)
- So sánh giữa Monolithic và Microservices
- Lab: [setup Centos box bằng Vagrant](module-1/vagrant-quickstart/)

### Module 2: [Building Microservices: Using an API Gateway](module-2/Building-Microservices-Using-an-API-Gateway.pptx)
- Hướng dẫn dùng API gateway để kết nối và bảo mật microservices
- Lab:
  - [Krakend playground](module-2/krakend-playground/)

### Module 3: [Building Microservices: Inter-Process Communication in a Microservices Architecture](module-3/Building-Microservices-Inter-Process-Communication-in-a-Microservices-Architecture.pptx)
- Giải thích các cơ chế IPC để kết nối các microservices: synchronous, asynchronous
- Phân biệt các khái niệm trong asynchronous: concurrency, parallel, blocking, non-blocking mode, event-loop
- Giải thích cơ chế messaging của RabbitMQ, MQTT, NATS, so sánh với Kafka.
- Lab:
  - Synchronous RPC:
     - [Protobuf](module-3/protobuf-quickstart/), [GRPC](module-3/grpc-quickstart/) (golang)
     - [REST api](module-3/restapi-quickstart/), [OpenAPI, Swagger](module-3/openapi-swagger-quickstart/) (golang)
  - Asynchronous RPC:
    - [rabbitmq quickstart](module-3/rabbitmq-quickstart/) (golang)
    - [grpc over nats](module-3/grpc-over-nats/) (golang)
    - [olso.messageing demo](module-3/oslo-messaging-demo/) (tìm hiểu core messaging của Openstack, python)
    - [Nameko quickstart](module-3/nameko-quickstart/) (microservices framework, python)
    - [MQTT IoT demo](module-3/mqtt-iot-demo/) (cách thiết kế MQTT topic trong IoT, golang)

### Module 4: [Service Discovery in a Microservices Architecture](module-4/Service-Discovery-in-a-Microservices-Architecture.pptx)
- Giải thích các cơ chế client-side, server-side discovery
- Lab:
  - [Dùng etcd cho service discovery](module-4/etcd-service-discovery-demo/) (golang)
  - [Dùng consul cho service discovery, centralized configuration](module-4/consul-service-discovery-demo/) (consul api, golang)
  - [viper demo](module-4/viper-demo/) (golang)

### Module 5: [Event-Driven Data Management for Microservices](module-5/Event-Driven-Data-Management-for-Microservices.pptx)
- Giải bài toán business business transaction trên môi trường phân tán microservices
- Lab:
  - [CQRS, event-sourcing over NATS JetStream](module-5/event-sourcing-cqrs-grpcs-nats-jetstream-demo/) (golang)

### Module 6: [Choosing a Microservices Deployment Strategy](module-6/Choosing-a-Microservices-Deployment-Strategy.pptx)
- Giới thiệu OpenTelemetry framwork để giám sát microservices
- Tìm hiểu các lựa chọn để triển khai microservices: physical server, VM, container, serverless
- Lab:
  - [Opentelemetry, Jaeger distributed tracing demo](module-6/opentelemetry-jaeger-distributed-tracing-demo/)
  - [Deploy wordpress bằng docker-compose và kubernetes](module-6/wordpress-docker-compose-kubernetes/)
  
### Module 7: [Refactoring a Monolith into Microservices](module-7/Refactoring-a-Monolith-into-Microservices.pptx)
- Tìm hiểu các chiến lược để migrate monolith sang microservices: ***Stop Digging, Split Frontend and Backend, Extract Services***

# Các vấn đề khác
- Khóa học có sử dụng tài liệu và source code của nhiều người, tất cả đều được ghi nhận trong các file README.md
- Bạn được phép sử dụng miễn phí tất cả các tài liệu của khóa học này. 
- Mọi thắc mắc vui lòng liên hệ tôi qua email tuanndd@gmail.com