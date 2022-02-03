package service

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cfabrica46/gokit-crud/database-app/models"
)

func TestOpenDB(t *testing.T) {
	for i, tt := range []struct {
		inDriver, inInfo string
		out              string
	}{
		{dbDriver, psqlInfo, ""},
		{"", psqlInfo, "unknown driver"},
		{dbDriver, "", "connection refused"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result string

			s := getServiceDB()
			err := s.OpenDB(tt.inDriver, tt.inInfo)
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
	for i, tt := range []struct {
		in  models.User
		out string
	}{
		{models.User{Username: "username", Password: "password", Email: "email"}, ""},
		{models.User{}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result string

			s := getServiceDB()

			err := s.OpenDB(dbDriver, psqlInfo)
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
				defer s.DeleteUserByUsername(tt.in.Username)
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
	for i, tt := range []struct {
		in  models.User
		out models.User
	}{
		{models.User{ID: -1}, models.User{}},
		{models.User{Username: "username", Password: "password", Email: "email"}, models.User{Username: "username", Password: "password", Email: "email"}},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			s := getServiceDB()

			err := s.OpenDB(dbDriver, psqlInfo)
			if err != nil {
				t.Error(err)
			}
			defer s.db.Close()

			if tt.in.ID != -1 {
				err := s.InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
				if err != nil {
					t.Error(err)
				}
				defer s.DeleteUserByUsername(tt.in.Username)
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
	for i, tt := range []struct {
		in  models.User
		out models.User
	}{
		{models.User{}, models.User{}},
		{models.User{Username: "username", Password: "password", Email: "email"}, models.User{Username: "username", Password: "password", Email: "email"}},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			s := getServiceDB()

			err := s.OpenDB(dbDriver, psqlInfo)
			if err != nil {
				t.Error(err)
			}
			defer s.db.Close()

			if tt.in.Username != "" {
				err := s.InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
				if err != nil {
					t.Error(err)
				}
				defer s.DeleteUserByUsername(tt.in.Username)
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
	for i, tt := range []struct {
		in  models.User
		out string
	}{
		{models.User{Username: "username", Password: "password", Email: "email"}, ""},
		{models.User{}, "database is closed"},
		{models.User{}, "duplicate key value"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result string

			s := getServiceDB()

			err := s.OpenDB(dbDriver, psqlInfo)
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
			defer s.DeleteUserByUsername(tt.in.Username)

			if !strings.Contains(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestDeleteUserbByUsername(t *testing.T) {
	for i, tt := range []struct {
		in              models.User
		outRowsAffected int
		outError        string
	}{
		{models.User{Username: "username", Password: "password", Email: "email"}, 1, ""},
		{models.User{}, 0, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var result string

			s := getServiceDB()

			err := s.OpenDB(dbDriver, psqlInfo)
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

			rowsAffected, err := s.DeleteUserByUsername(tt.in.Username)
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
