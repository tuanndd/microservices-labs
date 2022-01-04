package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"

	"demo/pb"
)

const (
	event     = "order.created"
	aggregate = "order"
	grpcUri   = "localhost:50051"
)

// rpc interface
type rpcClient interface {
	createOrder(order pb.OrderCreateCommand) error
}
type grpcClient struct {
}

// createOrder calls the CreateEvent RPC
func (gc grpcClient) createOrder(order pb.OrderCreateCommand) error {
	conn, err := grpc.Dial(grpcUri, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEventStoreClient(conn)
	orderJSON, _ := json.Marshal(order)

	event := &pb.Event{
		EventId:       uuid.NewV4().String(),
		EventType:     event,
		AggregateId:   order.OrderId,
		AggregateType: aggregate,
		EventData:     string(orderJSON),
		Stream:        "order",
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

type orderHandler struct {
	rpc rpcClient
}

// web handler> tao order
func (h orderHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	var order pb.OrderCreateCommand
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid Order Data", 500)
		return
	}

	log.Printf("---------- INPUT: r.body: %+v\n", r.Body)
	log.Printf("---------- OUTPUT: order: %+v\n", order)

	aggregateID := uuid.NewV4().String()
	order.OrderId = aggregateID
	order.Status = "Pending" // trang thai
	order.CreatedOn = time.Now().Unix()

	err = h.rpc.createOrder(order)
	if err != nil {
		log.Print(err)
		http.Error(w, "Failed to create Order", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	j, _ := json.Marshal(order)
	w.Write(j)
}

// Tao routes
func initRoutes() *mux.Router {
	router := mux.NewRouter()
	h := orderHandler{
		rpc: grpcClient{},
	}
	router.HandleFunc("/api/orders", h.createOrder).Methods("POST")
	return router
}

func main() {
	// Create the Server
	server := &http.Server{
		Addr:    ":3000",
		Handler: initRoutes(),
	}
	log.Println("HTTP Sever listening...")
	// Running the HTTP Server
	server.ListenAndServe()
}
