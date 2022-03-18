package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		inUsername, inPassword, inEmail string
		isError                         bool
		url                             string
		method                          string
	}{
		{"cesar", "01234", "cesar@email.com", false, "", ""},
		{"cesar", "01234", "cesar@email.com", true, "http://db:8080/user", http.MethodPost},
		{"cesar", "01234", "cesar@email.com", true, "http://db:8080/id/username", http.MethodGet},
		{"cesar", "01234", "cesar@email.com", true, "http://token:8080/generate", http.MethodPost},
		{"cesar", "01234", "cesar@email.com", true, "http://token:8080/token", http.MethodPost},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string
			var errorMessage string

			if tt.isError {
				errorMessage = "Error from web server"
			}

			testResp := struct {
				ID    int    `json:"id"`
				Token string `json:"token"`
				Err   string `json:"err"`
			}{
				ID:    1,
				Token: "token",
				Err:   errorMessage,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				response := &http.Response{Body: io.NopCloser(bytes.NewReader([]byte("{}")))}

				// log.Println(i, req.URL.String(), tt.url, req.URL.String() == tt.url)

				if req.URL.String() == tt.url {
					switch req.Method {
					case tt.method:
						response = &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(jsonData)))}
					}
				}

				return response, nil
			})

			svc := GetService(mock, "db", "8080", "token", "8080", "secret")

			_, err = svc.SignUp(tt.inUsername, tt.inPassword, tt.inEmail)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, errorMessage, resultErr)
		})
	}
}

func TestSignIn(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		inUsername, inPassword string
		isError                bool
		url                    string
		method                 string
	}{
		{"cesar", "01234", false, "", ""},
		{"cesar", "01234", true, "http://db:8080/user/username_password", http.MethodGet},
		{"cesar", "01234", true, "http://token:8080/generate", http.MethodPost},
		{"cesar", "01234", true, "http://token:8080/token", http.MethodPost},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string
			var errorMessage string

			if tt.isError {
				errorMessage = "Error from web server"
			}

			testResp := struct {
				User  dbapp.User
				Token string `json:"token"`
				Err   string `json:"err"`
			}{
				User: dbapp.User{
					ID:       1,
					Username: "cesar",
					Password: "01234",
					Email:    "cesar@email.com",
				},
				Token: "token",
				Err:   errorMessage,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				response := &http.Response{Body: io.NopCloser(bytes.NewReader([]byte("{}")))}

				if req.URL.String() == tt.url {
					switch req.Method {
					case tt.method:
						response = &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(jsonData)))}
					}
				}

				return response, nil
			})

			svc := GetService(mock, "db", "8080", "token", "8080", "secret")

			_, err = svc.SignIn(tt.inUsername, tt.inPassword)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, errorMessage, resultErr)
		})
	}
}

func TestLogOut(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		inToken  string
		outCheck bool

		isError bool
		url     string
		method  string
	}{
		{"token", true, false, "http://token:8080/check", http.MethodPost},
		{"token", true, true, "http://token:8080/check", http.MethodPost},
		{"token", false, true, "http://token:8080/check", http.MethodPost},
		{"token", true, true, "http://token:8080/token", http.MethodDelete},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string
			var errorMessage string

			if tt.isError {
				errorMessage = "Error from web server"
			}

			testResp := struct {
				Check bool   `json:"check"`
				Err   string `json:"err"`
			}{
				Check: tt.outCheck,
				Err:   errorMessage,
			}

			if !tt.outCheck {
				testResp.Err = ""
				errorMessage = "token not validate"
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
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

				log.Println(i, req.URL.String(), tt.url, req.Method, tt.method)

				if req.URL.String() == tt.url {
					switch req.Method {
					case tt.method:
						response = &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(jsonData)))}
					}
				}

				return response, nil
			})

			svc := GetService(mock, "db", "8080", "token", "8080", "secret")

			err = svc.LogOut(tt.inToken)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, errorMessage, resultErr)
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	for i, tt := range []struct {
		outErr string
	}{
		{""},
		{"Error from web server"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			testResp := struct {
				User []dbapp.User `json:"user"`
				Err  string       `json:"err"`
			}{
				User: []dbapp.User{{
					ID:       1,
					Username: "cesar",
					Password: "01234",
					Email:    "cesar@email.com",
				}},
				Err: tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))}, nil
			})

			svc := GetService(mock, "localhost", "8080", "localhost", "8080", "secret")

			_, err = svc.GetAllUsers()
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}

