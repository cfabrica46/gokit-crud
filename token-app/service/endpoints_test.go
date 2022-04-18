package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/cfabrica46/gokit-crud/token-app/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMakeGenerateTokenEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in  service.GenerateTokenRequest
		out string
	}{
		{service.GenerateTokenRequest{1, "cesar", "cesar@email.com", "secret"}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       idTest,
		"username": usernameTest,
		"email":    emailTest,
		"uuid":     uuid.NewString(),
	})

	tokenSigned, _ := token.SignedString([]byte(secretTest))

	for indx, tt := range []struct {
		in     service.ExtractTokenRequest
		outErr string
	}{
		{service.ExtractTokenRequest{tokenSigned, secretTest}, ""},
		{service.ExtractTokenRequest{
			"",
			secretTest,
		}, "token contains an invalid number of segments"},
	} {
		t.Run(fmt.Sprintf("%v", indx), func(t *testing.T) {
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

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestMakeSetTokenEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     service.SetTokenRequest
		outErr string
	}{
		{service.SetTokenRequest{"token"}, ""},
		{service.SetTokenRequest{""}, errRedisClosed},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			mr, err := miniredis.Run()
			if err != nil {
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			// Generate Conflict
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

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestMakeDeleteTokenEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     service.DeleteTokenRequest
		outErr string
	}{
		{service.DeleteTokenRequest{"token"}, ""},
		{service.DeleteTokenRequest{""}, errRedisClosed},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			mr, err := miniredis.Run()
			if err != nil {
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			// Generate Conflict
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

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestMakeCheckTokenEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     service.CheckTokenRequest
		outErr string
	}{
		{service.CheckTokenRequest{"token"}, ""},
		{service.CheckTokenRequest{""}, errRedisClosed},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			mr, err := miniredis.Run()
			if err != nil {
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := service.GetService(client)

			// Generate Conflict
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

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}
