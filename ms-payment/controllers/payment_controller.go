package controllers

import (
	"log"
	"payment/api"
	"payment/models/entity"
	"payment/models/web"
	"payment/repository"
	"payment/services"
	"payment/utils"

	"github.com/labstack/echo/v4"
)

type Payment interface {
	Create(ctx echo.Context) error
}

type PaymentImpl struct {
	XenditApi api.XenditApi
	Repo      repository.PaymentRepository
	Notif     services.NotificationService
}

func NewPaymentController(x api.XenditApi, r repository.PaymentRepository, n services.NotificationService) Payment {
	return &PaymentImpl{
		XenditApi: x,
		Repo:      r,
		Notif:     n,
	}
}

func (p *PaymentImpl) Create(c echo.Context) error {
	req := new(web.PaymentRequest)
	if err := c.Bind(req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "Invalid request body")
	}
	if err := c.Validate(req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "Invalid request body")
	}

	// Create payment to Xendit
	resp, err := p.XenditApi.CreteInvoice(req)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, err.Error())
	}

	// Save to database
	payment := entity.Payments{
		InvoiceID:   resp.InvoiceID,
		OrderID:     req.OrderID.Hex(),
		UserID:      req.UserID,
		Email:       req.Email,
		Amount:      req.Amount,
		Description: req.Description,
		Status:      resp.Status,
		CreatedAt:   resp.CreateAt,
		UpdatedAt:   resp.CreateAt,
	}
	err = p.Repo.Save(&payment)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, err.Error())
	}

	// Send invoice to email notification
	go func() {
		err = p.Notif.SendInvoice(resp)
		for err != nil {
			err = p.Notif.SendInvoice(resp)
		}

		log.Printf("[Success] Add invoice: '%v' to message broker", resp.InvoiceID)
	}()

	return c.JSON(201, resp)
}