func TestProfile(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		inToken  string
		outCheck bool

		isError bool
		url     string
		method  string
	}{
		{"token", true, false, "http://token:8080/check", http.MethodPost},
		{"token", true, true, "http://token:8080/check", http.MethodPost},
		{"token", false, true, "http://token:8080/check", http.MethodPost},
		{"token", true, true, "http://token:8080/extract", http.MethodPost},
		{"token", true, true, "http://db:8080/user/id", http.MethodGet},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string
			var errorMessage string

			if tt.isError {
				errorMessage = "Error from web server"
			}

			testResp := struct {
				User     dbapp.User `json:"user"`
				ID       int        `json:"id"`
				Username string     `json:"username"`
				Email    string     `json:"email"`
				Check    bool       `json:"check"`
				Err      string     `json:"err"`
			}{
				User: dbapp.User{
					ID:       1,
					Username: "cesar",
					Password: "01234",
					Email:    "cesar@email.com",
				},
				ID:       1,
				Username: "cesar",
				Email:    "cesar@email.com",
				Check:    tt.outCheck,
				Err:      errorMessage,
			}

			if !tt.outCheck {
				testResp.Err = ""
				errorMessage = "token not validate"
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
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

				log.Println(i, req.URL.String(), tt.url, req.Method, tt.method)

				if req.URL.String() == tt.url {
					switch req.Method {
					case tt.method:
						response = &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(jsonData)))}
					}
				}

				return response, nil
			})

			svc := GetService(mock, "db", "8080", "token", "8080", "secret")

			_, err = svc.Profile(tt.inToken)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, errorMessage, resultErr)
		})
	}
}

func TestDeleteAccount(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		inToken  string
		outCheck bool

		isError bool
		url     string
		method  string
	}{
		{"token", true, false, "http://token:8080/check", http.MethodPost},
		{"token", true, true, "http://token:8080/check", http.MethodPost},
		{"token", false, true, "http://token:8080/check", http.MethodPost},
		{"token", true, true, "http://token:8080/extract", http.MethodPost},
		{"token", true, true, "http://db:8080/user", http.MethodDelete},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string
			var errorMessage string

			if tt.isError {
				errorMessage = "Error from web server"
			}

			testResp := struct {
				User     dbapp.User `json:"user"`
				ID       int        `json:"id"`
				Username string     `json:"username"`
				Email    string     `json:"email"`
				Check    bool       `json:"check"`
				Err      string     `json:"err"`
			}{
				User: dbapp.User{
					ID:       1,
					Username: "cesar",
					Password: "01234",
					Email:    "cesar@email.com",
				},
				ID:       1,
				Username: "cesar",
				Email:    "cesar@email.com",
				Check:    tt.outCheck,
				Err:      errorMessage,
			}

			if !tt.outCheck {
				testResp.Err = ""
				errorMessage = "token not validate"
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
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

				log.Println(i, req.URL.String(), tt.url, req.Method, tt.method)

				if req.URL.String() == tt.url {
					switch req.Method {
					case tt.method:
						response = &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(jsonData)))}
					}
				}

				return response, nil
			})

			svc := GetService(mock, "db", "8080", "token", "8080", "secret")

			err = svc.DeleteAccount(tt.inToken)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, errorMessage, resultErr)
		})
	}
}
