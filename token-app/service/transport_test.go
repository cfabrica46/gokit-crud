package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cfabrica46/gokit-crud/token-app/service"
	"github.com/stretchr/testify/assert"
)

func TestDecodeRequest(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/generate"

	dataJSON, err := json.Marshal(service.GenerateTokenRequest{
		ID:       idTest,
		Username: usernameTest,
		Email:    emailTest,
		Secret:   secretTest,
	})
	if err != nil {
		assert.Error(t, err)
	}

	goodReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		assert.Error(t, err)
	}

	badReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name   string
		in     *http.Request
		outErr string
		out    service.GenerateTokenRequest
	}{
		{
			name: nameNoError,
			in:   goodReq,
			out: service.GenerateTokenRequest{
				ID:       idTest,
				Username: usernameTest,
				Email:    emailTest,
				Secret:   secretTest,
			},
			outErr: "",
		},
		{
			name:   "BadRequest",
			in:     badReq,
			out:    service.GenerateTokenRequest{},
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.GenerateTokenRequest
			var resultErr string

			r, err := service.DecodeRequest(service.GenerateTokenRequest{})(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			log.Println(r)

			result, ok := r.(service.GenerateTokenRequest)
			if !ok {
				if (tt.out != service.GenerateTokenRequest{}) {
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
