package routes

import (
	"os"
	"payment/api"
	"payment/repository"

	"payment/configs"
	"payment/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	db := configs.InitDB()

	xenditApi := api.NewXenditAPI(os.Getenv("XENDIT_API_KEY"))

	repository := repository.NewPaymentRepository(db)

	paymentController := controllers.NewPaymentController(xenditApi, repository)

	payment := e.Group("/payment")
	{
		payment.POST("", paymentController.Create)
	}
}
