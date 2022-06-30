package service_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/cfabrica46/gokit-crud/app/service"
	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	"github.com/stretchr/testify/assert"
)

const (
	idTest       int    = 1
	usernameTest string = "username"
	passwordTest string = "password"
	emailTest    string = "email@email.com"
	secretTest   string = "secret"

	urlTest       string = "localhost:8080"
	dbHostTest    string = "db"
	tokenHostTest string = "token"
	portTest      string = "8080"
	tokenTest     string = "token"

	nameNoError string = "NoError"
)

var (
	errWebServer        = errors.New("error from web server")
	errNotTypeIndicated = errors.New("response is not of the type indicated")
)

type errorHTTPComponents struct {
	errorURL, errorMethod string
}

type signUpTestStruct struct {
	name                            string
	inUsername, inPassword, inEmail string
	url                             string
	method                          string
	isError                         bool
	isErrorInsideRequest            bool
}

type profileTestStruct struct {
	name                 string
	inToken              string
	url                  string
	method               string
	outUser              dbapp.User
	outCheck             bool
	isError              bool
	isErrorInsideRequest bool
}

func newErrorHTTPComponets(url, method string) errorHTTPComponents {
	return errorHTTPComponents{
		errorURL:    url,
		errorMethod: method,
	}
}

func getSignUpTestEntity() []signUpTestStruct {
	return []signUpTestStruct{
		{
			name:       "NoError",
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			isError:    false,
			url:        "http://db:8080/user",
			method:     http.MethodPost,
		},
		{
			name:       "ErrorInsertUser",
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			isError:    true,
			url:        "http://db:8080/user",
			method:     http.MethodPost,
		},
		{
			name:                 "ErrorInsideInsertUser",
			inUsername:           usernameTest,
			inPassword:           passwordTest,
			inEmail:              emailTest,
			isError:              true,
			isErrorInsideRequest: true,
			url:                  "http://db:8080/user",
			method:               http.MethodPost,
		},
		{
			name:       "ErrorGetID",
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			isError:    true,
			url:        "http://db:8080/id/username",
			method:     http.MethodGet,
		},
		{
			name:                 "ErrorInsideGetID",
			inUsername:           usernameTest,
			inPassword:           passwordTest,
			inEmail:              emailTest,
			isError:              true,
			isErrorInsideRequest: true,
			url:                  "http://db:8080/id/username",
			method:               http.MethodGet,
		},
		{
			name:       "ErrorGenerate",
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			isError:    true,
			url:        "http://token:8080/generate",
			method:     http.MethodPost,
		},
		{
			name:       "ErrorSetToken",
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			isError:    true,
			url:        "http://token:8080/token",
			method:     http.MethodPost,
		},
		{
			name:                 "ErrorInsideSetToken",
			inUsername:           usernameTest,
			inPassword:           passwordTest,
			inEmail:              emailTest,
			isError:              true,
			isErrorInsideRequest: true,
			url:                  "http://token:8080/token",
			method:               http.MethodPost,
		},
	}
}

func getProfileTestEntity() []profileTestStruct {
	return []profileTestStruct{
		{
			name:    "NoError",
			inToken: tokenTest,
			outUser: dbapp.User{
				ID:       idTest,
				Username: usernameTest,
				Password: passwordTest,
				Email:    emailTest,
			},
			outCheck: true,
			isError:  false,
			url:      "http://db:8080/user/id",
			method:   http.MethodGet,
		},
		{
			name:     "ErrorCheckToken",
			inToken:  tokenTest,
			outUser:  dbapp.User{},
			outCheck: true,
			isError:  true,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			name:                 "ErrorInsideCheckToken",
			inToken:              tokenTest,
			outUser:              dbapp.User{},
			outCheck:             true,
			isError:              true,
			isErrorInsideRequest: true,
			url:                  "http://token:8080/check",
			method:               http.MethodPost,
		},
		{
			name:     "FalseCheckToken",
			inToken:  tokenTest,
			outUser:  dbapp.User{},
			outCheck: false,
			isError:  false,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			name:     "ErrorExtractToken",
			inToken:  tokenTest,
			outUser:  dbapp.User{},
			outCheck: true,
			isError:  true,
			url:      "http://token:8080/extract",
			method:   http.MethodPost,
		},
		{
			name:                 "ErrorInsideExtractToken",
			inToken:              tokenTest,
			outUser:              dbapp.User{},
			outCheck:             true,
			isError:              true,
			isErrorInsideRequest: true,
			url:                  "http://token:8080/extract",
			method:               http.MethodPost,
		},
		{
			name:     "ErrorGetID",
			inToken:  tokenTest,
			outUser:  dbapp.User{},
			outCheck: true,
			isError:  true,
			url:      "http://db:8080/user/id",
			method:   http.MethodGet,
		},
		{
			name:                 "ErrorInsideGetID",
			inToken:              tokenTest,
			outUser:              dbapp.User{},
			outCheck:             true,
			isError:              true,
			isErrorInsideRequest: true,
			url:                  "http://db:8080/user/id",
			method:               http.MethodGet,
		},
	}
}

