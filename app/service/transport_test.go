package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeSignUpRequest(t *testing.T) {
	url := "localhost:8080"

	dataJSON, err := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}{
		"cesar",
		"01234",
		"cesar@email.com",
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
		in     *http.Request
		out    SignUpRequest
		outErr string
	}{
		{goodReq, SignUpRequest{"cesar", "01234", "cesar@email.com"}, ""},
		{badReq, SignUpRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := DecodeSignUpRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(SignUpRequest)
			if !ok {
				if (tt.out != SignUpRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, tt.outErr, resultErr)
			assert.Equal(t, tt.out, result)
		})
	}
}

func TestDecodeSignInRequest(t *testing.T) {
	url := "localhost:8080"

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
		in     *http.Request
		out    SignInRequest
		outErr string
	}{
		{goodReq, SignInRequest{"cesar", "01234"}, ""},
		{badReq, SignInRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := DecodeSignInRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(SignInRequest)
			if !ok {
				if (tt.out != SignInRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, tt.outErr, resultErr)
			assert.Equal(t, tt.out, result)
		})
	}
}

func TestDecodeLogOutRequest(t *testing.T) {
	url := "localhost:8080"

	dataJSON, err := json.Marshal(struct {
		Token string `json:"token"`
	}{
		"token",
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
		in     *http.Request
		out    LogOutRequest
		outErr string
	}{
		{goodReq, LogOutRequest{"token"}, ""},
		{badReq, LogOutRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := DecodeLogOutRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(LogOutRequest)
			if !ok {
				if (tt.out != LogOutRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, tt.outErr, resultErr)
			assert.Equal(t, tt.out, result)
		})
	}
}

func TestDecodeGetAllUsersRequest(t *testing.T) {
	url := "localhost:8080"

	goodReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Error(err)
	}

	for i, tt := range []struct {
		in     *http.Request
		out    GetAllUsersRequest
		outErr string
	}{
		{goodReq, GetAllUsersRequest{}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := DecodeGetAllUsersRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(GetAllUsersRequest)
			if !ok {
				if (tt.out != GetAllUsersRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, tt.outErr, resultErr)
			assert.Equal(t, tt.out, result)
		})
	}
}

func TestDecodeProfileRequest(t *testing.T) {
	url := "localhost:8080"

	dataJSON, err := json.Marshal(struct {
		Token string `json:"token"`
	}{
		"token",
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
		in     *http.Request
		out    ProfileRequest
		outErr string
	}{
		{goodReq, ProfileRequest{"token"}, ""},
		{badReq, ProfileRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := DecodeProfileRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(ProfileRequest)
			if !ok {
				if (tt.out != ProfileRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, tt.outErr, resultErr)
			assert.Equal(t, tt.out, result)
		})
	}
}

func TestDecodeDeleteAccountRequest(t *testing.T) {
	url := "localhost:8080"

	dataJSON, err := json.Marshal(struct {
		Token string `json:"token"`
	}{
		"token",
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
		in     *http.Request
		out    DeleteAccountRequest
		outErr string
	}{
		{goodReq, DeleteAccountRequest{"token"}, ""},
		{badReq, DeleteAccountRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := DecodeDeleteAccountRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(DeleteAccountRequest)
			if !ok {
				if (tt.out != DeleteAccountRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, tt.outErr, resultErr)
			assert.Equal(t, tt.out, result)
		})
	}
}
