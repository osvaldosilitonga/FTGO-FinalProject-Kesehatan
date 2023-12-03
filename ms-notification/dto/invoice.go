package dto

import "time"

type Invoice struct {
	InvoiceUrl  string    `json:"invoice_url"`
	InvoiceID   string    `json:"invoice_id"`
	Amount      int       `json:"amount"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"created"`
	ExpairyDate time.Time `json:"expairy_date"`
	ExternalId  string    `json:"external_id"`
	PayerEmail  string    `json:"payer_email"`
}
