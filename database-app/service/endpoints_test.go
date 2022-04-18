package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cfabrica46/gokit-crud/database-app/service"
	"github.com/stretchr/testify/assert"
)

func TestMakeGetAllUsersEndpoint(t *testing.T) {
	for indx, tt := range []struct {
		outID                              int
		outUsername, outPassword, outEmail string
		inRequest                          service.GetAllUsersRequest
		outErr                             string
	}{
		{
			outID:       idTest,
			outUsername: usernameTest,
			// inPassword: passwordTest,
			outEmail:  emailTest,
			inRequest: service.GetAllUsersRequest{},
			outErr:    "",
		},
		{
			outID:       idTest,
			outUsername: usernameTest,
			// inPassword: passwordTest,
			outEmail:  emailTest,
			inRequest: service.GetAllUsersRequest{},
			outErr:    errDatabaseClosed,
		},
	} {
		t.Run(fmt.Sprintf("%v", indx), func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			// generate confict closing db
			if tt.outErr == errDatabaseClosed {
				db.Close()
			}

			svc := service.GetService(db)

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

			r, err := service.MakeGetAllUsersEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.GetAllUsersResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err)
		})
	}
}

func TestMakeGetUserByIDEndpoint(t *testing.T) {
	for indx, tt := range []struct {
		inID                            int
		inUsername, inPassword, inEmail string
		inRequest                       service.GetUserByIDRequest
		outErr                          string
	}{
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest:  service.GetUserByIDRequest{ID: idTest},
			outErr:     "",
		},
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest:  service.GetUserByIDRequest{},
			outErr:     errDatabaseClosed,
		},
	} {
		t.Run(fmt.Sprintf("%v", indx), func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			// generate confict closing db
			if tt.outErr == errDatabaseClosed {
				db.Close()
			}

			svc := service.GetService(db)

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

			r, err := service.MakeGetUserByIDEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.GetUserByIDResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err)
		})
	}
}

func TestMakeGetUserByUsernameAndPasswordEndpoint(t *testing.T) {
	for indx, tt := range []struct {
		inID                            int
		inUsername, inPassword, inEmail string
		inRequest                       service.GetUserByUsernameAndPasswordRequest
		outErr                          string
	}{
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest: service.GetUserByUsernameAndPasswordRequest{
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
			inRequest:  service.GetUserByUsernameAndPasswordRequest{},
			outErr:     errDatabaseClosed,
		},
	} {
		t.Run(fmt.Sprintf("%v", indx), func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			// generate confict closing db
			if tt.outErr == errDatabaseClosed {
				db.Close()
			}

			svc := service.GetService(db)

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

			r, err := service.MakeGetUserByUsernameAndPasswordEndpoint(svc)(
				context.TODO(),
				tt.inRequest,
			)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.GetUserByUsernameAndPasswordResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err)
		})
	}
}

func TestGetIDByUsernameEndpoint(t *testing.T) {
	for indx, tt := range []struct {
		inID       int
		inUsername string
		inRequest  service.GetIDByUsernameRequest
		outErr     string
	}{
		{
			inID:       idTest,
			inUsername: usernameTest,
			inRequest: service.GetIDByUsernameRequest{
				Username: usernameTest,
			},
			outErr: "",
		},
		{
			inID:       idTest,
			inUsername: usernameTest,
			inRequest:  service.GetIDByUsernameRequest{},
			outErr:     errDatabaseClosed,
		},
	} {
		t.Run(fmt.Sprintf("%v", indx), func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			// generate confict closing db
			if tt.outErr == errDatabaseClosed {
				db.Close()
			}

			svc := service.GetService(db)

			rows := sqlmock.NewRows([]string{"id"}).AddRow(tt.inID)

			mock.ExpectQuery("^SELECT id FROM users").WithArgs(tt.inUsername).WillReturnRows(rows)

			r, err := service.MakeGetIDByUsernameEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.GetIDByUsernameResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err)
		})
	}
}

func TestMakeInsertUserEndpoint(t *testing.T) {
	for indx, tt := range []struct {
		inUsername, inPassword, inEmail string
		inRequest                       service.InsertUserRequest
		outErr                          string
	}{
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest: service.InsertUserRequest{
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
			inRequest:  service.InsertUserRequest{},
			outErr:     errDatabaseClosed,
		},
	} {
		t.Run(fmt.Sprintf("%v", indx), func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			// generate confict closing db
			if tt.outErr == errDatabaseClosed {
				db.Close()
			}

			svc := service.GetService(db)

			mock.ExpectExec("^INSERT INTO users").
				WithArgs(
					tt.inUsername,
					tt.inPassword,
					tt.inEmail,
				).WillReturnResult(sqlmock.NewResult(0, 1))

			r, err := service.MakeInsertUserEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.InsertUserResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err)
		})
	}
}

func TestMakeDeleteUserEndpoint(t *testing.T) {
	for indx, tt := range []struct {
		inID      int
		inRequest service.DeleteUserRequest
		outErr    string
	}{
		{
			inID: idTest,
			inRequest: service.DeleteUserRequest{
				ID: idTest,
			},
			outErr: "",
		},
		{
			inID:      idTest,
			inRequest: service.DeleteUserRequest{},
			outErr:    errDatabaseClosed,
		},
	} {
		t.Run(fmt.Sprintf("%v", indx), func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			// generate confict closing db
			if tt.outErr == errDatabaseClosed {
				db.Close()
			}

			svc := service.GetService(db)

			mock.ExpectExec("^DELETE FROM users").
				WithArgs(tt.inID).WillReturnResult(sqlmock.NewResult(0, 1))

			r, err := service.MakeDeleteUserEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.DeleteUserResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			assert.Equal(t, tt.outErr, result.Err)
		})
	}
}
