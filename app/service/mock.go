package service

import (
	"errors"
	"net/http"
)

const (
	// userTest = dbapp.User{
	// 	ID:       1,
	// 	Username: "username",
	// 	Password: "password",
	// 	Email:    "email@email.com",
	// }.
	usernameTest = "username"
	passwordTest = "username"
	emailTest    = "username"

	urlTest   = "localhost:8080"
	portTest  = "8080"
	tokenTest = "token"
)

var (
	errWebServer        = errors.New("error from web server")
	errNotTypeIndicated = errors.New("response is not of the type indicated")
)

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
