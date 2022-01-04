package main

import (
	"log"
	gintrace "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin"
	option "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/api/global"
	"github.com/gin-gonic/gin"
	"jaeger-tracing-go-service/config"
	"jaeger-tracing-go-service/routes"
	"jaeger-tracing-go-service/tracer"
)

var tr = global.Tracer("jaeger-tracing-go-service")

func main() {
	fn := tracer.InitJaeger()
	defer fn()

	// Set client options
	config.Connect()

	// Init Router
	router := gin.Default()
	router.Use(gintrace.Middleware("jaeger-tracing-go-service", option.WithTracer(tr)))

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":8091"))
}

