package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
)

const (
	mqttUrl  = "tcp://localhost:1883"
	serverId = "SERVER"
)

type Device struct {
	Info string `json:"info"`
	Stat string `json:"stat"`
	Ret  string `json:"ret"`
}

var DEVICES = make(map[string]Device)

var messagePubHandler mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
	fmt.Printf("Message %s received on topic %s\n", m.Payload(), m.Topic())

	topic := m.Topic()

	if strings.HasPrefix(topic, "home/") {
		parts := strings.Split(topic, "/")

		payload := m.Payload()

		device := parts[1]
		code := parts[2]

		val, ok := DEVICES[device]
		if ok == false {
			val = Device{}
		}

		switch code {
		case "info":
			val.Info = string(payload)
		case "stat":
			val.Stat = string(payload)
		case "ret":
			val.Ret = string(payload)

		}

		DEVICES[device] = val
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

func main() {
	options := mqtt.NewClientOptions()

	options.AddBroker(mqttUrl)
	options.SetClientID(serverId)

	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()

	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	subscribeTopic(client, "home/+/info")
	subscribeTopic(client, "home/+/stat")
	subscribeTopic(client, "home/+/ret")

	r := gin.Default()

	r.GET("/info/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(200, gin.H{
			"result": DEVICES[id].Info,
		})
	})

	r.GET("/stat/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(200, gin.H{
			"result": DEVICES[id].Stat,
		})
	})

	r.GET("/ret/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(200, gin.H{
			"result": DEVICES[id].Ret,
		})
	})

	r.POST("/broadcast", func(c *gin.Context) {
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			// Handle error
		}

		client.Publish("home/$broadcast/alert", 0, false, data)
	})

	r.POST("/cmd/:id", func(c *gin.Context) {
		id := c.Param("id")

		topic := fmt.Sprintf("home/%s/cmd", id)

		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			// Handle error
		}

		client.Publish(topic, 0, false, data)

		c.JSON(200, gin.H{
			"result": "success",
		})
	})

	r.GET("/devices", func(c *gin.Context) {

		c.JSON(200, DEVICES)
	})

	r.Run()

	select {}
}
