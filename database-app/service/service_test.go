package service

import (
	"fmt"
	"strings"
	"testing"
)

func TestOpenDB(t *testing.T) {
	for i, tt := range []struct {
		inHost, inPort, inUsername, inPassword, inDBName, inSSLMode, inDriver string
		out                                                                   string
	}{
		{"localhost", "5431", "cfabrica46", "01234", "go_crud", "disable", "postgres", ""},
		{"localhost", "5431", "cfabrica46", "01234", "go_crud", "disable", "", "unknown driver"},
		{"localhost", "0", "cfabrica46", "01234", "go_crud", "", "postgres", "connection refused"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result string
			s := GetService(tt.inHost, tt.inPort, tt.inUsername, tt.inPassword, tt.inDBName, tt.inSSLMode, tt.inDriver)

			err := s.OpenDB()
			if err != nil {
				result = err.Error()
			} else {
				defer s.db.Close()
			}

			if !strings.Contains(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	host, port, username, password, dbName, sslMode, driver := "localhost", "5431", "cfabrica46", "01234", "go_crud", "disable", "postgres"

	for i, tt := range []struct {
		in  User
		out string
	}{
		{User{Username: "username", Password: "password", Email: "email"}, ""},
		{User{}, "database is closed"},
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

			// insert user
			if tt.out == "" {
				err := s.InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
				if err != nil {
					t.Error(err)
				}
				defer s.DeleteUser(tt.in.Username, tt.in.Password, tt.in.Email)
			}

			_, err = s.GetAllUsers()
			if err != nil {
				result = err.Error()
			}

			if !strings.Contains(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	host, port, username, password, dbName, sslMode, driver := "localhost", "5431", "cfabrica46", "01234", "go_crud", "disable", "postgres"

	for i, tt := range []struct {
		in  User
		out User
	}{
		{User{ID: -1}, User{}},
		{User{Username: "username", Password: "password", Email: "email"}, User{Username: "username", Password: "password", Email: "email"}},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			s := GetService(host, port, username, password, dbName, sslMode, driver)

			err := s.OpenDB()
			if err != nil {
				t.Error(err)
			}
			defer s.db.Close()

			if tt.in.ID != -1 {
				err := s.InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
				if err != nil {
					t.Error(err)
				}
				defer s.DeleteUser(tt.in.Username, tt.in.Password, tt.in.Email)
			}

			id, err := s.GetIDByUsername(tt.in.Username)
			if err != nil {
				t.Error(err)
			}

			user, err := s.GetUserByID(id)
			if err != nil {
				t.Error(err)
			}

			if user.Username != tt.out.Username {
				t.Errorf("want %v; got %v", tt.out, user)
			}

		})
	}
}

func TestGetUserByUsernameAndPassword(t *testing.T) {
	host, port, username, password, dbName, sslMode, driver := "localhost", "5431", "cfabrica46", "01234", "go_crud", "disable", "postgres"

	for i, tt := range []struct {
		in  User
		out User
	}{
		{User{}, User{}},
		{User{Username: "username", Password: "password", Email: "email"}, User{Username: "username", Password: "password", Email: "email"}},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			s := GetService(host, port, username, password, dbName, sslMode, driver)

			err := s.OpenDB()
			if err != nil {
				t.Error(err)
			}
			defer s.db.Close()

			if tt.in.Username != "" {
				err := s.InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
				if err != nil {
					t.Error(err)
				}
				defer s.DeleteUser(tt.in.Username, tt.in.Password, tt.in.Email)
			}

			user, err := s.GetUserByUsernameAndPassword(tt.in.Username, tt.in.Password)
			if err != nil {
				t.Error(err)
			}

			if user.Username != tt.out.Username {
				t.Errorf("want %v; got %v", tt.out, user)
			}

		})
	}
}

func TestInsertUser(t *testing.T) {
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
}

func TestDeleteUser(t *testing.T) {
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
}
