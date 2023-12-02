package entity

import "time"

type Payments struct {
	ID          uint
	InvoiceID   string
	OrderID     string
	UserID      int
	Email       string
	Amount      int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
