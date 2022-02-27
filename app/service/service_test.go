package service

import (
	"fmt"
	"testing"
)

func TestSignUp(t *testing.T) {
	for i, tt := range []struct {
		outErr string
	}{
		{""},
		{"sql: database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var resultErr string

			// db, mock, err := sqlmock.New()
			// if err != nil {
			// 	t.Error(err)
			// }
			// defer db.Close()

			// if tt.outErr == "sql: database is closed" {
			// 	db.Close()
			// }

			// svc := GetService(db)

			// rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(userTest.ID, userTest.Username, userTest.Password, userTest.Email)

			// mock.ExpectQuery("SELECT id, username, email FROM users").WillReturnRows(rows)

			// _, err = svc.GetAllUsers()
			// if err != nil {
			// 	resultErr = err.Error()
			// }

			// assert.Equal(t, tt.outErr, resultErr, "they should be equal")
		})
	}

}
