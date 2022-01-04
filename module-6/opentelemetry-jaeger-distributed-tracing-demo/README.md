# Mục đích
Thử nghiệm distributed tracing bằng opentelemetry và jaeger

![](figure-1.png)

# Hướng dẫn build & run
```bash
# cài đặt java-11 sdk

# build java app
cd jaeger-tracing-java-service
./gradlew build

# edit jaeger-tracing-frontend-service/src/components/*
# thay thế NODEJS_SERVICE_URL: http://[your-server-ip-address]:8081/ 

docker compose build
docker compose up
```

# Tài liệu tham khảo
- https://capgemini.github.io/development/Distributed-Tracing-with-OpenTelemetry-And-Jaeger/
