package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

type MockUserHandler struct{}

func (m *MockUserHandler) RegisterUser(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Registration successful"})
}

func (m *MockUserHandler) RegisterAdmin(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Admin registration successful"})
}

func (m *MockUserHandler) LoginUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"token": "mock-token", "id": 123, "role": "user"})
}

func TestRegisterUser(t *testing.T) {
	mockHandler := &MockUserHandler{}

	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"name": "Harry Maguire", "email": "maguire@email.com", "password": "maguiree", "address": "123 Street", "phone": "1234567890", "birthdate": "1990-01-01", "gender": "male"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	err := mockHandler.RegisterUser(c)
	if err != nil {
		t.Errorf("Failed to register user: %v", err)
	}

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, rec.Code)
	}
}

func TestRegisterAdmin(t *testing.T) {
	mockHandler := &MockUserHandler{}

	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"name": "Harry Maguire", "email": "maguire@email.com", "password": "maguiree", "address": "123 Street", "phone": "1234567890", "birthdate": "1990-01-01", "gender": "male"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	err := mockHandler.RegisterUser(c)
	if err != nil {
		t.Errorf("Failed to register user: %v", err)
	}

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, rec.Code)
	}
}

func TestLoginUser(t *testing.T) {
	mockHandler := &MockUserHandler{}

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email": "maguire@email.com", "password": "maguiree"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	err := mockHandler.LoginUser(c)
	if err != nil {
		t.Errorf("Login failed: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, rec.Code)
	}
}
