package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"

	"demo/pb"
)

const (
	subscribeChannel = "order.created"

	event     = "order.payment.debited"
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

	sub, _ := js.PullSubscribe(subscribeChannel, "paymentsvc")

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
			log.Println("----------  before ACK")
			msg.Ack()

			order := pb.OrderCreateCommand{}
			// Unmarshal JSON that represents the Order data
			err := json.Unmarshal(msg.Data, &order)
			if err != nil {
				log.Print(err)
				return
			}
			// Create OrderPaymentDebitedCommand from Order
			command := pb.OrderPaymentDebitedCommand{
				OrderId:    order.OrderId,
				CustomerId: order.CustomerId,
				Amount:     order.Amount,
			}
			log.Println("Payment has been debited from customer account for Order:", order.OrderId)
			if err := createPaymentDebitedCommand(command); err != nil {
				log.Println("error occured while executing the PaymentDebited command")
			}
		}
	}
}

// createPaymentDebitedCommand calls the event store RPC to create an event
// PaymentDebited command is created on Event Store
func createPaymentDebitedCommand(command pb.OrderPaymentDebitedCommand) error {

	conn, err := grpc.Dial(grpcUri, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewEventStoreClient(conn)
	paymentJSON, _ := json.Marshal(command)

	event := &pb.Event{
		EventId:       uuid.NewV4().String(),
		EventType:     event,
		AggregateId:   command.OrderId,
		AggregateType: aggregate,
		EventData:     string(paymentJSON),
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
