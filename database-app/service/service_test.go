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

			_, err = svc.GetUserByID(idTest)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

/* func TestGetUserByUsernameAndPassword(t *testing.T) {
	for i, tt := range []struct {
		inID                            int
		inUsername, inPassword, inEmail string
		outErr    string
		condition string
	}{
		{"", ""},
		{"", "no rows"},
		{"sql: database is closed", "close db"},
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

			rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(idTest, usernameTest, passwordTest, emailTest)

			if tt.condition == "no rows" {
				rows = sqlmock.NewRows([]string{"id", "username", "password", "email"})
			}

			mock.ExpectQuery("^SELECT id, username, password, email FROM users").WithArgs(usernameTest, passwordTest).WillReturnRows(rows)

			_, err = svc.GetUserByUsernameAndPassword(usernameTest, passwordTest)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
} */

/* func TestGetIDByUsername(t *testing.T) {
	for i, tt := range []struct {
		inID                            int
		inUsername, inPassword, inEmail string
		outErr    string
		condition string
	}{
		{"", ""},
		{"", "no rows"},
		{"sql: database is closed", "close db"},
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

			rows := sqlmock.NewRows([]string{"id"}).AddRow(idTest)

			if tt.condition == "no rows" {
				rows = sqlmock.NewRows([]string{"id"})
			}

			mock.ExpectQuery("^SELECT id FROM users").WithArgs(usernameTest).WillReturnRows(rows)

			_, err = svc.GetIDByUsername(usernameTest)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
} */

/* func TestInsertUser(t *testing.T) {
	for i, tt := range []struct {
		inID                            int
		inUsername, inPassword, inEmail string
		outErr    string
		condition string
	}{
		{"", ""},
		{"", "duplicate key"},
		{"sql: database is closed", "close db"},
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

			mock.ExpectExec("^INSERT INTO users").WithArgs(usernameTest, passwordTest, emailTest).WillReturnResult(sqlmock.NewResult(0, 1))

			err = svc.InsertUser(usernameTest, passwordTest, emailTest)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
} */

/* func TestDeleteUser(t *testing.T) {
	for i, tt := range []struct {
		inID                            int
		inUsername, inPassword, inEmail string
		outErr    string
		condition string
	}{
		{"", ""},
		{"sql: database is closed", "close db"},
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

			mock.ExpectExec("^DELETE FROM users").WithArgs(idTest).WillReturnResult(sqlmock.NewResult(0, 1))

			_, err = svc.DeleteUser(idTest)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
} */
