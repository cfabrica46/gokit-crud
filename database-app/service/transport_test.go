package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cfabrica46/gokit-crud/database-app/service"
	"github.com/stretchr/testify/assert"
)

func TestDecodeGetAllUsersRequest(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		in     *http.Request
		out    service.GetAllUsersRequest
		outErr string
	}{
		{
			name:   nameNoError,
			in:     &http.Request{},
			out:    service.GetAllUsersRequest{},
			outErr: "",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.GetAllUsersRequest
			var resultErr string

			r, err := service.DecodeGetAllUsersRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.GetAllUsersRequest)
			if !ok {
				t.Error("result is not of the type indicated")
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeGetUserByIDRequest(t *testing.T) {
	t.Parallel()

	dataJSON, err := json.Marshal(service.GetUserByIDRequest{ID: idTest})
	if err != nil {
		t.Error(err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, "localhost:8080", bytes.NewBuffer(dataJSON))
	if err != nil {
		t.Error(err)
	}

	badReq, err := http.NewRequest(http.MethodGet, "localhost:8080", bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Error(err)
	}

	for _, tt := range []struct {
		name   string
		in     *http.Request
		outErr string
		out    service.GetUserByIDRequest
	}{
		{
			name: nameNoError,
			in:   goodReq,
			out: service.GetUserByIDRequest{
				ID: idTest,
			},
			outErr: "",
		},
		{
			name: "ErrRequestBody",
			in:   badReq,
			out: service.GetUserByIDRequest{
				ID: 0,
			},
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.GetUserByIDRequest
			var resultErr string

			r, err := service.DecodeGetUserByIDRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.GetUserByIDRequest)
			if !ok {
				if (tt.out != service.GetUserByIDRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeGetUserByUsernameAndPasswordRequest(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/user/username_password"

	dataJSON, err := json.Marshal(service.GetUserByUsernameAndPasswordRequest{
		Username: usernameTest,
		Password: passwordTest,
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
		out    service.GetUserByUsernameAndPasswordRequest
		outErr string
	}{
		{
			name: nameNoError,
			in:   goodReq,
			out: service.GetUserByUsernameAndPasswordRequest{
				Username: usernameTest,
				Password: passwordTest,
			},
			outErr: "",
		},
		{
			name:   "ErrRequestBody",
			in:     badReq,
			out:    service.GetUserByUsernameAndPasswordRequest{},
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.GetUserByUsernameAndPasswordRequest
			var resultErr string

			r, err := service.DecodeGetUserByUsernameAndPasswordRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.GetUserByUsernameAndPasswordRequest)
			if !ok {
				if (tt.out != service.GetUserByUsernameAndPasswordRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeGetIDByUsernameRequest(t *testing.T) {
	t.Parallel()

	dataJSON, err := json.Marshal(service.GetIDByUsernameRequest{Username: usernameTest})
	if err != nil {
		t.Error(err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, "localhost:8080", bytes.NewBuffer(dataJSON))
	if err != nil {
		t.Error(err)
	}

	badReq, err := http.NewRequest(http.MethodGet, "localhost:8080", bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Error(err)
	}

	for _, tt := range []struct {
		name   string
		in     *http.Request
		outErr string
		out    service.GetIDByUsernameRequest
	}{
		{
			name: nameNoError,
			in:   goodReq,
			out: service.GetIDByUsernameRequest{
				Username: usernameTest,
			},
			outErr: "",
		},
		{
			name:   "ErrRequestBody",
			in:     badReq,
			out:    service.GetIDByUsernameRequest{},
			outErr: "",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.GetIDByUsernameRequest
			var resultErr string

			r, err := service.DecodeGetIDByUsernameRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.GetIDByUsernameRequest)
			if !ok {
				if (tt.out != service.GetIDByUsernameRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeInsertUserRequest(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/user"

	dataJSON, err := json.Marshal(service.InsertUserRequest{usernameTest, "0idTest234", emailTest})
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
		out    service.InsertUserRequest
		outErr string
	}{
		{
			name: nameNoError,
			in:   goodReq,
			out: service.InsertUserRequest{
				Username: usernameTest,
				Password: "0idTest234",
				Email:    emailTest,
			},
			outErr: "",
		},
		{
			name:   "ErrRequestBody",
			in:     badReq,
			out:    service.InsertUserRequest{},
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.InsertUserRequest
			var resultErr string

			r, err := service.DecodeInsertUserRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.InsertUserRequest)
			if !ok {
				if (tt.out != service.InsertUserRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDecodeDeleteUserRequest(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/user"

	dataJSON, err := json.Marshal(service.DeleteUserRequest{idTest})
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
		outErr string
		out    service.DeleteUserRequest
	}{
		{
			name: nameNoError,
			in:   goodReq,
			out: service.DeleteUserRequest{
				ID: idTest,
			},
			outErr: "",
		},
		{
			name:   "ErrRequestBody",
			in:     badReq,
			out:    service.DeleteUserRequest{},
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.DeleteUserRequest
			var resultErr string

			r, err := service.DecodeDeleteUserRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.DeleteUserRequest)
			if !ok {
				if (tt.out != service.DeleteUserRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
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
