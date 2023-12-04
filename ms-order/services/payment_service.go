package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"order/models/dto"
	"os"
)

type PaymentService interface {
	CreatePayment(data *dto.PaymentRequest) (*dto.PaymentResponse, int, error)
}

type PaymentServiceImpl struct {
}

func NewPaymentService() PaymentService {
	return &PaymentServiceImpl{}
}

func (p *PaymentServiceImpl) CreatePayment(data *dto.PaymentRequest) (*dto.PaymentResponse, int, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	baseUrl := os.Getenv("PAYMENT_SERVICE_BASE_URL")
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/", baseUrl), bytes.NewBuffer(d))
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

	response := dto.PaymentResponse{}

	err = json.Unmarshal([]byte(stringBody), &response)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &response, resp.StatusCode, nil
}
