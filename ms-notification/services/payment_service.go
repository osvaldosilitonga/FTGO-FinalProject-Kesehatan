package services

import (
	"log"
	"notification/configs"
	"notification/handlers"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PaymentService interface {
	InvoiceNotification()
}

type PaymentServiceImpl struct {
	RabbitCH *amqp.Channel
}

func NewPaymentService(rc *amqp.Channel) PaymentService {
	return &PaymentServiceImpl{
		RabbitCH: rc,
	}
}

func (p *PaymentServiceImpl) InvoiceNotification() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("Panic Recover, Error: %v", err)
		}
	}()

	q := configs.InitQueue(p.RabbitCH, "invoice")

	msgs, err := p.RabbitCH.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		panic(err)
	}

	var forever chan struct{}

	go func() {
		var data []byte

		for d := range msgs {
			data = d.Body
			err := handlers.InvoiceNotification(data)
			for err != nil {
				log.Println("Handlers Invoice Error: ", err)
				err = handlers.InvoiceNotification(data)
			}
		}
	}()

	<-forever
}
