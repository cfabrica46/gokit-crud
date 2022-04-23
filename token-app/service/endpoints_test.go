package service_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/cfabrica46/gokit-crud/token-app/service"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMakeGenerateTokenEndpoint(t *testing.T) {
	for indx, tt := range []struct {
		out string
		in  service.GenerateTokenRequest
	}{
		{
			in: service.GenerateTokenRequest{
				ID:       idTest,
				Username: usernameTest,
				Email:    emailTest,
				Secret:   secretTest,
			},
			out: "",
		},
	} {
		t.Run(fmt.Sprintf("%v", indx), func(t *testing.T) {
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
		{
			in: service.ExtractTokenRequest{
				tokenSigned,
				secretTest,
			},
			outErr: "",
		},
		{
			in: service.ExtractTokenRequest{
				"",
				secretTest,
			},
			outErr: "token contains an invalid number of segments",
		},
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
	for indx, tt := range []struct {
		outErr string
		in     service.SetTokenRequest
	}{
		{
			in:     service.SetTokenRequest{"token"},
			outErr: "",
		},
		{
			in:     service.SetTokenRequest{""},
			outErr: errRedisClosed,
		},
	} {
		t.Run(fmt.Sprintf("%v", indx), func(t *testing.T) {
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
	for indx, tt := range []struct {
		in     service.DeleteTokenRequest
		outErr string
	}{
		{
			in:     service.DeleteTokenRequest{"token"},
			outErr: "",
		},
		{
			in:     service.DeleteTokenRequest{""},
			outErr: errRedisClosed,
		},
	} {
		t.Run(strconv.Itoa(indx), func(t *testing.T) {
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
	for indx, tt := range []struct {
		in     service.CheckTokenRequest
		outErr string
	}{
		{
			in:     service.CheckTokenRequest{"token"},
			outErr: "",
		},
		{
			in:     service.CheckTokenRequest{""},
			outErr: errRedisClosed,
		},
	} {
		t.Run(fmt.Sprintf("%v", indx), func(t *testing.T) {
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
