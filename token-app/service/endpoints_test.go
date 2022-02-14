package service

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestMakeGenerateToken(t *testing.T) {
	for i, tt := range []struct {
		in  generateTokenRequest
		out string
	}{
		{generateTokenRequest{1, "cesar", "cesar@email.com", "secret"}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			svc := GetService()

			r, err := MakeGenerateToken(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(generateTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if result.Token == "" {
				t.Error("token its empty")
			}
		})
	}
}

func TestMakeExtractToken(t *testing.T) {
	for i, tt := range []struct {
		in  extractTokenRequest
		out string
	}{
		{extractTokenRequest{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNlc2FyQGVtYWlsLmNvbSIsImlkIjoxLCJ1c2VybmFtZSI6ImNlc2FyIiwidXVpZCI6IjcxNzFjZTU2LWIwMzYtNDEzMi1hMDljLWQyZmZiMzgzYjdjMSJ9.V_vEFyz6OAc5eOFgt589CC0OCFf72BU5MuBg2IRl4dg", "secret"}, ""},
		{extractTokenRequest{"", "secret"}, "token contains an invalid number of segments"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			svc := GetService()

			r, err := MakeExtractToken(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(extractTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}

func TestMakeSetToken(t *testing.T) {
	for i, tt := range []struct {
		in  setTokenRequest
		out string
	}{
		{setTokenRequest{"token"}, ""},
		{setTokenRequest{""}, "close"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			svc := GetService()

			// OpenDB
			err := svc.OpenDB()
			if err != nil {
				t.Error(err)
			}

			// Generate Conflict
			if tt.out == "close" {
				svc.db.Close()
			}

			r, err := MakeSetToken(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(setTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}

func TestMakeDeleteToken(t *testing.T) {
	for i, tt := range []struct {
		in  deleteTokenRequest
		out string
	}{
		{deleteTokenRequest{"token"}, ""},
		{deleteTokenRequest{""}, "close"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			svc := GetService()

			// OpenDB
			err := svc.OpenDB()
			if err != nil {
				t.Error(err)
			}

			// Generate Conflict
			if tt.out == "close" {
				svc.db.Close()
			}

			r, err := MakeDeleteToken(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(deleteTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}

func TestMakeCheckToken(t *testing.T) {
	for i, tt := range []struct {
		in  checkTokenRequest
		out string
	}{
		{checkTokenRequest{"token"}, ""},
		{checkTokenRequest{""}, "close"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			svc := GetService()

			// OpenDB
			err := svc.OpenDB()
			if err != nil {
				t.Error(err)
			}

			// Generate Conflict
			if tt.out == "close" {
				svc.db.Close()
			}

			r, err := MakeCheckToken(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(checkTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}
