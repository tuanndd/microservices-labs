package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"

	"demo/pb"
	"demo/repository"
)

const (
	clientID   = "order-query-store1"
	channel    = "order.created"
	queueGroup = "order-query-store-group"
)

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	sub, _ := js.PullSubscribe(channel, queueGroup, nats.PullMaxWaiting(128))
	ctx, cancel := context.WithTimeout(context.Background(), 3600*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		msgs, _ := sub.Fetch(1, nats.Context(ctx)) // batch_size = 1 messages
		for _, msg := range msgs {
			msg.Ack()
			order := pb.OrderCreateCommand{}
			err := json.Unmarshal(msg.Data, &order)
			if err == nil {
				// Handle the message
				log.Printf("Subscribed message from clientID - %s: %+v\n", clientID, order)
				queryRepository := repository.QueryStoreRepository{}
				// Perform data replication for query model into CockroachDB
				err := queryRepository.SyncOrderQueryModel(order)
				if err != nil {
					log.Printf("Error while replicating the query model %+v", err)
				}
			}
		}
	}
}
