package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/cfabrica46/gokit-crud/app/service"
	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	"github.com/stretchr/testify/assert"
)

const (
	nameErrorRequest string = "ErrorRequest"
)

type incorrectRequest struct {
	incorrect bool
}

func TestSignUpEndpoint(t *testing.T) {
	t.Parallel()

	infoServiceTest := service.InfoServices{
		DBHost:    dbHostTest,
		DBPort:    portTest,
		TokenHost: tokenHostTest,
		TokenPort: portTest,
		Secret:    secretTest,
	}

	for _, tt := range []struct {
		name     string
		in       any
		outToken string
		outErr   string
	}{
		{
			name: nameNoError,
			in: service.UsernamePasswordEmailRequest{
				Username: usernameTest,
				Password: passwordTest,
				Email:    emailTest,
			},
			outToken: tokenTest,
			outErr:   "",
		},
		{
			name: nameErrorRequest,
			in: incorrectRequest{
				incorrect: true,
			},
			outErr: "isn't of type",
		},
		{
			name:     "ErrorWebService",
			in:       service.UsernamePasswordEmailRequest{},
			outToken: "",
			outErr:   errWebServer.Error(),
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			testResp := struct {
				Token string `json:"token"`
				Err   string `json:"err"`
				ID    int    `json:"id"`
			}{
				ID:    idTest,
				Token: tt.outToken,
				Err:   tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				assert.Error(t, err)
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

			r, err := service.MakeSignUpEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.TokenErrorResponse)
			if !ok {
				if tt.name != nameErrorRequest {
					assert.Error(t, errNotTypeIndicated)
				}
			}

			if result.Err != "" {
				resultErr = result.Err
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			assert.Equal(t, tt.outToken, result.Token)
		})
	}
}

func TestSignInEndpoint(t *testing.T) {
	t.Parallel()

	infoServiceTest := service.InfoServices{
		DBHost:    dbHostTest,
		DBPort:    portTest,
		TokenHost: tokenHostTest,
		TokenPort: portTest,
		Secret:    secretTest,
	}

	for _, tt := range []struct {
		name     string
		in       any
		outToken string
		outErr   string
	}{
		{
			name: nameNoError,
			in: service.UsernamePasswordRequest{
				Username: usernameTest,
				Password: passwordTest,
			},
			outToken: tokenTest,
			outErr:   "",
		},
		{
			name: nameErrorRequest,
			in: incorrectRequest{
				incorrect: true,
			},
			outErr: "isn't of type",
		},
		{
			name:     "ErrorWebService",
			in:       service.UsernamePasswordRequest{},
			outToken: "",
			outErr:   errWebServer.Error(),
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			testResp := struct {
				Token string `json:"token"`
				Err   string `json:"err"`
				User  dbapp.User
			}{
				User: dbapp.User{
					ID:       idTest,
					Username: usernameTest,
					Password: passwordTest,
					Email:    emailTest,
				},
				Token: tt.outToken,
				Err:   tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				assert.Error(t, err)
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

			r, err := service.MakeSignInEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.TokenErrorResponse)
			if !ok {
				if tt.name != nameErrorRequest {
					assert.Error(t, errNotTypeIndicated)
				}
			}

			if result.Err != "" {
				resultErr = result.Err
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			assert.Equal(t, tt.outToken, result.Token)
		})
	}
}

func TestLogOutEndpoint(t *testing.T) {
	t.Parallel()

	infoServiceTest := service.InfoServices{
		DBHost:    dbHostTest,
		DBPort:    portTest,
		TokenHost: tokenHostTest,
		TokenPort: portTest,
		Secret:    secretTest,
	}

	for _, tt := range []struct {
		name   string
		in     any
		outErr string
	}{
		{
			name: nameNoError,
			in: service.TokenRequest{
				Token: tokenTest,
			},
			outErr: "",
		},
		{
			name: nameErrorRequest,
			in: incorrectRequest{
				incorrect: true,
			},
			outErr: "isn't of type",
		},
		{
			name:   "ErrorWebService",
			in:     service.TokenRequest{},
			outErr: errWebServer.Error(),
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			testResp := struct {
				Err   string `json:"err"`
				Check bool   `json:"check"`
			}{
				Check: true,
				Err:   tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				assert.Error(t, err)
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

			r, err := service.MakeLogOutEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.ErrorResponse)
			if !ok {
				if tt.name != nameErrorRequest {
					assert.Error(t, errNotTypeIndicated)
				}
			}

			if result.Err != "" {
				resultErr = result.Err
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}
		})
	}
}

