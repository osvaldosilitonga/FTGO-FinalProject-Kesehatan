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
)

func init() {
	initializers.LoadEnvFile()
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middlewares.LogrusConfig()))
	e.Validator = &initializers.CustomValidator{Validator: validator.New()}

	router.Router(e)

	gatewayPort := os.Getenv("GATEWAY_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", gatewayPort)))
}
