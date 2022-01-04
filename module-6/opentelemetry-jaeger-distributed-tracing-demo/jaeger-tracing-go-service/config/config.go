package config

import (
	"context"
	"os"
	"time"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"jaeger-tracing-go-service/controllers"
	mongotrace "go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver"
)

// func Connect(tracer trace.Tracer) {
func Connect() {
	// Database Config
	// clientOptions := options.Client().ApplyURI("mongodb://admin:admin@localhost:27017/admin?connect=direct")
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	clientOptions.Monitor = mongotrace.NewMonitor("jaeger-tracing-go-service")
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	db := client.Database(os.Getenv("DATABASE_NAME"))
	// db := client.Database("employees")
	controllers.EmployeesCollection(db)
	return
}


