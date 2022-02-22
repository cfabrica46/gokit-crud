package service

import (
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var u = User{
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
			log.SetFlags(log.Lshortfile)
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

			rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(u.ID, u.Username, u.Password, u.Email)

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

			rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(u.ID, u.Username, u.Password, u.Email)

			if tt.condition == "no rows" {
				rows = sqlmock.NewRows([]string{"id", "username", "password", "email"})
			}

			mock.ExpectQuery("^SELECT id, username, password, email FROM users").WithArgs(u.ID).WillReturnRows(rows)

			_, err = svc.GetUserByID(u.ID)
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

			rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(u.ID, u.Username, u.Password, u.Email)

			if tt.condition == "no rows" {
				rows = sqlmock.NewRows([]string{"id", "username", "password", "email"})
			}

			mock.ExpectQuery("^SELECT id, username, password, email FROM users").WithArgs(u.Username, u.Password).WillReturnRows(rows)

			_, err = svc.GetUserByUsernameAndPassword(u.Username, u.Password)
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

			rows := sqlmock.NewRows([]string{"id"}).AddRow(u.ID)

			if tt.condition == "no rows" {
				rows = sqlmock.NewRows([]string{"id"})
			}

			mock.ExpectQuery("^SELECT id FROM users").WithArgs(u.Username).WillReturnRows(rows)

			_, err = svc.GetIDByUsername(u.Username)
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

			/* if tt.condition == "duplicate key" {
				rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(u.ID, u.Username, u.Password, u.Email)
				mock.ExpectQuery("").WillReturnRows(rows)
			} */

			mock.ExpectPrepare("^INSERT INTO users").ExpectExec().WithArgs(u.Username, u.Password, u.Email).WillReturnResult(sqlmock.NewResult(0, 1))

			// .WithArgs(u.Username, u.Password, u.Email).WillReturnRows(rows)

			if tt.condition == "duplicate key" {
				rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(u.ID, u.Username, u.Password, u.Email)
				mock.ExpectPrepare("^INSERT INTO users").ExpectQuery().WillReturnRows(rows)
				/* err = svc.InsertUser(u.Username, u.Password, u.Email)
				if err != nil {
					resultErr = err.Error()
				} */
			}

			err = svc.InsertUser(u.Username, u.Password, u.Email)
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

			mock.ExpectPrepare("^DELETE FROM users").ExpectExec().WithArgs(u.Username, u.Password, u.Email).WillReturnResult(sqlmock.NewResult(0, 1))

			_, err = svc.DeleteUser(u.Username, u.Password, u.Email)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}
}

/* func TestInsertUser(t *testing.T) {
	host, port, username, password, dbName, sslMode, driver := "localhost", "5431", "cfabrica46", "01234", "go_crud", "disable", "postgres"

	for i, tt := range []struct {
		in  User
		out string
	}{
		{User{Username: "username", Password: "password", Email: "email"}, ""},
		{User{}, "database is closed"},
		{User{}, "duplicate key value"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result string
			s := GetService(host, port, username, password, dbName, sslMode, driver)

			err := s.OpenDB()
			if err != nil {
				t.Error(err)
			}
			defer s.db.Close()

			// generate confict closing db
			if tt.out == "database is closed" {
				err := s.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			// generate duplicate
			if tt.out == "duplicate key value" {
				err := s.InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
				if err != nil {
					result = err.Error()
				}
			}

			err = s.InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
			if err != nil {
				result = err.Error()
			}
			defer s.DeleteUser(tt.in.Username, tt.in.Password, tt.in.Email)

			if !strings.Contains(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
} */

/* func TestDeleteUser(t *testing.T) {
	host, port, username, password, dbName, sslMode, driver := "localhost", "5431", "cfabrica46", "01234", "go_crud", "disable", "postgres"

	for i, tt := range []struct {
		in              User
		outRowsAffected int
		outError        string
	}{
		{User{Username: "username", Password: "password", Email: "email"}, 1, ""},
		{User{}, 0, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result string
			s := GetService(host, port, username, password, dbName, sslMode, driver)

			err := s.OpenDB()
			if err != nil {
				t.Error(err)
			}
			defer s.db.Close()

			// generate confict closing db
			if tt.outError == "database is closed" {
				err := s.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			// insert user
			if tt.outRowsAffected == 1 {
				err := s.InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
				if err != nil {
					t.Error(err)
				}
			}

			rowsAffected, err := s.DeleteUser(tt.in.Username, tt.in.Password, tt.in.Email)
			if err != nil {
				result = err.Error()
			}

			if !strings.Contains(result, tt.outError) {
				t.Errorf("want %v; got %v", tt.outError, result)
			}

			if rowsAffected != tt.outRowsAffected {
				t.Errorf("want %v; got %v", tt.outRowsAffected, rowsAffected)
			}
		})
	}
} */
