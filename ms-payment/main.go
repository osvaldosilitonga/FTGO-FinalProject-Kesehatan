package main

import (
	"fmt"
	"log"
	"os"
	"payment/initializers"
	"payment/middlewares"
	"payment/routes"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middlewares.LogrusConfig()))
	e.Validator = &initializers.CustomValidator{Validator: validator.New()}

	routes.Routes(e)

	paymentPort := os.Getenv("PAYMENT_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", paymentPort)))
}
