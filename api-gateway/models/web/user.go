package web

// Request
type UsersLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UsersRegisterRequest struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password,omitempty" validate:"required,min=8"`
	Address   string `json:"address" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Birthdate string `json:"birthdate" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
}

// Response
type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// ------------------------------------------------------------

// HttpUserLogin is a struct to response login body from user service
type HttpUserLogin struct {
	ID      int    `json:"id"`
	Role    string `json:"role"`
	Token   string `json:"token"`
	Message string `json:"message"`
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
type HttpUserRegister struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

type Profile struct {
	UserID    int    `json:"user_id"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Birthdate string `json:"birthdate"`
	Gender    string `json:"gender"`
}
type HttpUserProfile struct {
	// Message string `json:"message"`
	// User    Profile `json:"user"`
	UserID    int    `json:"user_id"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Birthdate string `json:"birthdate"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Name      string `json:"name"`
}
