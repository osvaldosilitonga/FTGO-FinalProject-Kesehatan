package routes

import (
	"miniproject/handler"
	"miniproject/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	e.POST("/register", handler.RegisterUser)
	e.POST("/register/admin", handler.RegisterAdmin)
	e.POST("/login", handler.LoginUser)

	user := e.Group("/user")
	user.Use(middleware.RequireAuth)
	{
		user.GET("/profile", handler.GetUserProfile)
		user.PUT("/profile/update", handler.UpdateUserProfile)
		user.GET("/activities", handler.GetUserActivities)
	}

	admin := e.Group("/admin")
	admin.Use(middleware.RequireAuth)
	{
		admin.GET("/:id/activities", handler.GetActivitiesByUserID)
		admin.GET("/:id", handler.GetAllUserDataByID)
	}
}
