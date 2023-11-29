package handler

import (
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"gopkg.in/gomail.v2"
)

func sendRegistrationEmail(recipientEmail, userName, userAddress string, registrationTime time.Time) error {
	// Konfigurasi email
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Membuat pesan email
	m := gomail.NewMessage()
	m.SetHeader("From", "tankaruniadzikra1@gmail.com") // Ganti dengan alamat email pengirim
	m.SetHeader("To", recipientEmail)                  // Menggunakan alamat email penerima dari parameter fungsi
	m.SetHeader("Subject", "Selamat bergabung!")

	// Membuat isi email dengan informasi pendaftaran
	body := fmt.Sprintf("Selamat, pendaftaran Anda berhasil, %s!\n", userName)
	body += fmt.Sprintf("Alamat: %s\n", userAddress)
	body += fmt.Sprintf("Waktu Pendaftaran: %s\n", registrationTime.Format("2006-01-02 15:04:05"))
	m.SetBody("text/plain", body)

	// Mengirim email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	// Email terkirim dengan sukses
	return nil
}
