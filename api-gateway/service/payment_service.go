package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gateway/models/web"
	"io"
	"net/http"
	"os"
)

type Payment interface {
	CreatePayment(data *web.CreatePaymentRequest) (*web.CreatePaymentResponse, int, error)
	FindByInvoiceID(invoiceID string) (*web.Payments, int, error)
	FIndByOrderID(orderID string) (*web.Payments, int, error)
	FindByUserID(userID, queryPage, queryStatus string) (*[]web.Payments, int, error)
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

func (p *PaymentImpl) FindByInvoiceID(invoiceID string) (*web.Payments, int, error) {
	baseUrl := os.Getenv("PAYMENT_SERVICE_BASE_URL")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/payment/%s", baseUrl, invoiceID), nil)
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

	payment := web.Payments{}

	err = json.Unmarshal([]byte(stringBody), &payment)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &payment, resp.StatusCode, nil
}

func (p *PaymentImpl) FIndByOrderID(orderID string) (*web.Payments, int, error) {
	baseUrl := os.Getenv("PAYMENT_SERVICE_BASE_URL")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/payment/order/%s", baseUrl, orderID), nil)
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

	payment := web.Payments{}

	err = json.Unmarshal([]byte(stringBody), &payment)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &payment, resp.StatusCode, nil
}

func (p *PaymentImpl) FindByUserID(userID, queryPage, queryStatus string) (*[]web.Payments, int, error) {
	baseUrl := os.Getenv("PAYMENT_SERVICE_BASE_URL")

	// req, err := http.NewRequest("GET", fmt.Sprintf("%s/payment/user/%s", baseUrl, userID), nil)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/payment/user/%s?page=%s&status=%s", baseUrl, userID, queryPage, queryStatus), nil)
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

	type errorResponse struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
		Detail string `json:"detail"`
	}
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		stringBody := string(body)

		errRes := errorResponse{}

		err = json.Unmarshal([]byte(stringBody), &errRes)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return nil, resp.StatusCode, errors.New(errRes.Detail)
	}

	body, _ := io.ReadAll(resp.Body)
	stringBody := string(body)

	payment := []web.Payments{}

	err = json.Unmarshal([]byte(stringBody), &payment)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &payment, resp.StatusCode, nil
}
