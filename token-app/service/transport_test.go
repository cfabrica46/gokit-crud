package service_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cfabrica46/gokit-crud/token-app/service"
	"github.com/stretchr/testify/assert"
)

const (
	generateTokenRequestJSON = `{
		 "username": "username",
		 "email": "email@email.com",
		 "secret": "secret",
		 "id": 1
	}`

	extractTokenRequestJSON = `{
		"token": "token",
		 "secret": "secret"
	}`

	tokenRequestJSON = `{
		"token": "token"
	}`
)

func TestDecodeRequest(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/"

	generateTokenReq, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer([]byte(generateTokenRequestJSON)),
	)
	if err != nil {
		assert.Error(t, err)
	}

	extractTokenReq, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer([]byte(extractTokenRequestJSON)),
	)
	if err != nil {
		assert.Error(t, err)
	}

	tokenReq, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer([]byte(tokenRequestJSON)),
	)
	if err != nil {
		assert.Error(t, err)
	}

	badReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		inType      any
		in          *http.Request
		name        string
		outUsername string
		outEmail    string
		outToken    string
		outSecret   string
		outErr      string
		outID       int
	}{
		{
			name:        nameNoError + "GenerateToken",
			inType:      service.GenerateTokenRequest{},
			in:          generateTokenReq,
			outID:       idTest,
			outUsername: usernameTest,
			outEmail:    emailTest,
			outSecret:   secretTest,
			outErr:      "",
		},
		{
			name:      nameNoError + "ExtractToken",
			inType:    service.ExtractTokenRequest{},
			in:        extractTokenReq,
			outToken:  tokenTest,
			outSecret: secretTest,
			outErr:    "",
		},
		{
			name:     nameNoError + "Token",
			inType:   service.Token{},
			in:       tokenReq,
			outToken: tokenTest,
			outErr:   "",
		},
		{
			name:   "BadRequest",
			inType: service.GenerateTokenRequest{},
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
			case service.GenerateTokenRequest:
				r, err = service.DecodeRequest(resultType)(context.TODO(), tt.in)
				if err != nil {
					resultErr = err.Error()
				}
			case service.ExtractTokenRequest:
				r, err = service.DecodeRequest(resultType)(context.TODO(), tt.in)
				if err != nil {
					resultErr = err.Error()
				}
			case service.Token:
				r, err = service.DecodeRequest(resultType)(context.TODO(), tt.in)
				if err != nil {
					resultErr = err.Error()
				}
			default:
				assert.Fail(t, "Error to type inType")
			}

			switch result := r.(type) {
			case service.GenerateTokenRequest:
				assert.Equal(t, tt.outID, result.ID)
				assert.Equal(t, tt.outUsername, result.Username)
				assert.Equal(t, tt.outEmail, result.Email)
				assert.Equal(t, tt.outSecret, result.Secret)
				assert.Empty(t, resultErr)
			case service.ExtractTokenRequest:
				assert.Equal(t, tt.outToken, result.Token)
				assert.Equal(t, tt.outSecret, result.Secret)
				assert.Empty(t, resultErr)
			case service.Token:
				assert.Equal(t, tt.outToken, result.Token)
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
			name:   "ErrorEncode",
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
