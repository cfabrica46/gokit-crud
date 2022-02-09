package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDecodeGetAllUsersRequest(t *testing.T) {
	for i, tt := range []struct {
		in       *http.Request
		out      *getAllUsersRequest
		outError string
	}{
		{&http.Request{}, &getAllUsersRequest{}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := DecodeGetAllUsersRequest(context.TODO(), tt.in)
			if err != nil {
				if err.Error() != tt.outError {
					t.Errorf("want %v; got %v", tt.outError, err)
				}
			}
			if result != *tt.out {
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
		out      *getUserByIDRequest
		outError string
	}{
		{req, &getUserByIDRequest{}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := DecodeGetUserByIDRequest(context.TODO(), tt.in)
			if err != nil {
				if err.Error() != tt.outError {
					t.Errorf("want %v; got %v", tt.outError, err)
				}
			}
			if result != *tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeGetUserByUsernameAndPasswordRequest(t *testing.T) {
	url := "localhost:8080/user/username_password"

	dataJSON, err := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		"cesar",
		"01234",
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

	for i, tt := range []struct {
		in       *http.Request
		out      *getUserByUsernameAndPasswordRequest
		outError string
	}{
		{goodReq, &getUserByUsernameAndPasswordRequest{"cesar", "01234"}, ""},
		{badReq, nil, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := DecodeGetUserByUsernameAndPasswordRequest(context.TODO(), tt.in)
			if err != nil {
				if err.Error() != tt.outError {
					t.Errorf("want %v; got %v", tt.outError, err)
				}
			}

			if tt.out == nil {
				if result != nil {
					t.Errorf("want %v; got %v", tt.out, result)
				}
			} else {
				if result != *tt.out {
					t.Errorf("want %v; got %v", tt.out, result)
				}
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
		out      *getIDByUsernameRequest
		outError string
	}{
		{req, &getIDByUsernameRequest{}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := DecodeGetIDByUsernameRequest(context.TODO(), tt.in)
			if err != nil {
				if err.Error() != tt.outError {
					t.Errorf("want %v; got %v", tt.outError, err)
				}
			}
			if result != *tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeInsertUserRequest(t *testing.T) {
	url := "localhost:8080/user"

	dataJSON, err := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}{
		"cesar",
		"01234",
		"cesar@gmail.com",
	})
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
		out      *insertUserRequest
		outError string
	}{
		{goodReq, &insertUserRequest{"cesar", "01234", "cesar@gmail.com"}, ""},
		{badReq, nil, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := DecodeInsertUserRequest(context.TODO(), tt.in)
			if err != nil {
				if err.Error() != tt.outError {
					t.Errorf("want %v; got %v", tt.outError, err)
				}
			}

			if tt.out == nil {
				if result != nil {
					t.Errorf("want %v; got %v", tt.out, result)
				}
			} else {
				if result != *tt.out {
					t.Errorf("want %v; got %v", tt.out, result)
				}
			}
		})
	}
}

func TestDecodeDeleteUserRequest(t *testing.T) {
	url := "localhost:8080/user"

	dataJSON, err := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}{
		"cesar",
		"01234",
		"cesar@gmail.com",
	})
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
		out      *deleteUserRequest
		outError string
	}{
		{goodReq, &deleteUserRequest{"cesar", "01234", "cesar@gmail.com"}, ""},
		{badReq, nil, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := DecodeDeleteUserRequest(context.TODO(), tt.in)
			if err != nil {
				if err.Error() != tt.outError {
					t.Errorf("want %v; got %v", tt.outError, err)
				}
			}

			if tt.out == nil {
				if result != nil {
					t.Errorf("want %v; got %v", tt.out, result)
				}
			} else {
				if result != *tt.out {
					t.Errorf("want %v; got %v", tt.out, result)
				}
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
			err := EncodeResponse(context.TODO(), httptest.NewRecorder(), tt.in)
			if err != nil {
				if err.Error() == tt.outError {
					t.Errorf("want %v; got %v", tt.outError, err)
				}
			}

		})
	}
}
