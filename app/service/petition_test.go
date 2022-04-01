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
	tokenapp "github.com/cfabrica46/gokit-crud/token-app/service"
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
	goodResponseTest := dbapp.GetAllUsersResponse{
		Users: []dbapp.User{
			{
				ID:       idTest,
				Username: usernameTest,
				Password: passwordTest,
				Email:    emailTest,
			},
		},
	}

	badResponseTest := dbapp.GetAllUsersResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		t.Error(err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
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
			inResp:   goodJSONTest,
			outUsers: goodResponseTest.Users,
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
			inResp:   badJSONTest,
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
	goodResponseTest := dbapp.GetIDByUsernameResponse{
		ID: idTest,
	}

	badResponseTest := dbapp.GetIDByUsernameResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		t.Error(err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
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
			inResp:     goodJSONTest,
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
			inResp:     badJSONTest,
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
	goodResponseTest := dbapp.GetUserByIDResponse{
		User: dbapp.User{
			ID:       idTest,
			Username: usernameTest,
			Password: passwordTest,
			Email:    emailTest,
		},
	}

	badResponseTest := dbapp.GetUserByIDResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		t.Error(err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
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
			inResp:  goodJSONTest,
			outUser: goodResponseTest.User,
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
			inResp:  badJSONTest,
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

func TestPetitionGetUserByUsernameAndPassword(t *testing.T) {
	goodResponseTest := dbapp.GetUserByIDResponse{
		User: dbapp.User{
			ID:       idTest,
			Username: usernameTest,
			Password: passwordTest,
			Email:    emailTest,
		},
	}

	badResponseTest := dbapp.GetUserByIDResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		t.Error(err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		inURL      string
		inUsername string
		inPassword string
		inResp     []byte
		outUser    dbapp.User
		outErr     string
		isError    bool
	}{
		{
			inURL:      urlTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inResp:     goodJSONTest,
			outUser:    goodResponseTest.User,
			outErr:     "",
			isError:    false,
		},
		{
			inURL:      "%%",
			inUsername: usernameTest,
			inPassword: passwordTest,
			inResp:     []byte("{}"),
			outUser:    dbapp.User{},
			outErr:     `parse "%%": invalid URL escape "%%"`,
			isError:    true,
		},
		{
			inURL:      urlTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inResp:     []byte(""),
			outUser:    dbapp.User{},
			outErr:     "unexpected end of JSON input",
			isError:    true,
		},
		{
			inURL:      urlTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inResp:     badJSONTest,
			outUser:    dbapp.User{},
			outErr:     "error",
			isError:    true,
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

			resultUser, resultErr = service.PetitionGetUserByUsernameAndPassword(
				mock,
				table.inURL,
				dbapp.GetUserByUsernameAndPasswordRequest{
					Username: table.inUsername,
					Password: table.inPassword,
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

func TestPetitionInsertUser(t *testing.T) {
	goodResponseTest := dbapp.InsertUserResponse{}

	badResponseTest := dbapp.InsertUserResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		t.Error(err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		inURL      string
		inUsername string
		inPassword string
		inEmail    string
		inResp     []byte
		outErr     string
		isError    bool
	}{
		{
			inURL:      urlTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inResp:     goodJSONTest,
			outErr:     "",
			isError:    false,
		},
		{
			inURL:      "%%",
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inResp:     []byte("{}"),
			outErr:     `parse "%%": invalid URL escape "%%"`,
			isError:    true,
		},
		{
			inURL:      urlTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inResp:     []byte(""),
			outErr:     "unexpected end of JSON input",
			isError:    true,
		},
		{
			inURL:      urlTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inResp:     badJSONTest,
			outErr:     "error",
			isError:    true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			resultErr = service.PetitionInsertUser(
				mock,
				table.inURL,
				dbapp.InsertUserRequest{
					Username: table.inUsername,
					Password: table.inPassword,
					Email:    table.inEmail,
				},
			)

			if !table.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, table.outErr)
			}
		})
	}
}

func TestPetitionDeleteUser(t *testing.T) {
	goodResponseTest := dbapp.DeleteUserResponse{}

	badResponseTest := dbapp.DeleteUserResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		t.Error(err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		inURL   string
		inID    int
		inResp  []byte
		outErr  string
		isError bool
	}{
		{
			inURL:   urlTest,
			inID:    idTest,
			inResp:  goodJSONTest,
			outErr:  "",
			isError: false,
		},
		{
			inURL:   "%%",
			inID:    idTest,
			inResp:  []byte("{}"),
			outErr:  `parse "%%": invalid URL escape "%%"`,
			isError: true,
		},
		{
			inURL:   urlTest,
			inID:    idTest,
			inResp:  []byte(""),
			outErr:  "unexpected end of JSON input",
			isError: true,
		},
		{
			inURL:   urlTest,
			inID:    idTest,
			inResp:  badJSONTest,
			outErr:  "error",
			isError: true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			resultErr = service.PetitionDeleteUser(
				mock,
				table.inURL,
				dbapp.DeleteUserRequest{
					ID: table.inID,
				},
			)

			if !table.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, table.outErr)
			}
		})
	}
}

func TestPetitionGenerateToken(t *testing.T) {
	goodResponseTest := tokenapp.GenerateTokenResponse{
		Token: tokenTest,
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		inURL      string
		inID       int
		inUsername string
		inEmail    string
		inSecret   string
		inResp     []byte
		outToken   string
		outErr     string
		isError    bool
	}{
		{
			inURL:      urlTest,
			inID:       idTest,
			inUsername: usernameTest,
			inEmail:    emailTest,
			inSecret:   secretTest,
			inResp:     goodJSONTest,
			outToken:   tokenTest,
			outErr:     "",
			isError:    false,
		},
		{
			inURL:      "%%",
			inID:       idTest,
			inUsername: usernameTest,
			inEmail:    emailTest,
			inSecret:   secretTest,
			inResp:     []byte("{}"),
			outToken:   "",
			outErr:     `parse "%%": invalid URL escape "%%"`,
			isError:    true,
		},
		{
			inURL:      urlTest,
			inID:       idTest,
			inUsername: usernameTest,
			inEmail:    emailTest,
			inSecret:   secretTest,
			inResp:     []byte(""),
			outToken:   "",
			outErr:     "unexpected end of JSON input",
			isError:    true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultToken string
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			resultToken, resultErr = service.PetitionGenerateToken(
				mock,
				table.inURL,
				tokenapp.GenerateTokenRequest{
					ID:       table.inID,
					Username: table.inUsername,
					Email:    table.inEmail,
					Secret:   table.inSecret,
				},
			)

			if !table.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, table.outErr)
			}

			assert.Equal(t, table.outToken, resultToken)
		})
	}
}

