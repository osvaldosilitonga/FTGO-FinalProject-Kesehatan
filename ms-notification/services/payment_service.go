package services

import (
	"log"
	"notification/configs"
	"notification/handlers"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PaymentService interface {
	InvoiceNotification()
	PaidNotification()
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
		log.Fatal(err)
	}

	// var forever chan struct{}

	// go func() {
	// 	var data []byte

	for d := range msgs {
		log.Printf("\033[36mNEW MESSAGE (INVOICE):\033[0m %s", d.Body)

		data := d.Body
		err := handlers.InvoiceNotification(data)
		for err != nil {
			log.Println("Handlers Invoice Error: ", err)
			err = handlers.InvoiceNotification(data)
		}
	}

	// }()

	// <-forever
}

func (p *PaymentServiceImpl) PaidNotification() {
	q := configs.InitQueue(p.RabbitCH, "paid")

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
		log.Fatal(err)
	}

	// var forever chan struct{}

	// go func() {
	// 	var data []byte

	for d := range msgs {
		log.Printf("\033[36mNEW MESSAGE (PAID):\033[0m %s", d.Body)

		data := d.Body
		err := handlers.PaidNotification(data)
		for err != nil {
			log.Println("Handlers Paid Error: ", err)
			err = handlers.PaidNotification(data)
		}
	}
	// }()

	// <-forever
}
