package api

import (
	"context"
	"payment/models/web"

	xendit "github.com/xendit/xendit-go/v3"
	invoice "github.com/xendit/xendit-go/v3/invoice"
)

type XenditApi interface {
	CreteInvoice(d *web.PaymentRequest) (*web.InvoiceResponse, error)
}

type XenditApiImpl struct {
	Client *xendit.APIClient
}

func NewXenditAPI(key string) XenditApi {
	return &XenditApiImpl{
		Client: xendit.NewClient(key),
	}
}

func (x XenditApiImpl) CreteInvoice(d *web.PaymentRequest) (*web.InvoiceResponse, error) {

	createInvoiceRequest := invoice.CreateInvoiceRequest{
		ExternalId:  d.OrderID.Hex(),
		Amount:      float32(d.Amount),
		PayerEmail:  &d.Email,
		Description: &d.Description,
	}

	resp, _, err := x.Client.InvoiceApi.CreateInvoice(context.Background()).
		CreateInvoiceRequest(createInvoiceRequest).
		Execute()
	if err != nil {
		return nil, err
	}

	data := web.InvoiceResponse{
		InvoiceUrl:  resp.InvoiceUrl,
		InvoiceID:   *resp.Id,
		Status:      string(resp.Status),
		Description: *resp.Description,
		CreateAt:    resp.Created.Local(),
		ExpairyDate: resp.ExpiryDate.Local(),
		ExternalId:  resp.ExternalId,
		PayerEmail:  *resp.PayerEmail,
		Amount:      int(resp.Amount),
	}

	return &data, nil
}
