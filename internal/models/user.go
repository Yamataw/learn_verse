package models

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

// SignupRequest payload pour /signup
type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest payload pour /login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
