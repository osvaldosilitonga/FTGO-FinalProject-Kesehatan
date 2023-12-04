package services

import (
	"log"
	"notification/configs"
	"notification/handlers"

	amqp "github.com/rabbitmq/amqp091-go"
)

type UserService interface {
	UserNotification()
}

type UserServiceImpl struct {
	RabbitCH *amqp.Channel
}

func NewUserService(rc *amqp.Channel) UserService {
	return &UserServiceImpl{
		RabbitCH: rc,
	}
}

func (u *UserServiceImpl) UserNotification() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("Panic Recover, Error: %v", err)
		}
	}()

	q := configs.InitQueue(u.RabbitCH, "user")

	msgs, err := u.RabbitCH.Consume(
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
			err := handlers.UserNotification(data)
			for err != nil {
				log.Println("Handlers Invoice Error: ", err)
				err = handlers.UserNotification(data)
			}
		}
	}()

	<-forever
}
