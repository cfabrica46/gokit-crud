package service_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cfabrica46/gokit-crud/database-app/service"
	"github.com/stretchr/testify/assert"
)

const (
	idTest       int    = 1
	usernameTest string = "username"
	passwordTest string = "password"
	emailTest    string = "email@email.com"

	noRows  string = "no rows"
	closeDB string = "close db"

	errDatabaseClosed string = "sql: database is closed"
)

func TestGetAllUsers(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                               string
		outUsername, outPassword, outEmail string
		outErr                             string
		outID                              int
	}{
		{
			name:        "NoError",
			outID:       idTest,
			outUsername: usernameTest,
			// outPassword: passwordTest,.
			outEmail: emailTest,
			outErr:   "",
		},
		{
			name:        "ErrorDBClose",
			outID:       idTest,
			outUsername: usernameTest,
			// outPassword: passwordTest,.
			outEmail: emailTest,
			outErr:   "sql: database is closed",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			if tt.outErr == "sql: database is closed" {
				db.Close()
			}

			svc := service.GetService(db)

			rows := sqlmock.NewRows(
				[]string{
					"id",
					"username",
					// "password",.
					"email",
				}).AddRow(
				tt.outID,
				// "oli",.
				tt.outUsername,
				// tt.outPassword,.
				tt.outEmail,
			)

			mock.ExpectQuery("SELECT id, username, email FROM users").WillReturnRows(rows)

			_, err = svc.GetAllUsers()
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestGetUserByID(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                            string
		inUsername, inPassword, inEmail string
		outErr                          string
		condition                       string
		inID                            int
	}{
		{
			name:       "NoError",
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "",
			condition:  "",
		},
		{
			name:       "ErrorNoRows",
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "",
			condition:  noRows,
		},
		{
			name:       "ErrorDBClose",
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "sql: database is closed",
			condition:  closeDB,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			if tt.condition == closeDB {
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

			if tt.condition == noRows {
				rows = sqlmock.NewRows([]string{"id", "username", "password", "email"})
			}

			mock.ExpectQuery(
				"^SELECT id, username, password, email FROM users",
			).WithArgs(tt.inID).WillReturnRows(rows)

			_, err = svc.GetUserByID(tt.inID)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestGetUserByUsernameAndPassword(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                            string
		inUsername, inPassword, inEmail string
		outErr                          string
		condition                       string
		inID                            int
	}{
		{
			name:       "NoError",
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "",
			condition:  "",
		},
		{
			name:       "ErrorNoRows",
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "",
			condition:  noRows,
		},
		{
			name:       "ErrorDBClose",
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "sql: database is closed",
			condition:  closeDB,
		},
		// {"", ""},
		// {"", noRows},
		// {"sql: database is closed", closeDB},.
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			if tt.condition == closeDB {
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

			if tt.condition == noRows {
				rows = sqlmock.NewRows([]string{"id", "username", "password", "email"})
			}

			mock.ExpectQuery(
				"^SELECT id, username, password, email FROM users",
			).WithArgs(tt.inUsername, tt.inPassword).WillReturnRows(rows)

			_, err = svc.GetUserByUsernameAndPassword(tt.inUsername, tt.inPassword)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestGetIDByUsername(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name       string
		inUsername string
		outErr     string
		condition  string
		inID       int
	}{
		{
			name:       "NoError",
			inID:       idTest,
			inUsername: usernameTest,
			outErr:     "",
			condition:  "",
		},
		{
			name:       "ErrorNoRows",
			inID:       idTest,
			inUsername: usernameTest,
			outErr:     "",
			condition:  noRows,
		},
		{
			name:       "ErrorDBClose",
			inID:       idTest,
			inUsername: usernameTest,
			outErr:     "sql: database is closed",
			condition:  closeDB,
		},
		/* {"", ""},
		{"", noRows},
		{"sql: database is closed", closeDB}, */
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			if tt.condition == closeDB {
				db.Close()
			}

			svc := service.GetService(db)

			rows := sqlmock.NewRows([]string{"id"}).AddRow(tt.inID)

			if tt.condition == noRows {
				rows = sqlmock.NewRows([]string{"id"})
			}

			mock.ExpectQuery("^SELECT id FROM users").WithArgs(tt.inUsername).WillReturnRows(rows)

			_, err = svc.GetIDByUsername(tt.inUsername)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestInsertUser(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                            string
		inUsername, inPassword, inEmail string
		outErr                          string
		condition                       string
	}{
		{
			name:       "NoError",
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "",
			condition:  "",
		},
		{
			name:       "ErrorNoRows",
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "",
			condition:  noRows,
		},
		{
			name:       "ErrorDBClose",
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "sql: database is closed",
			condition:  closeDB,
		},
		// {"", ""},
		// {"", "duplicate key"},
		// {"sql: database is closed", closeDB},.
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			if tt.condition == closeDB {
				db.Close()
			}

			svc := service.GetService(db)

			mock.ExpectExec(
				"^INSERT INTO users",
			).WithArgs(
				tt.inUsername,
				tt.inPassword,
				tt.inEmail,
			).WillReturnResult(
				sqlmock.NewResult(0, 1),
			)

			err = svc.InsertUser(tt.inUsername, tt.inPassword, tt.inEmail)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestDeleteUser(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name      string
		outErr    string
		condition string
		inID      int
	}{
		{
			name:      "NoError",
			inID:      idTest,
			outErr:    "",
			condition: "",
		},
		{
			name:      "ErrorDBClose",
			inID:      idTest,
			outErr:    "sql: database is closed",
			condition: closeDB,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr string

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			if tt.condition == closeDB {
				db.Close()
			}

			svc := service.GetService(db)

			mock.ExpectExec(
				"^DELETE FROM users",
			).WithArgs(
				tt.inID,
			).WillReturnResult(
				sqlmock.NewResult(0, 1),
			)

			_, err = svc.DeleteUser(tt.inID)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}
