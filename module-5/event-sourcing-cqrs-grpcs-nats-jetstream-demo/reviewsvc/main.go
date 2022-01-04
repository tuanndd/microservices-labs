package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"

	"demo/pb"
	"demo/repository"
)

const (
	channel = "order.payment.debited"

	event     = "order.approved"
	aggregate = "order"
	stream    = "order"

	grpcUri = "localhost:50051"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	sub, _ := js.PullSubscribe(channel, "reviewsvc", nats.PullMaxWaiting(128))
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
			msg.Ack() // Manual ACK
			paymentDebited := pb.OrderPaymentDebitedCommand{}
			// Unmarshal JSON that represents the Order data
			err := json.Unmarshal(msg.Data, &paymentDebited)
			if err != nil {
				log.Print(err)
				return
			}
			// Handle the message
			repository := repository.QueryStoreRepository{}
			if err := repository.ChangeOrderStatus(paymentDebited.OrderId, "Approved"); err != nil {
				log.Println(err)
				return
			}
			log.Printf("Order approved for Order ID: %s for Customer: %s\n", paymentDebited.OrderId, paymentDebited.CustomerId)
			// Publish event to Event Store
			if err := createOrderApprovedCommand(paymentDebited.OrderId); err != nil {
				log.Println("error occured while executing the OrderApproved command")
			}
		}
	}
}

// createOrderApprovedCommand calls the event store RPC to create an event
// OrderApproved command is created on Event Store
func createOrderApprovedCommand(orderId string) error {

	conn, err := grpc.Dial(grpcUri, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewEventStoreClient(conn)

	event := &pb.Event{
		EventId:       uuid.NewV4().String(),
		EventType:     event,
		AggregateId:   orderId,
		AggregateType: aggregate,
		EventData:     "",
		Stream:        stream,
	}

	resp, err := client.CreateEvent(context.Background(), event)
	if err != nil {
		return fmt.Errorf("error from RPC server: %w", err)
	}
	if resp.IsSuccess {
		return nil
	}
	return errors.New("error from RPC server")

}
