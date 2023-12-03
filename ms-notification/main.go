package main

import (
	"log"
	"notification/configs"
	"notification/services"

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

	conn, rch := configs.InitRabbit()
	defer func() {
		conn.Close()
		rch.Close()
	}()

	log.Println("RabbitMQ Connected")

	paymentService := services.NewPaymentService(rch)

	paymentService.InvoiceNotification()
}
