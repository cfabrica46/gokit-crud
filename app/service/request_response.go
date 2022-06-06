package service

import dbapp "github.com/cfabrica46/gokit-crud/database-app/service"

// UsernamePasswordEmailRequest (string, string, string) (string, error).
type UsernamePasswordEmailRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UsernamePasswordRequest (string, string) (string, error).
type UsernamePasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TokenRequest (string) error.
type TokenRequest struct {
	Token string `json:"token"`
}

// EmptyRequest () ([]dbapp.User, error).
type EmptyRequest struct{}

// ---

// TokenErrorResponse (string, string, string) (string, error).
type TokenErrorResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

// UsersErrorResponse () ([]dbapp.User, error).
type UsersErrorResponse struct {
	Err   string       `json:"err,omitempty"`
	Users []dbapp.User `json:"users"`
}

// UserErrorResponse () (dbapp.User, error).
type UserErrorResponse struct {
	User dbapp.User `json:"user"`
	Err  string     `json:"err,omitempty"`
}

// ErrorResponse (string, string, string) (string, error).
type ErrorResponse struct {
	Err string `json:"err,omitempty"`
}
