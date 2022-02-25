package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestMakeGenerateTokenEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in  GenerateTokenRequest
		out string
	}{
		{GenerateTokenRequest{1, "cesar", "cesar@email.com", "secret"}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			mr, err := miniredis.Run()
			if err != nil {
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := GetService(client)

			r, err := MakeGenerateTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GenerateTokenResponse)
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
	for i, tt := range []struct {
		in     ExtractTokenRequest
		outErr string
	}{
		{ExtractTokenRequest{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNlc2FyQGVtYWlsLmNvbSIsImlkIjoxLCJ1c2VybmFtZSI6ImNlc2FyIiwidXVpZCI6IjcxNzFjZTU2LWIwMzYtNDEzMi1hMDljLWQyZmZiMzgzYjdjMSJ9.V_vEFyz6OAc5eOFgt589CC0OCFf72BU5MuBg2IRl4dg", "secret"}, ""},
		{ExtractTokenRequest{"", "secret"}, "token contains an invalid number of segments"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			mr, err := miniredis.Run()
			if err != nil {
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := GetService(client)

			r, err := MakeExtractTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(ExtractTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestMakeSetTokenEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     SetTokenRequest
		outErr string
	}{
		{SetTokenRequest{"token"}, ""},
		{SetTokenRequest{""}, "redis: client is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			mr, err := miniredis.Run()
			if err != nil {
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := GetService(client)

			// Generate Conflict
			if tt.outErr == "redis: client is closed" {
				svc.db.Close()
			}

			r, err := MakeSetTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(SetTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestMakeDeleteTokenEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     DeleteTokenRequest
		outErr string
	}{
		{DeleteTokenRequest{"token"}, ""},
		{DeleteTokenRequest{""}, "redis: client is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			mr, err := miniredis.Run()
			if err != nil {
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := GetService(client)

			// Generate Conflict
			if tt.outErr == "redis: client is closed" {
				svc.db.Close()
			}

			r, err := MakeDeleteTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(DeleteTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestMakeCheckTokenEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     CheckTokenRequest
		outErr string
	}{
		{CheckTokenRequest{"token"}, ""},
		{CheckTokenRequest{""}, "redis: client is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			mr, err := miniredis.Run()
			if err != nil {
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := GetService(client)

			// Generate Conflict
			if tt.outErr == "redis: client is closed" {
				svc.db.Close()
			}

			r, err := MakeCheckTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(CheckTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}
