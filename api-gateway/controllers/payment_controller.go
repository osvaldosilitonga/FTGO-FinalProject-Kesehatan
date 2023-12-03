package controllers

import "github.com/labstack/echo/v4"

type Payment interface {
	Create(ctx echo.Context) error
}

type PaymentImpl struct {
}

func NewPaymentController() Payment {
	return &PaymentImpl{}
}

func (p *PaymentImpl) Create(c echo.Context) error {
	// TODO: Check order id to order service

	// TODO: Send payment request to payment service

	return nil
}