func TestSignUp(t *testing.T) {
	t.Parallel()

	infoServiceTest := service.InfoServices{
		DBHost:    dbHostTest,
		DBPort:    portTest,
		TokenHost: tokenHostTest,
		TokenPort: portTest,
		Secret:    secretTest,
	}

	for _, tt := range getSignUpTestEntity() {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultToken string
			var resultErr error
			var tokenResponse, errorResponse string
			var mock *service.MockClient

			if tt.isError {
				errorResponse = errWebServer.Error()
			} else {
				tokenResponse = tokenTest
			}

			responseJSON := `{
						"token": "token",
						"id": 1
			}`

			if tt.isError {
				mock = service.NewMockClient(getIsErrorMock(
					tt.isErrorInsideRequest,
					newErrorHTTPComponets(tt.url, tt.method),
					responseJSON,
				))
			} else {
				mock = service.NewMockClient(getMock(
					responseJSON,
				))
			}

			svc := service.NewService(
				mock,
				&infoServiceTest,
			)

			resultToken, resultErr = svc.SignUp(tt.inUsername, tt.inPassword, tt.inEmail)

			if !tt.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, errorResponse)
			}
			assert.Equal(t, tokenResponse, resultToken)
		})
	}
}

func TestSignIn(t *testing.T) {
	t.Parallel()

	infoServiceTest := service.InfoServices{
		DBHost:    dbHostTest,
		DBPort:    portTest,
		TokenHost: tokenHostTest,
		TokenPort: portTest,
		Secret:    secretTest,
	}

	for _, tt := range []struct {
		name                   string
		inUsername, inPassword string
		url                    string
		method                 string
		isError                bool
		isErrorInsideRequest   bool
	}{
		{
			name:       "NoError",
			inUsername: usernameTest,
			inPassword: passwordTest,
			isError:    false,
			url:        "http://token:8080/generate",
			method:     http.MethodPost,
		},
		{
			name:       "ErrorGetUser",
			inUsername: usernameTest,
			inPassword: passwordTest,
			isError:    true,
			url:        "http://db:8080/user/username_password",
			method:     http.MethodGet,
		},
		{
			name:                 "ErrorInsideGetToken",
			inUsername:           usernameTest,
			inPassword:           passwordTest,
			isError:              true,
			isErrorInsideRequest: true,
			url:                  "http://db:8080/user/username_password",
			method:               http.MethodGet,
		},
		{
			name:       "ErrorGenerateToken",
			inUsername: usernameTest,
			inPassword: passwordTest,
			isError:    true,
			url:        "http://token:8080/generate",
			method:     http.MethodPost,
		},
		{
			name:       "ErrorSetToken",
			inUsername: usernameTest,
			inPassword: passwordTest,
			isError:    true,
			url:        "http://token:8080/token",
			method:     http.MethodPost,
		},
		{
			name:                 "ErrorInsideSetToken",
			inUsername:           usernameTest,
			inPassword:           passwordTest,
			isError:              true,
			isErrorInsideRequest: true,
			url:                  "http://token:8080/token",
			method:               http.MethodPost,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultToken string
			var resultErr error
			var tokenResponse, errorResponse string
			var mock *service.MockClient

			if tt.isError {
				errorResponse = errWebServer.Error()
			} else {
				tokenResponse = tokenTest
			}

			responseJSON := `{
					"token":"token",
					"user":{
						"id":       1,
						"username": "username",
						"password": "password",
						"email":    "email@email.com"
					}
			}`

			if tt.isError {
				mock = service.NewMockClient(getIsErrorMock(
					tt.isErrorInsideRequest,
					newErrorHTTPComponets(tt.url, tt.method),
					responseJSON,
				))
			} else {
				mock = service.NewMockClient(getMock(
					responseJSON,
				))
			}

			svc := service.NewService(
				mock,
				&infoServiceTest,
			)

			resultToken, resultErr = svc.SignIn(tt.inUsername, tt.inPassword)

			if !tt.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, errorResponse)
			}
			assert.Equal(t, tokenResponse, resultToken)
		})
	}
}

