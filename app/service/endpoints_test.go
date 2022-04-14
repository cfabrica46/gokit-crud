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
		in       service.UEP
		outToken string
		outErr   string
		isError  bool
	}{
		{
			in: service.UEP{
				Username: usernameTest,
				Password: passwordTest,
				Email:    emailTest,
			},
			outToken: tokenTest,
			outErr:   "",
			isError:  false,
		},
		{
			in:       service.UEP{},
			outToken: "",
			outErr:   errWebServer.Error(),
			isError:  true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			testResp := struct {
				ID    int    `json:"id"`
				Token string `json:"token"`
				Err   string `json:"err"`
			}{
				ID:    idTest,
				Token: table.outToken,
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

			result, ok := r.(service.TokenErr)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			if !table.isError {
				assert.Zero(t, result.Err)
			} else {
				assert.Contains(t, result.Err, table.outErr)
			}

			assert.Equal(t, table.outToken, result.Token)
		})
	}
}

func TestSignInEndpoint(t *testing.T) {
	for index, table := range []struct {
		in       service.UP
		outToken string
		outErr   string
		isError  bool
	}{
		{
			in: service.UP{
				Username: usernameTest,
				Password: passwordTest,
			},
			outToken: tokenTest,
			outErr:   "",
			isError:  false,
		},
		{
			in:       service.UP{},
			outToken: "",
			outErr:   errWebServer.Error(),
			isError:  true,
		},
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
				Token: table.outToken,
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

			result, ok := r.(service.TokenErr)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			if !table.isError {
				assert.Zero(t, result.Err)
			} else {
				assert.Contains(t, result.Err, table.outErr)
			}

			assert.Equal(t, table.outToken, result.Token)
		})
	}
}

func TestLogOutEndpoint(t *testing.T) {
	for index, table := range []struct {
		in      service.Token
		outErr  string
		isError bool
	}{
		{
			in: service.Token{
				Token: tokenTest,
			},
			outErr:  "",
			isError: false,
		},
		{
			in:      service.Token{},
			outErr:  errWebServer.Error(),
			isError: true,
		},
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

			result, ok := r.(service.Err)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			if !table.isError {
				assert.Zero(t, result.Err)
			} else {
				assert.Contains(t, result.Err, table.outErr)
			}
		})
	}
}

func TestGetAllUsersEndpoint(t *testing.T) {
	for index, table := range []struct {
		outUsers []dbapp.User
		outErr   string
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
			outErr:  "",
			isError: false,
		},
		{
			outUsers: nil,
			outErr:   errWebServer.Error(),
			isError:  true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			testResp := struct {
				Users []dbapp.User `json:"users"`
				Err   string       `json:"err"`
			}{
				Users: table.outUsers,
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

			r, err := service.MakeGetAllUsersEndpoint(svc)(context.TODO(), nil)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.UsersErr)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			if !table.isError {
				assert.Zero(t, result.Err)
			} else {
				assert.Contains(t, result.Err, table.outErr)
			}

			assert.Equal(t, table.outUsers, result.Users)
		})
	}
}

func TestProfileEndpoint(t *testing.T) {
	for index, table := range []struct {
		in      service.Token
		outUser dbapp.User
		outErr  string
		isError bool
	}{
		{
			in: service.Token{
				Token: tokenTest,
			},
			outUser: dbapp.User{
				ID:       idTest,
				Username: usernameTest,
				Password: passwordTest,
				Email:    emailTest,
			},
			outErr:  "",
			isError: false,
		},
		{
			in:      service.Token{},
			outUser: dbapp.User{},
			outErr:  errWebServer.Error(),
			isError: true,
		},
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
				User:     table.outUser,
				ID:       table.outUser.ID,
				Username: table.outUser.Username,
				Email:    table.outUser.Email,
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

			result, ok := r.(service.UserErr)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			if !table.isError {
				assert.Zero(t, result.Err)
			} else {
				assert.Contains(t, result.Err, table.outErr)
			}

			assert.Equal(t, table.outUser, result.User)
		})
	}
}

func TestDeleteAccountEndpoint(t *testing.T) {
	for index, table := range []struct {
		in      service.Token
		outErr  string
		isError bool
	}{
		{
			in: service.Token{
				Token: tokenTest,
			},
			outErr:  "",
			isError: false,
		},
		{
			in:      service.Token{},
			outErr:  errWebServer.Error(),
			isError: true,
		},
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

			result, ok := r.(service.Err)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			if !table.isError {
				assert.Zero(t, result.Err)
			} else {
				assert.Contains(t, result.Err, table.outErr)
			}
		})
	}
}
