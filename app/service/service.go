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
	return &Service{client, is.DBHost + ":" + is.DBPort, is.TokenHost + ":" + is.TokenPort, is.Secret}
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
		return "", fmt.Errorf("%w:%s", ErrWebServer, errorDBResponse.Err)
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
		return "", fmt.Errorf("%w:%s", ErrWebServer, errorDBResponse.Err)
	}

	return tokenResponse.Token, nil
}

/* // SignIn ...
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
} */

/* // LogOut ...
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
} */

/* // GetAllUsers  ...
func (s *Service) GetAllUsers() (users []dbapp.User, err error) {
	dbURL := fmt.Sprintf("%s:%s", s.dbHost, s.dbPort)

	users, err = PetitionWithoutBody(s.client, dbURL+"/users")
	if err != nil {
		return nil, err
	}

	return users, nil
} */

/* // Profile  ...
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
} */

/* // DeleteAccount  ...
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
