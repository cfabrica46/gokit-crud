package service_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cfabrica46/gokit-crud/database-app/service"
	"github.com/stretchr/testify/assert"
)

func TestMakeGetAllUsersEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                               string
		outUsername, outPassword, outEmail string
		outErr                             string
		inRequest                          service.GetAllUsersRequest
		outID                              int
	}{
		{
			name:        nameNoError,
			outID:       idTest,
			outUsername: usernameTest,
			outEmail:    emailTest,
			inRequest:   service.GetAllUsersRequest{},
			outErr:      "",
		},
		{
			name:        nameErrorDBClosed,
			outID:       idTest,
			outUsername: usernameTest,
			outEmail:    emailTest,
			inRequest:   service.GetAllUsersRequest{},
			outErr:      errDatabaseClosed,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				assert.Error(t, err)
			}
			defer db.Close()

			if tt.name == nameErrorDBClosed {
				db.Close()
			}

			svc := service.GetService(db)

			rows := sqlmock.NewRows(
				[]string{
					"id",
					"username",
					"email",
				}).AddRow(
				tt.outID,
				tt.outUsername,
				tt.outEmail,
			)

			mock.ExpectQuery("^SELECT id, username, email FROM users").WillReturnRows(rows)

			r, err := service.MakeGetAllUsersEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				assert.Error(t, err)
			}

			result, ok := r.(service.GetAllUsersResponse)
			if !ok {
				assert.Fail(t, "response is not of the type indicated")
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, result.Err, tt.outErr)
			}
		})
	}
}

func TestMakeGetUserByIDEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                            string
		inUsername, inPassword, inEmail string
		outErr                          string
		inRequest                       service.GetUserByIDRequest
		inID                            int
	}{
		{
			name:       nameNoError,
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest:  service.GetUserByIDRequest{ID: idTest},
			outErr:     "",
		},
		{
			name:       nameErrorDBClosed,
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest:  service.GetUserByIDRequest{},
			outErr:     errDatabaseClosed,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				assert.Error(t, err)
			}
			defer db.Close()

			if tt.name == nameErrorDBClosed {
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
				assert.Error(t, err)
			}

			result, ok := r.(service.GetUserByIDResponse)
			if !ok {
				assert.Fail(t, "response is not of the type indicated")
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, result.Err, tt.outErr)
			}
		})
	}
}

func TestMakeGetUserByUsernameAndPasswordEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                            string
		inUsername, inPassword, inEmail string
		outErr                          string
		inRequest                       service.GetUserByUsernameAndPasswordRequest
		inID                            int
	}{
		{
			name:       nameNoError,
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
			name:       nameErrorDBClosed,
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest:  service.GetUserByUsernameAndPasswordRequest{},
			outErr:     errDatabaseClosed,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				assert.Error(t, err)
			}
			defer db.Close()

			if tt.name == nameErrorDBClosed {
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
				assert.Error(t, err)
			}

			result, ok := r.(service.GetUserByUsernameAndPasswordResponse)
			if !ok {
				assert.Fail(t, "response is not of the type indicated")
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, result.Err, tt.outErr)
			}
		})
	}
}

func TestGetIDByUsernameEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name       string
		inUsername string
		outErr     string
		inRequest  service.GetIDByUsernameRequest
		inID       int
	}{
		{
			name:       nameNoError,
			inID:       idTest,
			inUsername: usernameTest,
			inRequest: service.GetIDByUsernameRequest{
				Username: usernameTest,
			},
			outErr: "",
		},
		{
			name:       nameErrorDBClosed,
			inID:       idTest,
			inUsername: usernameTest,
			inRequest:  service.GetIDByUsernameRequest{},
			outErr:     errDatabaseClosed,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				assert.Error(t, err)
			}
			defer db.Close()

			if tt.name == nameErrorDBClosed {
				db.Close()
			}

			svc := service.GetService(db)

			rows := sqlmock.NewRows([]string{"id"}).AddRow(tt.inID)

			mock.ExpectQuery("^SELECT id FROM users").WithArgs(tt.inUsername).WillReturnRows(rows)

			r, err := service.MakeGetIDByUsernameEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				assert.Error(t, err)
			}

			result, ok := r.(service.GetIDByUsernameResponse)
			if !ok {
				assert.Fail(t, "response is not of the type indicated")
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, result.Err, tt.outErr)
			}
		})
	}
}

func TestMakeInsertUserEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                            string
		inUsername, inPassword, inEmail string
		inRequest                       service.InsertUserRequest
		outErr                          string
	}{
		{
			name:       nameNoError,
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
			name:       nameErrorDBClosed,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inRequest:  service.InsertUserRequest{},
			outErr:     errDatabaseClosed,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, mock, err := sqlmock.New()
			if err != nil {
				assert.Error(t, err)
			}
			defer db.Close()

			if tt.name == nameErrorDBClosed {
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
				assert.Error(t, err)
			}

			result, ok := r.(service.InsertUserResponse)
			if !ok {
				assert.Fail(t, "response is not of the type indicated")
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, result.Err, tt.outErr)
			}
		})
	}
}

func TestMakeDeleteUserEndpoint(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name      string
		outErr    string
		inRequest service.DeleteUserRequest
		inID      int
	}{
		{
			name: nameNoError,
			inID: idTest,
			inRequest: service.DeleteUserRequest{
				ID: idTest,
			},
			outErr: "",
		},
		{
			name:      nameErrorDBClosed,
			inID:      idTest,
			inRequest: service.DeleteUserRequest{},
			outErr:    errDatabaseClosed,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, mock, err := sqlmock.New()
			if err != nil {
				assert.Error(t, err)
			}
			defer db.Close()

			if tt.name == nameErrorDBClosed {
				db.Close()
			}

			svc := service.GetService(db)

			mock.ExpectExec("^DELETE FROM users").
				WithArgs(tt.inID).WillReturnResult(sqlmock.NewResult(0, 1))

			r, err := service.MakeDeleteUserEndpoint(svc)(context.TODO(), tt.inRequest)
			if err != nil {
				assert.Error(t, err)
			}

			result, ok := r.(service.DeleteUserResponse)
			if !ok {
				assert.Fail(t, "response is not of the type indicated")
			}

			if tt.name == nameNoError {
				assert.Empty(t, result.Err)
			} else {
				assert.Contains(t, result.Err, tt.outErr)
			}
		})
	}
}
