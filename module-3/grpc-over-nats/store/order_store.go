package store

import (
	pb "demo/order"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// OrderStore provides CRUD operations against the collection "orders"
type OrderStore struct {
}

// CreateOrder inserts the value of struct Order into collection.
func (store OrderStore) CreateOrder(order *pb.Order) error {
	col := mgoDb.Collection("orders")

	_, err := col.InsertOne(mgoCtx, order)

	return err
}

// GetOrders returns all documents from the collection.
func (store OrderStore) GetOrders() []pb.Order {
	var orders []pb.Order

	col := mgoDb.Collection("orders")
	cur, _ := col.Find(mgoCtx, bson.D{})
	defer cur.Close(mgoCtx)

	for cur.Next(mgoCtx) {
		data := &pb.Order{}

		if err := cur.Decode(data); err != nil {
			log.Fatalf("Error decode: %v", err)
		}

		log.Printf("GetOrders(i): %v ---------- ", data)

		orders = append(orders, *data)
	}

	return orders
}
