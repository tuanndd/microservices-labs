// gRPC API for Event Store
package main

import (
	"context"
	"log"
	"net"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"

	"demo/pb"
	"demo/repository"
)

const (
	port           = ":50051"
	streamName     = "order"
	streamSubjects = "order.>"
)

type server struct {
	// NAT stream context
	pb.UnimplementedEventStoreServer
	nats.JetStreamContext
}

// CreateEvent RPC creates a new Event into EventStoreRepository
// and publish an event to NATS Streaming
func (s *server) CreateEvent(ctx context.Context, in *pb.Event) (*pb.Response, error) {
	// Luu event vao DB
	// Persist data into EventStoreRepository database
	command := repository.EventStoreRepository{}
	// Persist events as immutable logs into CockroachDB
	err := command.CreateEvent(in)
	if err != nil {
		return nil, err
	}
	// Publish event on NATS Streaming Server
	go publishEvent(s.JetStreamContext, in)
	return &pb.Response{IsSuccess: true}, nil
}

// GetEvents RPC gets events from EventStoreRepository by given AggregateId
func (s *server) GetEvents(ctx context.Context, in *pb.EventFilter) (*pb.EventResponse, error) {
	eventStore := repository.EventStoreRepository{}
	events := eventStore.GetEvents(in)
	return &pb.EventResponse{Events: events}, nil
}

// publishEvent publishes an event via NATS Streaming server
func publishEvent(js nats.JetStreamContext, event *pb.Event) {
	channel := event.EventType
	eventMsg := []byte(event.EventData)
	// Publish message on subject (channel)
	js.Publish(channel, eventMsg)
	log.Println("Published message on channel: " + channel)
}

func createStream(js nats.JetStreamContext) error {
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Println(err)
	}

	if stream == nil {
		log.Printf("creating stream %q and subject %q", streamName, streamSubjects)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,               // stream name
			Subjects: []string{streamSubjects}, // subjects
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	js, _ := nc.JetStream()
	err = createStream(js)
	if err != nil {
		log.Fatal(err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()

	pb.RegisterEventStoreServer(s, &server{JetStreamContext: js})
	s.Serve(lis)
}
