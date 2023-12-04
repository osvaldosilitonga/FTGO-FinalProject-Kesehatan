package routes

import (
	"miniproject/handler"
	"miniproject/middleware"

	"github.com/labstack/echo/v4"
	amqp "github.com/rabbitmq/amqp091-go"
)

func RegisterRoutes(e *echo.Echo, rch *amqp.Channel) {

	// UserNotificationService := services.NewUserNotificationService(rch)

	userHandler := handler.NewUserHandler(rch)

	e.POST("/register", userHandler.RegisterUser)
	e.POST("/register/admin", userHandler.RegisterAdmin)
	e.POST("/login", userHandler.LoginUser)
	// e.POST("/register", handler.RegisterUser)
	// e.POST("/register/admin", handler.RegisterAdmin)
	// e.POST("/login", handler.LoginUser)

	user := e.Group("/user")
	user.Use(middleware.RequireAuth)
	{
		user.GET("/profile", handler.GetUserProfile) //ngambil dari parameter id, jadi handler diganti juga
		user.PUT("/profile/update", handler.UpdateUserProfile)
		user.GET("/activities", handler.GetUserActivities)
	}

	admin := e.Group("/admin")
	admin.Use(middleware.RequireAuth)
	{
		admin.GET("/:id/activities", handler.GetActivitiesByUserID)
		admin.GET("/:id", handler.GetUserByID)
	}
}
