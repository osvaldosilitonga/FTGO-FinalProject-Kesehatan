package web

import "time"

type XenditCallbackBody struct {
	ID                 string    `json:"id"`
	ExternalID         string    `json:"external_id"`
	UserID             string    `json:"user_id"`
	PaymentMethod      string    `json:"payment_method"`
	Status             string    `json:"status"`
	MerchantName       string    `json:"merchant_name"`
	Amount             int       `json:"amount"`
	PaidAmount         int       `json:"paid_amount"`
	BankCode           string    `json:"bank_code"`
	PaidAt             string    `json:"paid_at"`
	PayerEmail         string    `json:"payer_email"`
	Description        string    `json:"description"`
	Currency           string    `json:"currency"`
	PaymentChannel     string    `json:"payment_channel"`
	PaymentDestination string    `json:"payment_destination"`
	Created            time.Time `json:"created"`
	Updated            time.Time `json:"updated"`
}
