package handler

import (
	"net/http"
	"time"

	"miniproject/config"
	"miniproject/entity"

	"github.com/labstack/echo/v4"
)

func CreateUserProfile(c echo.Context) error {
	userID := c.Get("user").(int)

	newProfileData := new(entity.UserProfile)
	if err := c.Bind(newProfileData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	userProfile := entity.UserProfile{
		UserID:    userID,
		Address:   newProfileData.Address,
		Phone:     newProfileData.Phone,
		Birthdate: newProfileData.Birthdate,
		Gender:    newProfileData.Gender,
	}

	if err := config.DB.Create(&userProfile).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to create a user profile"})
	}

	userActivity := entity.UserActivity{
		UserID:      userID,
		Method:      c.Request().Method,
		Description: "Create User Profile",
		Date:        time.Now(),
	}
	if err := config.DB.Create(&userActivity).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to save user activity"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "User profile created successfully", "user_profile": userProfile})
}

func GetUserProfile(c echo.Context) error {
	userID := c.Get("user").(int)

	var userProfile entity.UserProfile
	if err := config.DB.Where("user_id = ?", userID).First(&userProfile).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "User profile not found"})
	}

	userActivity := entity.UserActivity{
		UserID:      userID,
		Method:      c.Request().Method,
		Description: "Get User Profile",
		Date:        time.Now(),
	}
	if err := config.DB.Create(&userActivity).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to save user activity"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"user_profile": userProfile})
}

func UpdateUserProfile(c echo.Context) error {
	userID := c.Get("user").(int)

	updatedProfileData := new(entity.UserProfile)
	if err := c.Bind(updatedProfileData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	var userProfile entity.UserProfile
	if err := config.DB.Where("user_id = ?", userID).First(&userProfile).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "User profile not found"})
	}

	userProfile.Address = updatedProfileData.Address
	userProfile.Phone = updatedProfileData.Phone
	userProfile.Birthdate = updatedProfileData.Birthdate
	userProfile.Gender = updatedProfileData.Gender

	config.DB.Save(&userProfile)

	userActivity := entity.UserActivity{
		UserID:      userID,
		Method:      c.Request().Method,
		Description: "Update User Profile",
		Date:        time.Now(),
	}
	if err := config.DB.Create(&userActivity).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to save user activity"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "User profile updated successfully", "user_profile": userProfile})
}
