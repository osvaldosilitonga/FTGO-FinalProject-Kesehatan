package controllers

import (
	"gateway/service"
	"gateway/utils"

	"github.com/labstack/echo/v4"
)

type Payment interface {
	Create(ctx echo.Context) error
	FindByInvoiceID(ctx echo.Context) error
	FindByOrderID(ctx echo.Context) error
	FindByUserID(ctx echo.Context) error
}

type PaymentImpl struct {
	PaymentService service.Payment
}

func NewPaymentController(ps service.Payment) Payment {
	return &PaymentImpl{
		PaymentService: ps,
	}
}

func (p *PaymentImpl) Create(c echo.Context) error {
	// TODO: Check order id to order service

	// TODO: Send payment request to payment service

	return nil
}

func (p *PaymentImpl) FindByInvoiceID(c echo.Context) error {
	invoiceID := c.Param("id")

	// make request to payment service
	resp, code, err := p.PaymentService.FindByInvoiceID(invoiceID)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, err.Error())
	}
	if code != 200 {
		return utils.HttpCodeError(c, code, resp.Message)
	}

	return utils.SuccessMessage(c, &utils.ApiOk, resp)
}

func (p *PaymentImpl) FindByOrderID(ctx echo.Context) error {
	orderID := ctx.Param("id")

	// make request to payment service
	resp, code, err := p.PaymentService.FIndByOrderID(orderID)
	if err != nil {
		return utils.ErrorMessage(ctx, &utils.ApiInternalServer, err.Error())
	}
	if code != 200 {
		return utils.HttpCodeError(ctx, code, resp.Message)
	}

	return utils.SuccessMessage(ctx, &utils.ApiOk, resp)
}

func (p *PaymentImpl) FindByUserID(ctx echo.Context) error {
	userID := ctx.Param("id")

	queryPage := ctx.QueryParam("page")
	if queryPage == "" {
		queryPage = "1"
	}

	queryStatus := ctx.QueryParam("status")
	if len(queryStatus) == 0 {
		queryStatus = "ALL"
	}

	// make request to payment service
	resp, code, err := p.PaymentService.FindByUserID(userID, queryPage, queryStatus)
	if code != 200 {
		return utils.HttpCodeError(ctx, code, err.Error())
	}
	if err != nil {
		return utils.ErrorMessage(ctx, &utils.ApiInternalServer, err.Error())
	}

	return utils.SuccessMessage(ctx, &utils.ApiOk, resp)
}
