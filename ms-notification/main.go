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
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Fatalf("Panic Recover, Error: %v", err)
	// 	}
	// }()

	conn, rch := configs.InitRabbit()
	defer conn.Close()

	log.Println("RabbitMQ Connected")

	paymentService := services.NewPaymentService(rch)

	go paymentService.InvoiceNotification()
	go paymentService.PaidNotification()

	// goroutine
	// wg := new(sync.WaitGroup)

	// wg.Add(1)
	// go paymentService.InvoiceNotification()
	// defer wg.Done()

	// wg.Add(1)
	// go paymentService.PaidNotification()
	// defer wg.Done()

	// wg.Wait()

	var forever chan struct{}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
