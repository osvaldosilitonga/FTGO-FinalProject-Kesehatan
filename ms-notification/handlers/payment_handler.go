package handlers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"html/template"
	"log"
	"notification/dto"

	"os"

	"gopkg.in/gomail.v2"
)

func InvoiceNotification(msg []byte) error {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("Panic Recover From handlers invoice notification, Error: %v", err)
		}
	}()

	// get data
	var data dto.Invoice
	if err := json.Unmarshal(msg, &data); err != nil {
		log.Fatalf("error unmarshal: %v", err)
		return err
	}

	var body bytes.Buffer
	t, err := template.ParseFiles("./template/invoice.html")
	t.Execute(&body, data)
	if err != nil {
		log.Fatalf("error parse template: %v", err)
		return err
	}

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
		log.Fatalf("error send email: %v", err)
		return err
	}

	return nil
}

func PaidNotification(msg []byte) error {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("Panic Recover From handlers paid notification, Error: %v", err)
		}
	}()

	// get data
	data := dto.Paid{}
	if err := json.Unmarshal(msg, &data); err != nil {
		return err
	}

	var body bytes.Buffer
	t, err := template.ParseFiles("./template/paid.html")
	t.Execute(&body, data)

	if err != nil {
		return err
	}

	email := os.Getenv("EMAIL_HOST_USER")
	password := os.Getenv("EMAIL_HOST_PASSWORD")

	// Set Header
	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", data.Email)
	m.SetHeader("Subject", "WellnessLink Invoice")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, email, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
