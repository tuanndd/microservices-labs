package store

import (
	pb "demo/order"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// Event provides CRUD operations against the collection "events"
type EventStore struct {
}

// CreateEvent inserts the value of struct EventStore into collection.
func (store EventStore) CreateEvent(event *pb.EventStore) error {
	col := mgoDb.Collection("events")

	_, err := col.InsertOne(mgoCtx, event)

	return err
}

// GetEvents returns all documents from the collection.
func (store EventStore) GetEvents() []pb.EventStore {
	var events []pb.EventStore

	col := mgoDb.Collection("events")
	cur, _ := col.Find(mgoCtx, bson.D{})
	defer cur.Close(mgoCtx)

	for cur.Next(mgoCtx) {
		data := &pb.EventStore{}

		if err := cur.Decode(data); err != nil {
			log.Fatalf("Error decode: %v", err)
		}

		log.Printf("GetEvents(i): %v ---------- ", data)

		events = append(events, *data)
	}

	return events
}
