package web

// Request
type UsersLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Response
type LoginResponse struct {
	Token string `json:"token"`
}

// HttpUserLogin is a struct to response login body from user service
type HttpUserLogin struct {
	ID      int    `json:"id"`
	Role    string `json:"role"`
	Token   string `json:"token"`
	Message string `json:"message"`
}
