package service_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/cfabrica46/gokit-crud/app/service"
	"github.com/stretchr/testify/assert"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
)

func TestMakePetition(t *testing.T) {
	for index, table := range []struct {
		inURL, inMethod string
		inBody          interface{}
		out             []byte
		outErr          string
		isError         bool
	}{
		{
			inURL:    urlTest,
			inMethod: http.MethodGet,
			inBody:   []byte("body"),
			out:      []byte("body"),
			outErr:   "",
			isError:  false,
		},
		{
			inURL:    urlTest,
			inMethod: http.MethodGet,
			inBody:   func() {},
			out:      []byte(nil),
			outErr:   "json: unsupported type: func()",
			isError:  true,
		},
		{
			inURL:    "%%",
			inMethod: http.MethodGet,
			inBody:   []byte("body"),
			out:      []byte(nil),
			outErr:   `parse "%%": invalid URL escape "%%"`,
			isError:  true,
		},
		{
			inURL:    urlTest,
			inMethod: http.MethodGet,
			inBody:   []byte("body"),
			out:      []byte(nil),
			outErr:   errWebServer.Error(),
			isError:  true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var result []byte
			var resultErr error

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

			result, resultErr = service.MakePetition(
				mock,
				table.inURL,
				table.inMethod,
				table.inBody,
			)

			if !table.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, table.outErr)
			}

			assert.Equal(t, table.out, result)
		})
	}
}

func TestPetitionGetAllUsers(t *testing.T) {
	usersTest := []dbapp.User{
		{
			ID:       idTest,
			Username: usernameTest,
			Password: passwordTest,
			Email:    emailTest,
		},
	}

	responseTest := struct {
		Users []dbapp.User
	}{
		Users: usersTest,
	}

	jsonUsersTest, err := json.Marshal(responseTest)
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		inURL    string
		inResp   []byte
		outUsers []dbapp.User
		outErr   string
		isError  bool
	}{
		{
			inURL:    urlTest,
			inResp:   jsonUsersTest,
			outUsers: usersTest,
			outErr:   "",
			isError:  false,
		},
		{
			inURL:    "%%",
			inResp:   []byte("{}"),
			outUsers: nil,
			outErr:   `parse "%%": invalid URL escape "%%"`,
			isError:  true,
		},
		{
			inURL:    urlTest,
			inResp:   []byte(""),
			outUsers: nil,
			outErr:   "unexpected end of JSON input",
			isError:  true,
		},
		{
			inURL:    urlTest,
			inResp:   []byte(`{"err":"error"}`),
			outUsers: nil,
			outErr:   "error",
			isError:  true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultUsers []dbapp.User
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			resultUsers, resultErr = service.PetitionGetAllUsers(mock, table.inURL)

			if !table.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, table.outErr)
			}

			assert.Equal(t, table.outUsers, resultUsers)
		})
	}
}

func TestPetitionGetIDByUsername(t *testing.T) {
	responseTest := struct {
		ID int
	}{
		ID: idTest,
	}

	jsonResponseTest, err := json.Marshal(responseTest)
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		inURL, inUsername string
		inResp            []byte
		outID             int
		outErr            string
		isError           bool
	}{
		{
			inURL:      urlTest,
			inUsername: usernameTest,
			inResp:     jsonResponseTest,
			outID:      idTest,
			outErr:     "",
			isError:    false,
		},
		{
			inURL:      "%%",
			inUsername: usernameTest,
			inResp:     []byte("{}"),
			outID:      0,
			outErr:     `parse "%%": invalid URL escape "%%"`,
			isError:    true,
		},
		{
			inURL:      urlTest,
			inUsername: usernameTest,
			inResp:     []byte(""),
			outID:      0,
			outErr:     "unexpected end of JSON input",
			isError:    true,
		},
		{
			inURL:      urlTest,
			inUsername: usernameTest,
			inResp:     []byte(`{"err":"error"}`),
			outID:      0,
			outErr:     "error",
			isError:    true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultID int
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			resultID, resultErr = service.PetitionGetIDByUsername(
				mock,
				table.inURL,
				dbapp.GetIDByUsernameRequest{
					Username: table.inUsername,
				},
			)

			if !table.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, table.outErr)
			}

			assert.Equal(t, table.outID, resultID)
		})
	}
}

func TestPetitionGetUserByID(t *testing.T) {
	responseTest := struct {
		User dbapp.User
	}{
		User: dbapp.User{
			ID:       idTest,
			Username: usernameTest,
			Password: passwordTest,
			Email:    emailTest,
		},
	}

	jsonResponseTest, err := json.Marshal(responseTest)
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		inURL   string
		inID    int
		inResp  []byte
		outUser dbapp.User
		outErr  string
		isError bool
	}{
		{
			inURL:   urlTest,
			inID:    idTest,
			inResp:  jsonResponseTest,
			outUser: responseTest.User,
			outErr:  "",
			isError: false,
		},
		{
			inURL:   "%%",
			inID:    idTest,
			inResp:  []byte("{}"),
			outUser: dbapp.User{},
			outErr:  `parse "%%": invalid URL escape "%%"`,
			isError: true,
		},
		{
			inURL:   urlTest,
			inID:    idTest,
			inResp:  []byte(""),
			outUser: dbapp.User{},
			outErr:  "unexpected end of JSON input",
			isError: true,
		},
		{
			inURL:   urlTest,
			inID:    idTest,
			inResp:  []byte(`{"err":"error"}`),
			outUser: dbapp.User{},
			outErr:  "error",
			isError: true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultUser dbapp.User
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			resultUser, resultErr = service.PetitionGetUserByID(
				mock,
				table.inURL,
				dbapp.GetUserByIDRequest{
					ID: table.inID,
				},
			)

			if !table.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, table.outErr)
			}

			assert.Equal(t, table.outUser, resultUser)
		})
	}
}

