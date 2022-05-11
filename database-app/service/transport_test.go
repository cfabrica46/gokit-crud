package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cfabrica46/gokit-crud/database-app/service"
)

func TestDecodeGetAllUsersRequest(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name     string
		in       *http.Request
		out      service.GetAllUsersRequest
		outError string
	}{
		{
			name:     "NoError",
			in:       &http.Request{},
			out:      service.GetAllUsersRequest{},
			outError: "",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.GetAllUsersRequest
			var resultErr string

			r, err := service.DecodeGetAllUsersRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.GetAllUsersRequest)
			if !ok {
				t.Error("result is not of the type indicated")
			}

			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, resultErr)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeGetUserByIDRequest(t *testing.T) {
	t.Parallel()

	dataJSON, err := json.Marshal(service.GetUserByIDRequest{ID: idTest})
	if err != nil {
		t.Error(err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, "localhost:8080", bytes.NewBuffer(dataJSON))
	if err != nil {
		t.Error(err)
	}

	badReq, err := http.NewRequest(http.MethodGet, "localhost:8080", bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Error(err)
	}

	for _, tt := range []struct {
		name     string
		in       *http.Request
		outError string
		out      service.GetUserByIDRequest
	}{
		{
			name: "NoError",
			in:   goodReq,
			out: service.GetUserByIDRequest{
				ID: idTest,
			},
			outError: "",
		},
		{
			name: "ErrRequestBody",
			in:   badReq,
			out: service.GetUserByIDRequest{
				ID: 0,
			},
			outError: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.GetUserByIDRequest
			var resultErr string

			r, err := service.DecodeGetUserByIDRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.GetUserByIDRequest)
			if !ok {
				if (tt.out != service.GetUserByIDRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, resultErr)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeGetUserByUsernameAndPasswordRequest(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/user/username_password"

	dataJSON, err := json.Marshal(service.GetUserByUsernameAndPasswordRequest{
		Username: usernameTest,
		Password: passwordTest,
	})
	if err != nil {
		t.Error(err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		t.Error(err)
	}

	badReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Error(err)
	}

	for _, tt := range []struct {
		name     string
		in       *http.Request
		out      service.GetUserByUsernameAndPasswordRequest
		outError string
	}{
		{
			name: "NoError",
			in:   goodReq,
			out: service.GetUserByUsernameAndPasswordRequest{
				Username: usernameTest,
				Password: passwordTest,
			},
			outError: "",
		},
		{
			name:     "ErrRequestBody",
			in:       badReq,
			out:      service.GetUserByUsernameAndPasswordRequest{},
			outError: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.GetUserByUsernameAndPasswordRequest
			var resultErr string

			r, err := service.DecodeGetUserByUsernameAndPasswordRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.GetUserByUsernameAndPasswordRequest)
			if !ok {
				if (tt.out != service.GetUserByUsernameAndPasswordRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, resultErr)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeGetIDByUsernameRequest(t *testing.T) {
	t.Parallel()

	dataJSON, err := json.Marshal(service.GetIDByUsernameRequest{Username: usernameTest})
	if err != nil {
		t.Error(err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, "localhost:8080", bytes.NewBuffer(dataJSON))
	if err != nil {
		t.Error(err)
	}

	badReq, err := http.NewRequest(http.MethodGet, "localhost:8080", bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Error(err)
	}

	for _, tt := range []struct {
		name     string
		in       *http.Request
		outError string
		out      service.GetIDByUsernameRequest
	}{
		{
			name: "NoError",
			in:   goodReq,
			out: service.GetIDByUsernameRequest{
				Username: usernameTest,
			},
			outError: "",
		},
		{
			name:     "ErrRequestBody",
			in:       badReq,
			out:      service.GetIDByUsernameRequest{},
			outError: "",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.GetIDByUsernameRequest
			var resultErr string

			r, err := service.DecodeGetIDByUsernameRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.GetIDByUsernameRequest)
			if !ok {
				if (tt.out != service.GetIDByUsernameRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, resultErr)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeInsertUserRequest(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/user"

	dataJSON, err := json.Marshal(service.InsertUserRequest{usernameTest, "0idTest234", emailTest})
	if err != nil {
		t.Error(err)
	}

	goodReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		t.Error(err)
	}

	badReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Error(err)
	}

	for _, tt := range []struct {
		name     string
		in       *http.Request
		out      service.InsertUserRequest
		outError string
	}{
		{
			name: "NoError",
			in:   goodReq,
			out: service.InsertUserRequest{
				Username: usernameTest,
				Password: "0idTest234",
				Email:    emailTest,
			},
			outError: "",
		},
		{
			name:     "ErrRequestBody",
			in:       badReq,
			out:      service.InsertUserRequest{},
			outError: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.InsertUserRequest
			var resultErr string

			r, err := service.DecodeInsertUserRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.InsertUserRequest)
			if !ok {
				if (tt.out != service.InsertUserRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, resultErr)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeDeleteUserRequest(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/user"

	dataJSON, err := json.Marshal(service.DeleteUserRequest{idTest})
	if err != nil {
		t.Error(err)
	}

	goodReq, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		t.Error(err)
	}

	badReq, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Error(err)
	}

	for _, tt := range []struct {
		name     string
		in       *http.Request
		outError string
		out      service.DeleteUserRequest
	}{
		{
			name: "NoError",
			in:   goodReq,
			out: service.DeleteUserRequest{
				ID: idTest,
			},
			outError: "",
		},
		{
			name:     "ErrRequestBody",
			in:       badReq,
			out:      service.DeleteUserRequest{},
			outError: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.DeleteUserRequest
			var resultErr string

			r, err := service.DecodeDeleteUserRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.DeleteUserRequest)
			if !ok {
				if (tt.out != service.DeleteUserRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, result)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestEncodeResponse(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name     string
		in       string
		outError string
	}{
		{
			name:     "NoError",
			in:       "test",
			outError: "",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var resultErr string

			err := service.EncodeResponse(context.TODO(), httptest.NewRecorder(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, resultErr)
			}
		})
	}
}
