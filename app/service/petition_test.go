package service_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/cfabrica46/gokit-crud/app/service"
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
		{
			urlTest,
			http.MethodGet,
			[]byte("body"),
			[]byte("body"),
			"",
		},
		{
			urlTest,
			http.MethodGet,
			func() {},
			[]byte(nil),
			"json: unsupported type: func()",
		},
		{
			"%%",
			http.MethodGet,
			[]byte("body"),
			[]byte(nil),
			`parse "%%": invalid URL escape "%%"`,
		},
		{
			urlTest,
			http.MethodGet,
			[]byte("body"),
			[]byte(nil),
			errWebServer.Error(),
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, i), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.out)),
				}, nil
			})

			if tt.outErr == errWebServer.Error() {
				mock = service.NewMockClient(func(req *http.Request) (*http.Response, error) {
					return nil, errWebServer
				})
			}

			result, err := service.MakePetition(mock, tt.inURL, tt.inMethod, tt.inBody)
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
		{urlTest, []byte("{}"), ""},
		{"%%", []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, i), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			_, err := service.PetitionGetAllUsers(mock, tt.inURL)
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
		{urlTest, usernameTest, []byte("{}"), ""},
		{"%%", usernameTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, usernameTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, usernameTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, i), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			_, err := service.PetitionGetIDByUsername(
				mock,
				tt.inURL,
				dbapp.GetIDByUsernameRequest{
					Username: tt.inUsername,
				},
			)
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
		{urlTest, idTest, []byte("{}"), ""},
		{"%%", idTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, idTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, idTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, i), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			_, err := service.PetitionGetUserByID(mock, tt.inURL, dbapp.GetUserByIDRequest{ID: tt.inID})
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
		{
			urlTest,
			usernameTest,
			passwordTest,
			[]byte("{}"),
			"",
		},
		{
			"%%",
			usernameTest,
			passwordTest,
			[]byte("{}"),
			`parse "%%": invalid URL escape "%%"`,
		},
		{
			urlTest,
			usernameTest,
			passwordTest,
			[]byte(""),
			"unexpected end of JSON input",
		},
		{
			urlTest,
			usernameTest,
			passwordTest,
			[]byte(`{"err":"error"}`),
			"error",
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, i), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			_, err := service.PetitionGetUserByUsernameAndPassword(
				mock,
				tt.inURL,
				dbapp.GetUserByUsernameAndPasswordRequest{
					Username: tt.inUsername,
					Password: tt.inPassword,
				},
			)
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
		{
			urlTest,
			usernameTest,
			passwordTest,
			emailTest,
			[]byte("{}"),
			"",
		},
		{
			"%%",
			usernameTest,
			passwordTest,
			emailTest,
			[]byte("{}"),
			`parse "%%": invalid URL escape "%%"`,
		},
		{
			urlTest,
			usernameTest,
			passwordTest,
			emailTest,
			[]byte(""),
			"unexpected end of JSON input",
		},
		{
			urlTest,
			usernameTest,
			passwordTest,
			emailTest,
			[]byte(`{"err":"error"}`),
			"error",
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, i), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			err := service.PetitionInsertUser(
				mock,
				tt.inURL,
				dbapp.InsertUserRequest{
					Username: tt.inUsername,
					Password: tt.inPassword,
					Email:    tt.inEmail,
				},
			)
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
		{urlTest, idTest, []byte("{}"), ""},
		{"%%", idTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, idTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, idTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, i), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			err := service.PetitionDeleteUser(mock, tt.inURL, dbapp.DeleteUserRequest{ID: tt.inID})
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
		{
			urlTest,
			usernameTest,
			emailTest,
			secretTest,
			idTest,
			[]byte("{}"),
			"",
		},
		{
			"%%",
			usernameTest,
			emailTest,
			secretTest,
			idTest,
			[]byte("{}"),
			`parse "%%": invalid URL escape "%%"`,
		},
		{
			urlTest,
			usernameTest,
			emailTest,
			secretTest,
			idTest,
			[]byte(""),
			"unexpected end of JSON input",
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, i), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			_, err := service.PetitionGenerateToken(
				mock,
				tt.inURL,
				tokenapp.GenerateTokenRequest{
					ID:       tt.inID,
					Username: tt.inUsername,
					Email:    tt.inEmail,
					Secret:   tt.inSecret,
				},
			)
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
		{urlTest, tokenTest, secretTest, []byte("{}"), ""},
		{"%%", tokenTest, secretTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, tokenTest, secretTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, tokenTest, secretTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, i), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			_, _, _, err := service.PetitionExtractToken(
				mock,
				tt.inURL,
				tokenapp.ExtractTokenRequest{
					Token:  tt.inToken,
					Secret: tt.inSecret,
				},
			)
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
		{urlTest, tokenTest, []byte("{}"), ""},
		{"%%", tokenTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, tokenTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, tokenTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, i), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			err := service.PetitionSetToken(mock, tt.inURL, tokenapp.SetTokenRequest{Token: tt.inToken})
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
		{urlTest, tokenTest, []byte("{}"), ""},
		{"%%", tokenTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, tokenTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, tokenTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, i), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			err := service.PetitionDeleteToken(
				mock,
				tt.inURL,
				tokenapp.DeleteTokenRequest{
					Token: tt.inToken,
				},
			)
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
		{urlTest, tokenTest, []byte("{}"), ""},
		{"%%", tokenTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, tokenTest, []byte(""), "unexpected end of JSON input"},
		// {urlTest, tokenTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, i), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			_, err := service.PetitionCheckToken(
				mock,
				tt.inURL,
				tokenapp.CheckTokenRequest{
					Token: tt.inToken,
				},
			)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}
