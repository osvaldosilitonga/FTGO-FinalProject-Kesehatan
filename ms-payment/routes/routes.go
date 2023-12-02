package routes

import (
	"os"
	"payment/api"
	"payment/hooks"
	"payment/repository"
	"payment/services"

	"payment/configs"
	"payment/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	db := configs.InitDB()
	repository := repository.NewPaymentRepository(db)

	xenditApi := api.NewXenditAPI(os.Getenv("XENDIT_API_KEY"))
	paymentController := controllers.NewPaymentController(xenditApi, repository)
	payment := e.Group("/payment")
	{
		payment.POST("", paymentController.Create)
	}

	triggerService := services.NewTriggerApiGateway()
	xenditHook := hooks.NewXenditHook(repository, triggerService)
	hook := e.Group("/hook")
	{
		hook.POST("/xendit/invoice/paid", xenditHook.InvoicePaidHook)
	}
}
