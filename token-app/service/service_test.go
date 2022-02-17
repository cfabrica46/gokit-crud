package service

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
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
			svc := GetService("localhost", "6379")

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
		{"", nil, 0, "", "", ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultID int
			var resultUsername, resultEmail, resultErr string
			var err error

			svc := GetService("localhost", "6379")
			resultID, resultUsername, resultEmail, err = svc.ExtractToken(tt.inToken, tt.inSecret)
			if err != nil {
				resultErr = err.Error()
			}

			if resultID != tt.outID && resultUsername != tt.outUsername && resultEmail != tt.outEmail {
				t.Errorf("want %v %v %v; got %v %v %v", tt.outErr, tt.outUsername, tt.outEmail, resultID, resultUsername, resultEmail)
			}

			if !strings.Contains(resultErr, tt.outErr) {
				t.Errorf("want %v; got %v", tt.outErr, resultErr)
			}
		})
	}
}

func TestSetToken(t *testing.T) {
	for i, tt := range []struct {
		in  string
		out string
	}{
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNlc2FyQGVtYWlsLmNvbSIsImlkIjoxLCJ1c2VybmFtZSI6ImNlc2FyIiwidXVpZCI6IjcxNzFjZTU2LWIwMzYtNDEzMi1hMDljLWQyZmZiMzgzYjdjMSJ9.V_vEFyz6OAc5eOFgt589CC0OCFf72BU5MuBg2IRl4dg", ""},
		{"", "close"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result string
			svc := GetService("localhost", "6379")

			// OpenDB
			err := svc.OpenDB()
			if err != nil {
				t.Error(err)
			}

			// Generate Conflict
			if tt.out == "close" {
				svc.db.Close()
			}

			err = svc.SetToken(tt.in)
			if err != nil {
				result = err.Error()
			}

			if !strings.Contains(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDeleteToken(t *testing.T) {
	for i, tt := range []struct {
		in  string
		out string
	}{
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNlc2FyQGVtYWlsLmNvbSIsImlkIjoxLCJ1c2VybmFtZSI6ImNlc2FyIiwidXVpZCI6IjcxNzFjZTU2LWIwMzYtNDEzMi1hMDljLWQyZmZiMzgzYjdjMSJ9.V_vEFyz6OAc5eOFgt589CC0OCFf72BU5MuBg2IRl4dg", ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result string
			svc := GetService("localhost", "6379")

			// OpenDB
			err := svc.OpenDB()
			if err != nil {
				t.Error(err)
			}

			err = svc.DeleteToken(tt.in)
			if err != nil {
				result = err.Error()
			}

			if !strings.Contains(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
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
		{"", false, "close"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultCheck bool
			var resultErr string
			svc := GetService("localhost", "6379")

			// OpenDB
			err := svc.OpenDB()
			if err != nil {
				t.Error(err)
			}

			// insert
			if tt.in != "" {
				err = svc.SetToken(tt.in)
				if err != nil {
					t.Error(err)
				}
			}

			// Generate Conflict
			if tt.outErr == "close" {
				svc.db.Close()
			}

			resultCheck, err = svc.CheckToken(tt.in)
			if err != nil {
				resultErr = err.Error()
			}

			if resultCheck != tt.outCheck {
				t.Errorf("want %v; got %v", tt.outCheck, resultCheck)
			}

			if !strings.Contains(resultErr, tt.outErr) {
				t.Errorf("want %v; got %v", tt.outErr, resultErr)
			}
		})
	}
}

func TestKeyFunc(t *testing.T) {
	for i, tt := range []struct {
		inSecret  []byte
		outSecret []byte
		outErr    string
	}{
		{[]byte("secret"), []byte("secret"), "Unexpected signing method"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var result []byte
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
