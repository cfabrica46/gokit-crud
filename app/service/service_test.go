package service_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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
)

var (
	errWebServer        = errors.New("error from web server")
	errNotTypeIndicated = errors.New("response is not of the type indicated")
)

func TestSignUp(t *testing.T) {
	t.Parallel()

	infoServiceTest := service.InfoServices{
		DBHost:    dbHostTest,
		DBPort:    portTest,
		TokenHost: tokenHostTest,
		TokenPort: portTest,
		Secret:    secretTest,
	}

	for _, tt := range []struct {
		name                            string
		inUsername, inPassword, inEmail string
		url                             string
		method                          string
		isError                         bool
	}{
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			isError:    false,
			url:        "http://token:8080/generate",
			method:     http.MethodPost,
		},
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			isError:    true,
			url:        "http://db:8080/user",
			method:     http.MethodPost,
		},
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			isError:    true,
			url:        "http://db:8080/id/username",
			method:     http.MethodGet,
		},
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			isError:    true,
			url:        "http://token:8080/generate",
			method:     http.MethodPost,
		},
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			isError:    true,
			url:        "http://token:8080/token",
			method:     http.MethodPost,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultToken string
			var resultErr error
			var tokenResponse, errorResponse string

			if tt.isError {
				errorResponse = errWebServer.Error()
			} else {
				tokenResponse = tokenTest
			}

			testResp := struct {
				Token string `json:"token"`
				Err   string `json:"err"`
				ID    int    `json:"id"`
			}{
				ID:    idTest,
				Token: tokenResponse,
				Err:   errorResponse,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				response := &http.Response{Body: io.NopCloser(bytes.NewReader([]byte("{}")))}

				if req.URL.String() == tt.url {
					if req.Method == tt.method {
						response = &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(bytes.NewReader(jsonData)),
						}
					}
				}

				return response, nil
			})

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
	}{
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			isError:    false,
			url:        "http://token:8080/generate",
			method:     http.MethodPost,
		},
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			isError:    true,
			url:        "http://db:8080/user/username_password",
			method:     http.MethodGet,
		},
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			isError:    true,
			url:        "http://token:8080/generate",
			method:     http.MethodPost,
		},
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			isError:    true,
			url:        "http://token:8080/token",
			method:     http.MethodPost,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultToken string
			var resultErr error
			var tokenResponse, errorResponse string

			if tt.isError {
				errorResponse = errWebServer.Error()
			} else {
				tokenResponse = tokenTest
			}

			testResp := struct {
				User  dbapp.User `json:"user"`
				Token string     `json:"token"`
				Err   string     `json:"err"`
			}{
				User: dbapp.User{
					ID:       idTest,
					Username: usernameTest,
					Password: passwordTest,
					Email:    emailTest,
				},
				Token: tokenResponse,
				Err:   errorResponse,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				response := &http.Response{Body: io.NopCloser(bytes.NewReader([]byte("{}")))}

				if req.URL.String() == tt.url {
					if req.Method == tt.method {
						response = &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(bytes.NewReader(jsonData)),
						}
					}
				}

				return response, nil
			})

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
		name     string
		inToken  string
		url      string
		method   string
		outCheck bool
		isError  bool
	}{
		{
			inToken:  tokenTest,
			outCheck: true,
			isError:  false,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			inToken:  tokenTest,
			outCheck: true,
			isError:  true,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			inToken:  tokenTest,
			outCheck: false,
			isError:  true,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			inToken:  tokenTest,
			outCheck: true,
			isError:  true,
			url:      "http://token:8080/token",
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

			if tt.isError {
				errorResponse = errWebServer.Error()
			}

			testResp := struct {
				Err   string `json:"err"`
				Check bool   `json:"check"`
			}{
				Check: tt.outCheck,
				Err:   errorResponse,
			}

			if !tt.outCheck {
				testResp.Err = ""
				errorResponse = service.ErrTokenNotValid.Error()
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				testCheck := struct {
					Check bool `json:"check"`
				}{
					Check: tt.outCheck,
				}

				jsonCheck, err := json.Marshal(testCheck)
				if err != nil {
					t.Error(err)
				}

				response := &http.Response{Body: io.NopCloser(bytes.NewReader(jsonCheck))}

				if req.URL.String() == tt.url {
					if req.Method == tt.method {
						response = &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(bytes.NewReader(jsonData)),
						}
					}
				}

				return response, nil
			})

			svc := service.NewService(
				mock,
				&infoServiceTest,
			)

			resultErr = svc.LogOut(tt.inToken)

			if !tt.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, errorResponse)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	t.Parallel()

	log.SetFlags(log.Lshortfile)

	for _, tt := range []struct {
		name     string
		url      string
		method   string
		outUsers []dbapp.User
		isError  bool
	}{
		{
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
			outUsers: nil,
			isError:  true,
			url:      "http://db:8080/users",
			method:   http.MethodGet,
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

			if tt.isError {
				errorResponse = errWebServer.Error()
			}

			testResp := struct {
				Err   string       `json:"err"`
				Users []dbapp.User `json:"users"`
			}{
				Users: []dbapp.User{{
					ID:       idTest,
					Username: usernameTest,
					Password: passwordTest,
					Email:    emailTest,
				}},
				Err: errorResponse,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(jsonData)),
				}, nil
			})

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

	for _, tt := range []struct {
		name     string
		inToken  string
		url      string
		method   string
		outUser  dbapp.User
		outCheck bool
		isError  bool
	}{
		{
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
			inToken:  tokenTest,
			outUser:  dbapp.User{},
			outCheck: true,
			isError:  true,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			inToken:  tokenTest,
			outUser:  dbapp.User{},
			outCheck: false,
			isError:  true,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			inToken:  tokenTest,
			outUser:  dbapp.User{},
			outCheck: true,
			isError:  true,
			url:      "http://token:8080/extract",
			method:   http.MethodPost,
		},
		{
			inToken:  tokenTest,
			outUser:  dbapp.User{},
			outCheck: true,
			isError:  true,
			url:      "http://db:8080/user/id",
			method:   http.MethodGet,
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

			var resultUser dbapp.User
			var resultErr error
			var errorResponse string

			if tt.isError {
				errorResponse = errWebServer.Error()
			}

			testResp := struct {
				Username string     `json:"username"`
				Email    string     `json:"email"`
				Err      string     `json:"err"`
				User     dbapp.User `json:"user"`
				ID       int        `json:"id"`
				Check    bool       `json:"check"`
			}{
				User:     tt.outUser,
				ID:       tt.outUser.ID,
				Username: tt.outUser.Username,
				Email:    tt.outUser.Email,
				Check:    tt.outCheck,
				Err:      errorResponse,
			}

			if !tt.outCheck {
				testResp.Err = ""
				errorResponse = service.ErrTokenNotValid.Error()
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				testCheck := struct {
					Check bool `json:"check"`
				}{
					Check: tt.outCheck,
				}

				jsonCheck, err := json.Marshal(testCheck)
				if err != nil {
					t.Error(err)
				}

				response := &http.Response{Body: io.NopCloser(bytes.NewReader(jsonCheck))}

				if req.URL.String() == tt.url {
					if req.Method == tt.method {
						response = &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(bytes.NewReader((jsonData))),
						}
					}
				}

				return response, nil
			})

			svc := service.NewService(
				mock,
				&infoServiceTest,
			)

			resultUser, resultErr = svc.Profile(tt.inToken)

			if !tt.isError {
				assert.Nil(t, resultErr)
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
		name     string
		inToken  string
		url      string
		method   string
		outCheck bool
		isError  bool
	}{
		{
			inToken:  tokenTest,
			outCheck: true,
			isError:  false,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			inToken:  tokenTest,
			outCheck: true,
			isError:  true,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},

		{
			inToken:  tokenTest,
			outCheck: false,
			isError:  true,
			url:      "http://token:8080/check",
			method:   http.MethodPost,
		},
		{
			inToken:  tokenTest,
			outCheck: true,
			isError:  true,
			url:      "http://token:8080/extract",
			method:   http.MethodPost,
		},
		{
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

			if tt.isError {
				errorResponse = errWebServer.Error()
			}

			testResp := struct {
				Username string     `json:"username"`
				Email    string     `json:"email"`
				Err      string     `json:"err"`
				User     dbapp.User `json:"user"`
				Check    bool       `json:"check"`
				ID       int        `json:"id"`
			}{
				User: dbapp.User{
					ID:       idTest,
					Username: usernameTest,
					Password: passwordTest,
					Email:    emailTest,
				},
				ID:       idTest,
				Username: usernameTest,
				Email:    emailTest,
				Check:    tt.outCheck,
				Err:      errorResponse,
			}

			if !tt.outCheck {
				testResp.Err = ""
				errorResponse = service.ErrTokenNotValid.Error()
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				testCheck := struct {
					Check bool `json:"check"`
				}{
					Check: tt.outCheck,
				}

				jsonCheck, err := json.Marshal(testCheck)
				if err != nil {
					t.Error(err)
				}

				response := &http.Response{Body: io.NopCloser(bytes.NewReader(jsonCheck))}

				if req.URL.String() == tt.url {
					if req.Method == tt.method {
						response = &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(bytes.NewReader(jsonData)),
						}
					}
				}

				return response, nil
			})

			svc := service.NewService(
				mock,
				&infoServiceTest,
			)

			resultErr = svc.DeleteAccount(tt.inToken)

			if !tt.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, errorResponse)
			}
		})
	}
}
