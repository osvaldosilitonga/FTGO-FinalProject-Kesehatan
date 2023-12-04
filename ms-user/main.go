package main

import (
	"log"
	"miniproject/config"
	"miniproject/initializers"
	"miniproject/routes"
	"os"
	"payment/configs"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	conn, rch := configs.InitRabbit()
	defer func() {
		conn.Close()
		rch.Close()
	}()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &initializers.CustomValidator{Validator: validator.New()}

	// Database initialization
	config.InitDB()

	// Routes
	routes.RegisterRoutes(e)

	// Start the server
	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))

}
