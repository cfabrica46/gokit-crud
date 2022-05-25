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

func TestMakeGenerateTokenEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		out  string
		in   service.GenerateTokenRequest
	}{
		{
			name: nameNoError,
			in: service.GenerateTokenRequest{
				ID:       idTest,
				Username: usernameTest,
				Email:    emailTest,
				Secret:   secretTest,
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mr, err := miniredis.Run()
			if err != nil {
				assert.Error(t, err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			r, err := service.MakeGenerateTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				assert.Error(t, err)
			}

			result, ok := r.(service.Token)
			if !ok {
				assert.Fail(t, "response is not of the type indicated")
			}

			assert.NotEmpty(t, result.Token)
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
		in     service.ExtractTokenRequest
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

			mr, err := miniredis.Run()
			if err != nil {
				assert.Error(t, err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			r, err := service.MakeExtractTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				assert.Error(t, err)
			}

			result, ok := r.(service.ExtractTokenResponse)
			if !ok {
				assert.Fail(t, "response is not of the type indicated")
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, result.Err, tt.outErr)
			}
		})
	}
}

func TestMakeManageTokenEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name    string
		outErr  string
		inState service.State
		in      service.Token
	}{
		{
			name:    nameNoError,
			in:      service.Token{"token"},
			inState: service.NewSetTokenState(),
			outErr:  "",
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
				assert.Error(t, err)
			}

			result, ok := r.(service.ErrorResponse)
			if !ok {
				assert.Fail(t, "response is not of the type indicated")
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, result.Err, tt.outErr)
			}
		})
	}
}

func TestMakeCheckTokenEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		in     service.Token
		outErr string
	}{
		{
			name:   nameNoError,
			in:     service.Token{"token"},
			outErr: "",
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
				assert.Error(t, err)
			}

			result, ok := r.(service.CheckTokenResponse)
			if !ok {
				assert.Fail(t, "response is not of the type indicated")
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, result.Err, tt.outErr)
			}
		})
	}
}
