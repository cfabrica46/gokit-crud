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
		in       service.SignUpRequest
		outToken string
		outErr   string
		isError  bool
	}{
		{
			in: service.SignUpRequest{
				Username: usernameTest,
				Password: passwordTest,
				Email:    emailTest,
			},
			outToken: tokenTest,
			outErr:   "",
			isError:  false,
		},
		{
			in:       service.SignUpRequest{},
			outToken: "",
			outErr:   errWebServer.Error(),
			isError:  true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			testResp := struct {
				Token string `json:"token"`
				Err   string `json:"err"`
				ID    int    `json:"id"`
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

			result, ok := r.(service.SignUpResponse)
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
		in       service.SignInRequest
		outToken string
		outErr   string
		isError  bool
	}{
		{
			in: service.SignInRequest{
				Username: usernameTest,
				Password: passwordTest,
			},
			outToken: tokenTest,
			outErr:   "",
			isError:  false,
		},
		{
			in:       service.SignInRequest{},
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

			result, ok := r.(service.SignInResponse)
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
		in      service.LogOutRequest
		outErr  string
		isError bool
	}{
		{
			in: service.LogOutRequest{
				Token: tokenTest,
			},
			outErr:  "",
			isError: false,
		},
		{
			in:      service.LogOutRequest{},
			outErr:  errWebServer.Error(),
			isError: true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			testResp := struct {
				Err   string `json:"err"`
				Check bool   `json:"check"`
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
		outErr   string
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
				Err   string       `json:"err"`
				Users []dbapp.User `json:"users"`
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

			result, ok := r.(service.GetAllUsersResponse)
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
		in      service.ProfileRequest
		outUser dbapp.User
		outErr  string
		isError bool
	}{
		{
			in: service.ProfileRequest{
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
			in:      service.ProfileRequest{},
			outUser: dbapp.User{},
			outErr:  errWebServer.Error(),
			isError: true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			testResp := struct {
				Username string     `json:"username"`
				Email    string     `json:"email"`
				Err      string     `json:"err"`
				User     dbapp.User `json:"user"`
				ID       int        `json:"id"`
				Check    bool       `json:"check"`
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

			result, ok := r.(service.ProfileResponse)
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
		in      service.DeleteAccountRequest
		outErr  string
		isError bool
	}{
		{
			in: service.DeleteAccountRequest{
				Token: tokenTest,
			},
			outErr:  "",
			isError: false,
		},
		{
			in:      service.DeleteAccountRequest{},
			outErr:  errWebServer.Error(),
			isError: true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			testResp := struct {
				Username string `json:"username"`
				Email    string `json:"email"`
				Err      string `json:"err"`
				ID       int    `json:"id"`
				Check    bool   `json:"check"`
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

			if !table.isError {
				assert.Zero(t, result.Err)
			} else {
				assert.Contains(t, result.Err, table.outErr)
			}
		})
	}
}

// func TestDeleteAccountEndpoint(t *testing.T) {
// 	for index, table := range []struct {
// 		in     service.DeleteAccountRequest
// 		outErr string
// 	}{
// 		{service.DeleteAccountRequest{}, ""},
// 		{service.DeleteAccountRequest{}, errWebServer.Error()},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			testResp := struct {
// 				ID       int    `json:"id"`
// 				Username string `json:"username"`
// 				Email    string `json:"email"`
// 				Check    bool   `json:"check"`
// 				Err      string `json:"err"`
// 			}{
// 				ID:       idTest,
// 				Username: usernameTest,
// 				Email:    emailTest,
// 				Check:    true,
// 				Err:      table.outErr,
// 			}

// 			jsonData, err := json.Marshal(testResp)
// 			if err != nil {
// 				t.Error(err)
// 			}

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				return &http.Response{
// 					StatusCode: http.StatusOK,
// 					Body:       ioutil.NopCloser(bytes.NewReader(jsonData)),
// 				}, nil
// 			})

// 			svc := service.NewService(
// 				mock,
// 				dbHostTest,
// 				portTest,
// 				tokenHostTest,
// 				portTest,
// 				secretTest,
// 			)

// 			r, err := service.MakeDeleteAccountEndpoint(svc)(context.TODO(), table.in)
// 			if err != nil {
// 				t.Error(err)
// 			}

// 			result, ok := r.(service.DeleteAccountResponse)
// 			if !ok {
// 				t.Error(errNotTypeIndicated)
// 			}

// 			assert.Equal(t, table.outErr, result.Err)
// 		})
// 	}
// }
