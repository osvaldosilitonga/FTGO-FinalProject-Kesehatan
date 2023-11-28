package routes

import (
	"miniproject/handler"
	"miniproject/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	e.POST("/register", handler.RegisterUser)
	e.POST("/login", handler.LoginUser)

	user := e.Group("/user")
	user.Use(middleware.RequireAuth)
	{
		user.GET("/profile", handler.GetUserProfile)
		user.PUT("/profile/update", handler.UpdateUserProfile)
		user.GET("/activities", handler.GetUserActivities)
	}

	e.GET("/admin/:id/activities", handler.GetActivitiesByUserID, middleware.AdminOnlyMiddleware)
}
