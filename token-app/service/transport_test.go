package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDecodeGenerateToken(t *testing.T) {
	url := "localhost:8080/generate"

	dataJSON, err := json.Marshal(struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Secret   string `json:"secret"`
	}{
		1,
		"cesar",
		"cesar@email.com",
		"secret",
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
		out      *generateTokenRequest
		outError string
	}{
		{goodReq, &generateTokenRequest{1, "cesar", "cesar@email.com", "secret"}, ""},
		{badReq, nil, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result interface{}
			var resultErr string

			result, err = DecodeGenerateTokenRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, resultErr)
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

func TestDecodeExtractToken(t *testing.T) {
	url := "localhost:8080/extract"

	dataJSON, err := json.Marshal(struct {
		Token  string `json:"token"`
		Secret string `json:"secret"`
	}{
		"token",
		"secret",
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
		out      *extractTokenRequest
		outError string
	}{
		{goodReq, &extractTokenRequest{"token", "secret"}, ""},
		{badReq, nil, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result interface{}
			var resultErr string

			result, err = DecodeExtractTokenRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, resultErr)
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

func TestDecodeSetToken(t *testing.T) {
	url := "localhost:8080/token"

	dataJSON, err := json.Marshal(struct {
		Token string `json:"token"`
	}{
		"token",
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
		out      *setTokenRequest
		outError string
	}{
		{goodReq, &setTokenRequest{"token"}, ""},
		{badReq, nil, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result interface{}
			var resultErr string

			result, err = DecodeSetTokenRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, resultErr)
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

func TestDecodeDeleteToken(t *testing.T) {
	url := "localhost:8080/token"

	dataJSON, err := json.Marshal(struct {
		Token string `json:"token"`
	}{
		"token",
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
		out      *deleteTokenRequest
		outError string
	}{
		{goodReq, &deleteTokenRequest{"token"}, ""},
		{badReq, nil, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result interface{}
			var resultErr string

			result, err = DecodeDeleteTokenRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, resultErr)
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

func TestDecodeCheckToken(t *testing.T) {
	url := "localhost:8080/Check"

	dataJSON, err := json.Marshal(struct {
		Token string `json:"token"`
	}{
		"token",
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
		out      *checkTokenRequest
		outError string
	}{
		{goodReq, &checkTokenRequest{"token"}, ""},
		{badReq, nil, "EOF"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result interface{}
			var resultErr string

			result, err = DecodeCheckTokenRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, resultErr)
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
			var resultErr string

			err := EncodeResponse(context.TODO(), httptest.NewRecorder(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			if !strings.Contains(resultErr, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, resultErr)
			}
		})
	}
}
