package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/models/web"
	"io"
	"net/http"
	"os"
)

type Payment interface {
	CreatePayment(data *web.CreatePaymentRequest) (*web.CreatePaymentResponse, int, error)
}

type PaymentImpl struct {
}

func NewPaymentService() Payment {
	return &PaymentImpl{}
}

func (p *PaymentImpl) CreatePayment(data *web.CreatePaymentRequest) (*web.CreatePaymentResponse, int, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	baseUrl := os.Getenv("PAYMENT_SERVICE_BASE_URL")
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/payment", baseUrl), bytes.NewBuffer(d))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	stringBody := string(body)

	payment := web.CreatePaymentResponse{}

	err = json.Unmarshal([]byte(stringBody), &payment)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &payment, resp.StatusCode, nil
}
