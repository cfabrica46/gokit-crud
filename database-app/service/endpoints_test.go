package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMakeGetAllUsersEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     GetAllUsersRequest
		outErr string
	}{
		{GetAllUsersRequest{}, ""},
		{GetAllUsersRequest{}, "sql: database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			// generate confict closing db
			if tt.outErr == "sql: database is closed" {
				db.Close()
			}

			svc := GetService(db)

			rows := sqlmock.NewRows(
				[]string{
					"id",
					"username",
					"password",
					"email",
				}).AddRow(
				idTest,
				usernameTest,
				passwordTest,
				emailTest,
			)

			mock.ExpectQuery("^SELECT id, username, email FROM users").WillReturnRows(rows)

			r, err := MakeGetAllUsersEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetAllUsersResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestMakeGetUserByIDEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     GetUserByIDRequest
		outErr string
	}{
		{GetUserByIDRequest{1}, ""},
		{GetUserByIDRequest{}, "sql: database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			// generate confict closing db
			if tt.outErr == "sql: database is closed" {
				db.Close()
			}

			svc := GetService(db)

			rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(idTest, usernameTest, passwordTest, emailTest)

			mock.ExpectQuery("^SELECT id, username, password, email FROM users").WithArgs(idTest).WillReturnRows(rows)

			r, err := MakeGetUserByIDEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetUserByIDResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestMakeGetUserByUsernameAndPasswordEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     GetUserByUsernameAndPasswordRequest
		outErr string
	}{
		{GetUserByUsernameAndPasswordRequest{"cesar", "01234"}, ""},
		{GetUserByUsernameAndPasswordRequest{}, "sql: database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			// generate confict closing db
			if tt.outErr == "sql: database is closed" {
				db.Close()
			}

			svc := GetService(db)

			rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(idTest, usernameTest, passwordTest, emailTest)

			mock.ExpectQuery("^SELECT id, username, password, email FROM users").WithArgs(usernameTest, passwordTest).WillReturnRows(rows)

			r, err := MakeGetUserByUsernameAndPasswordEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetUserByUsernameAndPasswordResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestGetIDByUsernameEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     GetIDByUsernameRequest
		outErr string
	}{
		{GetIDByUsernameRequest{"cesar"}, ""},
		{GetIDByUsernameRequest{}, "sql: database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			// generate confict closing db
			if tt.outErr == "sql: database is closed" {
				db.Close()
			}

			svc := GetService(db)

			rows := sqlmock.NewRows([]string{"id"}).AddRow(idTest)

			mock.ExpectQuery("^SELECT id FROM users").WithArgs(usernameTest).WillReturnRows(rows)

			r, err := MakeGetIDByUsernameEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetIDByUsernameResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestMakeInsertUserEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     InsertUserRequest
		outErr string
	}{
		{InsertUserRequest{"cesar", "01234", "cesar@email.com"}, ""},
		{InsertUserRequest{}, "sql: database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			// generate confict closing db
			if tt.outErr == "sql: database is closed" {
				db.Close()
			}

			svc := GetService(db)

			mock.ExpectExec("^INSERT INTO users").WithArgs(usernameTest, passwordTest, emailTest).WillReturnResult(sqlmock.NewResult(0, 1))

			r, err := MakeInsertUserEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(InsertUserResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}

func TestMakeDeleteUserEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in     DeleteUserRequest
		outErr string
	}{
		{DeleteUserRequest{1}, ""},
		{DeleteUserRequest{}, "sql: database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			// generate confict closing db
			if tt.outErr == "sql: database is closed" {
				db.Close()
			}

			svc := GetService(db)

			mock.ExpectExec("^DELETE FROM users").WithArgs(idTest).WillReturnResult(sqlmock.NewResult(0, 1))

			r, err := MakeDeleteUserEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(DeleteUserResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err, "they should be equal")
		})
	}
}
