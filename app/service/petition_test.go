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
	for index, table := range []struct {
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
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.out)),
				}, nil
			})

			if table.outErr == errWebServer.Error() {
				mock = service.NewMockClient(func(req *http.Request) (*http.Response, error) {
					return nil, errWebServer
				})
			}

			result, err := service.MakePetition(mock, table.inURL, table.inMethod, table.inBody)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
			assert.Equal(t, table.out, result)
		})
	}
}

func TestPetitionGetAllUsers(t *testing.T) {
	for index, table := range []struct {
		inURL  string
		inResp []byte
		outErr string
	}{
		{urlTest, []byte("{}"), ""},
		{"%%", []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			_, err := service.PetitionGetAllUsers(mock, table.inURL)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
		})
	}
}

func TestPetitionGetIDByUsername(t *testing.T) {
	for index, table := range []struct {
		inURL, inUsername string
		inResp            []byte
		outErr            string
	}{
		{urlTest, usernameTest, []byte("{}"), ""},
		{"%%", usernameTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, usernameTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, usernameTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			_, err := service.PetitionGetIDByUsername(
				mock,
				table.inURL,
				dbapp.GetIDByUsernameRequest{
					Username: table.inUsername,
				},
			)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
		})
	}
}

func TestPetitionGetUserByID(t *testing.T) {
	for index, table := range []struct {
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
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			_, err := service.PetitionGetUserByID(
				mock,
				table.inURL,
				dbapp.GetUserByIDRequest{
					ID: table.inID,
				},
			)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
		})
	}
}

func TestPetitionGetUserByUsernameAndPassword(t *testing.T) {
	for index, table := range []struct {
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
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			_, err := service.PetitionGetUserByUsernameAndPassword(
				mock,
				table.inURL,
				dbapp.GetUserByUsernameAndPasswordRequest{
					Username: table.inUsername,
					Password: table.inPassword,
				},
			)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
		})
	}
}

func TestPetitionInsertUser(t *testing.T) {
	for index, table := range []struct {
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
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			err := service.PetitionInsertUser(
				mock,
				table.inURL,
				dbapp.InsertUserRequest{
					Username: table.inUsername,
					Password: table.inPassword,
					Email:    table.inEmail,
				},
			)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
		})
	}
}

func TestPetitionDeleteUser(t *testing.T) {
	for index, table := range []struct {
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
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			err := service.PetitionDeleteUser(
				mock,
				table.inURL,
				dbapp.DeleteUserRequest{
					ID: table.inID,
				},
			)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
		})
	}
}

func TestPetitionGenerateToken(t *testing.T) {
	for index, table := range []struct {
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
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			_, err := service.PetitionGenerateToken(
				mock,
				table.inURL,
				tokenapp.GenerateTokenRequest{
					ID:       table.inID,
					Username: table.inUsername,
					Email:    table.inEmail,
					Secret:   table.inSecret,
				},
			)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
		})
	}
}

func TestPetitionExtractToken(t *testing.T) {
	for index, table := range []struct {
		inURL, inToken, inSecret string
		inResp                   []byte
		outErr                   string
	}{
		{urlTest, tokenTest, secretTest, []byte("{}"), ""},
		{"%%", tokenTest, secretTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, tokenTest, secretTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, tokenTest, secretTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			_, _, _, err := service.PetitionExtractToken(
				mock,
				table.inURL,
				tokenapp.ExtractTokenRequest{
					Token:  table.inToken,
					Secret: table.inSecret,
				},
			)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
		})
	}
}

func TestPetitionSetAndDeleteToken(t *testing.T) {
	for index, table := range []struct {
		inURL, inToken string
		inResp         []byte
		outErr         string
	}{
		{urlTest, tokenTest, []byte("{}"), ""},
		{"%%", tokenTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, tokenTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, tokenTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			err := service.PetitionSetToken(
				mock,
				table.inURL,
				tokenapp.SetTokenRequest{
					Token: table.inToken,
				},
			)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
		})
	}
}

func TestPetitionDeleteToken(t *testing.T) {
	for index, table := range []struct {
		inURL, inToken string
		inResp         []byte
		outErr         string
	}{
		{urlTest, tokenTest, []byte("{}"), ""},
		{"%%", tokenTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, tokenTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, tokenTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			err := service.PetitionDeleteToken(
				mock,
				table.inURL,
				tokenapp.DeleteTokenRequest{
					Token: table.inToken,
				},
			)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
		})
	}
}

func TestPetitionCheckToken(t *testing.T) {
	for index, table := range []struct {
		inURL, inToken string
		inResp         []byte
		outErr         string
	}{
		{urlTest, tokenTest, []byte("{}"), ""},
		{"%%", tokenTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, tokenTest, []byte(""), "unexpected end of JSON input"},
		// {urlTest, tokenTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr string

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			_, err := service.PetitionCheckToken(
				mock,
				table.inURL,
				tokenapp.CheckTokenRequest{
					Token: table.inToken,
				},
			)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
		})
	}
}
