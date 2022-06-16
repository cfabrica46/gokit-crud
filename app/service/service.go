package service

import (
	"errors"
	"fmt"
	"net/http"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	tokenapp "github.com/cfabrica46/gokit-crud/token-app/service"
)

var (
	ErrResponse      = errors.New("error to response")
	ErrTokenNotValid = errors.New("token not validate")
	ErrWebServer     = errors.New("error from web server")
)

type InfoServices struct {
	DBHost    string
	DBPort    string
	TokenHost string
	TokenPort string
	Secret    string
}

type serviceInterface interface {
	SignUp(string, string, string) (string, error)
	SignIn(string, string) (string, error)
	LogOut(string) error
	GetAllUsers() ([]dbapp.User, error)
	Profile(string) (dbapp.User, error)
	DeleteAccount(string) error
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Service ...
type Service struct {
	client                    HttpClient
	dbHost, tokenHost, secret string
}

// NewService ...
func NewService(client HttpClient, is *InfoServices) *Service {
	return &Service{client, "http://" + is.DBHost + ":" + is.DBPort, "http://" + is.TokenHost + ":" + is.TokenPort, is.Secret}
}

// SignUp ...
func (s *Service) SignUp(username, password, email string) (token string, err error) {
	var errorDBResponse dbapp.ErrorResponse
	var idResponse dbapp.IDErrorResponse
	var tokenResponse tokenapp.Token
	var errorTokenResponse tokenapp.ErrorResponse

	if err = RequestFunc(
		s.client,
		dbapp.UsernamePasswordEmailRequest{
			Username: username,
			Password: password,
			Email:    email,
		},
		s.dbHost+"/user",
		http.MethodPost,
		&errorDBResponse,
	); err != nil {
		return "", err
	}

	if errorDBResponse.Err != "" {
		return "", fmt.Errorf("%w:%s", ErrWebServer, errorDBResponse.Err)
	}

	if err = RequestFunc(
		s.client,
		dbapp.UsernameRequest{
			Username: username,
		},
		s.dbHost+"/id/username",
		http.MethodGet,
		&idResponse,
	); err != nil {
		return "", err
	}

	if idResponse.Err != "" {
		return "", fmt.Errorf("%w:%s", ErrWebServer, idResponse.Err)
	}

	if err = RequestFunc(
		s.client,
		tokenapp.IDUsernameEmailSecretRequest{
			ID:       idResponse.ID,
			Username: username,
			Email:    email,
			Secret:   s.secret,
		},
		s.tokenHost+"/generate",
		http.MethodPost,
		&tokenResponse,
	); err != nil {
		return "", err
	}

	if err = RequestFunc(
		s.client,
		tokenapp.Token{
			Token: tokenResponse.Token,
		},
		s.tokenHost+"/token",
		http.MethodPost,
		&errorTokenResponse,
	); err != nil {
		return "", err
	}

	if errorTokenResponse.Err != "" {
		return "", fmt.Errorf("%w:%s", ErrWebServer, errorTokenResponse.Err)
	}

	return tokenResponse.Token, nil
}

// SignIn ...
func (s *Service) SignIn(username, password string) (token string, err error) {
	var userErrorResponse dbapp.UserErrorResponse
	var tokenResponse tokenapp.Token
	var errorResponse tokenapp.ErrorResponse

	if err = RequestFunc(
		s.client,
		dbapp.UsernamePasswordRequest{
			Username: username,
			Password: password,
		},
		s.dbHost+"/user/username_password",
		http.MethodGet,
		&userErrorResponse,
	); err != nil {
		return "", err
	}

	if userErrorResponse.Err != "" {
		return "", fmt.Errorf("%w:%s", ErrWebServer, userErrorResponse.Err)
	}

	if err = RequestFunc(
		s.client,
		tokenapp.IDUsernameEmailSecretRequest{
			ID:       userErrorResponse.User.ID,
			Username: userErrorResponse.User.Username,
			Email:    userErrorResponse.User.Email,
			Secret:   s.secret,
		},
		s.tokenHost+"/generate",
		http.MethodPost,
		&tokenResponse,
	); err != nil {
		return "", err
	}

	if err = RequestFunc(
		s.client,
		tokenapp.Token{
			Token: tokenResponse.Token,
		},
		s.tokenHost+"/token",
		http.MethodPost,
		&errorResponse,
	); err != nil {
		return "", err
	}

	if errorResponse.Err != "" {
		return "", fmt.Errorf("%w:%s", ErrWebServer, errorResponse.Err)
	}

	return tokenResponse.Token, nil
}

// LogOut ...
func (s *Service) LogOut(token string) (err error) {
	var checkErrorResponse tokenapp.CheckErrResponse
	var errorResponse tokenapp.ErrorResponse

	if err = RequestFunc(
		s.client,
		tokenapp.Token{
			Token: token,
		},
		s.tokenHost+"/check",
		http.MethodPost,
		&checkErrorResponse,
	); err != nil {
		return err
	}

	if checkErrorResponse.Err != "" {
		return fmt.Errorf("%w:%s", ErrWebServer, checkErrorResponse.Err)
	}

	if !checkErrorResponse.Check {
		err = ErrTokenNotValid

		return err
	}

	if err = RequestFunc(
		s.client,
		tokenapp.Token{
			Token: token,
		},
		s.tokenHost+"/token",
		http.MethodDelete,
		&errorResponse,
	); err != nil {
		return err
	}

	if errorResponse.Err != "" {
		return fmt.Errorf("%w:%s", ErrWebServer, errorResponse.Err)
	}

	return nil
}

// GetAllUsers  ...
func (s *Service) GetAllUsers() (users []dbapp.User, err error) {
	var usersErrorResponse dbapp.UsersErrorResponse

	if err = RequestFuncWithoutBody(
		s.client,
		s.dbHost+"/users",
		http.MethodGet,
		&usersErrorResponse,
	); err != nil {
		return nil, err
	}

	if usersErrorResponse.Err != "" {
		return nil, fmt.Errorf("%w:%s", ErrWebServer, usersErrorResponse.Err)
	}

	return usersErrorResponse.Users, nil
}

// Profile  ...
func (s *Service) Profile(token string) (user dbapp.User, err error) {
	var checkErrorResponse tokenapp.CheckErrResponse
	var idUsernameEmailErrResponse tokenapp.IDUsernameEmailErrResponse
	var userErrorResponse dbapp.UserErrorResponse

	if err = RequestFunc(
		s.client,
		tokenapp.Token{
			Token: token,
		},
		s.tokenHost+"/check",
		http.MethodPost,
		&checkErrorResponse,
	); err != nil {
		return dbapp.User{}, err
	}

	if checkErrorResponse.Err != "" {
		return dbapp.User{}, fmt.Errorf("%w:%s", ErrWebServer, checkErrorResponse.Err)
	}

	if !checkErrorResponse.Check {
		err = ErrTokenNotValid

		return dbapp.User{}, err
	}

	if err = RequestFunc(
		s.client,
		tokenapp.TokenSecretRequest{
			Token:  token,
			Secret: s.secret,
		},
		s.tokenHost+"/extract",
		http.MethodPost,
		&idUsernameEmailErrResponse,
	); err != nil {
		return dbapp.User{}, err
	}

	if idUsernameEmailErrResponse.Err != "" {
		return dbapp.User{}, fmt.Errorf("%w:%s", ErrWebServer, idUsernameEmailErrResponse.Err)
	}

	if err = RequestFunc(
		s.client,
		dbapp.IDRequest{
			ID: idUsernameEmailErrResponse.ID,
		},
		s.dbHost+"/user/id",
		http.MethodGet,
		&userErrorResponse,
	); err != nil {
		return dbapp.User{}, err
	}

	if userErrorResponse.Err != "" {
		return dbapp.User{}, fmt.Errorf("%w:%s", ErrWebServer, userErrorResponse.Err)
	}

	return userErrorResponse.User, nil
}

// DeleteAccount  ...
func (s *Service) DeleteAccount(token string) (err error) {
	var checkErrorResponse tokenapp.CheckErrResponse
	var idUsernameEmailErrResponse tokenapp.IDUsernameEmailErrResponse
	var errorResponse dbapp.ErrorResponse

	if err = RequestFunc(
		s.client,
		tokenapp.Token{
			Token: token,
		},
		s.tokenHost+"/check",
		http.MethodPost,
		&checkErrorResponse,
	); err != nil {
		return err
	}

	if checkErrorResponse.Err != "" {
		return fmt.Errorf("%w:%s", ErrWebServer, checkErrorResponse.Err)
	}

	if !checkErrorResponse.Check {
		err = ErrTokenNotValid

		return err
	}

	if err = RequestFunc(
		s.client,
		tokenapp.TokenSecretRequest{
			Token:  token,
			Secret: s.secret,
		},
		s.tokenHost+"/extract",
		http.MethodPost,
		&idUsernameEmailErrResponse,
	); err != nil {
		return err
	}

	if idUsernameEmailErrResponse.Err != "" {
		return fmt.Errorf("%w:%s", ErrWebServer, idUsernameEmailErrResponse.Err)
	}

	if err = RequestFunc(
		s.client,
		dbapp.IDRequest{
			ID: idUsernameEmailErrResponse.ID,
		},
		s.dbHost+"/user",
		http.MethodDelete,
		&errorResponse,
	); err != nil {
		return err
	}

	return nil
}
