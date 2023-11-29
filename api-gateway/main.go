package main

import (
	"gateway/initializers"
	"gateway/router"

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
	e.Use(middleware.Logger())
	e.Validator = &initializers.CustomValidator{Validator: validator.New()}

	router.Router(e)

	e.Logger.Fatal(e.Start(":8080"))
}
