package configs

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbit() (*amqp.Connection, *amqp.Channel) {
	username := os.Getenv("RABBIT_USERNAME")
	password := os.Getenv("RABBIT_PASSWORD")
	host := os.Getenv("RABBIT_HOST	")
	port := os.Getenv("RABBIT_PORT")

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port))
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return conn, ch
}

func InitQueue(ch *amqp.Channel, queueName string) amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	return q
}
