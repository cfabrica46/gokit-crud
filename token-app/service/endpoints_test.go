package service_test

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/cfabrica46/gokit-crud/token-app/service"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type incorrectRequest struct {
	incorrect bool
}

func TestMakeGenerateTokenEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		in     any
		name   string
		outErr string
	}{
		{
			name: nameNoError,
			in: service.GenerateTokenRequest{
				ID:       idTest,
				Username: usernameTest,
				Email:    emailTest,
				Secret:   secretTest,
			},
			outErr: "",
		},
		{
			name: nameErrorRequest,
			in: incorrectRequest{
				incorrect: true,
			},
			outErr: "isn't of type",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			mr, err := miniredis.Run()
			if err != nil {
				assert.Error(t, err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			r, err := service.MakeGenerateTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.Token)
			if !ok {
				if tt.name != nameErrorRequest {
					assert.Fail(t, "response is not of the type indicated")
				}
			}

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
				assert.NotEmpty(t, result.Token)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
				assert.Empty(t, result.Token)
			}
		})
	}
}

func TestMakeExtractTokenEndpoint(t *testing.T) {
	t.Parallel()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       idTest,
		"username": usernameTest,
		"email":    emailTest,
		"uuid":     uuid.NewString(),
	})

	tokenSigned, err := token.SignedString([]byte(secretTest))
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name   string
		in     any
		outErr string
	}{
		{
			name: nameNoError,
			in: service.ExtractTokenRequest{
				tokenSigned,
				secretTest,
			},
			outErr: "",
		},
		{
			name: nameErrorRequest,
			in: incorrectRequest{
				incorrect: true,
			},
			outErr: "isn't of type",
		},
		{
			name: "ErrorNotValidToken",
			in: service.ExtractTokenRequest{
				"",
				secretTest,
			},
			outErr: "token contains an invalid number of segments",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			mr, err := miniredis.Run()
			if err != nil {
				assert.Error(t, err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			r, err := service.MakeExtractTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.ExtractTokenResponse)
			if !ok {
				if tt.name != nameErrorRequest {
					assert.Fail(t, "response is not of the type indicated")
				}
			} else {
				resultErr = result.Err
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}
		})
	}
}

func TestMakeManageTokenEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		inState service.State
		in      any
		name    string
		outErr  string
	}{
		{
			name:    nameNoError,
			in:      service.Token{"token"},
			inState: service.NewSetTokenState(),
			outErr:  "",
		},
		{
			name: nameErrorRequest,
			in: incorrectRequest{
				incorrect: true,
			},
			outErr: "isn't of type",
		},
		{
			name:    nameErrorRedisClose,
			in:      service.Token{""},
			inState: service.NewSetTokenState(),
			outErr:  errRedisClosed,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			mr, err := miniredis.Run()
			if err != nil {
				assert.Error(t, err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			if tt.name == nameErrorRedisClose {
				svc.DB.Close()
			}

			r, err := service.MakeManageTokenEndpoint(svc, tt.inState)(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.ErrorResponse)
			if !ok {
				if tt.name != nameErrorRequest {
					assert.Fail(t, "response is not of the type indicated")
				}
			} else {
				resultErr = result.Err
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}
		})
	}
}

func TestMakeCheckTokenEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		in     any
		outErr string
	}{
		{
			name:   nameNoError,
			in:     service.Token{"token"},
			outErr: "",
		},
		{
			name: nameErrorRequest,
			in: incorrectRequest{
				incorrect: true,
			},
			outErr: "isn't of type",
		},
		{
			name:   nameErrorRedisClose,
			in:     service.Token{""},
			outErr: errRedisClosed,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			mr, err := miniredis.Run()
			if err != nil {
				assert.Error(t, err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			if tt.name == nameErrorRedisClose {
				svc.DB.Close()
			}

			r, err := service.MakeCheckTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.CheckTokenResponse)
			if !ok {
				if tt.name != nameErrorRequest {
					assert.Fail(t, "response is not of the type indicated")
				}
			} else {
				resultErr = result.Err
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, resultErr, tt.outErr)
			}
		})
	}
}
