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
		outID                              int
		outUsername, outPassword, outEmail string
		inRequest                          GetAllUsersRequest
		outErr                             string
	}{
		{
			outID:       idTest,
			outUsername: usernameTest,
			// inPassword: passwordTest,
			outEmail:  emailTest,
			inRequest: GetAllUsersRequest{},
			outErr:    "",
		},
		{
			outID:       idTest,
			outUsername: usernameTest,
			// inPassword: passwordTest,
			outEmail:  emailTest,
			inRequest: GetAllUsersRequest{},
			outErr:    "sql: database is closed",
		},
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
					// "password",
					"email",
				}).AddRow(
				tt.outID,
				tt.outUsername,
				// tt.inPassword,
				tt.outEmail,
			)

			mock.ExpectQuery("^SELECT id, username, email FROM users").WillReturnRows(rows)

			r, err := MakeGetAllUsersEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetAllUsersResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err)
		})
	}
}

func TestMakeGetUserByIDEndpoint(t *testing.T) {
	for i, tt := range []struct {
		inID                            int
		inUsername, inPassword, inEmail string
		inRequest                       GetUserByIDRequest
		outErr                          string
	}{
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest:  GetUserByIDRequest{ID: idTest},
			outErr:     "",
		},
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest:  GetUserByIDRequest{},
			outErr:     "sql: database is closed",
		},
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
				tt.inID,
				tt.inUsername,
				tt.inPassword,
				tt.inEmail,
			)

			mock.ExpectQuery("^SELECT id, username, password, email FROM users").
				WithArgs(tt.inID).WillReturnRows(rows)

			r, err := MakeGetUserByIDEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetUserByIDResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err)
		})
	}
}

func TestMakeGetUserByUsernameAndPasswordEndpoint(t *testing.T) {
	for i, tt := range []struct {
		inID                            int
		inUsername, inPassword, inEmail string
		inRequest                       GetUserByUsernameAndPasswordRequest
		outErr                          string
	}{
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest: GetUserByUsernameAndPasswordRequest{
				Username: usernameTest,
				Password: passwordTest,
			},
			outErr: "",
		},
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest:  GetUserByUsernameAndPasswordRequest{},
			outErr:     "sql: database is closed",
		},
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
				tt.inID,
				tt.inUsername,
				tt.inPassword,
				tt.inEmail,
			)

			mock.ExpectQuery("^SELECT id, username, password, email FROM users").
				WithArgs(tt.inUsername, tt.inPassword).WillReturnRows(rows)

			r, err := MakeGetUserByUsernameAndPasswordEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetUserByUsernameAndPasswordResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err)
		})
	}
}

func TestGetIDByUsernameEndpoint(t *testing.T) {
	for i, tt := range []struct {
		inID       int
		inUsername string
		inRequest  GetIDByUsernameRequest
		outErr     string
	}{
		{
			inID:       idTest,
			inUsername: usernameTest,
			inRequest: GetIDByUsernameRequest{
				Username: usernameTest,
			},
			outErr: "",
		},
		{
			inID:       idTest,
			inUsername: usernameTest,
			inRequest:  GetIDByUsernameRequest{},
			outErr:     "sql: database is closed",
		},
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

			rows := sqlmock.NewRows([]string{"id"}).AddRow(tt.inID)

			mock.ExpectQuery("^SELECT id FROM users").WithArgs(tt.inUsername).WillReturnRows(rows)

			r, err := MakeGetIDByUsernameEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(GetIDByUsernameResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err)
		})
	}
}

func TestMakeInsertUserEndpoint(t *testing.T) {
	for i, tt := range []struct {
		inUsername, inPassword, inEmail string
		inRequest                       InsertUserRequest
		outErr                          string
	}{
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest: InsertUserRequest{
				Username: usernameTest,
				Password: passwordTest,
				Email:    emailTest,
			},
			outErr: "",
		},
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest:  InsertUserRequest{},
			outErr:     "sql: database is closed",
		},
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

			mock.ExpectExec("^INSERT INTO users").
				WithArgs(tt.inUsername, tt.inPassword, tt.inEmail).WillReturnResult(sqlmock.NewResult(0, 1))

			r, err := MakeInsertUserEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(InsertUserResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err)
		})
	}
}

func TestMakeDeleteUserEndpoint(t *testing.T) {
	for i, tt := range []struct {
		inID      int
		inRequest DeleteUserRequest
		outErr    string
	}{
		{
			inID: idTest,
			inRequest: DeleteUserRequest{
				ID: idTest,
			},
			outErr: "",
		},
		{
			inID:      idTest,
			inRequest: DeleteUserRequest{},
			outErr:    "sql: database is closed",
		},
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

			mock.ExpectExec("^DELETE FROM users").
				WithArgs(tt.inID).WillReturnResult(sqlmock.NewResult(0, 1))

			r, err := MakeDeleteUserEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(DeleteUserResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err)
		})
	}
}
