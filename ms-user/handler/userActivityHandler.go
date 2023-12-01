package handler

import (
	"miniproject/config"
	"miniproject/entity"
	"net/http"
	"strconv"

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

func GetAllUserDataByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid user ID"})
	}

	// Fetch user data from the database by ID
	user, err := getUserByIDFromDB(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to fetch user data"})
	}

	// Return user data as JSON response
	return c.JSON(http.StatusOK, map[string]interface{}{"user": user})
}

func getUserByIDFromDB(userID int) (*entity.User, error) {
	var user entity.User
	result := config.DB.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
