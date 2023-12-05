package entity

type AuthUser struct {
	ID    int    `json:"id"`
	Role  string `json:"role"`
	Email string `json:"email"`
}
