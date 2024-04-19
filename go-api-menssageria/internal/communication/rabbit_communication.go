package communication

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SendDataToRabbitMQ(payload *[]byte) error {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to connect to the channel")

	query, err := ch.QueueDeclare("go-products", false, false, false, false, nil)
	failOnError(err, "error to declare queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch.PublishWithContext(ctx, "", query.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        *payload,
	})

	if err != nil {
		return err
	}
	fmt.Println("Product Insertion Logged")
	return nil

}
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
