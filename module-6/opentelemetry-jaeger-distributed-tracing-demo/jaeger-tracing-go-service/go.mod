module jaeger-tracing-go-service

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.1
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.4.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin v0.11.0
	go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver v0.11.0
	go.opentelemetry.io/otel v0.11.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.11.0
	go.opentelemetry.io/otel/sdk v0.11.0
)
