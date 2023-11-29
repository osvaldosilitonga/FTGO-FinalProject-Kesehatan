package handler

import (
	"miniproject/config"
	"miniproject/entity"
	"miniproject/middleware"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c echo.Context) error {
	input := new(struct {
		entity.User
		Address   string `json:"address"`
		Phone     string `json:"phone"`
		Birthdate string `json:"birthdate"`
		Gender    string `json:"gender"`
	})
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to hash password"})
	}

	user := entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to register user"})
	}

	userProfile := entity.UserProfile{
		UserID:    user.ID,
		Address:   input.Address,
		Phone:     input.Phone,
		Birthdate: input.Birthdate,
		Gender:    input.Gender,
	}

	if err := config.DB.Create(&userProfile).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to create user profile"})
	}

	registrationTime := time.Now()
	if err := sendRegistrationEmail(user.Email, user.Name, userProfile.Address, registrationTime); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to send registration email"})
	}

	user.Password = ""

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Registration successful", "user": user})
}

func LoginUser(c echo.Context) error {

	input := new(entity.User)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	var user entity.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Invalid credentials"})
	}

	tokenString, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to generate token"})
	}

	user.Password = ""

	return c.JSON(http.StatusOK, map[string]interface{}{"token": tokenString, "message": "Login success"})
}
