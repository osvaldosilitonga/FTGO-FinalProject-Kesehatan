package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
