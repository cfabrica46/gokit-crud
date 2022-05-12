package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cfabrica46/gokit-crud/token-app/service"
)

func TestDecodeGenerateToken(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/generate"

	dataJSON, err := json.Marshal(service.GenerateTokenRequest{
		ID:       idTest,
		Username: usernameTest,
		Email:    emailTest,
		Secret:   secretTest,
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

	for _, tt := range []struct {
		name   string
		in     *http.Request
		outErr string
		out    service.GenerateTokenRequest
	}{
		{
			name: "NoError",
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

			r, err := service.DecodeGenerateTokenRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.GenerateTokenRequest)
			if !ok {
				if (tt.out != service.GenerateTokenRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if !strings.Contains(resultErr, tt.outErr) {
				t.Errorf("want %v; got %v", tt.outErr, resultErr)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeExtractToken(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/extract"

	dataJSON, err := json.Marshal(service.ExtractTokenRequest{tokenTest, secretTest})
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

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.ExtractTokenRequest
		outErr string
	}{
		{
			name: "NoError",
			in:   goodReq,
			out: service.ExtractTokenRequest{
				Token:  tokenTest,
				Secret: secretTest,
			},
			outErr: "",
		},
		{
			name: "BadRequest",
			in:   badReq,
			out:  service.ExtractTokenRequest{}, outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.ExtractTokenRequest
			var resultErr string

			r, err := service.DecodeExtractTokenRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.ExtractTokenRequest)
			if !ok {
				if (tt.out != service.ExtractTokenRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if !strings.Contains(resultErr, tt.outErr) {
				t.Errorf("want %v; got %v", tt.outErr, resultErr)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeSetToken(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/token"

	dataJSON, err := json.Marshal(service.SetTokenRequest{tokenTest})
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

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.SetTokenRequest
		outErr string
	}{
		{
			name: "NoError",
			in:   goodReq,
			out: service.SetTokenRequest{
				Token: tokenTest,
			},
			outErr: "",
		},
		{
			name:   "BadRequest",
			in:     badReq,
			out:    service.SetTokenRequest{},
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.SetTokenRequest
			var resultErr string

			r, err := service.DecodeSetTokenRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.SetTokenRequest)
			if !ok {
				if (tt.out != service.SetTokenRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if !strings.Contains(resultErr, tt.outErr) {
				t.Errorf("want %v; got %v", tt.outErr, resultErr)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeDeleteToken(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/token"

	dataJSON, err := json.Marshal(service.DeleteTokenRequest{tokenTest})
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

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.DeleteTokenRequest
		outErr string
	}{
		{
			name: "NoError",
			in:   goodReq,
			out: service.DeleteTokenRequest{
				Token: tokenTest,
			},
			outErr: "",
		},
		{
			name:   "BadRequest",
			in:     badReq,
			out:    service.DeleteTokenRequest{},
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.DeleteTokenRequest
			var resultErr string

			r, err := service.DecodeDeleteTokenRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.DeleteTokenRequest)
			if !ok {
				if (tt.out != service.DeleteTokenRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if !strings.Contains(resultErr, tt.outErr) {
				t.Errorf("want %v; got %v", tt.outErr, resultErr)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeCheckToken(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/Check"

	dataJSON, err := json.Marshal(service.CheckTokenRequest{tokenTest})
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

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.CheckTokenRequest
		outErr string
	}{
		{
			name:   "NoError",
			in:     goodReq,
			out:    service.CheckTokenRequest{tokenTest},
			outErr: "",
		},
		{
			name:   "BadRequest",
			in:     badReq,
			out:    service.CheckTokenRequest{},
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.CheckTokenRequest
			var resultErr string

			r, err := service.DecodeCheckTokenRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.CheckTokenRequest)
			if !ok {
				if (tt.out != service.CheckTokenRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if !strings.Contains(resultErr, tt.outErr) {
				t.Errorf("want %v; got %v", tt.outErr, resultErr)
			}
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
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
			if !strings.Contains(resultErr, tt.outErr) {
				t.Errorf("want %v; got %v", tt.outErr, resultErr)
			}
		})
	}
}
