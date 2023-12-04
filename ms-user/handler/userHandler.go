package handler

import (
	"miniproject/config"
	"miniproject/dto"
	"miniproject/entity"
	"miniproject/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c echo.Context) error {
	input := dto.UserRegisterRequest{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}
	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to hash password"})
	}

	user := entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "user",
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

	user.Password = ""

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Registration successful", "user": user})
}

func RegisterAdmin(c echo.Context) error {
	input := dto.UserRegisterRequest{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}
	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to hash password"})
	}

	user := entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "admin",
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

	user.Password = ""

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Registration successful", "user": user})
}

func LoginUser(c echo.Context) error {
	input := new(dto.UserLoginRequest)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	var user entity.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "Email not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid password"})
	}

	tokenString, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenString,
		"id":    user.ID,
		"role":  user.Role,
	})
}
