package services

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"

	"encoding/json"
	"payment/models/entity"
)

type TriggerApiGateway interface {
	TriggerPaymentUpdate(data *entity.Payments) error
}

type TriggerApiGatewayImpl struct{}

func NewTriggerApiGateway() TriggerApiGateway {
	return &TriggerApiGatewayImpl{}
}

func (t *TriggerApiGatewayImpl) TriggerPaymentUpdate(data *entity.Payments) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	baseUrl := os.Getenv("API_GATEWAY_BASE_URL")
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/paymentupdate", baseUrl), bytes.NewBuffer(d))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Failed to trigger payment update")
	}
	defer resp.Body.Close()

	return nil
}
