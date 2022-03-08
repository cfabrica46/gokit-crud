package service

import (
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var userTest = User{
	ID:       1,
	Username: "cesar",
	Password: "01234",
	Email:    "cesar@email.com",
}

func TestGetAllUsers(t *testing.T) {
	for i, tt := range []struct {
		outErr string
	}{
		{""},
		{"sql: database is closed"},
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

			rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(userTest.ID, userTest.Username, userTest.Password, userTest.Email)

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

			rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(userTest.ID, userTest.Username, userTest.Password, userTest.Email)

			if tt.condition == "no rows" {
				rows = sqlmock.NewRows([]string{"id", "username", "password", "email"})
			}

			mock.ExpectQuery("^SELECT id, username, password, email FROM users").WithArgs(userTest.ID).WillReturnRows(rows)

			_, err = svc.GetUserByID(userTest.ID)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestGetUserByUsernameAndPassword(t *testing.T) {
	for i, tt := range []struct {
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

			rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(userTest.ID, userTest.Username, userTest.Password, userTest.Email)

			if tt.condition == "no rows" {
				rows = sqlmock.NewRows([]string{"id", "username", "password", "email"})
			}

			mock.ExpectQuery("^SELECT id, username, password, email FROM users").WithArgs(userTest.Username, userTest.Password).WillReturnRows(rows)

			_, err = svc.GetUserByUsernameAndPassword(userTest.Username, userTest.Password)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestGetIDByUsername(t *testing.T) {
	for i, tt := range []struct {
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

			rows := sqlmock.NewRows([]string{"id"}).AddRow(userTest.ID)

			if tt.condition == "no rows" {
				rows = sqlmock.NewRows([]string{"id"})
			}

			mock.ExpectQuery("^SELECT id FROM users").WithArgs(userTest.Username).WillReturnRows(rows)

			_, err = svc.GetIDByUsername(userTest.Username)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestInsertUser(t *testing.T) {
	for i, tt := range []struct {
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

			mock.ExpectPrepare("^INSERT INTO users").ExpectExec().WithArgs(userTest.Username, userTest.Password, userTest.Email).WillReturnResult(sqlmock.NewResult(0, 1))

			err = svc.InsertUser(userTest.Username, userTest.Password, userTest.Email)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

func TestDeleteUser(t *testing.T) {
	for i, tt := range []struct {
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

			mock.ExpectPrepare("^DELETE FROM users").ExpectExec().WithArgs(userTest.ID).WillReturnResult(sqlmock.NewResult(0, 1))

			_, err = svc.DeleteUser(userTest.ID)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}
