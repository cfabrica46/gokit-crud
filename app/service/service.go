package service

import (
	"net/http"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	_ "github.com/lib/pq"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type serviceInterface interface {
	SignUp(string, string, string, string) (string, error)
	SignIn(string, string, string) (string, error)
	LogOut(string) error
	ViewEveryUsers() ([]dbapp.User, error)
	DeleteAccount(string) error
}

type service struct {
	dbHost, dbPort, tokenHost, tokenPort string
	client                               httpClient
}

func GetService(dbHost, dbPort, tokenHost, tokenPort string, client httpClient) *service {
	return &service{dbHost, dbPort, tokenHost, tokenPort, client}
}

func (s service) SignUp(username, password, email, secret string) (token string, err error) {
	var dbURL = s.dbHost + ":" + s.dbPort
	var tokenURL = s.tokenHost + ":" + s.tokenPort

	err = petitionInsertUser(s.client, dbURL+"/user", username, password, email)
	if err != nil {
		return
	}

	id, err := petitionGetIDByUsername(s.client, dbURL+"/id/"+username)
	if err != nil {
		return
	}

	token, err = petitionGenerateToken(s.client, tokenURL+"/token", id, username, email, secret)
	if err != nil {
		return
	}

	err = petitionSetToken(s.client, tokenURL+"/token", token)
	if err != nil {
		return
	}
	return
}

func (s service) SignIn(username, password, secret string) (token string, err error) {
	var dbURL = s.dbHost + ":" + s.dbPort
	var tokenURL = s.tokenHost + ":" + s.tokenPort

	user, err := petitionGetUserByUsernameAndPassword(s.client, dbURL+"/user/username_password", username, password)
	if err != nil {
		return
	}

	token, err = petitionGenerateToken(s.client, tokenURL+"/token", user.ID, user.Username, user.Password, user.Email)
	if err != nil {
		return
	}
	return
}

func (s service) LogOut(token string) (err error) {
	var tokenURL = s.tokenHost + ":" + s.tokenPort

	err = petitionCheckToken(s.client, tokenURL+"/check", token)
	if err != nil {
		return
	}

	err = petitionDeleteToken(s.client, tokenURL+"/token", token)
	if err != nil {
		return
	}
	return
}
