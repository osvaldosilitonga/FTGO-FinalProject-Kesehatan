package services

import (
	"context"
	"encoding/json"
	"miniproject/dto"
	"payment/configs"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type UserNotificationService interface {
	SendUserRegistration(data *dto.UserRegister) error
}

type UserNotificationServiceImpl struct {
	RabbitCH *amqp.Channel
}

func NewUserNotificationService(rc *amqp.Channel) UserNotificationService {
	return &UserNotificationServiceImpl{
		RabbitCH: rc,
	}
}

func (u *UserNotificationServiceImpl) SendUserRegistration(data *dto.UserRegister) error {
	q := configs.InitQueue(u.RabbitCH, "user")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = u.RabbitCH.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json", // Changed to "application/json" for JSON data
			Body:        jsonData,
		})
	if err != nil {
		return err
	}

	return nil
}
