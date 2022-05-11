package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cfabrica46/gokit-crud/app/service"
	"github.com/stretchr/testify/assert"
)

func TestDecodeSignUpRequest(t *testing.T) {
	t.Parallel()

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

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.SignUpRequest
		outErr string
	}{
		{
			name: "NoError",
			in:   goodReq,
			out: service.SignUpRequest{
				Username: usernameTest,
				Password: passwordTest,
				Email:    emailTest,
			},
			outErr: "",
		},
		{
			name:   "ErrorBadRequest",
			in:     badReq,
			out:    service.SignUpRequest{},
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result interface{}
			var resultErr string

			r, err := service.DecodeSignUpRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.SignUpRequest)
			if !ok {
				if (tt.out != service.SignUpRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, tt.outErr, resultErr)
			assert.Equal(t, tt.out, result)
		})
	}
}

func TestDecodeSignInRequest(t *testing.T) {
	t.Parallel()

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

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.SignInRequest
		outErr string
	}{
		{
			name: "NoError",
			in:   goodReq,
			out: service.SignInRequest{
				usernameTest, passwordTest,
			},
			outErr: "",
		},
		{
			name:   "ErrorBadRequest",
			in:     badReq,
			out:    service.SignInRequest{},
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result interface{}
			var resultErr string

			r, err := service.DecodeSignInRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.SignInRequest)
			if !ok {
				if (tt.out != service.SignInRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, tt.outErr, resultErr)
			assert.Equal(t, tt.out, result)
		})
	}
}

func TestDecodeLogOutRequest(t *testing.T) {
	t.Parallel()

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

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.LogOutRequest
		outErr string
	}{
		{
			name:   "NoError",
			in:     goodReq,
			out:    service.LogOutRequest{},
			outErr: "",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result interface{}
			var resultErr string

			r, err := service.DecodeLogOutRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.LogOutRequest)
			if !ok {
				if (tt.out != service.LogOutRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, tt.outErr, resultErr)
			assert.Equal(t, tt.out, result)
		})
	}
}

func TestDecodeGetAllUsersRequest(t *testing.T) {
	t.Parallel()

	url := urlTest

	goodReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Error(err)
	}

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.GetAllUsersRequest
		outErr string
	}{
		{
			name:   "NoError",
			in:     goodReq,
			out:    service.GetAllUsersRequest{},
			outErr: "",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result interface{}
			var resultErr string

			r, err := service.DecodeGetAllUsersRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.GetAllUsersRequest)
			if !ok {
				if (tt.out != service.GetAllUsersRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, tt.outErr, resultErr)
			assert.Equal(t, tt.out, result)
		})
	}
}

func TestDecodeProfileRequest(t *testing.T) {
	t.Parallel()

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

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.ProfileRequest
		outErr string
	}{
		{
			name:   "NoError",
			in:     goodReq,
			out:    service.ProfileRequest{},
			outErr: "",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result interface{}
			var resultErr string

			r, err := service.DecodeProfileRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.ProfileRequest)
			if !ok {
				if (tt.out != service.ProfileRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, tt.outErr, resultErr)
			assert.Equal(t, tt.out, result)
		})
	}
}

func TestDecodeDeleteAccountRequest(t *testing.T) {
	t.Parallel()

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

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.DeleteAccountRequest
		outErr string
	}{
		{
			name:   "NoError",
			in:     goodReq,
			out:    service.DeleteAccountRequest{},
			outErr: "",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result interface{}
			var resultErr string

			r, err := service.DecodeDeleteAccountRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.DeleteAccountRequest)
			if !ok {
				if (tt.out != service.DeleteAccountRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, tt.outErr, resultErr)
			assert.Equal(t, tt.out, result)
		})
	}
}

func TestEncodeResponse(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		in     interface{}
		outErr string
	}{
		{
			name:   "NoError",
			in:     "test",
			outErr: "",
		},
		{
			name:   "ErrorBadEncode",
			in:     func() {},
			outErr: "json: unsupported type: func()",
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

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}
