package hooks

import (
	"gateway/models/web"
	"gateway/service"
	"net/http"
	"os"

	orderPb "gateway/internal/order"

	"github.com/labstack/echo/v4"
)

type XenditHooks interface {
	InvoiceHooks(c echo.Context) error
}

type XenditHooksImpl struct {
	OrderService   service.Order
	PaymentService service.Payment
}

func NewXenditHooks(o service.Order, p service.Payment) XenditHooks {
	return &XenditHooksImpl{
		OrderService:   o,
		PaymentService: p,
	}
}

func (x *XenditHooksImpl) InvoiceHooks(c echo.Context) error {
	cbToken := c.Request().Header.Get("x-callback-token")
	if cbToken != os.Getenv("XENDIT_WEBHOOK_VERIFICATION_TOKEN") {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "token validation error",
		})
	}

	body := web.XenditCallbackBody{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "binding body error",
		})
	}

	// Update order status
	order := &orderPb.UpdateOrderStatusRequest{
		OrderId: body.ExternalID,
		Status:  body.Status,
	}

	_, err := x.OrderService.UpdateStatus(c.Request().Context(), order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "update order status error",
		})
	}

	// Update payment status
	paidRequest := &web.PaidRequest{
		Status:        body.Status,
		PaymentMethod: body.PaymentMethod,
		MerchantName:  body.MerchantName,
	}

	err = x.PaymentService.Paid(body.ExternalID, paidRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "update payment status error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{})
}
