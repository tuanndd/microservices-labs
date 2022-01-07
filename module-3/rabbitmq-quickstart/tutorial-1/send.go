package main

import (
	"log"

	"github.com/streadway/amqp"

	"demo/config"
)

// print loi
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial(config.RABBITMQ_URL)
	failOnError(err, "Loi ket noi RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel() // mo channel
	failOnError(err, "Loi mo channnel")
	defer ch.Close()

	q, err := ch.QueueDeclare( // khai bao queue
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-waits
		nil,     // arguments
	)

	failOnError(err, "Loi khai bao queue")

	body := "Hello"

	err = ch.Publish( // publish msg
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Loi goi message")

	log.Printf(" [x] Sent %s\n", body)
}
