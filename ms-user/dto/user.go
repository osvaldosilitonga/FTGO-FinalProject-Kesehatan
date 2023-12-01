package dto

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=8"`
}

type UserRegisterRequest struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password,omitempty" validate:"required,min=8"`
	Address   string `json:"address" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Birthdate string `json:"birthdate" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
}
