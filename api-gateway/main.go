package main

import (
	"fmt"
	"gateway/initializers"
	"gateway/middlewares"
	"gateway/router"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title WellnessLink by Cap-OT - API Documentation
// @version BETA
// @description The Health and Pharmaceutical Sales API provides an integrated solution for health and drug sales business systems.

// @contact.name WellnessLink
// @contact.url www.welnesslink.com
// @contact.email wellnesslink.ot@gmail.com

// @host localhost:8080
// @BasePath /api/v1
// @license.name Apache 2.0

func init() {
	initializers.LoadEnvFile()
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middlewares.LogrusConfig()))
	e.Validator = &initializers.CustomValidator{Validator: validator.New()}

	router.Router(e)

	e.GET("/swagger/*.html", echoSwagger.WrapHandler)

	gatewayPort := os.Getenv("GATEWAY_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", gatewayPort)))
}
