package service_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cfabrica46/gokit-crud/app/service"
	"github.com/stretchr/testify/assert"
)

const (
	usernamePasswordEmailRequestJSON = `{
		 "username": "username",
		 "password": "password",
		 "email": "email@email.com"
	}`

	usernamePasswordRequestJSON = `{
		 "username": "username",
		 "password": "password"
	}`
)

func TestDecodeRequestWithoutBody(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		out    any
		in     *http.Request
		name   string
		outErr string
	}{
		{
			name:   nameNoError + "GetAllRequest",
			in:     nil,
			outErr: "",
			out:    service.EmptyRequest{},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r, err := service.DecodeRequestWithoutBody()(context.TODO(), tt.in)

			assert.Empty(t, err)
			assert.Equal(t, tt.out, r)
		})
	}
}

func TestDecodeRequestWithBody(t *testing.T) {
	t.Parallel()

	usernamePasswordEmailReq, err := http.NewRequest(
		http.MethodPost,
		urlTest,
		bytes.NewBuffer([]byte(usernamePasswordEmailRequestJSON)),
	)
	if err != nil {
		assert.Error(t, err)
	}

	usernamePasswordReq, err := http.NewRequest(
		http.MethodPost,
		urlTest,
		bytes.NewBuffer([]byte(usernamePasswordRequestJSON)),
	)
	if err != nil {
		assert.Error(t, err)
	}

	badReq, err := http.NewRequest(http.MethodPost, urlTest, bytes.NewBuffer([]byte{}))
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		inType      any
		in          *http.Request
		name        string
		outUsername string
		outPassword string
		outEmail    string
		outErr      string
		outID       int
	}{
		{
			name:        nameNoError + "UsernamePasswordEmailRequest",
			inType:      service.UsernamePasswordEmailRequest{},
			in:          usernamePasswordEmailReq,
			outUsername: usernameTest,
			outPassword: passwordTest,
			outEmail:    emailTest,
			outErr:      "",
		},
		{
			name:        nameNoError + "UsernamePasswordRequest",
			inType:      service.UsernamePasswordRequest{},
			in:          usernamePasswordReq,
			outUsername: usernameTest,
			outPassword: passwordTest,
			outErr:      "",
		},
		{
			name:   "BadRequest",
			inType: service.UsernamePasswordEmailRequest{},
			in:     badReq,
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			var r any

			switch resultType := tt.inType.(type) {
			case service.UsernamePasswordEmailRequest:
				r, err = service.DecodeRequestWithBody(resultType)(context.TODO(), tt.in)
				if err != nil {
					resultErr = err.Error()
				}
			case service.UsernamePasswordRequest:
				r, err = service.DecodeRequestWithBody(resultType)(context.TODO(), tt.in)
				if err != nil {
					resultErr = err.Error()
				}
			default:
				assert.Fail(t, "Error to type inType")
			}

			switch result := r.(type) {
			case service.UsernamePasswordEmailRequest:
				assert.Equal(t, tt.outUsername, result.Username)
				assert.Equal(t, tt.outPassword, result.Password)
				assert.Equal(t, tt.outEmail, result.Email)
				assert.Empty(t, resultErr)
			case service.UsernamePasswordRequest:
				assert.Equal(t, tt.outUsername, result.Username)
				assert.Equal(t, tt.outPassword, result.Password)
				assert.Empty(t, resultErr)
			default:
				if tt.name != nameNoError {
					assert.Contains(t, resultErr, tt.outErr)
				}
			}
		})
	}
}

func TestDecodeRequestWithHeader(t *testing.T) {
	t.Parallel()

	okReq, err := http.NewRequest(
		http.MethodPost,
		urlTest,
		bytes.NewBuffer([]byte{}),
	)
	if err != nil {
		assert.Error(t, err)
	}

	okReq.Header.Set("Authorization", "token")

	badReq, err := http.NewRequest(http.MethodPost, urlTest, bytes.NewBuffer([]byte{}))
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		inType   service.TokenRequest
		in       *http.Request
		name     string
		outErr   string
		outToken string
		outID    int
	}{
		{
			name:     nameNoError,
			inType:   service.TokenRequest{},
			in:       okReq,
			outToken: tokenTest,
			outErr:   "",
		},
		{
			name:   "BadRequest",
			inType: service.TokenRequest{},
			in:     badReq,
			outErr: "failed to get header",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			r, err := service.DecodeRequestWithHeader(tt.inType)(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.TokenRequest)
			if tt.name == nameNoError {
				if !ok {
					assert.Fail(t, "Error to type inType")
				}
			}

			if tt.name == nameNoError {
				assert.Equal(t, tt.outToken, result.Token)
				assert.Nil(t, err)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}
		})
	}
}

func TestEncodeResponse(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		in     any
		outErr string
	}{
		{
			name:   nameNoError,
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

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}
		})
	}
}

/* import (
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
		assert.Error(t, err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		assert.Error(t, err)
	}

	badReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		assert.Error(t, err)
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
					assert.Fail(t, "result is not of the type indicated")
				}
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

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
		assert.Error(t, err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		assert.Error(t, err)
	}

	badReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		assert.Error(t, err)
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
					assert.Fail(t, "result is not of the type indicated")
				}
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

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
		assert.Error(t, err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		assert.Error(t, err)
	}

	goodReq.Header.Add("Authorization", tokenTest)

	badReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.LogOutRequest
		outErr string
	}{
		{
			name: "NoError",
			in:   goodReq,
			out: service.LogOutRequest{
				Token: tokenTest,
			},
			outErr: "",
		},
		{
			name:   "ErrorNotHeader",
			in:     badReq,
			out:    service.LogOutRequest{},
			outErr: "failed to get header",
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
					assert.Fail(t, "result is not of the type indicated")
				}
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			assert.Equal(t, tt.out, result)
		})
	}
}

func TestDecodeGetAllUsersRequest(t *testing.T) {
	t.Parallel()

	url := urlTest

	goodReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		assert.Error(t, err)
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
					assert.Fail(t, "result is not of the type indicated")
				}
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

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
		assert.Error(t, err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		assert.Error(t, err)
	}

	goodReq.Header.Add("Authorization", tokenTest)

	badReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.ProfileRequest
		outErr string
	}{
		{
			name: "NoError",
			in:   goodReq,
			out: service.ProfileRequest{
				Token: tokenTest,
			},
			outErr: "",
		},
		{
			name:   "ErrorNotHeader",
			in:     badReq,
			out:    service.ProfileRequest{},
			outErr: "failed to get header",
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
					assert.Fail(t, "result is not of the type indicated")
				}
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

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
		assert.Error(t, err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		assert.Error(t, err)
	}

	goodReq.Header.Add("Authorization", tokenTest)

	badReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.DeleteAccountRequest
		outErr string
	}{
		{
			name: "NoError",
			in:   goodReq,
			out: service.DeleteAccountRequest{
				Token: tokenTest,
			},
			outErr: "",
		},
		{
			name:   "ErrorNotHeader",
			in:     badReq,
			out:    service.DeleteAccountRequest{},
			outErr: "failed to get header",
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
					assert.Fail(t, "result is not of the type indicated")
				}
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

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

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}
		})
	}
} */
