package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/cfabrica46/gokit-crud/app/service"
	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	"github.com/stretchr/testify/assert"
)

func TestSignUpEndpoint(t *testing.T) {
	for index, table := range []struct {
		in     service.SignUpRequest
		outErr string
	}{
		{
			service.SignUpRequest{
				Username: usernameTest,
				Password: passwordTest,
				Email:    emailTest,
			},
			"",
		},
		{
			service.SignUpRequest{},
			errWebServer.Error(),
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			testResp := struct {
				ID    int    `json:"id"`
				Token string `json:"token"`
				Err   string `json:"err"`
			}{
				ID:    idTest,
				Token: tokenTest,
				Err:   table.outErr,
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
				dbHostTest,
				portTest,
				tokenHostTest,
				portTest,
				secretTest,
			)

			r, err := service.MakeSignUpEndpoint(svc)(context.TODO(), table.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.SignUpResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			assert.Equal(t, table.outErr, result.Err)
		})
	}
}

func TestSignInEndpoint(t *testing.T) {
	for index, table := range []struct {
		in     service.SignInRequest
		outErr string
	}{
		{service.SignInRequest{Username: usernameTest, Password: passwordTest}, ""},
		{service.SignInRequest{}, errWebServer.Error()},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			testResp := struct {
				User  dbapp.User
				Token string `json:"token"`
				Err   string `json:"err"`
			}{
				User: dbapp.User{
					ID:       idTest,
					Username: usernameTest,
					Password: passwordTest,
					Email:    emailTest,
				},
				Token: tokenTest,
				Err:   table.outErr,
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
				dbHostTest,
				portTest,
				tokenHostTest,
				portTest,
				secretTest,
			)

			r, err := service.MakeSignInEndpoint(svc)(context.TODO(), table.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.SignInResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			assert.Equal(t, table.outErr, result.Err)
		})
	}
}

func TestLogOutEndpoint(t *testing.T) {
	for index, table := range []struct {
		in     service.LogOutRequest
		outErr string
	}{
		{service.LogOutRequest{Token: tokenTest}, ""},
		{service.LogOutRequest{}, errWebServer.Error()},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			testResp := struct {
				Check bool   `json:"check"`
				Err   string `json:"err"`
			}{
				Check: true,
				Err:   table.outErr,
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
				dbHostTest,
				portTest,
				tokenHostTest,
				portTest,
				secretTest,
			)

			r, err := service.MakeLogOutEndpoint(svc)(context.TODO(), table.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.LogOutResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			assert.Equal(t, table.outErr, result.Err)
		})
	}
}

func TestGetAllUsersEndpoint(t *testing.T) {
	for index, table := range []struct {
		in     service.GetAllUsersRequest
		outErr string
	}{
		{service.GetAllUsersRequest{}, ""},
		{service.GetAllUsersRequest{}, errWebServer.Error()},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			testResp := struct {
				User []dbapp.User `json:"user"`
				Err  string       `json:"err"`
			}{
				User: []dbapp.User{{
					ID:       idTest,
					Username: usernameTest,
					Password: passwordTest,
					Email:    emailTest,
				}},
				Err: table.outErr,
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
				dbHostTest,
				portTest,
				tokenHostTest,
				portTest,
				secretTest,
			)

			r, err := service.MakeGetAllUsersEndpoint(svc)(context.TODO(), table.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.GetAllUsersResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			assert.Equal(t, table.outErr, result.Err, "they should be equal")
		})
	}
}

func TestProfileEndpoint(t *testing.T) {
	for index, table := range []struct {
		in     service.ProfileRequest
		outErr string
	}{
		{service.ProfileRequest{}, ""},
		{service.ProfileRequest{}, errWebServer.Error()},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			testResp := struct {
				User     dbapp.User `json:"user"`
				ID       int        `json:"id"`
				Username string     `json:"username"`
				Email    string     `json:"email"`
				Check    bool       `json:"check"`
				Err      string     `json:"err"`
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
				Check:    true,
				Err:      table.outErr,
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
				dbHostTest,
				portTest,
				tokenHostTest,
				portTest,
				secretTest,
			)

			r, err := service.MakeProfileEndpoint(svc)(context.TODO(), table.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.ProfileResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			assert.Equal(t, table.outErr, result.Err)
		})
	}
}

func TestDeleteAccountEndpoint(t *testing.T) {
	for index, table := range []struct {
		in     service.DeleteAccountRequest
		outErr string
	}{
		{service.DeleteAccountRequest{}, ""},
		{service.DeleteAccountRequest{}, errWebServer.Error()},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			testResp := struct {
				ID       int    `json:"id"`
				Username string `json:"username"`
				Email    string `json:"email"`
				Check    bool   `json:"check"`
				Err      string `json:"err"`
			}{
				ID:       idTest,
				Username: usernameTest,
				Email:    emailTest,
				Check:    true,
				Err:      table.outErr,
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
				dbHostTest,
				portTest,
				tokenHostTest,
				portTest,
				secretTest,
			)

			r, err := service.MakeDeleteAccountEndpoint(svc)(context.TODO(), table.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.DeleteAccountResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			assert.Equal(t, table.outErr, result.Err)
		})
	}
}
