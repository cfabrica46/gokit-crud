package service

import (
	"errors"
	"fmt"
	"net/http"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
)

var ErrResponse = errors.New("error to response")

type InfoServices struct {
	DBHost    string
	DBPort    string
	TokenHost string
	TokenPort string
	Secret    string
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Service ...
type Service struct {
	client                                       httpClient
	dbHost, dbPort, tokenHost, tokenPort, secret string
}

// NewService ...
func NewService(client httpClient, is *InfoServices) *Service {
	return &Service{client, is.DBHost, is.DBPort, is.TokenHost, is.TokenPort, is.Secret}
}

// GetIDByUsername ...
func (s *Service) GetIDByUsername(username string) (id int, err error) {
	dbDomain := fmt.Sprintf("%s:%s", s.dbHost, s.dbPort)

	resp, err := DoRequest(NewMRGetIDByUsername(s.client, dbDomain, username))
	if err != nil {
		return 0, err
	}

	r, _ := resp.(dbapp.IDErrorResponse)

	if r.Err != "" {
		return 0, fmt.Errorf("%w:%s", ErrResponse, r.Err)
	}

	return r.ID, nil
}

/* import (
	"errors"
	"fmt"
	"net/http"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	tokenapp "github.com/cfabrica46/gokit-crud/token-app/service"
)

const "%s:%s" = "http://%s:%s"

// ErrTokenNotValid ...
var ErrTokenNotValid = errors.New("token not validate")

type InfoServices struct {
	DBHost    string
	DBPort    string
	TokenHost string
	TokenPort string
	Secret    string
}

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
func NewService(client httpClient, is *InfoServices) *Service {
	return &Service{client, is.DBHost, is.DBPort, is.TokenHost, is.TokenPort, is.Secret}
}

// SignUp ...
func (s *Service) SignUp(username, password, email string) (token string, err error) {
	dbURL := fmt.Sprintf("%s:%s", s.dbHost, s.dbPort)
	tokenURL := fmt.Sprintf("%s:%s", s.tokenHost, s.tokenPort)

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
		return "", err
	}

	userID, err := PetitionGetIDByUsername(
		s.client,
		dbURL+"/id/username",
		dbapp.GetIDByUsernameRequest{
			Username: username,
		},
	)
	if err != nil {
		return "", err
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
		return "", err
	}

	err = PetitionSetToken(
		s.client,
		tokenURL+"/token",
		tokenapp.SetTokenRequest{
			Token: token,
		},
	)
	if err != nil {
		return "", err
	}

	return token, nil
}

// SignIn ...
func (s *Service) SignIn(username, password string) (token string, err error) {
	dbURL := fmt.Sprintf("%s:%s", s.dbHost, s.dbPort)
	tokenURL := fmt.Sprintf("%s:%s", s.tokenHost, s.tokenPort)

	user, err := PetitionGetUserByUsernameAndPassword(
		s.client,
		dbURL+"/user/username_password",
		dbapp.GetUserByUsernameAndPasswordRequest{
			Username: username,
			Password: password,
		},
	)
	if err != nil {
		return "", err
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
		return "", err
	}

	err = PetitionSetToken(
		s.client,
		tokenURL+"/token",
		tokenapp.SetTokenRequest{
			Token: token,
		},
	)
	if err != nil {
		return "", err
	}

	return token, nil
}

// LogOut ...
func (s *Service) LogOut(token string) (err error) {
	tokenURL := fmt.Sprintf("%s:%s", s.tokenHost, s.tokenPort)

	check, err := PetitionCheckToken(
		s.client,
		tokenURL+"/check",
		tokenapp.CheckTokenRequest{
			Token: token,
		},
	)
	if err != nil {
		return err
	}

	if !check {
		err = ErrTokenNotValid

		return err
	}

	err = PetitionDeleteToken(
		s.client,
		tokenURL+"/token",
		tokenapp.DeleteTokenRequest{
			Token: token,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// GetAllUsers  ...
func (s *Service) GetAllUsers() (users []dbapp.User, err error) {
	dbURL := fmt.Sprintf("%s:%s", s.dbHost, s.dbPort)

	users, err = PetitionWithoutBody(s.client, dbURL+"/users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Profile  ...
func (s *Service) Profile(token string) (user dbapp.User, err error) {
	dbURL := fmt.Sprintf("%s:%s", s.dbHost, s.dbPort)
	tokenURL := fmt.Sprintf("%s:%s", s.tokenHost, s.tokenPort)

	check, err := PetitionCheckToken(
		s.client,
		tokenURL+"/check",
		tokenapp.CheckTokenRequest{
			Token: token,
		},
	)
	if err != nil {
		return dbapp.User{}, err
	}

	if !check {
		err = ErrTokenNotValid

		return dbapp.User{}, err
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
		return dbapp.User{}, err
	}

	user, err = PetitionGetUserByID(
		s.client,
		dbURL+"/user/id",
		dbapp.GetUserByIDRequest{
			ID: userID,
		},
	)
	if err != nil {
		return dbapp.User{}, err
	}

	return user, nil
}

// DeleteAccount  ...
func (s *Service) DeleteAccount(token string) (err error) {
	dbURL := fmt.Sprintf("%s:%s", s.dbHost, s.dbPort)
	tokenURL := fmt.Sprintf("%s:%s", s.tokenHost, s.tokenPort)

	check, err := PetitionCheckToken(
		s.client,
		tokenURL+"/check",
		tokenapp.CheckTokenRequest{
			Token: token,
		},
	)
	if err != nil {
		return err
	}

	if !check {
		err = ErrTokenNotValid

		return err
	}

	userID, _, _, err := PetitionExtractToken(s.client,
		tokenURL+"/extract",
		tokenapp.ExtractTokenRequest{
			Token:  token,
			Secret: s.secret,
		},
	)
	if err != nil {
		return err
	}

	err = PetitionDeleteUser(s.client, dbURL+"/user", dbapp.DeleteUserRequest{ID: userID})
	if err != nil {
		return err
	}

	return err
} */
