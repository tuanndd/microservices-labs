package main

import (
	pb "demo/order"
	"demo/store"
	"log"
	"runtime"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
)

const subject = "Order.>"

func main() {
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	// subscribe to subject
	natsConnection.Subscribe(subject, func(msg *nats.Msg) {
		eventStore := pb.EventStore{}
		err := proto.Unmarshal(msg.Data, &eventStore)
		if err == nil {
			log.Printf("Receive message in EventStore service: %+v\n", eventStore)
			store := store.EventStore{}
			store.CreateEvent(&eventStore)
			log.Println("Inserted event into Event store")
		}
	})

	// keep the connection alive
	runtime.Goexit()
}
