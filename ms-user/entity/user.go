package entity

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserProfile struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Birthdate string `json:"birthdate"`
	Gender    string `json:"gender"`
}

type UserActivity struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Method      string    `json:"method"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
