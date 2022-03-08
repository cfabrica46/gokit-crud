package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDecodeGetAllUsersRequest(t *testing.T) {
	for i, tt := range []struct {
		in       *http.Request
		out      GetAllUsersRequest
		outError string
	}{
		{&http.Request{}, GetAllUsersRequest{}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result GetAllUsersRequest
			var resultErr string

			r, err := DecodeGetAllUsersRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(GetAllUsersRequest)
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
	dataJSON, err := json.Marshal(GetUserByIDRequest{ID: 1})
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
		out      GetUserByIDRequest
		outError string
	}{
		{goodReq, GetUserByIDRequest{ID: 1}, ""},
		{badReq, GetUserByIDRequest{ID: 0}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result GetUserByIDRequest
			var resultErr string

			r, err := DecodeGetUserByIDRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(GetUserByIDRequest)
			if !ok {
				if (tt.out != GetUserByIDRequest{}) {
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
	dataJSON, err := json.Marshal(GetUserByUsernameAndPasswordRequest{"cesar", "01234"})
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
		out      GetUserByUsernameAndPasswordRequest
		outError string
	}{
		{goodReq, GetUserByUsernameAndPasswordRequest{"cesar", "01234"}, ""},
		{badReq, GetUserByUsernameAndPasswordRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result GetUserByUsernameAndPasswordRequest
			var resultErr string

			r, err := DecodeGetUserByUsernameAndPasswordRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(GetUserByUsernameAndPasswordRequest)
			if !ok {
				if (tt.out != GetUserByUsernameAndPasswordRequest{}) {
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

	dataJSON, err := json.Marshal(GetIDByUsernameRequest{Username: "cesar"})
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
		out      GetIDByUsernameRequest
		outError string
	}{
		{goodReq, GetIDByUsernameRequest{Username: "cesar"}, ""},
		{badReq, GetIDByUsernameRequest{}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result GetIDByUsernameRequest
			var resultErr string

			r, err := DecodeGetIDByUsernameRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(GetIDByUsernameRequest)
			if !ok {
				if (tt.out != GetIDByUsernameRequest{}) {
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

	dataJSON, err := json.Marshal(InsertUserRequest{"cesar", "01234", "cesar@email.com"})
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
		out      InsertUserRequest
		outError string
	}{
		{goodReq, InsertUserRequest{"cesar", "01234", "cesar@email.com"}, ""},
		{badReq, InsertUserRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result InsertUserRequest
			var resultErr string

			r, err := DecodeInsertUserRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(InsertUserRequest)
			if !ok {
				if (tt.out != InsertUserRequest{}) {
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

	dataJSON, err := json.Marshal(DeleteUserRequest{1})
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
		out      DeleteUserRequest
		outError string
	}{
		{goodReq, DeleteUserRequest{1}, ""},
		{badReq, DeleteUserRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result DeleteUserRequest
			var resultErr string

			r, err := DecodeDeleteUserRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(DeleteUserRequest)
			if !ok {
				if (tt.out != DeleteUserRequest{}) {
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

			err := EncodeResponse(context.TODO(), httptest.NewRecorder(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, resultErr)
			}
		})
	}
}