func TestGetAllUsersEndpoint(t *testing.T) {
	t.Parallel()

	infoServiceTest := service.InfoServices{
		DBHost:    dbHostTest,
		DBPort:    portTest,
		TokenHost: tokenHostTest,
		TokenPort: portTest,
		Secret:    secretTest,
	}

	for _, tt := range []struct {
		name     string
		outErr   string
		outUsers []dbapp.User
	}{
		{
			name: nameNoError,
			outUsers: []dbapp.User{
				{
					ID:       idTest,
					Username: usernameTest,
					Password: passwordTest,
					Email:    emailTest,
				},
			},
			outErr: "",
		},
		{
			name:     "ErrorWebService",
			outUsers: nil,
			outErr:   errWebServer.Error(),
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			testResp := struct {
				Err   string       `json:"err"`
				Users []dbapp.User `json:"users"`
			}{
				Users: tt.outUsers,
				Err:   tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				assert.Error(t, err)
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

			r, err := service.MakeGetAllUsersEndpoint(svc)(context.TODO(), nil)
			if err != nil {
				assert.Error(t, err)
			}

			result, ok := r.(service.UsersErrorResponse)
			if !ok {
				assert.Error(t, errNotTypeIndicated)
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, result.Err, tt.outErr)
			}

			assert.Equal(t, tt.outUsers, result.Users)
		})
	}
}

func TestProfileEndpoint(t *testing.T) {
	t.Parallel()

	infoServiceTest := service.InfoServices{
		DBHost:    dbHostTest,
		DBPort:    portTest,
		TokenHost: tokenHostTest,
		TokenPort: portTest,
		Secret:    secretTest,
	}

	for _, tt := range []struct {
		in      any
		name    string
		outErr  string
		outUser dbapp.User
	}{
		{
			name: nameNoError,
			in: service.TokenRequest{
				Token: tokenTest,
			},
			outUser: dbapp.User{
				ID:       idTest,
				Username: usernameTest,
				Password: passwordTest,
				Email:    emailTest,
			},
			outErr: "",
		},
		{
			name: nameErrorRequest,
			in: incorrectRequest{
				incorrect: true,
			},
			outErr: "isn't of type",
		},
		{
			name:    "ErrorWebService",
			in:      service.TokenRequest{},
			outUser: dbapp.User{},
			outErr:  errWebServer.Error(),
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

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
				Check:    true,
				Err:      tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				assert.Error(t, err)
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

			r, err := service.MakeProfileEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.UserErrorResponse)
			if !ok {
				if tt.name != nameErrorRequest {
					assert.Error(t, errNotTypeIndicated)
				}
			}

			if result.Err != "" {
				resultErr = result.Err
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			assert.Equal(t, tt.outUser, result.User)
		})
	}
}

func TestDeleteAccountEndpoint(t *testing.T) {
	t.Parallel()

	infoServiceTest := service.InfoServices{
		DBHost:    dbHostTest,
		DBPort:    portTest,
		TokenHost: tokenHostTest,
		TokenPort: portTest,
		Secret:    secretTest,
	}

	for _, tt := range []struct {
		name   string
		in     any
		outErr string
	}{
		{
			name: nameNoError,
			in: service.TokenRequest{
				Token: tokenTest,
			},
			outErr: "",
		},
		{
			name: nameErrorRequest,
			in: incorrectRequest{
				incorrect: true,
			},
			outErr: "isn't of type",
		},
		{
			name:   "ErrorWebService",
			in:     service.TokenRequest{},
			outErr: errWebServer.Error(),
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

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
				Err:      tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				assert.Error(t, err)
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

			r, err := service.MakeDeleteAccountEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.ErrorResponse)
			if !ok {
				if tt.name != nameErrorRequest {
					assert.Error(t, errNotTypeIndicated)
				}
			}

			if result.Err != "" {
				resultErr = result.Err
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}
		})
	}
}
