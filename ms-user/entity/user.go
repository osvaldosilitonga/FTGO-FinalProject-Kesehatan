package entity

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password,omitempty" validate:"required,min=8"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserProfile struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Address   string `json:"address" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Birthdate string `json:"birthdate" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
}

type UserActivity struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Method      string    `json:"method"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
