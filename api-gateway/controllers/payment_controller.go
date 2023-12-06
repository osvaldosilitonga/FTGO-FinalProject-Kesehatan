package controllers

import (
	"gateway/service"
	"gateway/utils"
	"strconv"

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

	return nil
}

func (p *PaymentImpl) FindByInvoiceID(c echo.Context) error {
	invoiceID := c.Param("id")

	userId := c.Get("id").(int)
	role := c.Get("role").(string)

	// make request to payment service
	resp, code, err := p.PaymentService.FindByInvoiceID(invoiceID)
	if code != 200 {
		return utils.HttpCodeError(c, code, resp.Message)
	}
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, err.Error())
	}

	if resp.UserID != userId && role != "admin" {
		return utils.ErrorMessage(c, &utils.ApiForbidden, "You are not authorized to access this resource")
	}

	return utils.SuccessMessage(c, &utils.ApiOk, resp)
}

func (p *PaymentImpl) FindByOrderID(c echo.Context) error {
	orderID := c.Param("id")

	userId := c.Get("id").(int)
	role := c.Get("role").(string)

	// make request to payment service
	resp, code, err := p.PaymentService.FIndByOrderID(orderID)
	if code != 200 {
		return utils.HttpCodeError(c, code, resp.Message)
	}
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, err.Error())
	}

	if resp.UserID != userId && role != "admin" {
		return utils.ErrorMessage(c, &utils.ApiForbidden, "You are not authorized to access this resource")
	}

	return utils.SuccessMessage(c, &utils.ApiOk, resp)
}

func (p *PaymentImpl) FindByUserID(c echo.Context) error {
	userID := c.Param("id")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	userId := c.Get("id").(int)
	role := c.Get("role").(string)

	if userId != userIDInt && role != "admin" {
		return utils.ErrorMessage(c, &utils.ApiForbidden, "You are not authorized to access this resource")
	}

	queryPage := c.QueryParam("page")
	if queryPage == "" {
		queryPage = "1"
	}

	queryStatus := c.QueryParam("status")
	if len(queryStatus) == 0 {
		queryStatus = "ALL"
	}

	// make request to payment service
	resp, code, err := p.PaymentService.FindByUserID(userID, queryPage, queryStatus)
	if code != 200 {
		return utils.HttpCodeError(c, code, err.Error())
	}
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiOk, resp)
}
