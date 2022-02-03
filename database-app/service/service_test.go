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
				result = err.Error()
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
				result = err.Error()
			}
			defer s.db.Close()

			// generate confict closing db
			if tt.outError == "database is closed" {
				err := s.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

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
