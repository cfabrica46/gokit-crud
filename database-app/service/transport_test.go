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
		out      getAllUsersRequest
		outError string
	}{
		{&http.Request{}, getAllUsersRequest{}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result getAllUsersRequest
			var resultErr string

			r, err := DecodeGetAllUsersRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(getAllUsersRequest)
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
	id := 1
	url := fmt.Sprintf("localhost:8080/user/%d", id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Error(err)
	}

	for i, tt := range []struct {
		in       *http.Request
		out      getUserByIDRequest
		outError string
	}{
		{req, getUserByIDRequest{}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result getUserByIDRequest
			var resultErr string

			r, err := DecodeGetUserByIDRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(getUserByIDRequest)
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

func TestDecodeGetUserByUsernameAndPasswordRequest(t *testing.T) {
	url := "localhost:8080/user/username_password"
	dataJSON, err := json.Marshal(User{Username: "cesar", Password: "01234"})
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
		out      getUserByUsernameAndPasswordRequest
		outError string
	}{
		{goodReq, getUserByUsernameAndPasswordRequest{"cesar", "01234"}, ""},
		{badReq, getUserByUsernameAndPasswordRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result getUserByUsernameAndPasswordRequest
			var resultErr string

			r, err := DecodeGetUserByUsernameAndPasswordRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(getUserByUsernameAndPasswordRequest)
			if !ok {
				if (tt.out != getUserByUsernameAndPasswordRequest{}) {
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
	username := "cesar"
	url := fmt.Sprintf("localhost:8080/id/%s", username)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Error(err)
	}

	for i, tt := range []struct {
		in       *http.Request
		out      getIDByUsernameRequest
		outError string
	}{
		{req, getIDByUsernameRequest{}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result getIDByUsernameRequest
			var resultErr string

			r, err := DecodeGetIDByUsernameRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(getIDByUsernameRequest)
			if !ok {
				if (tt.out != getIDByUsernameRequest{}) {
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

	dataJSON, err := json.Marshal(User{Username: "cesar", Password: "01234", Email: "cesar@email.com"})
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
		out      insertUserRequest
		outError string
	}{
		{goodReq, insertUserRequest{"cesar", "01234", "cesar@email.com"}, ""},
		{badReq, insertUserRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result insertUserRequest
			var resultErr string

			r, err := DecodeInsertUserRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(insertUserRequest)
			if !ok {
				if (tt.out != insertUserRequest{}) {
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

	dataJSON, err := json.Marshal(User{Username: "cesar", Password: "01234", Email: "cesar@email.com"})
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
		out      deleteUserRequest
		outError string
	}{
		{goodReq, deleteUserRequest{"cesar", "01234", "cesar@email.com"}, ""},
		{badReq, deleteUserRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result deleteUserRequest
			var resultErr string

			r, err := DecodeDeleteUserRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(deleteUserRequest)
			if !ok {
				if (tt.out != deleteUserRequest{}) {
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
