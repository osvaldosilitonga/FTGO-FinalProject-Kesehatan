package handler

import (
	"miniproject/config"
	"miniproject/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserActivities(c echo.Context) error {
	userID := c.Get("user").(int)

	var activities []entity.UserActivity
	if err := config.DB.Where("user_id = ?", userID).Find(&activities).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to retrieve user activity"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"user_activities": activities})
}

func GetActivitiesByUserID(c echo.Context) error {
	userID := c.Param("id")

	var userActivities []entity.UserActivity
	if err := config.DB.Where("user_id = ?", userID).Find(&userActivities).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "User activity not found"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"user_activities": userActivities})
}
