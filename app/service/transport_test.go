package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cfabrica46/gokit-crud/app/service"
	"github.com/stretchr/testify/assert"
)

func TestDecodeSignUpRequest(t *testing.T) {
	url := urlTest

	dataJSON, err := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}{
		usernameTest,
		passwordTest,
		emailTest,
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

	for index, table := range []struct {
		in     *http.Request
		out    service.SignUpRequest
		outErr string
	}{
		{goodReq, service.SignUpRequest{usernameTest, passwordTest, emailTest}, ""},
		{badReq, service.SignUpRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := service.DecodeSignUpRequest(context.TODO(), table.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.SignUpRequest)
			if !ok {
				if (table.out != service.SignUpRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, table.outErr, resultErr)
			assert.Equal(t, table.out, result)
		})
	}
}

func TestDecodeSignInRequest(t *testing.T) {
	url := urlTest

	dataJSON, err := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		usernameTest,
		passwordTest,
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

	for index, table := range []struct {
		in     *http.Request
		out    service.SignInRequest
		outErr string
	}{
		{goodReq, service.SignInRequest{usernameTest, passwordTest}, ""},
		{badReq, service.SignInRequest{}, "EOF"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := service.DecodeSignInRequest(context.TODO(), table.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.SignInRequest)
			if !ok {
				if (table.out != service.SignInRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, table.outErr, resultErr)
			assert.Equal(t, table.out, result)
		})
	}
}

func TestDecodeLogOutRequest(t *testing.T) {
	url := urlTest

	dataJSON, err := json.Marshal(struct {
		Token string `json:"token"`
	}{
		tokenTest,
	})
	if err != nil {
		t.Error(err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		in     *http.Request
		out    service.LogOutRequest
		outErr string
	}{
		{goodReq, service.LogOutRequest{}, ""},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := service.DecodeLogOutRequest(context.TODO(), table.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.LogOutRequest)
			if !ok {
				if (table.out != service.LogOutRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, table.outErr, resultErr)
			assert.Equal(t, table.out, result)
		})
	}
}

func TestDecodeGetAllUsersRequest(t *testing.T) {
	url := urlTest

	goodReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		in     *http.Request
		out    service.GetAllUsersRequest
		outErr string
	}{
		{goodReq, service.GetAllUsersRequest{}, ""},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := service.DecodeGetAllUsersRequest(context.TODO(), table.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.GetAllUsersRequest)
			if !ok {
				if (table.out != service.GetAllUsersRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, table.outErr, resultErr)
			assert.Equal(t, table.out, result)
		})
	}
}

func TestDecodeProfileRequest(t *testing.T) {
	url := urlTest

	dataJSON, err := json.Marshal(struct {
		Token string `json:"token"`
	}{
		tokenTest,
	})
	if err != nil {
		t.Error(err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		in     *http.Request
		out    service.ProfileRequest
		outErr string
	}{
		{goodReq, service.ProfileRequest{}, ""},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := service.DecodeProfileRequest(context.TODO(), table.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.ProfileRequest)
			if !ok {
				if (table.out != service.ProfileRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, table.outErr, resultErr)
			assert.Equal(t, table.out, result)
		})
	}
}

func TestDecodeDeleteAccountRequest(t *testing.T) {
	url := urlTest

	dataJSON, err := json.Marshal(struct {
		Token string `json:"token"`
	}{
		tokenTest,
	})
	if err != nil {
		t.Error(err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		in     *http.Request
		out    service.DeleteAccountRequest
		outErr string
	}{
		{goodReq, service.DeleteAccountRequest{}, ""},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := service.DecodeDeleteAccountRequest(context.TODO(), table.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.DeleteAccountRequest)
			if !ok {
				if (table.out != service.DeleteAccountRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, table.outErr, resultErr)
			assert.Equal(t, table.out, result)
		})
	}
}

func TestEncodeResponse(t *testing.T) {
	for index, table := range []struct {
		in     interface{}
		outErr string
	}{
		{"test", ""},
		{func() {}, "json: unsupported type: func()"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr string

			err := service.EncodeResponse(context.TODO(), httptest.NewRecorder(), table.in)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
		})
	}
}