// func TestPetitionGetUserByUsernameAndPassword(t *testing.T) {
// 	for index, table := range []struct {
// 		inURL, inUsername, inPassword string
// 		inResp                        []byte
// 		outErr                        string
// 	}{
// 		{
// 			urlTest,
// 			usernameTest,
// 			passwordTest,
// 			[]byte("{}"),
// 			"",
// 		},
// 		{
// 			"%%",
// 			usernameTest,
// 			passwordTest,
// 			[]byte("{}"),
// 			`parse "%%": invalid URL escape "%%"`,
// 		},
// 		{
// 			urlTest,
// 			usernameTest,
// 			passwordTest,
// 			[]byte(""),
// 			"unexpected end of JSON input",
// 		},
// 		{
// 			urlTest,
// 			usernameTest,
// 			passwordTest,
// 			[]byte(`{"err":"error"}`),
// 			"error",
// 		},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultErr string

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				return &http.Response{
// 					StatusCode: http.StatusOK,
// 					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
// 				}, nil
// 			})

// 			_, err := service.PetitionGetUserByUsernameAndPassword(
// 				mock,
// 				table.inURL,
// 				dbapp.GetUserByUsernameAndPasswordRequest{
// 					Username: table.inUsername,
// 					Password: table.inPassword,
// 				},
// 			)
// 			if err != nil {
// 				resultErr = err.Error()
// 			}

// 			assert.Equal(t, table.outErr, resultErr)
// 		})
// 	}
// }

// func TestPetitionInsertUser(t *testing.T) {
// 	for index, table := range []struct {
// 		inURL, inUsername, inPassword, inEmail string
// 		inResp                                 []byte
// 		outErr                                 string
// 	}{
// 		{
// 			urlTest,
// 			usernameTest,
// 			passwordTest,
// 			emailTest,
// 			[]byte("{}"),
// 			"",
// 		},
// 		{
// 			"%%",
// 			usernameTest,
// 			passwordTest,
// 			emailTest,
// 			[]byte("{}"),
// 			`parse "%%": invalid URL escape "%%"`,
// 		},
// 		{
// 			urlTest,
// 			usernameTest,
// 			passwordTest,
// 			emailTest,
// 			[]byte(""),
// 			"unexpected end of JSON input",
// 		},
// 		{
// 			urlTest,
// 			usernameTest,
// 			passwordTest,
// 			emailTest,
// 			[]byte(`{"err":"error"}`),
// 			"error",
// 		},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultErr string

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				return &http.Response{
// 					StatusCode: http.StatusOK,
// 					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
// 				}, nil
// 			})

// 			err := service.PetitionInsertUser(
// 				mock,
// 				table.inURL,
// 				dbapp.InsertUserRequest{
// 					Username: table.inUsername,
// 					Password: table.inPassword,
// 					Email:    table.inEmail,
// 				},
// 			)
// 			if err != nil {
// 				resultErr = err.Error()
// 			}

// 			assert.Equal(t, table.outErr, resultErr)
// 		})
// 	}
// }

// func TestPetitionDeleteUser(t *testing.T) {
// 	for index, table := range []struct {
// 		inURL  string
// 		inID   int
// 		inResp []byte
// 		outErr string
// 	}{
// 		{urlTest, idTest, []byte("{}"), ""},
// 		{"%%", idTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
// 		{urlTest, idTest, []byte(""), "unexpected end of JSON input"},
// 		{urlTest, idTest, []byte(`{"err":"error"}`), "error"},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultErr string

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				return &http.Response{
// 					StatusCode: http.StatusOK,
// 					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
// 				}, nil
// 			})

// 			err := service.PetitionDeleteUser(
// 				mock,
// 				table.inURL,
// 				dbapp.DeleteUserRequest{
// 					ID: table.inID,
// 				},
// 			)
// 			if err != nil {
// 				resultErr = err.Error()
// 			}

// 			assert.Equal(t, table.outErr, resultErr)
// 		})
// 	}
// }

// func TestPetitionGenerateToken(t *testing.T) {
// 	for index, table := range []struct {
// 		inURL, inUsername, inEmail, inSecret string
// 		inID                                 int
// 		inResp                               []byte
// 		outErr                               string
// 	}{
// 		{
// 			urlTest,
// 			usernameTest,
// 			emailTest,
// 			secretTest,
// 			idTest,
// 			[]byte("{}"),
// 			"",
// 		},
// 		{
// 			"%%",
// 			usernameTest,
// 			emailTest,
// 			secretTest,
// 			idTest,
// 			[]byte("{}"),
// 			`parse "%%": invalid URL escape "%%"`,
// 		},
// 		{
// 			urlTest,
// 			usernameTest,
// 			emailTest,
// 			secretTest,
// 			idTest,
// 			[]byte(""),
// 			"unexpected end of JSON input",
// 		},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultErr string

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				return &http.Response{
// 					StatusCode: http.StatusOK,
// 					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
// 				}, nil
// 			})

// 			_, err := service.PetitionGenerateToken(
// 				mock,
// 				table.inURL,
// 				tokenapp.GenerateTokenRequest{
// 					ID:       table.inID,
// 					Username: table.inUsername,
// 					Email:    table.inEmail,
// 					Secret:   table.inSecret,
// 				},
// 			)
// 			if err != nil {
// 				resultErr = err.Error()
// 			}

// 			assert.Equal(t, table.outErr, resultErr)
// 		})
// 	}
// }

// func TestPetitionExtractToken(t *testing.T) {
// 	for index, table := range []struct {
// 		inURL, inToken, inSecret string
// 		inResp                   []byte
// 		outErr                   string
// 	}{
// 		{urlTest, tokenTest, secretTest, []byte("{}"), ""},
// 		{"%%", tokenTest, secretTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
// 		{urlTest, tokenTest, secretTest, []byte(""), "unexpected end of JSON input"},
// 		{urlTest, tokenTest, secretTest, []byte(`{"err":"error"}`), "error"},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultErr string

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				return &http.Response{
// 					StatusCode: http.StatusOK,
// 					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
// 				}, nil
// 			})

// 			_, _, _, err := service.PetitionExtractToken(
// 				mock,
// 				table.inURL,
// 				tokenapp.ExtractTokenRequest{
// 					Token:  table.inToken,
// 					Secret: table.inSecret,
// 				},
// 			)
// 			if err != nil {
// 				resultErr = err.Error()
// 			}

// 			assert.Equal(t, table.outErr, resultErr)
// 		})
// 	}
// }

// func TestPetitionSetAndDeleteToken(t *testing.T) {
// 	for index, table := range []struct {
// 		inURL, inToken string
// 		inResp         []byte
// 		outErr         string
// 	}{
// 		{urlTest, tokenTest, []byte("{}"), ""},
// 		{"%%", tokenTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
// 		{urlTest, tokenTest, []byte(""), "unexpected end of JSON input"},
// 		{urlTest, tokenTest, []byte(`{"err":"error"}`), "error"},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultErr string

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				return &http.Response{
// 					StatusCode: http.StatusOK,
// 					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
// 				}, nil
// 			})

// 			err := service.PetitionSetToken(
// 				mock,
// 				table.inURL,
// 				tokenapp.SetTokenRequest{
// 					Token: table.inToken,
// 				},
// 			)
// 			if err != nil {
// 				resultErr = err.Error()
// 			}

// 			assert.Equal(t, table.outErr, resultErr)
// 		})
// 	}
// }

// func TestPetitionDeleteToken(t *testing.T) {
// 	for index, table := range []struct {
// 		inURL, inToken string
// 		inResp         []byte
// 		outErr         string
// 	}{
// 		{urlTest, tokenTest, []byte("{}"), ""},
// 		{"%%", tokenTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
// 		{urlTest, tokenTest, []byte(""), "unexpected end of JSON input"},
// 		{urlTest, tokenTest, []byte(`{"err":"error"}`), "error"},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultErr string

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				return &http.Response{
// 					StatusCode: http.StatusOK,
// 					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
// 				}, nil
// 			})

// 			err := service.PetitionDeleteToken(
// 				mock,
// 				table.inURL,
// 				tokenapp.DeleteTokenRequest{
// 					Token: table.inToken,
// 				},
// 			)
// 			if err != nil {
// 				resultErr = err.Error()
// 			}

// 			assert.Equal(t, table.outErr, resultErr)
// 		})
// 	}
// }

// func TestPetitionCheckToken(t *testing.T) {
// 	for index, table := range []struct {
// 		inURL, inToken string
// 		inResp         []byte
// 		outErr         string
// 	}{
// 		{urlTest, tokenTest, []byte("{}"), ""},
// 		{"%%", tokenTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
// 		{urlTest, tokenTest, []byte(""), "unexpected end of JSON input"},
// 		// {urlTest, tokenTest, []byte(`{"err":"error"}`), "error"},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultErr string

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				return &http.Response{
// 					StatusCode: http.StatusOK,
// 					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
// 				}, nil
// 			})

// 			_, err := service.PetitionCheckToken(
// 				mock,
// 				table.inURL,
// 				tokenapp.CheckTokenRequest{
// 					Token: table.inToken,
// 				},
// 			)
// 			if err != nil {
// 				resultErr = err.Error()
// 			}

// 			assert.Equal(t, table.outErr, resultErr)
// 		})
// 	}
// }
