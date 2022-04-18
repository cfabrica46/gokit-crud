package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cfabrica46/gokit-crud/database-app/service"
)

func TestDecodeGetAllUsersRequest(t *testing.T) {
	for i, tt := range []struct {
		in       *http.Request
		out      service.GetAllUsersRequest
		outError string
	}{
		{&http.Request{}, service.GetAllUsersRequest{}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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
	dataJSON, err := json.Marshal(service.GetUserByIDRequest{ID: 1})
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

	for i, tt := range []struct {
		in       *http.Request
		out      service.GetUserByIDRequest
		outError string
	}{
		{goodReq, service.GetUserByIDRequest{ID: 1}, ""},
		{badReq, service.GetUserByIDRequest{ID: 0}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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
	url := "localhost:8080/user/username_password"

	dataJSON, err := json.Marshal(service.GetUserByUsernameAndPasswordRequest{"cesar", "01234"})
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

	for i, tt := range []struct {
		in       *http.Request
		out      service.GetUserByUsernameAndPasswordRequest
		outError string
	}{
		{goodReq, service.GetUserByUsernameAndPasswordRequest{"cesar", "01234"}, ""},
		{badReq, service.GetUserByUsernameAndPasswordRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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
	dataJSON, err := json.Marshal(service.GetIDByUsernameRequest{Username: "cesar"})
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

	for i, tt := range []struct {
		in       *http.Request
		out      service.GetIDByUsernameRequest
		outError string
	}{
		{goodReq, service.GetIDByUsernameRequest{Username: "cesar"}, ""},
		{badReq, service.GetIDByUsernameRequest{}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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
	url := "localhost:8080/user"

	dataJSON, err := json.Marshal(service.InsertUserRequest{"cesar", "01234", "cesar@email.com"})
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

	for i, tt := range []struct {
		in       *http.Request
		out      service.InsertUserRequest
		outError string
	}{
		{goodReq, service.InsertUserRequest{"cesar", "01234", "cesar@email.com"}, ""},
		{badReq, service.InsertUserRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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
	url := "localhost:8080/user"

	dataJSON, err := json.Marshal(service.DeleteUserRequest{1})
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

	for i, tt := range []struct {
		in       *http.Request
		out      service.DeleteUserRequest
		outError string
	}{
		{goodReq, service.DeleteUserRequest{1}, ""},
		{badReq, service.DeleteUserRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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
	for i, tt := range []struct {
		in       string
		outError string
	}{
		{"test", ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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
