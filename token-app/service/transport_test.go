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
	//nolint:gosec
	generateTokenRequestJSON = `{
		 "username": "username",
		 "email": "email@email.com",
		 "secret": "secret",
		 "id": 1
	 }`

	//nolint:gosec
	extractTokenRequestJSON = `{
		"token": "token",
		 "secret": "secret"
	}`

	//nolint:gosec
	tokenRequestJSON = `{
		"token": "token"
	}`
)

func TestDecodeRequest(t *testing.T) {
	t.Parallel()

	generateTokenReq, err := http.NewRequest(
		http.MethodPost,
		urlTest,
		bytes.NewBuffer([]byte(generateTokenRequestJSON)),
	)
	if err != nil {
		assert.Error(t, err)
	}

	extractTokenReq, err := http.NewRequest(
		http.MethodPost,
		urlTest,
		bytes.NewBuffer([]byte(extractTokenRequestJSON)),
	)
	if err != nil {
		assert.Error(t, err)
	}

	tokenReq, err := http.NewRequest(
		http.MethodPost,
		urlTest,
		bytes.NewBuffer([]byte(tokenRequestJSON)),
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
		outEmail    string
		outToken    string
		outSecret   string
		outErr      string
		outID       int
	}{
		{
			name:        nameNoError + "GenerateToken",
			inType:      service.IDUsernameEmailSecretRequest{},
			in:          generateTokenReq,
			outID:       idTest,
			outUsername: usernameTest,
			outEmail:    emailTest,
			outSecret:   secretTest,
			outErr:      "",
		},
		{
			name:      nameNoError + "ExtractToken",
			inType:    service.TokenSecretRequest{},
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
			inType: service.IDUsernameEmailSecretRequest{},
			in:     badReq,
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			var req any

			switch resultType := tt.inType.(type) {
			case service.IDUsernameEmailSecretRequest:
				req, err = service.DecodeRequest(resultType)(context.TODO(), tt.in)
				if err != nil {
					resultErr = err.Error()
				}

				result, ok := req.(service.IDUsernameEmailSecretRequest)
				if ok {
					assert.Equal(t, tt.outID, result.ID)
					assert.Equal(t, tt.outUsername, result.Username)
					assert.Equal(t, tt.outEmail, result.Email)
					assert.Equal(t, tt.outSecret, result.Secret)
					assert.Contains(t, resultErr, tt.outErr)
				} else {
					assert.NotNil(t, err)
				}

			case service.TokenSecretRequest:
				req, err = service.DecodeRequest(resultType)(context.TODO(), tt.in)
				if err != nil {
					resultErr = err.Error()
				}

				result, ok := req.(service.TokenSecretRequest)
				assert.True(t, ok)

				assert.Equal(t, tt.outToken, result.Token)
				assert.Equal(t, tt.outSecret, result.Secret)
				assert.Contains(t, resultErr, tt.outErr)

			case service.Token:
				req, err = service.DecodeRequest(resultType)(context.TODO(), tt.in)
				if err != nil {
					resultErr = err.Error()
				}

				result, ok := req.(service.Token)
				assert.True(t, ok)

				assert.Equal(t, tt.outToken, result.Token)
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