func TestLogOut(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                 string
		inToken              string
		url                  string
		method               string
		outCheck             bool
		isError              bool
		isErrorInsideRequest bool
	}{
		{
			name:     "NoError",
			inToken:  tokenTest,
			outCheck: true,
			isError:  false,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			name:     "ErrorCheckToken",
			inToken:  tokenTest,
			outCheck: true,
			isError:  true,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			name:                 "ErrorInsideCheckToken",
			inToken:              tokenTest,
			outCheck:             true,
			isError:              true,
			isErrorInsideRequest: true,
			url:                  "http://token:8080/check",
			method:               http.MethodPost,
		},
		{
			name:     "FalseCheckToken",
			inToken:  tokenTest,
			outCheck: false,
			isError:  false,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			name:     "ErrorDeleteToken",
			inToken:  tokenTest,
			outCheck: true,
			isError:  true,
			url:      "http://token:8080/token",
			method:   http.MethodDelete,
		},
		{
			name:                 "ErrorDeleteToken",
			inToken:              tokenTest,
			outCheck:             true,
			isError:              true,
			isErrorInsideRequest: true,
			url:                  "http://token:8080/token",
			method:               http.MethodDelete,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr error
			var errorResponse string
			var mock *service.MockClient

			infoServiceTest := service.InfoServices{
				DBHost:    dbHostTest,
				DBPort:    portTest,
				TokenHost: tokenHostTest,
				TokenPort: portTest,
				Secret:    secretTest,
			}

			if tt.isError {
				errorResponse = errWebServer.Error()
			}

			responseJSON := fmt.Sprintf(`{
					"check": %t
				}`, tt.outCheck)

			if tt.isError {
				mock = service.NewMockClient(getIsErrorMock(
					tt.isErrorInsideRequest,
					newErrorHTTPComponets(tt.url, tt.method),
					responseJSON,
				))
			} else {
				mock = service.NewMockClient(getMock(
					responseJSON,
				))
			}

			svc := service.NewService(
				mock,
				&infoServiceTest,
			)

			resultErr = svc.LogOut(tt.inToken)

			if !tt.isError {
				if tt.outCheck {
					assert.Nil(t, resultErr)
				} else {
					assert.ErrorContains(t, resultErr, errorResponse)
				}
			} else {
				assert.ErrorContains(t, resultErr, errorResponse)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                 string
		url                  string
		method               string
		outUsers             []dbapp.User
		isError              bool
		isErrorInsideRequest bool
	}{
		{
			name: "NoError",
			outUsers: []dbapp.User{
				{
					ID:       idTest,
					Username: usernameTest,
					Password: passwordTest,
					Email:    emailTest,
				},
			},
			isError: false,
			url:     "http://db:8080/users",
			method:  http.MethodGet,
		},
		{
			name:     "ErrorGetAllUsers",
			outUsers: nil,
			isError:  true,
			url:      "http://db:8080/users",
			method:   http.MethodGet,
		},
		{
			name:                 "ErrorInsideGetAllUsers",
			outUsers:             nil,
			isError:              true,
			isErrorInsideRequest: true,
			url:                  "http://db:8080/users",
			method:               http.MethodGet,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			infoServiceTest := service.InfoServices{
				DBHost:    dbHostTest,
				DBPort:    portTest,
				TokenHost: tokenHostTest,
				TokenPort: portTest,
				Secret:    secretTest,
			}

			var resultUsers []dbapp.User
			var resultErr error
			var errorResponse string
			var mock *service.MockClient

			if tt.isError {
				errorResponse = errWebServer.Error()
			}

			responseJSON := `{
				"users":[
					{
						"username":"username",
						"password":"password",
						"email":"email@email.com",
						"id":1
					}
				]
			}`

			if tt.isError {
				mock = service.NewMockClient(getIsErrorMock(
					tt.isErrorInsideRequest,
					newErrorHTTPComponets(tt.url, tt.method),
					responseJSON,
				))
			} else {
				mock = service.NewMockClient(getMock(
					responseJSON,
				))
			}

			svc := service.NewService(
				mock,
				&infoServiceTest,
			)

			resultUsers, resultErr = svc.GetAllUsers()

			if !tt.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, errorResponse)
			}
			assert.Equal(t, tt.outUsers, resultUsers)
		})
	}
}

