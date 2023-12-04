package web

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreatePaymentRequest struct {
	OrderID     primitive.ObjectID `json:"order_id"`
	UserID      int                `json:"user_id"`
	Email       string             `json:"email"`
	Amount      int                `json:"amount"`
	Description string             `json:"description"`
}

type CreatePaymentResponse struct {
	InvoiceUrl  string    `json:"invoice_url"`
	InvoiceID   string    `json:"invoice_id"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"created"`
	ExpairyDate time.Time `json:"expairy_date"`
	ExternalId  string    `json:"external_id"`
	PayerEmail  string    `json:"payer_email"`
	Amount      int       `json:"amount"`
}
