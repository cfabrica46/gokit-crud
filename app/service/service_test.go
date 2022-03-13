package service

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		outErr                          string
	}{
		{"cesar", "01234", "cesar@email.com", ""},
		{"cesar", "01234", "cesar@email.com", "Error from web server"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			testResp := struct {
				ID    int    `json:"id"`
				Token string `json:"token"`
				Err   string `json:"err"`
			}{
				ID:    1,
				Token: "token",
				Err:   tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))}, nil
			})

			svc := GetService("localhost", "8080", "localhost", "8080", "secret", mock)

			_, err = svc.SignUp(tt.inUsername, tt.inPassword, tt.inEmail)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}

func TestSignIn(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		inUsername, inPassword string
		outErr                 string
	}{
		{"cesar", "01234", ""},
		{"cesar", "01234", "Error from web server"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

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
				Err:   tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))}, nil
			})

			svc := GetService("localhost", "8080", "localhost", "8080", "secret", mock)

			_, err = svc.SignIn(tt.inUsername, tt.inPassword)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}

func TestLogOut(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		inToken  string
		outCheck bool
		outErr   string
	}{
		{"token", true, ""},
		{"token", true, "Error from web server"},
		{"token", false, "token not validate"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			testResp := struct {
				Check bool   `json:"check"`
				Err   string `json:"err"`
			}{
				Check: tt.outCheck,
				Err:   tt.outErr,
			}

			if tt.outErr == "token not validate" {
				testResp.Err = ""
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))}, nil
			})

			svc := GetService("localhost", "8080", "localhost", "8080", "secret", mock)

			err = svc.LogOut(tt.inToken)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
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

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))}, nil
			})

			svc := GetService("localhost", "8080", "localhost", "8080", "secret", mock)

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
		outErr   string
	}{
		{"token", true, ""},
		{"token", true, "Error from web server"},
		{"token", false, "token not validate"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

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
				Err:      tt.outErr,
			}

			if tt.outErr == "token not validate" {
				testResp.Err = ""
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))}, nil
			})

			svc := GetService("localhost", "8080", "localhost", "8080", "secret", mock)

			_, err = svc.Profile(tt.inToken)
			if err != nil {
				resultErr = err.Error()
			}

			// if tt.outErr == "token not validate" {
			// 	testResp.Err = tt.outErr
			// }

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}

func TestDeleteAccount(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		inToken  string
		outCheck bool
		outErr   string
	}{
		{"token", true, ""},
		{"token", true, "Error from web server"},
		{"token", false, "token not validate"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			testResp := struct {
				ID       int    `json:"id"`
				Username string `json:"username"`
				Email    string `json:"email"`
				Check    bool   `json:"check"`
				Err      string `json:"err"`
			}{
				ID:       1,
				Username: "cesar",
				Email:    "cesar@email.com",
				Check:    tt.outCheck,
				Err:      tt.outErr,
			}

			if tt.outErr == "token not validate" {
				testResp.Err = ""
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))}, nil
			})

			svc := GetService("localhost", "8080", "localhost", "8080", "secret", mock)

			err = svc.DeleteAccount(tt.inToken)
			if err != nil {
				resultErr = err.Error()
			}

			// if tt.outErr == "token not validate" {
			// 	testResp.Err = tt.outErr
			// }

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}
