package controllers

import (
	"log"
	"payment/api"
	"payment/models/entity"
	"payment/models/web"
	"payment/repository"
	"payment/services"
	"payment/utils"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Payment interface {
	Create(ctx echo.Context) error
	FindByInvoiceID(ctx echo.Context) error
	FindByOrderID(ctx echo.Context) error
	FindByUserID(ctx echo.Context) error
	Cancel(ctx echo.Context) error
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

func (p *PaymentImpl) FindByInvoiceID(c echo.Context) error {
	invoiceId := c.Param("id")

	payment, err := p.Repo.FindByInvoiceID(invoiceId)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiNotFound, err.Error())
	}

	return c.JSON(200, payment)
}

func (p *PaymentImpl) FindByOrderID(c echo.Context) error {
	orderId := c.Param("id")

	payment, err := p.Repo.FindByOrderID(orderId)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiNotFound, err.Error())
	}

	return c.JSON(200, payment)
}

func (p *PaymentImpl) FindByUserID(c echo.Context) error {
	orderId := c.Param("id")
	id, err := strconv.Atoi(orderId)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	status := strings.ToUpper(c.QueryParam("status"))
	if len(status) == 0 {
		status = "ALL"
	}

	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err)
	}

	payment, err := p.Repo.FindByUserID(id, pageInt, status)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiNotFound, err.Error())
	}

	if len(*payment) == 0 {
		return utils.ErrorMessage(c, &utils.ApiNotFound, "Data not found")
	}

	return c.JSON(200, payment)
}

func (p *PaymentImpl) Cancel(c echo.Context) error {
	orderId := c.Param("id")

	payment, err := p.Repo.FindByOrderID(orderId)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiNotFound, err.Error())
	}

	if payment.Status == "PAID" {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "Invoice already paid")
	}

	payment.Status = "CANCEL"
	payment.UpdatedAt = time.Now()

	err = p.Repo.Update(orderId, payment)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, err.Error())
	}

	return c.JSON(200, payment)
}
