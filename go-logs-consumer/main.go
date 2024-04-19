package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@rabbitMQ:5672/")
	failOnError(err, "error to connect")
	defer connection.Close()

	ch, err := connection.Channel()
	failOnError(err, "Error in the channel")
	q, err := ch.QueueDeclare(
		"go-products", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)

	msg, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	var forever chan bool

	for d := range msg {
		log.Printf("Menssage Recieved: %s", d.Body)
	}
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
