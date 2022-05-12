package service_test

import (
	"context"
	"strings"
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
			out: "",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mr, err := miniredis.Run()
			if err != nil {
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			r, err := service.MakeGenerateTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.GenerateTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if result.Token == "" {
				t.Error("token its empty")
			}

			assert.NotEqual(t, tt.out, result.Token, "they shouldn't be equal")
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
		t.Error(err)
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
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			r, err := service.MakeExtractTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.ExtractTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if tt.outErr != "" {
				if !strings.Contains(result.Err, tt.outErr) {
					t.Errorf("want %v; got %v", tt.outErr, result.Err)
				}
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
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			// Generate Conflict.
			if tt.outErr == errRedisClosed {
				svc.DB.Close()
			}

			r, err := service.MakeSetTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.SetTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if tt.outErr != "" {
				if !strings.Contains(result.Err, tt.outErr) {
					t.Errorf("want %v; got %v", tt.outErr, result.Err)
				}
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
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			// Generate Conflict.
			if tt.outErr == errRedisClosed {
				svc.DB.Close()
			}

			r, err := service.MakeDeleteTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.DeleteTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Contains(t, result.Err, tt.outErr)
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
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			// Generate Conflict.
			if tt.outErr == errRedisClosed {
				svc.DB.Close()
			}

			r, err := service.MakeCheckTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.CheckTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if tt.outErr != "" {
				if !strings.Contains(result.Err, tt.outErr) {
					t.Errorf("want %v; got %v", tt.outErr, result.Err)
				}
			}
		})
	}
}
