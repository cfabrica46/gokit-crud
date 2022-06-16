package service_test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/cfabrica46/gokit-crud/app/service"
	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	"github.com/stretchr/testify/assert"
)

func TestRequestFunc(t *testing.T) {
	t.Parallel()

	mockOK := service.NewMockClient(func(_ *http.Request) (*http.Response, error) {
		response := &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(strings.NewReader(`{
				"id": 1
			}`)),
		}

		return response, nil
	})

	mockNotOK := service.NewMockClient(func(_ *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("%w: error", errWebServer)
	})

	for _, tt := range []struct {
		Response *dbapp.IDErrorResponse
		client   service.HttpClient
		body     any
		name     string
		inURL    string
		inMethod string
		outErr   string
	}{
		{
			name:   "NoError",
			client: mockOK,
			body: dbapp.UsernameRequest{
				Username: usernameTest,
			},
			inURL:    "localhost:8080",
			inMethod: http.MethodPost,
			Response: &dbapp.IDErrorResponse{},
			outErr:   "",
		},
		{
			name:     "ErrorMarshal",
			client:   mockOK,
			body:     func() {},
			inURL:    "localhost:8080",
			inMethod: http.MethodPost,
			Response: &dbapp.IDErrorResponse{},
			outErr:   "error to make petition",
		},
		{
			name:   "ErrorURL",
			client: mockOK,
			body: dbapp.UsernameRequest{
				Username: usernameTest,
			},
			inURL:    "%%",
			inMethod: http.MethodPost,
			Response: &dbapp.IDErrorResponse{},
			outErr:   "error to make petition",
		},
		{
			name:   "ErrorService",
			client: mockNotOK,
			body: dbapp.UsernameRequest{
				Username: usernameTest,
			},
			inURL:    "localhost:8080",
			inMethod: http.MethodPost,
			Response: &dbapp.IDErrorResponse{},
			outErr:   "error to make petition",
		},
		{
			name:   "ErrorDecode",
			client: mockOK,
			body: dbapp.UsernameRequest{
				Username: usernameTest,
			},
			inURL:    "localhost:8080",
			inMethod: http.MethodPost,
			Response: nil,
			outErr:   "failed to decode request",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := service.RequestFunc(tt.client, tt.body, tt.inURL, tt.inMethod, tt.Response)

			if tt.name == "NoError" {
				assert.Nil(t, err)
				assert.Equal(t, idTest, tt.Response.ID)
				assert.Zero(t, tt.Response.Err)
			} else {
				assert.NotNil(t, err)
				assert.ErrorContains(t, err, tt.outErr)
			}
		})
	}
}

func TestRequestFuncWithoutBody(t *testing.T) {
	t.Parallel()

	mockOK := service.NewMockClient(func(_ *http.Request) (*http.Response, error) {
		response := &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(strings.NewReader(`{
				"users": [
					{
						"id":1,
						"username":"username",
						"password":"password",
						"email":"email@email.com"
					}
				]
			}`)),
		}

		return response, nil
	})

	mockNotOK := service.NewMockClient(func(_ *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("%w: error", errWebServer)
	})

	for _, tt := range []struct {
		Response *dbapp.UsersErrorResponse
		client   service.HttpClient
		name     string
		inURL    string
		inMethod string
		outErr   string
	}{
		{
			name:     "NoError",
			client:   mockOK,
			inURL:    "localhost:8080",
			inMethod: http.MethodPost,
			Response: &dbapp.UsersErrorResponse{},
			outErr:   "",
		},
		{
			name:     "ErrorURL",
			client:   mockOK,
			inURL:    "%%",
			inMethod: http.MethodPost,
			Response: &dbapp.UsersErrorResponse{},
			outErr:   "error to make petition",
		},
		{
			name:     "ErrorService",
			client:   mockNotOK,
			inURL:    "localhost:8080",
			inMethod: http.MethodPost,
			Response: &dbapp.UsersErrorResponse{},
			outErr:   "error to make petition",
		},
		{
			name:     "ErrorDecode",
			client:   mockOK,
			inURL:    "localhost:8080",
			inMethod: http.MethodPost,
			Response: nil,
			outErr:   "failed to decode request",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := service.RequestFuncWithoutBody(tt.client, tt.inURL, tt.inMethod, tt.Response)

			if tt.name == "NoError" {
				assert.Nil(t, err)
				assert.NotNil(t, idTest, tt.Response.Users)
				assert.Zero(t, tt.Response.Err)
			} else {
				assert.NotNil(t, err)
				assert.ErrorContains(t, err, tt.outErr)
			}
		})
	}
}
