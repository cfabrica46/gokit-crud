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
				assert.Fail(t, "result is not of the type indicated")
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

func TestDecodeGetUserByIDRequest(t *testing.T) {
	t.Parallel()

	dataJSON, err := json.Marshal(service.GetUserByIDRequest{ID: idTest})
	if err != nil {
		assert.Error(t, err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, "localhost:8080", bytes.NewBuffer(dataJSON))
	if err != nil {
		assert.Error(t, err)
	}

	badReq, err := http.NewRequest(http.MethodGet, "localhost:8080", bytes.NewBuffer([]byte{}))
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name   string
		in     *http.Request
		outErr string
		outID  int
	}{
		{
			name:   nameNoError,
			in:     goodReq,
			outID:  idTest,
			outErr: "",
		},
		{
			name:   "ErrRequestBody",
			in:     badReq,
			outID:  0,
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result service.GetUserByIDRequest
			var resultID int
			var resultErr string

			r, err := service.DecodeGetUserByIDRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.GetUserByIDRequest)
			if !ok {
				if tt.name == nameNoError {
					assert.Fail(t, "result is not of the type indicated")
				}
			}

			resultID = result.ID

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			assert.Equal(t, tt.outID, resultID)
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
		name        string
		in          *http.Request
		outUsername string
		outPassword string
		outErr      string
	}{
		{
			name:        nameNoError,
			in:          goodReq,
			outUsername: usernameTest,
			outPassword: passwordTest,
			outErr:      "",
		},
		{
			name:        "ErrRequestBody",
			in:          badReq,
			outUsername: "",
			outPassword: "",
			outErr:      "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultUsername, resultPassword string
			var resultErr string

			r, err := service.DecodeGetUserByUsernameAndPasswordRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.GetUserByUsernameAndPasswordRequest)
			if !ok {
				if tt.name == nameNoError {
					assert.Fail(t, "result is not of the type indicated")
				}
			}

			resultUsername = result.Username
			resultPassword = result.Password

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			assert.Equal(t, tt.outUsername, resultUsername)
			assert.Equal(t, tt.outPassword, resultPassword)
		})
	}
}

func TestDecodeGetIDByUsernameRequest(t *testing.T) {
	t.Parallel()

	dataJSON, err := json.Marshal(service.GetIDByUsernameRequest{Username: usernameTest})
	if err != nil {
		assert.Error(t, err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, "localhost:8080", bytes.NewBuffer(dataJSON))
	if err != nil {
		assert.Error(t, err)
	}

	badReq, err := http.NewRequest(http.MethodGet, "localhost:8080", bytes.NewBuffer([]byte{}))
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name        string
		in          *http.Request
		outErr      string
		outUsername string
	}{
		{
			name:        nameNoError,
			in:          goodReq,
			outUsername: usernameTest,
			outErr:      "",
		},
		{
			name:        "ErrRequestBody",
			in:          badReq,
			outUsername: "",
			outErr:      "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultUsername string
			var resultErr string

			r, err := service.DecodeGetIDByUsernameRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.GetIDByUsernameRequest)
			if !ok {
				if tt.name == nameNoError {
					assert.Fail(t, "result is not of the type indicated")
				}
			}

			resultUsername = result.Username

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			assert.Equal(t, tt.outUsername, resultUsername)
		})
	}
}

func TestDecodeInsertUserRequest(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/user"

	dataJSON, err := json.Marshal(service.InsertUserRequest{usernameTest, passwordTest, emailTest})
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
		name        string
		in          *http.Request
		outUsername string
		outPassword string
		outEmail    string
		outErr      string
	}{
		{
			name:        nameNoError,
			in:          goodReq,
			outUsername: usernameTest,
			outPassword: passwordTest,
			outEmail:    emailTest,
			outErr:      "",
		},
		{
			name:        "ErrRequestBody",
			in:          badReq,
			outUsername: "",
			outPassword: "",
			outEmail:    "",
			outErr:      "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultUsername, resultPassword, resultEmail string
			var resultErr string

			r, err := service.DecodeInsertUserRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.InsertUserRequest)
			if !ok {
				if tt.name == nameNoError {
					assert.Fail(t, "result is not of the type indicated")
				}
			}

			resultUsername = result.Username
			resultPassword = result.Password
			resultEmail = result.Email

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			assert.Equal(t, tt.outUsername, resultUsername)
			assert.Equal(t, tt.outPassword, resultPassword)
			assert.Equal(t, tt.outEmail, resultEmail)
		})
	}
}

func TestDecodeDeleteUserRequest(t *testing.T) {
	t.Parallel()

	url := "localhost:8080/user"

	dataJSON, err := json.Marshal(service.DeleteUserRequest{idTest})
	if err != nil {
		assert.Error(t, err)
	}

	goodReq, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		assert.Error(t, err)
	}

	badReq, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name   string
		in     *http.Request
		outErr string
		outID  int
	}{
		{
			name:   nameNoError,
			in:     goodReq,
			outID:  idTest,
			outErr: "",
		},
		{
			name:   "ErrRequestBody",
			in:     badReq,
			outID:  0,
			outErr: "EOF",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultID int
			var resultErr string

			r, err := service.DecodeDeleteUserRequest(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}
			result, ok := r.(service.DeleteUserRequest)
			if !ok {
				if tt.name == nameNoError {
					assert.Fail(t, "result is not of the type indicated")
				}
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}

			resultID = result.ID

			assert.Equal(t, tt.outID, resultID)
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
