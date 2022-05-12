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
			name: "NoError",
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

			result, ok := r.(service.GenerateTokenResponse)
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
			name: "NoError",
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

			if tt.name == "NoError" {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, result.Err, tt.outErr)
			}
		})
	}
}

func TestMakeSetTokenEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		outErr string
		in     service.SetTokenRequest
	}{
		{
			name:   "NoError",
			in:     service.SetTokenRequest{"token"},
			outErr: "",
		},
		{
			name:   "ErrorRedisClose",
			in:     service.SetTokenRequest{""},
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

			// Generate Conflict.
			if tt.outErr == errRedisClosed {
				svc.DB.Close()
			}

			r, err := service.MakeSetTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				assert.Error(t, err)
			}

			result, ok := r.(service.SetTokenResponse)
			if !ok {
				assert.Fail(t, "response is not of the type indicated")
			}

			if tt.name == "NoError" {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, result.Err, tt.outErr)
			}
		})
	}
}

func TestMakeDeleteTokenEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		in     service.DeleteTokenRequest
		outErr string
	}{
		{
			name:   "NoError",
			in:     service.DeleteTokenRequest{"token"},
			outErr: "",
		},
		{
			name:   "ErrorRedisClose",
			in:     service.DeleteTokenRequest{""},
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

			// Generate Conflict.
			if tt.outErr == errRedisClosed {
				svc.DB.Close()
			}

			r, err := service.MakeDeleteTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				assert.Error(t, err)
			}

			result, ok := r.(service.DeleteTokenResponse)
			if !ok {
				assert.Fail(t, "response is not of the type indicated")
			}

			if tt.name == "NoError" {
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
		in     service.CheckTokenRequest
		outErr string
	}{
		{
			name:   "NoError",
			in:     service.CheckTokenRequest{"token"},
			outErr: "",
		},
		{
			name:   "ErrorRedisClose",
			in:     service.CheckTokenRequest{""},
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

			// Generate Conflict.
			if tt.outErr == errRedisClosed {
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

			if tt.name == "NoError" {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, result.Err, tt.outErr)
			}
		})
	}
}