func TestProfile(t *testing.T) {
	t.Parallel()

	for _, tt := range getProfileTestEntity() {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			infoServiceTest := service.InfoServices{
				DBHost:    dbHostTest,
				DBPort:    portTest,
				TokenHost: tokenHostTest,
				TokenPort: portTest,
				Secret:    secretTest,
			}

			var resultUser dbapp.User
			var resultErr error
			var errorResponse string
			var mock *service.MockClient

			if tt.isError {
				errorResponse = errWebServer.Error()
			}

			responseJSON := fmt.Sprintf(`{
					"user":{
						"username":"username",
						"password":"password",
						"email":"email@email.com",
						"id":1
					},
					"id":1,
					"username":"usename",
					"email":"email@email.com",
					"check":%t
				}`, tt.outCheck)

			if tt.isError {
				mock = service.NewMockClient(getIsErrorMock(
					tt.isErrorInsideRequest,
					newErrorHTTPComponets(tt.url, tt.method),
					responseJSON,
				))
			} else {
				mock = service.NewMockClient(getMock(
					responseJSON,
				))
			}

			svc := service.NewService(
				mock,
				&infoServiceTest,
			)

			resultUser, resultErr = svc.Profile(tt.inToken)

			if !tt.isError {
				if tt.outCheck {
					assert.Nil(t, resultErr)
				} else {
					assert.ErrorContains(t, resultErr, errorResponse)
				}
			} else {
				assert.ErrorContains(t, resultErr, errorResponse)
			}

			assert.Equal(t, tt.outUser, resultUser)
		})
	}
}

func TestDeleteAccount(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                 string
		inToken              string
		url                  string
		method               string
		outCheck             bool
		isError              bool
		isErrorInsideRequest bool
	}{
		{
			name:     "NoError",
			inToken:  tokenTest,
			outCheck: true,
			isError:  false,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			name:     "ErrorCheckToken",
			inToken:  tokenTest,
			outCheck: true,
			isError:  true,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			name:                 "ErrorInsideCheckToken",
			inToken:              tokenTest,
			outCheck:             true,
			isError:              true,
			isErrorInsideRequest: true,
			url:                  "http://token:8080/check",
			method:               http.MethodPost,
		},
		{
			name:     "FalseCheckToken",
			inToken:  tokenTest,
			outCheck: false,
			isError:  false,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			name:     "ErrorExtractToken",
			inToken:  tokenTest,
			outCheck: true,
			isError:  true,
			url:      "http://token:8080/extract",
			method:   http.MethodPost,
		},
		{
			name:                 "ErrorInsideExtractToken",
			inToken:              tokenTest,
			outCheck:             true,
			isError:              true,
			isErrorInsideRequest: true,
			url:                  "http://token:8080/extract",
			method:               http.MethodPost,
		},
		{
			name:     "ErrorDeleteToken",
			inToken:  tokenTest,
			outCheck: true,
			isError:  true,
			url:      "http://db:8080/user",
			method:   http.MethodDelete,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			infoServiceTest := service.InfoServices{
				DBHost:    dbHostTest,
				DBPort:    portTest,
				TokenHost: tokenHostTest,
				TokenPort: portTest,
				Secret:    secretTest,
			}

			var resultErr error
			var errorResponse string
			var mock *service.MockClient

			if tt.isError {
				errorResponse = errWebServer.Error()
			}

			responseJSON := fmt.Sprintf(`{
					"user":{
						"username":"username",
						"password":"password",
						"email":"email@email.com",
						"id":1
					},
					"id":1,
					"username":"usename",
					"email":"email@email.com",
					"check":%t
				}`, tt.outCheck)

			if tt.isError {
				mock = service.NewMockClient(getIsErrorMock(
					tt.isErrorInsideRequest,
					newErrorHTTPComponets(tt.url, tt.method),
					responseJSON,
				))
			} else {
				mock = service.NewMockClient(getMock(
					responseJSON,
				))
			}

			svc := service.NewService(
				mock,
				&infoServiceTest,
			)

			resultErr = svc.DeleteAccount(tt.inToken)

			if !tt.isError {
				if tt.outCheck {
					assert.Nil(t, resultErr)
				} else {
					assert.ErrorContains(t, resultErr, errorResponse)
				}
			} else {
				assert.ErrorContains(t, resultErr, errorResponse)
			}
		})
	}
}

func getMock(jsonResponse string) func(*http.Request) (*http.Response, error) {
	return func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			Body: io.NopCloser(strings.NewReader(jsonResponse)),
		}, nil
	}
}

//nolint:revive
func getIsErrorMock(
	isErrorInsideRequest bool,
	errorHTTPComponents errorHTTPComponents,
	jsonResponse string,
) func(*http.Request) (*http.Response, error) {
	return func(r *http.Request) (*http.Response, error) {
		if r.URL.String() == errorHTTPComponents.errorURL && r.Method == errorHTTPComponents.errorMethod {
			if isErrorInsideRequest {
				return &http.Response{
					Body: io.NopCloser(bytes.NewReader([]byte(`{
										"err":"error"
									}`),
					)),
				}, nil
			}

			return nil, errWebServer
		}

		return &http.Response{
			Body: io.NopCloser(strings.NewReader(jsonResponse)),
		}, nil
	}
}
