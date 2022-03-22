package service

import (
	"net/http"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
)

var userTest = dbapp.User{
	ID:       1,
	Username: "cesar",
	Password: "01234",
	Email:    "cesar@email.com",
}

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
