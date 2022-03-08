package service

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakePetition(t *testing.T) {
	for i, tt := range []struct {
		inURL, inMethod string
		inBody          interface{}
		out             []byte
		outErr          string
	}{
		{"localhost:8080", http.MethodGet, []byte("body"), []byte("body"), ""},
		{"localhost:8080", http.MethodGet, func() {}, []byte(nil), "json: unsupported type: func()"},
		{"%%", http.MethodGet, []byte("body"), []byte(nil), `parse "%%": invalid URL escape "%%"`},
		{"localhost:8080", http.MethodGet, []byte("body"), []byte(nil), "Error from web server"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.out))}, nil
			})

			if tt.outErr == "Error from web server" {
				mock = getMockClient(func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("Error from web server")
				})
			}

			result, err := makePetition(mock, tt.inURL, tt.inMethod, tt.inBody)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
			assert.Equal(t, tt.out, result)
		})
	}
}

func TestPetitionGetIDByUsername(t *testing.T) {
	for i, tt := range []struct {
		inURL, inUsername string
		inResp            []byte
		outErr            string
	}{
		{"localhost:8080", "cesar", []byte("{}"), ""},
		{"%%", "cesar", []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{"localhost:8080", "cesar", []byte(""), "unexpected end of JSON input"},
		{"localhost:8080", "cesar", []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			_, err := petitionGetIDByUsername(mock, tt.inURL)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}

func TestPetitionGetUserByUsernameAndPassword(t *testing.T) {
	for i, tt := range []struct {
		inURL, inUsername, inPassword string
		inResp                        []byte
		outErr                        string
	}{
		{"localhost:8080", "cesar", "01234", []byte("{}"), ""},
		{"%%", "cesar", "01234", []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{"localhost:8080", "cesar", "01234", []byte(""), "unexpected end of JSON input"},
		{"localhost:8080", "cesar", "01234", []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			_, err := petitionGetUserByUsernameAndPassword(mock, tt.inURL, tt.inUsername, tt.inPassword)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}

func TestPetitionInsertUser(t *testing.T) {
	for i, tt := range []struct {
		inURL, inUsername, inPassword, inEmail string
		inResp                                 []byte
		outErr                                 string
	}{
		{"localhost:8080", "cesar", "01234", "cesar@email.com", []byte("{}"), ""},
		{"%%", "cesar", "01234", "cesar@email.com", []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{"localhost:8080", "cesar", "01234", "cesar@email.com", []byte(""), "unexpected end of JSON input"},
		{"localhost:8080", "cesar", "01234", "cesar@email.com", []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			err := petitionInsertUser(mock, tt.inURL, tt.inUsername, tt.inPassword, tt.inEmail)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}

func TestPetitionGenerateToken(t *testing.T) {
	for i, tt := range []struct {
		inURL, inUsername, inEmail, inSecret string
		inID                                 int
		inResp                               []byte
		outErr                               string
	}{
		{"localhost:8080", "cesar", "cesar@email.com", "secret", 1, []byte("{}"), ""},
		{"%%", "cesar", "cesar@email.com", "secret", 1, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{"localhost:8080", "cesar", "cesar@email.com", "secret", 1, []byte(""), "unexpected end of JSON input"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			_, err := petitionGenerateToken(mock, tt.inURL, tt.inID, tt.inUsername, tt.inEmail, tt.inSecret)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}
