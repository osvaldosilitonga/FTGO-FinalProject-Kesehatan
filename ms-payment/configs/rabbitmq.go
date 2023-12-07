package configs

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbit() (*amqp.Connection, *amqp.Channel) {
	rabbitUrl := os.Getenv("RABBIT_URL")

	conn, err := amqp.Dial(rabbitUrl)
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
