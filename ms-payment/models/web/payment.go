package web

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentRequest struct {
	OrderID     primitive.ObjectID `json:"order_id" validate:"required"`
	UserID      int                `json:"user_id" validate:"required"`
	Email       string             `json:"email" validate:"required"`
	Amount      int                `json:"amount" validate:"required"`
	Description string             `json:"description" validate:"required"`
}

type InvoiceResponse struct {
	InvoiceUrl  string    `json:"invoice_url"`
	InvoiceID   string    `json:"invoice_id"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"created"`
	ExpairyDate time.Time `json:"expairy_date"`
	ExternalId  string    `json:"external_id"`
	PayerEmail  string    `json:"payer_email"`
}
