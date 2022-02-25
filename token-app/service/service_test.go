package service

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	for i, tt := range []struct {
		inID                int
		inUsername, inEmail string
		inSecret            []byte
		outToken, outErr    string
	}{
		{1, "cesar", "cesar@email.com", []byte("secret"), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.", ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result string

			mr, err := miniredis.Run()
			if err != nil {
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := GetService(client)

			result = svc.GenerateToken(tt.inID, tt.inUsername, tt.inEmail, tt.inSecret)

			if !strings.Contains(result, tt.outToken) {
				t.Errorf("want %v; got %v", tt.outToken, result)
			}
		})
	}
}

func TestExtractToken(t *testing.T) {
	for i, tt := range []struct {
		inToken                       string
		inSecret                      []byte
		outID                         int
		outUsername, outEmail, outErr string
	}{
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNlc2FyQGVtYWlsLmNvbSIsImlkIjoxLCJ1c2VybmFtZSI6ImNlc2FyIiwidXVpZCI6IjcxNzFjZTU2LWIwMzYtNDEzMi1hMDljLWQyZmZiMzgzYjdjMSJ9.V_vEFyz6OAc5eOFgt589CC0OCFf72BU5MuBg2IRl4dg", []byte("secret"), 1, "cesar", "cesar@email.com", ""},
		{"", nil, 0, "", "", "token contains an invalid number of segments"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultID int
			var resultUsername, resultEmail, resultErr string

			mr, err := miniredis.Run()
			if err != nil {
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := GetService(client)

			resultID, resultUsername, resultEmail, err = svc.ExtractToken(tt.inToken, tt.inSecret)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outID, resultID, "they should be equal")
			assert.Equal(t, tt.outUsername, resultUsername, "they should be equal")
			assert.Equal(t, tt.outEmail, resultEmail, "they should be equal")
			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestSetToken(t *testing.T) {
	for i, tt := range []struct {
		in     string
		outErr string
	}{
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNlc2FyQGVtYWlsLmNvbSIsImlkIjoxLCJ1c2VybmFtZSI6ImNlc2FyIiwidXVpZCI6IjcxNzFjZTU2LWIwMzYtNDEzMi1hMDljLWQyZmZiMzgzYjdjMSJ9.V_vEFyz6OAc5eOFgt589CC0OCFf72BU5MuBg2IRl4dg", ""},
		{"", "redis: client is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

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

			err = svc.SetToken(tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestDeleteToken(t *testing.T) {
	for i, tt := range []struct {
		in     string
		outErr string
	}{
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNlc2FyQGVtYWlsLmNvbSIsImlkIjoxLCJ1c2VybmFtZSI6ImNlc2FyIiwidXVpZCI6IjcxNzFjZTU2LWIwMzYtNDEzMi1hMDljLWQyZmZiMzgzYjdjMSJ9.V_vEFyz6OAc5eOFgt589CC0OCFf72BU5MuBg2IRl4dg", ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mr, err := miniredis.Run()
			if err != nil {
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := GetService(client)

			err = svc.DeleteToken(tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestCheckToken(t *testing.T) {
	for i, tt := range []struct {
		in       string
		outCheck bool
		outErr   string
	}{
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNlc2FyQGVtYWlsLmNvbSIsImlkIjoxLCJ1c2VybmFtZSI6ImNlc2FyIiwidXVpZCI6IjcxNzFjZTU2LWIwMzYtNDEzMi1hMDljLWQyZmZiMzgzYjdjMSJ9.V_vEFyz6OAc5eOFgt589CC0OCFf72BU5MuBg2IRl4dg", true, ""},
		{"", false, ""},
		{"", false, "redis: client is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultCheck bool
			var resultErr string

			mr, err := miniredis.Run()
			if err != nil {
				t.Error(err)
			}

			client := redis.NewClient(&redis.Options{Addr: mr.Addr()})

			svc := GetService(client)

			// insert
			if tt.in != "" {
				err = svc.SetToken(tt.in)
				if err != nil {
					t.Error(err)
				}
			}

			// Generate Conflict
			if tt.outErr == "redis: client is closed" {
				svc.db.Close()
			}

			resultCheck, err = svc.CheckToken(tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outCheck, resultCheck, "they should be equal")
			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestKeyFunc(t *testing.T) {
	for i, tt := range []struct {
		inSecret  []byte
		outSecret []byte
		outErr    string
	}{
		{[]byte("secret"), []byte("secret"), "unexpected signing method: PS256"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			kf := keyFunc(tt.inSecret)

			//generateToken
			token := jwt.NewWithClaims(jwt.SigningMethodPS256, jwt.MapClaims{
				"id":       1,
				"username": "cesar",
				"email":    "cesar@email.com",
				"uuid":     uuid.NewString(),
			})

			_, err := kf(token)
			if err != nil {
				resultErr = err.Error()
			}

			if !strings.Contains(resultErr, tt.outErr) {
				t.Errorf("want %v; got %v", tt.outErr, resultErr)
			}
		})
	}
}
