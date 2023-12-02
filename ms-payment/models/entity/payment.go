package entity

import "time"

type Payments struct {
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
