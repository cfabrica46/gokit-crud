package service

import dbapp "github.com/cfabrica46/gokit-crud/database-app/service"

// SignUp(string, string, string) (string, error)
type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type SignUpResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

// SignIn(string, string) (string, error)
type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

// LogOut(string) error
type LogOutRequest struct {
	Token string `json:"token"`
}

type LogOutResponse struct {
	Err string `json:"err,omitempty"`
}

// GetAllUsers() ([]dbapp.User, error)
type GetAllUsersRequest struct {
}

type GetAllUsersResponse struct {
	Users []dbapp.User `json:"users"`
	Err   string       `json:"err,omitempty"`
}

// DeleteAccount(string) error
type DeleteAccountRequest struct {
	Token string `json:"token"`
}

type DeleteAccountResponse struct {
	Err string `json:"err,omitempty"`
}
