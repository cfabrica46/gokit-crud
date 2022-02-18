package service

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestMakeGenerateTokenEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in  GenerateTokenRequest
		out string
	}{
		{GenerateTokenRequest{1, "cesar", "cesar@email.com", "secret"}, ""},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			svc := GetService("localhost", "6379")

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
		})
	}
}

func TestMakeExtractTokenEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in  ExtractTokenRequest
		out string
	}{
		{ExtractTokenRequest{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNlc2FyQGVtYWlsLmNvbSIsImlkIjoxLCJ1c2VybmFtZSI6ImNlc2FyIiwidXVpZCI6IjcxNzFjZTU2LWIwMzYtNDEzMi1hMDljLWQyZmZiMzgzYjdjMSJ9.V_vEFyz6OAc5eOFgt589CC0OCFf72BU5MuBg2IRl4dg", "secret"}, ""},
		{ExtractTokenRequest{"", "secret"}, "token contains an invalid number of segments"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			svc := GetService("localhost", "6379")

			r, err := MakeExtractTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(ExtractTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}

func TestMakeSetTokenEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in  SetTokenRequest
		out string
	}{
		{SetTokenRequest{"token"}, ""},
		{SetTokenRequest{""}, "close"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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

			r, err := MakeSetTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(SetTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}

func TestMakeDeleteTokenEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in  DeleteTokenRequest
		out string
	}{
		{DeleteTokenRequest{"token"}, ""},
		{DeleteTokenRequest{""}, "close"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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

			r, err := MakeDeleteTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(DeleteTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}

func TestMakeCheckTokenEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in  CheckTokenRequest
		out string
	}{
		{CheckTokenRequest{"token"}, ""},
		{CheckTokenRequest{""}, "close"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
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

			r, err := MakeCheckTokenEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(CheckTokenResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}
