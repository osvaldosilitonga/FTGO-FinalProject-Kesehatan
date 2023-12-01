package dto

type UserLoginRequest struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Birthdate string `json:"birthdate"`
	Gender    string `json:"gender"`
}
