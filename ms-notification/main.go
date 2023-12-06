package main

import (
	"log"
	"notification/configs"
	"notification/services"
	"sync"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("Panic Recover, Error: %v", err)
		}
	}()

	_, rch := configs.InitRabbit()

	log.Println("RabbitMQ Connected")

	paymentService := services.NewPaymentService(rch)
	userService := services.NewUserService(rch)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		paymentService.InvoiceNotification()
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		paymentService.PaidNotification()
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		userService.UserNotification()
		wg.Done()
	}()

	wg.Wait()
}
