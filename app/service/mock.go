package service

import (
	"net/http"
)

type myDoFunc func(req *http.Request) (*http.Response, error)

type MockClient struct {
	doFunc myDoFunc
}

func NewMockClient(d myDoFunc) *MockClient {
	return &MockClient{d}
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.doFunc(req)
}
