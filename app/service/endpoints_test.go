package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	"github.com/stretchr/testify/assert"
)

func TestSignUpEndpoint(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		in     SignUpRequest
		outErr string
	}{
		{
			SignUpRequest{
				Username: userTest.Username,
				Password: userTest.Password,
				Email:    userTest.Email,
			},
			"",
		},
		{
			SignUpRequest{},
			errWebServer.Error(),
		},
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

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonData))),
				}, nil
			})

			svc := NewService(mock, "localhost", "8080", "localhost", "8080", "secret")

			r, err := MakeSignUpEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(SignUpResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestSignInEndpoint(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		in     SignInRequest
		outErr string
	}{
		{SignInRequest{Username: userTest.Username, Password: userTest.Password}, ""},
		{SignInRequest{}, errWebServer.Error()},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			testResp := struct {
				User  dbapp.User
				Token string `json:"token"`
				Err   string `json:"err"`
			}{
				User: dbapp.User{
					ID:       1,
					Username: userTest.Username,
					Password: userTest.Password,
					Email:    userTest.Email,
				},
				Token: "token",
				Err:   tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonData))),
				}, nil
			})

			svc := NewService(mock, "localhost", "8080", "localhost", "8080", "secret")

			r, err := MakeSignInEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(SignInResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestLogOutEndpoint(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		in     LogOutRequest
		outErr string
	}{
		{LogOutRequest{Token: "token"}, ""},
		{LogOutRequest{}, errWebServer.Error()},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			testResp := struct {
				Check bool   `json:"check"`
				Err   string `json:"err"`
			}{
				Check: true,
				Err:   tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonData))),
				}, nil
			})

			svc := NewService(mock, "localhost", "8080", "localhost", "8080", "secret")

			r, err := MakeLogOutEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(LogOutResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestGetAllUsersEndpoint(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		in     GetAllUsersRequest
		outErr string
	}{
		{GetAllUsersRequest{}, ""},
		{GetAllUsersRequest{}, errWebServer.Error()},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			testResp := struct {
				User []dbapp.User `json:"user"`
				Err  string       `json:"err"`
			}{
				User: []dbapp.User{{
					ID:       1,
					Username: userTest.Username,
					Password: userTest.Password,
					Email:    userTest.Email,
				}},
				Err: tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonData))),
				}, nil
			})

			svc := NewService(mock, "localhost", "8080", "localhost", "8080", "secret")

			r, err := MakeGetAllUsersEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetAllUsersResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestProfileEndpoint(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		in     ProfileRequest
		outErr string
	}{
		{ProfileRequest{}, ""},
		{ProfileRequest{}, errWebServer.Error()},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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
					Username: userTest.Username,
					Password: userTest.Password,
					Email:    userTest.Email,
				},
				ID:       1,
				Username: userTest.Username,
				Email:    userTest.Email,
				Check:    true,
				Err:      tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonData))),
				}, nil
			})

			svc := NewService(mock, "localhost", "8080", "localhost", "8080", "secret")

			r, err := MakeProfileEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(ProfileResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestDeleteAccountEndpoint(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	for i, tt := range []struct {
		in     DeleteAccountRequest
		outErr string
	}{
		{DeleteAccountRequest{}, ""},
		{DeleteAccountRequest{}, errWebServer.Error()},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			testResp := struct {
				ID       int    `json:"id"`
				Username string `json:"username"`
				Email    string `json:"email"`
				Check    bool   `json:"check"`
				Err      string `json:"err"`
			}{
				ID:       1,
				Username: userTest.Username,
				Email:    userTest.Email,
				Check:    true,
				Err:      tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonData))),
				}, nil
			})

			svc := NewService(mock, "localhost", "8080", "localhost", "8080", "secret")

			r, err := MakeDeleteAccountEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(DeleteAccountResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}
