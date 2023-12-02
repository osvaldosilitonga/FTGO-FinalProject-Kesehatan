package hooks

import (
	"os"
	"payment/models/web"
	"payment/repository"
	"payment/services"
	"payment/utils"

	"github.com/labstack/echo/v4"
)

type XenditHook interface {
	InvoicePaidHook(c echo.Context) error
}

type XenditHookImpl struct {
	Repo    repository.PaymentRepository
	Trigger services.TriggerApiGateway
}

func NewXenditHook(r repository.PaymentRepository, st services.TriggerApiGateway) XenditHook {
	return &XenditHookImpl{
		Repo:    r,
		Trigger: st,
	}
}

func (x *XenditHookImpl) InvoicePaidHook(c echo.Context) error {
	cbToken := c.Request().Header.Get("x-callback-token")
	if cbToken != os.Getenv("XENDIT_WEBHOOK_VERIFICATION_TOKEN") {
		return utils.ErrorMessage(c, &utils.ApiUnauthorized, "Invalid callback token")
	}

	body := web.XenditCallbackBody{}
	if err := c.Bind(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "Invalid request body")
	}

	payment, err := x.Repo.UpdateFromXendit(&body)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, "Failed to update data")
	}

	err = x.Trigger.TriggerPaymentUpdate(payment)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, "Failed to trigger payment update")
	}

	return utils.SuccessMessage(c, &utils.ApiOk, "Success")
}
