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

type Payments struct {
	Message       string    `json:"-"`
	ID            uint      `json:"id"`
	InvoiceID     string    `json:"invoice_id"`
	OrderID       string    `json:"order_id"`
	UserID        int       `json:"user_id"`
	Email         string    `json:"email"`
	Amount        int       `json:"amount"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	MerchantName  string    `json:"merchant_name"`
	Currency      string    `json:"currency"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PaymentList struct {
	Message string `json:"-"`
	Payment []Payments
}