func TestPetitionExtractToken(t *testing.T) {
	goodResponseTest := tokenapp.ExtractTokenResponse{
		ID:       idTest,
		Username: usernameTest,
		Email:    emailTest,
	}

	badResponseTest := tokenapp.ExtractTokenResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		t.Error(err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		inURL       string
		inToken     string
		inSecret    string
		inResp      []byte
		outID       int
		outUsername string
		outEmail    string
		outErr      string
		isError     bool
	}{
		{
			inURL:       urlTest,
			inToken:     tokenTest,
			inSecret:    secretTest,
			inResp:      goodJSONTest,
			outID:       idTest,
			outUsername: usernameTest,
			outEmail:    emailTest,
			outErr:      "",
			isError:     false,
		},
		{
			inURL:       "%%",
			inToken:     tokenTest,
			inSecret:    secretTest,
			inResp:      []byte("{}"),
			outID:       0,
			outUsername: "",
			outEmail:    "",
			outErr:      `parse "%%": invalid URL escape "%%"`,
			isError:     true,
		},
		{
			inURL:       urlTest,
			inToken:     tokenTest,
			inSecret:    secretTest,
			inResp:      []byte(""),
			outID:       0,
			outUsername: "",
			outEmail:    "",
			outErr:      "unexpected end of JSON input",
			isError:     true,
		},
		{
			inURL:       urlTest,
			inToken:     tokenTest,
			inSecret:    secretTest,
			inResp:      badJSONTest,
			outID:       0,
			outUsername: "",
			outEmail:    "",
			outErr:      "error",
			isError:     true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultID int
			var resultUsername, resultEmail string
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			resultID, resultUsername, resultEmail, resultErr = service.PetitionExtractToken(
				mock,
				table.inURL,
				tokenapp.ExtractTokenRequest{
					Token:  table.inToken,
					Secret: table.inSecret,
				},
			)

			if !table.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, table.outErr)
			}

			assert.Equal(t, table.outID, resultID)
			assert.Equal(t, table.outUsername, resultUsername)
			assert.Equal(t, table.outEmail, resultEmail)
		})
	}
}

