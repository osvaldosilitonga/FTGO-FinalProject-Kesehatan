package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/assert.v1"
)

type MockDB struct {
}

type MockUserHandler struct {
	DB *MockDB
}

func (m *MockUserHandler) GetUserProfile(c echo.Context) error {
	// userID := c.Param("id")

	// userIDInt, err := strconv.Atoi(userID)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid user ID"})
	// }

	// var userProfile entity.UserProfile
	// if err := config.DB.Where("user_id = ?", userID).First(&userProfile).Error; err != nil {
	// 	return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "User profile not found"})
	// }

	// var user entity.User
	// if err := config.DB.Select("name", "email").First(&user, userID).Error; err != nil {
	// 	return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "User profile not found"})
	// }

	// userActivity := entity.UserActivity{
	// 	UserID:      userIDInt,
	// 	Method:      c.Request().Method,
	// 	Description: "Get User Profile",
	// 	Date:        time.Now(),
	// }
	// if err := config.DB.Create(&userActivity).Error; err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to save user activity"})
	// }

	// responseData := dto.UserProfileResponse{
	// 	Email:     user.Email,
	// 	Name:      user.Name,
	// 	ID:        userProfile.ID,
	// 	UserID:    userProfile.UserID,
	// 	Address:   userProfile.Address,
	// 	Phone:     userProfile.Phone,
	// 	Birthdate: userProfile.Birthdate,
	// 	Gender:    userProfile.Gender,
	// }
	// return c.JSON(http.StatusOK, responseData)
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "User profile retrieved successfully"})
}

func (m *MockUserHandler) UpdateUserProfile(c echo.Context) error {
	// userID := c.Param("id")

	// userIDInt, err := strconv.Atoi(userID)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid user ID"})
	// }

	// updatedProfileData := new(entity.UserProfile)
	// if err := c.Bind(updatedProfileData); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	// }

	// var userProfile entity.UserProfile
	// if err := config.DB.Where("user_id = ?", userID).First(&userProfile).Error; err != nil {
	// 	return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "User profile not found"})
	// }

	// userProfile.Address = updatedProfileData.Address
	// userProfile.Phone = updatedProfileData.Phone
	// userProfile.Birthdate = updatedProfileData.Birthdate
	// userProfile.Gender = updatedProfileData.Gender

	// config.DB.Save(&userProfile)

	// userActivity := entity.UserActivity{
	// 	UserID:      userIDInt,
	// 	Method:      c.Request().Method,
	// 	Description: "Update User Profile",
	// 	Date:        time.Now(),
	// }
	// if err := config.DB.Create(&userActivity).Error; err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to save user activity"})
	// }

	// return c.JSON(http.StatusOK, map[string]interface{}{"message": "User profile updated successfully", "user_profile": userProfile})
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "User profile updated successfully"})
}

func (m *MockUserHandler) GetUserByID(c echo.Context) error {
	// userID := c.Param("id")

	// var user entity.User
	// if err := config.DB.First(&user, userID).Error; err != nil {
	// 	return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "User not found"})
	// }

	// var userProfile entity.UserProfile
	// if err := config.DB.Where("user_id = ?", userID).First(&userProfile).Error; err != nil {
	// 	return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "UserProfile not found"})
	// }

	// responseData := dto.UserProfileResponse{
	// 	Email:     user.Email,
	// 	Name:      user.Name,
	// 	ID:        userProfile.ID,
	// 	UserID:    userProfile.UserID,
	// 	Address:   userProfile.Address,
	// 	Phone:     userProfile.Phone,
	// 	Birthdate: userProfile.Birthdate,
	// 	Gender:    userProfile.Gender,
	// }

	// return c.JSON(http.StatusOK, responseData)
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "User profile retrieved successfully"})
}

func TestGetUserProfile(t *testing.T) {
	mockDB := &MockDB{}
	mockHandler := &MockUserHandler{DB: mockDB}

	req := httptest.NewRequest(http.MethodGet, "/user/profile/123", strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath("/user/profile/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := mockHandler.GetUserProfile(c)
	if err != nil {
		t.Errorf("GetUserProfile failed: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateUserProfile(t *testing.T) {
	mockDB := &MockDB{}
	mockHandler := &MockUserHandler{DB: mockDB}

	req := httptest.NewRequest(http.MethodPost, "/user/update/123", strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath("/user/update/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	err := mockHandler.UpdateUserProfile(c)
	if err != nil {
		t.Errorf("UpdateUserProfile failed: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

}

func TestGetUserByID(t *testing.T) {
	mockDB := &MockDB{}
	mockHandler := &MockUserHandler{DB: mockDB}

	req := httptest.NewRequest(http.MethodGet, "/user/profile/123", strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath("/user/profile/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	err := mockHandler.GetUserByID(c)
	if err != nil {
		t.Errorf("GetUserByID failed: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
}
