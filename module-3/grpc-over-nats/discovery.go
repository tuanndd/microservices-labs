package main

import (
	"log"
	"runtime"

	pb "demo/order"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
)

var orderServiceUri string

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Config file not found")
	}

	orderServiceUri = viper.GetString("discovery.orderservice")
}

func main() {
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	// log.Printf("CONNECT> %v ---------- %v", natsConnection, err)

	natsConnection.Subscribe("Discovery.OrderService", func(msg *nats.Msg) {
		orderServiceDiscovery := pb.ServiceDiscovery{
			OrderServiceUri: orderServiceUri,
		}
		data, err := proto.Marshal(&orderServiceDiscovery)
		// log.Printf("MSG: %v ---------- %v", data, err)

		if err == nil {
			natsConnection.Publish(msg.Reply, data)
		}
	})

	runtime.Goexit()
}
