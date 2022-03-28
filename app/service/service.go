package service

import (
	"errors"
	"fmt"
	"net/http"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	tokenapp "github.com/cfabrica46/gokit-crud/token-app/service"
)

const schemaURL = "http://%s:%s"

var ErrTokenNotValid = errors.New("token not validate")

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
	client                                       httpClient
	dbHost, dbPort, tokenHost, tokenPort, secret string
}

// NewService ...
func NewService(client httpClient, dbHost, dbPort, tokenHost, tokenPort, secret string) *Service {
	return &Service{client, dbHost, dbPort, tokenHost, tokenPort, secret}
}

// SignUp ...
func (s Service) SignUp(username, password, email string) (token string, err error) {
	dbURL := fmt.Sprintf(schemaURL, s.dbHost, s.dbPort)
	tokenURL := fmt.Sprintf(schemaURL, s.tokenHost, s.tokenPort)

	err = PetitionInsertUser(
		s.client,
		dbURL+"/user",
		dbapp.InsertUserRequest{
			Username: username,
			Password: password,
			Email:    email,
		},
	)
	if err != nil {
		return
	}

	userID, err := PetitionGetIDByUsername(
		s.client,
		dbURL+"/id/username",
		dbapp.GetIDByUsernameRequest{
			Username: username,
		},
	)
	if err != nil {
		return
	}

	token, err = PetitionGenerateToken(s.client,
		tokenURL+"/generate",
		tokenapp.GenerateTokenRequest{
			ID:       userID,
			Username: username,
			Email:    email,
			Secret:   s.secret,
		},
	)
	if err != nil {
		return
	}

	err = PetitionSetToken(
		s.client,
		tokenURL+"/token",
		tokenapp.SetTokenRequest{
			Token: token,
		},
	)
	if err != nil {
		return
	}

	return token, err
}

// SignIn ...
func (s Service) SignIn(username, password string) (token string, err error) {
	dbURL := fmt.Sprintf(schemaURL, s.dbHost, s.dbPort)
	tokenURL := fmt.Sprintf(schemaURL, s.tokenHost, s.tokenPort)

	user, err := PetitionGetUserByUsernameAndPassword(
		s.client,
		dbURL+"/user/username_password",
		dbapp.GetUserByUsernameAndPasswordRequest{
			Username: username,
			Password: password,
		},
	)
	if err != nil {
		return
	}

	token, err = PetitionGenerateToken(
		s.client,
		tokenURL+"/generate",
		tokenapp.GenerateTokenRequest{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Secret:   s.secret,
		},
	)
	if err != nil {
		return
	}

	err = PetitionSetToken(
		s.client,
		tokenURL+"/token",
		tokenapp.SetTokenRequest{
			Token: token,
		},
	)
	if err != nil {
		return
	}

	return token, err
}

// LogOut ...
func (s Service) LogOut(token string) (err error) {
	tokenURL := fmt.Sprintf(schemaURL, s.tokenHost, s.tokenPort)

	check, err := PetitionCheckToken(
		s.client,
		tokenURL+"/check",
		tokenapp.CheckTokenRequest{
			Token: token,
		},
	)
	if err != nil {
		return
	}

	if !check {
		err = ErrTokenNotValid

		return
	}

	err = PetitionDeleteToken(
		s.client,
		tokenURL+"/token",
		tokenapp.DeleteTokenRequest{
			Token: token,
		},
	)
	if err != nil {
		return
	}

	return err
}

// GetAllUsers  ...
func (s Service) GetAllUsers() (users []dbapp.User, err error) {
	dbURL := fmt.Sprintf(schemaURL, s.dbHost, s.dbPort)

	users, err = PetitionGetAllUsers(s.client, dbURL+"/users")
	if err != nil {
		return
	}

	return
}

// Profile  ...
func (s Service) Profile(token string) (user dbapp.User, err error) {
	dbURL := fmt.Sprintf(schemaURL, s.dbHost, s.dbPort)
	tokenURL := fmt.Sprintf(schemaURL, s.tokenHost, s.tokenPort)

	check, err := PetitionCheckToken(
		s.client,
		tokenURL+"/check",
		tokenapp.CheckTokenRequest{
			Token: token,
		},
	)
	if err != nil {
		return
	}

	if !check {
		err = ErrTokenNotValid

		return
	}

	userID, _, _, err := PetitionExtractToken(
		s.client,
		tokenURL+"/extract",
		tokenapp.ExtractTokenRequest{
			Token:  token,
			Secret: s.secret,
		},
	)
	if err != nil {
		return
	}

	user, err = PetitionGetUserByID(
		s.client,
		dbURL+"/user/id",
		dbapp.GetUserByIDRequest{
			ID: userID,
		},
	)
	if err != nil {
		return
	}

	return user, err
}

// DeleteAccount  ...
func (s Service) DeleteAccount(token string) (err error) {
	dbURL := fmt.Sprintf(schemaURL, s.dbHost, s.dbPort)
	tokenURL := fmt.Sprintf(schemaURL, s.tokenHost, s.tokenPort)

	check, err := PetitionCheckToken(
		s.client,
		tokenURL+"/check",
		tokenapp.CheckTokenRequest{
			Token: token,
		},
	)
	if err != nil {
		return
	}

	if !check {
		err = ErrTokenNotValid

		return
	}

	userID, _, _, err := PetitionExtractToken(s.client,
		tokenURL+"/extract",
		tokenapp.ExtractTokenRequest{
			Token:  token,
			Secret: s.secret,
		},
	)
	if err != nil {
		return
	}

	err = PetitionDeleteUser(s.client, dbURL+"/user", dbapp.DeleteUserRequest{ID: userID})
	if err != nil {
		return
	}

	return err
}
