package service

import (
	"errors"
	"net/http"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
)

var userTest = dbapp.User{
	ID:       1,
	Username: "username",
	Password: "password",
	Email:    "email@email.com",
}

var errWebService = errors.New("error from web server")

type myDoFunc func(req *http.Request) (*http.Response, error)

type mockClient struct {
	doFunc myDoFunc
}

func newMockClient(d myDoFunc) *mockClient {
	return &mockClient{d}
}

func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	return m.doFunc(req)
}