func TestPetitionSetToken(t *testing.T) {
	goodResponseTest := tokenapp.SetTokenResponse{}

	badResponseTest := tokenapp.SetTokenResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		t.Error(err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		inURL   string
		inToken string
		inResp  []byte
		outErr  string
		isError bool
	}{
		{
			inURL:   urlTest,
			inToken: tokenTest,
			inResp:  goodJSONTest,
			outErr:  "",
			isError: false,
		},
		{
			inURL:   "%%",
			inToken: tokenTest,
			inResp:  []byte("{}"),
			outErr:  `parse "%%": invalid URL escape "%%"`,
			isError: true,
		},
		{
			inURL:   urlTest,
			inToken: tokenTest,
			inResp:  []byte(""),
			outErr:  "unexpected end of JSON input",
			isError: true,
		},
		{
			inURL:   urlTest,
			inToken: tokenTest,
			inResp:  badJSONTest,
			outErr:  "error",
			isError: true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			resultErr = service.PetitionSetToken(
				mock,
				table.inURL,
				tokenapp.SetTokenRequest{
					Token: table.inToken,
				},
			)

			if !table.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, table.outErr)
			}
		})
	}
}

func TestPetitionDeleteToken(t *testing.T) {
	goodResponseTest := tokenapp.DeleteTokenResponse{}

	badResponseTest := tokenapp.DeleteTokenResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		t.Error(err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		inURL   string
		inToken string
		inResp  []byte
		outErr  string
		isError bool
	}{
		{
			inURL:   urlTest,
			inToken: tokenTest,
			inResp:  goodJSONTest,
			outErr:  "",
			isError: false,
		},
		{
			inURL:   "%%",
			inToken: tokenTest,
			inResp:  []byte("{}"),
			outErr:  `parse "%%": invalid URL escape "%%"`,
			isError: true,
		},
		{
			inURL:   urlTest,
			inToken: tokenTest,
			inResp:  []byte(""),
			outErr:  "unexpected end of JSON input",
			isError: true,
		},
		{
			inURL:   urlTest,
			inToken: tokenTest,
			inResp:  badJSONTest,
			outErr:  "error",
			isError: true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			resultErr = service.PetitionDeleteToken(
				mock,
				table.inURL,
				tokenapp.DeleteTokenRequest{
					Token: table.inToken,
				},
			)

			if !table.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, table.outErr)
			}
		})
	}
}

func TestPetitionCheckToken(t *testing.T) {
	goodResponseTest := tokenapp.CheckTokenResponse{
		Check: true,
	}

	badResponseTest := tokenapp.DeleteTokenResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		t.Error(err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		inURL    string
		inToken  string
		inResp   []byte
		outCheck bool
		outErr   string
		isError  bool
	}{
		{
			inURL:    urlTest,
			inToken:  tokenTest,
			inResp:   goodJSONTest,
			outCheck: true,
			outErr:   "",
			isError:  false,
		},
		{
			inURL:    "%%",
			inToken:  tokenTest,
			inResp:   []byte("{}"),
			outCheck: false,
			outErr:   `parse "%%": invalid URL escape "%%"`,
			isError:  true,
		},
		{
			inURL:    urlTest,
			inToken:  tokenTest,
			inResp:   []byte(""),
			outCheck: false,
			outErr:   "unexpected end of JSON input",
			isError:  true,
		},
		{
			inURL:    urlTest,
			inToken:  tokenTest,
			inResp:   badJSONTest,
			outCheck: false,
			outErr:   "error",
			isError:  true,
		},
	} {
		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
			var resultCheck bool
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(table.inResp)),
				}, nil
			})

			resultCheck, resultErr = service.PetitionCheckToken(
				mock,
				table.inURL,
				tokenapp.CheckTokenRequest{
					Token: table.inToken,
				},
			)

			if !table.isError {
				assert.Nil(t, resultErr)
			} else {
				assert.ErrorContains(t, resultErr, table.outErr)
			}

			assert.Equal(t, table.outCheck, resultCheck)
		})
	}
}

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
