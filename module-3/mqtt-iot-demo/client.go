package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	mqttUrl = "tcp://localhost:1883"
)

var clientId string

type Message struct {
	topic   string
	payload interface{}
	retain  bool
	qos     byte
}

type Command struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

var messagePubHandler mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
	fmt.Printf("Message %s received on topic %s\n", m.Payload(), m.Topic())

	topic := m.Topic()

	if topic == "home/$broadcast/alert" {
		sendMessage(c, Message{
			topic:   fmt.Sprintf("home/%s/ret", clientId),
			payload: fmt.Sprintf("%s-WARNING-%d", clientId, rand.Int()),
		})
	} else if strings.HasPrefix(topic, fmt.Sprintf("home/%s/", clientId)) {
		parts := strings.Split(topic, "/")

		code := parts[2]

		switch code {
		case "cmd":
			command := Command{}
			err := json.Unmarshal(m.Payload(), &command)
			if err != nil {
				log.Printf(err.Error())
			}

			sendMessage(c, Message{
				topic:   fmt.Sprintf("home/%s/ret", clientId),
				payload: fmt.Sprintf("%s-REPLY-%d-CMD-%s", clientId, rand.Int(), command.Id),
			})
		}
	}
}

var connectHandler mqtt.OnConnectHandler = func(c mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(c mqtt.Client, e error) {
	fmt.Printf("Connection Lost: %s\n", e.Error())
}

func subscribeTopic(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()

	fmt.Printf("subscribed to topic %s\n", topic)
}

func sendMessage(client mqtt.Client, message Message) {
	token := client.Publish(message.topic, message.qos, message.retain, message.payload)
	token.Wait()
}

func main() {
	clientId = os.Args[1]

	options := mqtt.NewClientOptions()

	options.AddBroker(mqttUrl)
	options.SetClientID(clientId)

	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()

	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	subscribeTopic(client, fmt.Sprintf("home/%s/cmd", clientId))
	subscribeTopic(client, "home/$broadcast/#")

	sendMessage(client,
		Message{
			topic:   fmt.Sprintf("home/%s/info", clientId),
			payload: fmt.Sprintf("%s-INFO-%d", clientId, rand.Int()),
			retain:  true,
		})

	for {
		sendMessage(client, Message{
			topic:   fmt.Sprintf("home/%s/stat", clientId),
			payload: fmt.Sprintf("%s-STAT-%d", clientId, rand.Int()),
		})
		time.Sleep(time.Second * 5)
	}
}
