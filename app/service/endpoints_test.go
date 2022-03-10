package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	"github.com/stretchr/testify/assert"
)

func TestSignUpEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     SignUpRequest
		outErr string
	}{
		{SignUpRequest{Username: "cesar", Password: "01234", Email: "cesar@email.com"}, ""},
		{SignUpRequest{}, "Error from web server"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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

			r, err := MakeSignUpEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(SignUpResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestSignInEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     SignInRequest
		outErr string
	}{
		{SignInRequest{Username: "cesar", Password: "01234"}, ""},
		{SignInRequest{}, "Error from web server"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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

			r, err := MakeSignInEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(SignInResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestLogOutEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     LogOutRequest
		outErr string
	}{
		{LogOutRequest{Token: "token"}, ""},
		{LogOutRequest{}, "Error from web server"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			testResp := struct {
				Err string `json:"err"`
			}{
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

			r, err := MakeLogOutEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(LogOutResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestGetAllUsersEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     GetAllUsersRequest
		outErr string
	}{
		{GetAllUsersRequest{}, ""},
		{GetAllUsersRequest{}, "Error from web server"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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

			r, err := MakeGetAllUsersEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetAllUsersResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestProfileEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     ProfileRequest
		outErr string
	}{
		{ProfileRequest{}, ""},
		{ProfileRequest{}, "Error from web server"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			testResp := struct {
				User     dbapp.User `json:"user"`
				ID       int        `json:"id"`
				Username string     `json:"username"`
				Email    string     `json:"email"`
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
				Err:      tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))}, nil
			})

			svc := GetService("localhost", "8080", "localhost", "8080", "secret", mock)

			r, err := MakeProfileEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(ProfileResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestDeleteAccountEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     DeleteAccountRequest
		outErr string
	}{
		{DeleteAccountRequest{}, ""},
		{DeleteAccountRequest{}, "Error from web server"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			testResp := struct {
				ID       int    `json:"id"`
				Username string `json:"username"`
				Email    string `json:"email"`
				Err      string `json:"err"`
			}{
				ID:       1,
				Username: "cesar",
				Email:    "cesar@email.com",
				Err:      tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))}, nil
			})

			svc := GetService("localhost", "8080", "localhost", "8080", "secret", mock)

			r, err := MakeDeleteAccountEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(DeleteAccountResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}
