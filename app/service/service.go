package service

import (
	"net/http"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	tokenapp "github.com/cfabrica46/gokit-crud/token-app/service"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type serviceInterface interface {
	SignUp(string, string, string) (string, error)
	SignIn(string, string) (string, error)
	LogOut(string) error
	GetAllUsers() ([]dbapp.User, error)
	Profile(string) (dbapp.User, error)
	DeleteAccount(string) error
}

// Service ...
type Service struct {
	dbHost, dbPort, tokenHost, tokenPort, secret string
	client                                       httpClient
}

// GetService ...
func GetService(dbHost, dbPort, tokenHost, tokenPort, secret string, client httpClient) *Service {
	return &Service{dbHost, dbPort, tokenHost, tokenPort, secret, client}
}

// SignUp ...
func (s Service) SignUp(username, password, email string) (token string, err error) {
	var dbURL = s.dbHost + ":" + s.dbPort
	var tokenURL = s.tokenHost + ":" + s.tokenPort

	err = petitionInsertUser(s.client, dbURL+"/user", dbapp.InsertUserRequest{Username: username, Password: password, Email: email})
	if err != nil {
		return
	}

	id, err := petitionGetIDByUsername(s.client, dbURL+"/id/username", dbapp.GetIDByUsernameRequest{Username: username})
	if err != nil {
		return
	}

	token, err = petitionGenerateToken(s.client, tokenURL+"/token", tokenapp.GenerateTokenRequest{ID: id, Username: username, Email: email, Secret: s.secret})
	if err != nil {
		return
	}

	err = petitionSetToken(s.client, tokenURL+"/token", tokenapp.SetTokenRequest{Token: token})
	if err != nil {
		return
	}
	return
}

// SignIn ...
func (s Service) SignIn(username, password string) (token string, err error) {
	var dbURL = s.dbHost + ":" + s.dbPort
	var tokenURL = s.tokenHost + ":" + s.tokenPort

	user, err := petitionGetUserByUsernameAndPassword(s.client, dbURL+"/user/username_password", dbapp.GetUserByUsernameAndPasswordRequest{Username: username, Password: password})
	if err != nil {
		return
	}

	token, err = petitionGenerateToken(s.client, tokenURL+"/token", tokenapp.GenerateTokenRequest{ID: user.ID, Username: user.Username, Email: user.Email, Secret: s.secret})
	if err != nil {
		return
	}

	err = petitionSetToken(s.client, tokenURL+"/token", tokenapp.SetTokenRequest{Token: token})
	if err != nil {
		return
	}
	return
}

// LogOut ...
func (s Service) LogOut(token string) (err error) {
	var tokenURL = s.tokenHost + ":" + s.tokenPort

	err = petitionCheckToken(s.client, tokenURL+"/check", tokenapp.CheckTokenRequest{Token: token})
	if err != nil {
		return
	}

	err = petitionDeleteToken(s.client, tokenURL+"/token", tokenapp.DeleteTokenRequest{Token: token})
	if err != nil {
		return
	}
	return
}

//GetAllUsers  ...
func (s Service) GetAllUsers() (users []dbapp.User, err error) {
	var dbURL = s.dbHost + ":" + s.dbPort

	users, err = petitionGetAllUsers(s.client, dbURL+"/users")
	if err != nil {
		return
	}
	return
}

//Profile  ...
func (s Service) Profile(token string) (user dbapp.User, err error) {
	var dbURL = s.dbHost + ":" + s.dbPort
	var tokenURL = s.tokenHost + ":" + s.tokenPort

	err = petitionCheckToken(s.client, tokenURL+"/check", tokenapp.CheckTokenRequest{Token: token})
	if err != nil {
		return
	}

	id, _, _, err := petitionExtractToken(s.client, tokenURL+"/extract", tokenapp.ExtractTokenRequest{Token: token, Secret: s.secret})
	if err != nil {
		return
	}

	user, err = petitionGetUserByID(s.client, dbURL+"/user/id", dbapp.GetUserByIDRequest{ID: id})
	if err != nil {
		return
	}
	return
}

//DeleteAccount  ...
func (s Service) DeleteAccount(token string) (err error) {
	var dbURL = s.dbHost + ":" + s.dbPort
	var tokenURL = s.tokenHost + ":" + s.tokenPort

	err = petitionCheckToken(s.client, tokenURL+"/check", tokenapp.CheckTokenRequest{Token: token})
	if err != nil {
		return
	}

	id, _, _, err := petitionExtractToken(s.client, tokenURL+"/extract", tokenapp.ExtractTokenRequest{Token: token, Secret: s.secret})
	if err != nil {
		return
	}

	err = petitionDeleteUser(s.client, dbURL+"/user", dbapp.DeleteUserRequest{ID: id})
	if err != nil {
		return
	}
	return
}
