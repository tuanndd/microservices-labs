package store

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoDb *mongo.Database
var mgoClient *mongo.Client
var mgoCtx context.Context

func init() {
	mgoCtx = context.Background()

	var err error
	// Connect den server MongoDB
	mgoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost"))

	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}

	mgoDb = mgoClient.Database("local")
}
