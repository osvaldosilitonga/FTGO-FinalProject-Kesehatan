package handlers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"notification/dto"

	"os"

	"gopkg.in/gomail.v2"
)

func InvoiceNotification(msg []byte) error {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("Panic Recover, Error: %v", err)
		}
	}()

	log.Println("Invoice Notification")

	// get data
	log.Println("json unmarshal")
	data := dto.Invoice{}
	if err := json.Unmarshal(msg, &data); err != nil {
		return err
	}
	log.Println("json unmarshal success")

	log.Println(data)

	// get html
	log.Println("get html")
	var body bytes.Buffer
	t, err := template.ParseFiles("./template/invoice.html")
	t.Execute(&body, data)
	log.Println("get html success")

	if err != nil {
		fmt.Println("error parsing template")
		return err
	}

	// port, err := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	// if err != nil {
	// 	return err
	// }
	email := os.Getenv("EMAIL_HOST_USER")
	password := os.Getenv("EMAIL_HOST_PASSWORD")

	// Set Header
	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", data.PayerEmail)
	m.SetHeader("Subject", "WellnessLink Invoice")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, email, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
