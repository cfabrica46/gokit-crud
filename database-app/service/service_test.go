package service

import (
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

const (
	idTest       int    = 1
	usernameTest string = "username"
	passwordTest string = "password"
	emailTest    string = "email@email.com"
)

func TestGetAllUsers(t *testing.T) {
	for i, tt := range []struct {
		inID                            int
		inUsername, inPassword, inEmail string
		outErr                          string
	}{
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "",
		},
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "sql: database is closed",
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

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
	for i, tt := range []struct {
		inID                            int
		inUsername, inPassword, inEmail string
		outErr                          string
		condition                       string
	}{
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "",
			condition:  "",
		},
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "",
			condition:  "no rows",
		},
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "sql: database is closed",
			condition:  "close db",
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			log.SetFlags(log.Lshortfile)
			var resultErr string

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			if tt.condition == "close db" {
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

			if tt.condition == "no rows" {
				rows = sqlmock.NewRows([]string{"id", "username", "password", "email"})
			}

			mock.ExpectQuery("^SELECT id, username, password, email FROM users").WithArgs(tt.inID).WillReturnRows(rows)

			_, err = svc.GetUserByID(tt.inID)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestGetUserByUsernameAndPassword(t *testing.T) {
	for i, tt := range []struct {
		inID                            int
		inUsername, inPassword, inEmail string
		outErr                          string
		condition                       string
	}{
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "",
			condition:  "",
		},
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "",
			condition:  "no rows",
		},
		{
			inID:       idTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "sql: database is closed",
			condition:  "close db",
		},
		// {"", ""},
		// {"", "no rows"},
		// {"sql: database is closed", "close db"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			log.SetFlags(log.Lshortfile)
			var resultErr string

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			if tt.condition == "close db" {
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

			if tt.condition == "no rows" {
				rows = sqlmock.NewRows([]string{"id", "username", "password", "email"})
			}

			mock.ExpectQuery("^SELECT id, username, password, email FROM users").WithArgs(tt.inUsername, tt.inPassword).WillReturnRows(rows)

			_, err = svc.GetUserByUsernameAndPassword(tt.inUsername, tt.inPassword)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestGetIDByUsername(t *testing.T) {
	for i, tt := range []struct {
		inID       int
		inUsername string
		outErr     string
		condition  string
	}{
		{
			inID:       idTest,
			inUsername: usernameTest,
			outErr:     "",
			condition:  "",
		},
		{
			inID:       idTest,
			inUsername: usernameTest,
			outErr:     "",
			condition:  "no rows",
		},
		{
			inID:       idTest,
			inUsername: usernameTest,
			outErr:     "sql: database is closed",
			condition:  "close db",
		},
		/* {"", ""},
		{"", "no rows"},
		{"sql: database is closed", "close db"}, */
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			log.SetFlags(log.Lshortfile)
			var resultErr string

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			if tt.condition == "close db" {
				db.Close()
			}

			svc := GetService(db)

			rows := sqlmock.NewRows([]string{"id"}).AddRow(tt.inID)

			if tt.condition == "no rows" {
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
	for i, tt := range []struct {
		inUsername, inPassword, inEmail string
		outErr                          string
		condition                       string
	}{
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "",
			condition:  "",
		},
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "",
			condition:  "no rows",
		},
		{
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			outErr:     "sql: database is closed",
			condition:  "close db",
		},
		// {"", ""},
		// {"", "duplicate key"},
		// {"sql: database is closed", "close db"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			log.SetFlags(log.Lshortfile)
			var resultErr string

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			if tt.condition == "close db" {
				db.Close()
			}

			svc := GetService(db)

			mock.ExpectExec("^INSERT INTO users").WithArgs(tt.inUsername, tt.inPassword, tt.inEmail).WillReturnResult(sqlmock.NewResult(0, 1))

			err = svc.InsertUser(tt.inUsername, tt.inPassword, tt.inEmail)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestDeleteUser(t *testing.T) {
	for i, tt := range []struct {
		inID int
		// inUsername, inPassword, inEmail string
		outErr    string
		condition string
	}{
		{
			inID:      idTest,
			outErr:    "",
			condition: "",
		},
		// {
		// 	inUsername: usernameTest,
		// 	inPassword: passwordTest,
		// 	inEmail:    emailTest,
		// 	outErr:     "",
		// 	condition:  "no rows",
		// },
		{
			inID:      idTest,
			outErr:    "sql: database is closed",
			condition: "close db",
		},
		// {"", ""},
		// {"sql: database is closed", "close db"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			log.SetFlags(log.Lshortfile)
			var resultErr string

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			if tt.condition == "close db" {
				db.Close()
			}

			svc := GetService(db)

			mock.ExpectExec("^DELETE FROM users").WithArgs(tt.inID).WillReturnResult(sqlmock.NewResult(0, 1))

			_, err = svc.DeleteUser(tt.inID)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}
