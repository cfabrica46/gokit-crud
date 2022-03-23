package service

import (
	"bytes"
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
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.out)),
				}, nil
			})

			if tt.outErr == errWebServer.Error() {
				mock = newMockClient(func(req *http.Request) (*http.Response, error) {
					return nil, errWebServer
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
		{urlTest, []byte("{}"), ""},
		{"%%", []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
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
		{urlTest, userTest.Username, []byte("{}"), ""},
		{"%%", userTest.Username, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, userTest.Username, []byte(""), "unexpected end of JSON input"},
		{urlTest, userTest.Username, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			_, err := petitionGetIDByUsername(
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
		{urlTest, userTest.ID, []byte("{}"), ""},
		{"%%", userTest.ID, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, userTest.ID, []byte(""), "unexpected end of JSON input"},
		{urlTest, userTest.ID, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
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
		{
			urlTest,
			userTest.Username,
			userTest.Password,
			[]byte("{}"),
			"",
		},
		{
			"%%",
			userTest.Username,
			userTest.Password,
			[]byte("{}"),
			`parse "%%": invalid URL escape "%%"`,
		},
		{
			urlTest,
			userTest.Username,
			userTest.Password,
			[]byte(""),
			"unexpected end of JSON input",
		},
		{
			urlTest,
			userTest.Username,
			userTest.Password,
			[]byte(`{"err":"error"}`),
			"error",
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			_, err := petitionGetUserByUsernameAndPassword(
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
			userTest.Username,
			userTest.Password,
			userTest.Email,
			[]byte("{}"),
			"",
		},
		{
			"%%",
			userTest.Username,
			userTest.Password,
			userTest.Email,
			[]byte("{}"),
			`parse "%%": invalid URL escape "%%"`,
		},
		{
			urlTest,
			userTest.Username,
			userTest.Password,
			userTest.Email,
			[]byte(""),
			"unexpected end of JSON input",
		},
		{
			urlTest,
			userTest.Username,
			userTest.Password,
			userTest.Email,
			[]byte(`{"err":"error"}`),
			"error",
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			err := petitionInsertUser(
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
		{urlTest, userTest.ID, []byte("{}"), ""},
		{"%%", userTest.ID, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, userTest.ID, []byte(""), "unexpected end of JSON input"},
		{urlTest, userTest.ID, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
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
		{
			urlTest,
			userTest.Username,
			userTest.Email,
			"secret",
			userTest.ID,
			[]byte("{}"),
			"",
		},
		{
			"%%",
			userTest.Username,
			userTest.Email,
			"secret",
			userTest.ID,
			[]byte("{}"),
			`parse "%%": invalid URL escape "%%"`,
		},
		{
			urlTest,
			userTest.Username,
			userTest.Email,
			"secret",
			userTest.ID,
			[]byte(""),
			"unexpected end of JSON input",
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			_, err := petitionGenerateToken(
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
		{urlTest, tokenTest, "secret", []byte("{}"), ""},
		{"%%", tokenTest, "secret", []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, tokenTest, "secret", []byte(""), "unexpected end of JSON input"},
		{urlTest, tokenTest, "secret", []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			_, _, _, err := petitionExtractToken(
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
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
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
		{urlTest, tokenTest, []byte("{}"), ""},
		{"%%", tokenTest, []byte("{}"), `parse "%%": invalid URL escape "%%"`},
		{urlTest, tokenTest, []byte(""), "unexpected end of JSON input"},
		{urlTest, tokenTest, []byte(`{"err":"error"}`), "error"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			err := petitionDeleteToken(
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
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			mock := newMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			_, err := petitionCheckToken(
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
