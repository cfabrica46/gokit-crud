package service

import dbapp "github.com/cfabrica46/gokit-crud/database-app/service"

// SignUpRequest (string, string, string) (string, error).
type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// SignUpResponse (string, string, string) (string, error).
type SignUpResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

// SignInRequest (string, string) (string, error).
type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignInResponse (string, string, string) (string, error).
type SignInResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

// LogOutRequest (string) error.
type LogOutRequest struct {
	Token string `json:"token"`
}

// LogOutResponse (string, string, string) (string, error).
type LogOutResponse struct {
	Err string `json:"err,omitempty"`
}

// GetAllUsersRequest () ([]dbapp.User, error).
type GetAllUsersRequest struct{}

// GetAllUsersResponse () ([]dbapp.User, error).
type GetAllUsersResponse struct {
	Err   string       `json:"err,omitempty"`
	Users []dbapp.User `json:"users"`
}

// ProfileRequest (string) (dbapp.User, error).
type ProfileRequest struct {
	Token string `json:"token"`
}

// ProfileResponse () (dbapp.User, error).
type ProfileResponse struct {
	User dbapp.User `json:"user"`
	Err  string     `json:"err,omitempty"`
}

// DeleteAccountRequest (string) error.
type DeleteAccountRequest struct {
	Token string `json:"token"`
}

// DeleteAccountResponse () ([]dbapp.User, error).
type DeleteAccountResponse struct {
	Err string `json:"err,omitempty"`
}
