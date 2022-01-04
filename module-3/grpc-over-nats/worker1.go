package main

import (
	pb "demo/order"
	"log"
	"runtime"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
)

const (
	queue   = "Order.OrdersCreatedQueue"
	subject = "Order.OrderCreated"
)

func main() {
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	// subscribe to subject
	natsConnection.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		eventStore := pb.EventStore{}
		err := proto.Unmarshal(msg.Data, &eventStore)
		if err == nil {
			log.Printf("Subscribed message in Worker 1: %+v\n", eventStore)
		}
	})

	// keep the connection alive
	runtime.Goexit()
}
