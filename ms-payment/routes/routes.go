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
	amqp "github.com/rabbitmq/amqp091-go"
)

func Routes(e *echo.Echo, rch *amqp.Channel) {
	db := configs.InitDB()

	repository := repository.NewPaymentRepository(db)

	xenditApi := api.NewXenditAPI(os.Getenv("XENDIT_API_KEY"))
	notificationService := services.NewNotificationService(rch)
	paymentController := controllers.NewPaymentController(xenditApi, repository, notificationService)
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
