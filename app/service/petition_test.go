package service

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	tokenapp "github.com/cfabrica46/gokit-crud/token-app/service"
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

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.out))}, nil
			})

			if tt.outErr == "Error from web server" {
				mock = newMockClient(func(req *http.Request) (*http.Response, error) {
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

func TestPetitionGetAllUsers(t *testing.T) {
	for i, tt := range []struct {
		inURL  string
		inResp []byte
		outErr string
	}{
		{"localhost:8080", []byte("{}"), ""},
		{"%%", []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{"localhost:8080", []byte(""), "unexpected end of JSON input"},
		{"localhost:8080", []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			_, err := petitionGetAllUsers(mock, tt.inURL)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
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

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			_, err := petitionGetIDByUsername(mock, tt.inURL, dbapp.GetIDByUsernameRequest{Username: tt.inUsername})
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}

func TestPetitionGetUserByID(t *testing.T) {
	for i, tt := range []struct {
		inURL  string
		inID   int
		inResp []byte
		outErr string
	}{
		{"localhost:8080", 1, []byte("{}"), ""},
		{"%%", 1, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{"localhost:8080", 1, []byte(""), "unexpected end of JSON input"},
		{"localhost:8080", 1, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			_, err := petitionGetUserByID(mock, tt.inURL, dbapp.GetUserByIDRequest{ID: tt.inID})
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

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			_, err := petitionGetUserByUsernameAndPassword(mock, tt.inURL, dbapp.GetUserByUsernameAndPasswordRequest{Username: tt.inUsername, Password: tt.inPassword})
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

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			err := petitionInsertUser(mock, tt.inURL, dbapp.InsertUserRequest{Username: tt.inUsername, Password: tt.inPassword, Email: tt.inEmail})
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}

func TestPetitionDeleteUser(t *testing.T) {
	for i, tt := range []struct {
		inURL  string
		inID   int
		inResp []byte
		outErr string
	}{
		{"localhost:8080", 1, []byte("{}"), ""},
		{"%%", 1, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{"localhost:8080", 1, []byte(""), "unexpected end of JSON input"},
		{"localhost:8080", 1, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			err := petitionDeleteUser(mock, tt.inURL, dbapp.DeleteUserRequest{ID: tt.inID})
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

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			_, err := petitionGenerateToken(mock, tt.inURL, tokenapp.GenerateTokenRequest{ID: tt.inID, Username: tt.inUsername, Email: tt.inEmail, Secret: tt.inSecret})
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}

func TestPetitionExtractToken(t *testing.T) {
	for i, tt := range []struct {
		inURL, inToken, inSecret string
		inResp                   []byte
		outErr                   string
	}{
		{"localhost:8080", "token", "secret", []byte("{}"), ""},
		{"%%", "token", "secret", []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{"localhost:8080", "token", "secret", []byte(""), "unexpected end of JSON input"},
		{"localhost:8080", "token", "secret", []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			_, _, _, err := petitionExtractToken(mock, tt.inURL, tokenapp.ExtractTokenRequest{Token: tt.inToken, Secret: tt.inSecret})
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}

func TestPetitionSetToken(t *testing.T) {
	for i, tt := range []struct {
		inURL, inToken string
		inResp         []byte
		outErr         string
	}{
		{"localhost:8080", "token", []byte("{}"), ""},
		{"%%", "token", []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{"localhost:8080", "token", []byte(""), "unexpected end of JSON input"},
		{"localhost:8080", "token", []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			err := petitionSetToken(mock, tt.inURL, tokenapp.SetTokenRequest{Token: tt.inToken})
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}

func TestPetitionDeleteToken(t *testing.T) {
	for i, tt := range []struct {
		inURL, inToken string
		inResp         []byte
		outErr         string
	}{
		{"localhost:8080", "token", []byte("{}"), ""},
		{"%%", "token", []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{"localhost:8080", "token", []byte(""), "unexpected end of JSON input"},
		{"localhost:8080", "token", []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			err := petitionDeleteToken(mock, tt.inURL, tokenapp.DeleteTokenRequest{Token: tt.inToken})
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}

func TestPetitionCheckToken(t *testing.T) {
	for i, tt := range []struct {
		inURL, inToken string
		inResp         []byte
		outErr         string
	}{
		{"localhost:8080", "token", []byte("{}"), ""},
		{"%%", "token", []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{"localhost:8080", "token", []byte(""), "unexpected end of JSON input"},
		// {"localhost:8080", "token", []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(tt.inResp))}, nil
			})

			_, err := petitionCheckToken(mock, tt.inURL, tokenapp.CheckTokenRequest{Token: tt.inToken})
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}
