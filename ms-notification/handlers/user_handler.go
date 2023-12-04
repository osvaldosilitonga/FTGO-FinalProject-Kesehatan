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

func UserNotification(msg []byte) error {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("Panic Recover, Error: %v", err)
		}
	}()

	data := dto.UserRegister{}
	if err := json.Unmarshal(msg, &data); err != nil {
		return err
	}

	var body bytes.Buffer
	t, err := template.ParseFiles("./template/user.html")
	if err != nil {
		return err
	}

	err = t.Execute(&body, data)
	if err != nil {
		return err
	}

	email := os.Getenv("EMAIL_HOST_USER")
	password := os.Getenv("EMAIL_HOST_PASSWORD")

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", data.Email)
	m.SetHeader("Subject", "User Registration")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, email, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
