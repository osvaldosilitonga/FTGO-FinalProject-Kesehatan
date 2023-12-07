package services

import (
	"context"
	"encoding/json"
	"payment/configs"
	"payment/models/web"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type NotificationService interface {
	SendInvoice(data *web.InvoiceResponse) error
	SendPaid(data *web.PaidNotification) error
}

type NotificationServiceImpl struct {
	RabbitCH *amqp.Channel
}

func NewNotificationService(rc *amqp.Channel) NotificationService {
	return &NotificationServiceImpl{
		RabbitCH: rc,
	}
}

func (n *NotificationServiceImpl) SendInvoice(data *web.InvoiceResponse) error {
	q := configs.InitQueue(n.RabbitCH, "invoice")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = n.RabbitCH.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        jsonData,
		})
	if err != nil {
		return err
	}

	return nil
}

func (n *NotificationServiceImpl) SendPaid(data *web.PaidNotification) error {
	localDate := data.PaidAt.Local()
	data.PaidAt = localDate

	q := configs.InitQueue(n.RabbitCH, "paid")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = n.RabbitCH.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        jsonData,
		})
	if err != nil {
		return err
	}

	return nil
}
