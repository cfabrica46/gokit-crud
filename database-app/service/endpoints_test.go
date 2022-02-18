package service

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestMakeGetAllUsersEndpoint(t *testing.T) {
	host, port, username, password, dbName, sslMode, driver := "localhost", "5431", "cfabrica46", "01234", "go_crud", "disable", "postgres"

	for i, tt := range []struct {
		in  GetAllUsersRequest
		out string
	}{
		{GetAllUsersRequest{}, ""},
		{GetAllUsersRequest{}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var resultErr string
			svc := GetService(host, port, username, password, dbName, sslMode, driver)

			//OpenDB
			err := svc.OpenDB()
			if err != nil {
				t.Error(err)
			}
			defer svc.db.Close()

			// generate confict closing db
			if tt.out == "database is closed" {
				err := svc.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			r, err := MakeGetAllUsersEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetAllUsersResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}

func TestMakeGetUserByIDEndpoint(t *testing.T) {
	host, port, username, password, dbName, sslMode, driver := "localhost", "5431", "cfabrica46", "01234", "go_crud", "disable", "postgres"

	for i, tt := range []struct {
		in  GetUserByIDRequest
		out string
	}{
		{GetUserByIDRequest{1}, ""},
		{GetUserByIDRequest{}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var resultErr string
			svc := GetService(host, port, username, password, dbName, sslMode, driver)

			//OpenDB
			err := svc.OpenDB()
			if err != nil {
				t.Error(err)
			}
			defer svc.db.Close()

			// generate confict closing db
			if tt.out == "database is closed" {
				err := svc.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			r, err := MakeGetUserByIDEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetUserByIDResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}

func TestMakeGetUserByUsernameAndPasswordEndpoint(t *testing.T) {
	host, port, username, password, dbName, sslMode, driver := "localhost", "5431", "cfabrica46", "01234", "go_crud", "disable", "postgres"

	for i, tt := range []struct {
		in  GetUserByUsernameAndPasswordRequest
		out string
	}{
		{GetUserByUsernameAndPasswordRequest{"cesar", "01234"}, ""},
		{GetUserByUsernameAndPasswordRequest{}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var resultErr string
			svc := GetService(host, port, username, password, dbName, sslMode, driver)

			//OpenDB
			err := svc.OpenDB()
			if err != nil {
				t.Error(err)
			}
			defer svc.db.Close()

			// generate confict closing db
			if tt.out == "database is closed" {
				err := svc.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			r, err := MakeGetUserByUsernameAndPasswordEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetUserByUsernameAndPasswordResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}

func TestGetIDByUsernameEndpoint(t *testing.T) {
	host, port, username, password, dbName, sslMode, driver := "localhost", "5431", "cfabrica46", "01234", "go_crud", "disable", "postgres"

	for i, tt := range []struct {
		in  GetIDByUsernameRequest
		out string
	}{
		{GetIDByUsernameRequest{"cesar"}, ""},
		{GetIDByUsernameRequest{}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var resultErr string
			svc := GetService(host, port, username, password, dbName, sslMode, driver)

			//OpenDB
			err := svc.OpenDB()
			if err != nil {
				t.Error(err)
			}
			defer svc.db.Close()

			// generate confict closing db
			if tt.out == "database is closed" {
				err := svc.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			r, err := MakeGetIDByUsernameEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetIDByUsernameResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}

func TestMakeInsertUserEndpoint(t *testing.T) {
	host, port, username, password, dbName, sslMode, driver := "localhost", "5431", "cfabrica46", "01234", "go_crud", "disable", "postgres"

	for i, tt := range []struct {
		in  InsertUserRequest
		out string
	}{
		{InsertUserRequest{}, ""},
		{InsertUserRequest{}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var resultErr string
			svc := GetService(host, port, username, password, dbName, sslMode, driver)

			//OpenDB
			err := svc.OpenDB()
			if err != nil {
				t.Error(err)
			}
			defer svc.db.Close()

			// generate confict closing db
			if tt.out == "database is closed" {
				err := svc.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			r, err := MakeInsertUserEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(InsertUserResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}

func TestMakeDeleteUserEndpoint(t *testing.T) {
	host, port, username, password, dbName, sslMode, driver := "localhost", "5431", "cfabrica46", "01234", "go_crud", "disable", "postgres"

	for i, tt := range []struct {
		in  DeleteUserRequest
		out string
	}{
		{DeleteUserRequest{}, ""},
		{DeleteUserRequest{}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var resultErr string
			svc := GetService(host, port, username, password, dbName, sslMode, driver)

			//OpenDB
			err := svc.OpenDB()
			if err != nil {
				t.Error(err)
			}
			defer svc.db.Close()

			// generate confict closing db
			if tt.out == "database is closed" {
				err := svc.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			r, err := MakeDeleteUserEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(DeleteUserResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
}
