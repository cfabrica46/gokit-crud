package service_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cfabrica46/gokit-crud/database-app/service"
	"github.com/stretchr/testify/assert"
)

const (
	idRequestJSON = `{
		 "id": 1
	}`

	usernameRequestJSON = `{
		 "username": "username",
	}`

	usernamePasswordRequestJSON = `{
		 "username": "username",
		 "password": "password",
	}`

	usernamePasswordEmailRequestJSON = `{
		 "username": "username",
		 "password": "password",
		 "email": "email@email.com",
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
			name:   nameNoError + "IDRequest",
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

func TestDecodeRequest(t *testing.T) {
	t.Parallel()

	idReq, err := http.NewRequest(
		http.MethodPost,
		urlTest,
		bytes.NewBuffer([]byte(idRequestJSON)),
	)
	if err != nil {
		assert.Error(t, err)
	}

	usernameReq, err := http.NewRequest(
		http.MethodPost,
		urlTest,
		bytes.NewBuffer([]byte(usernameRequestJSON)),
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

	usernamePasswordEmailReq, err := http.NewRequest(
		http.MethodPost,
		urlTest,
		bytes.NewBuffer([]byte(usernameRequestJSON)),
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
			name:   nameNoError + "IDRequest",
			inType: service.IDRequest{},
			in:     idReq,
			outID:  idTest,
			outErr: "",
		},
		{
			name:        nameNoError + "UsernameRequest",
			inType:      service.UsernameRequest{},
			in:          usernameReq,
			outUsername: usernameTest,
			outErr:      "",
		},
		{
			name:   nameNoError + "UsernamePasswordRequest",
			inType: service.UsernamePasswordRequest{},
			in:     usernamePasswordReq,
			outErr: "",
		},
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
			name:   "BadRequest",
			inType: service.IDRequest{},
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
			case service.IDRequest:
				r, err = service.DecodeRequest(resultType)(context.TODO(), tt.in)
				if err != nil {
					resultErr = err.Error()
				}
			case service.UsernameRequest:
				r, err = service.DecodeRequest(resultType)(context.TODO(), tt.in)
				if err != nil {
					resultErr = err.Error()
				}
			case service.UsernamePasswordRequest:
				r, err = service.DecodeRequest(resultType)(context.TODO(), tt.in)
				if err != nil {
					resultErr = err.Error()
				}
			case service.UsernamePasswordEmailRequest:
				r, err = service.DecodeRequest(resultType)(context.TODO(), tt.in)
				if err != nil {
					resultErr = err.Error()
				}
			default:
				assert.Fail(t, "Error to type inType")
			}

			switch result := r.(type) {
			case service.IDRequest:
				assert.Equal(t, tt.outID, result.ID)
				assert.Empty(t, resultErr)
			case service.UsernameRequest:
				assert.Equal(t, tt.outUsername, result.Username)
				assert.Empty(t, resultErr)
			case service.UsernamePasswordRequest:
				assert.Equal(t, tt.outUsername, result.Username)
				assert.Equal(t, tt.outPassword, result.Password)
				assert.Empty(t, resultErr)
			case service.UsernamePasswordEmailRequest:
				assert.Equal(t, tt.outUsername, result.Username)
				assert.Equal(t, tt.outPassword, result.Password)
				assert.Equal(t, tt.outEmail, result.Email)
				assert.Empty(t, resultErr)
			default:
				if tt.name != nameNoError {
					assert.Contains(t, resultErr, tt.outErr)
				}
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
