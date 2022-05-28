package service_test

/* import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/cfabrica46/gokit-crud/app/service"
	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	tokenapp "github.com/cfabrica46/gokit-crud/token-app/service"
	"github.com/stretchr/testify/assert"
)

func TestMakePetition(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name            string
		inURL, inMethod string
		outErr          string
		inBody          interface{}
		out             []byte
	}{
		{
			name:     "NoError",
			inURL:    urlTest,
			inMethod: http.MethodGet,
			inBody:   []byte("body"),
			out:      []byte("body"),
			outErr:   "",
		},
		{
			name:     "ErrorBadBodyRequest",
			inURL:    urlTest,
			inMethod: http.MethodGet,
			inBody:   func() {},
			out:      []byte(nil),
			outErr:   "json: unsupported type: func()",
		},
		{
			name:     "ErrorBadURL",
			inURL:    "%%",
			inMethod: http.MethodGet,
			inBody:   []byte("body"),
			out:      []byte(nil),
			outErr:   `parse "%%": invalid URL escape "%%"`,
		},
		{
			name:     "ErrorWebService",
			inURL:    urlTest,
			inMethod: http.MethodGet,
			inBody:   []byte("body"),
			out:      []byte(nil),
			outErr:   errWebServer.Error(),
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var result []byte
			var resultErr error

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

			result, resultErr = service.MakePetition(
				mock,
				tt.inURL,
				tt.inMethod,
				tt.inBody,
			)

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				log.Println(resultErr)
				log.Println(tt.outErr)
				assert.Contains(t, resultErr.Error(), tt.outErr)
			}

			assert.Equal(t, tt.out, result)
		})
	}
}

func TestPetitionGetAllUsers(t *testing.T) {
	t.Parallel()

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
		assert.Error(t, err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name     string
		inURL    string
		outErr   string
		inResp   []byte
		outUsers []dbapp.User
	}{
		{
			name:     "NoError",
			inURL:    urlTest,
			inResp:   goodJSONTest,
			outUsers: goodResponseTest.Users,
			outErr:   "",
		},
		{
			name:     "ErrorBadURL",
			inURL:    "%%",
			inResp:   []byte("{}"),
			outUsers: nil,
			outErr:   `parse "%%": invalid URL escape "%%"`,
		},
		{
			name:     "ErrorNoJSON",
			inURL:    urlTest,
			inResp:   []byte(""),
			outUsers: nil,
			outErr:   "unexpected end of JSON input",
		},
		{
			name:     "ErrorBadJSON",
			inURL:    urlTest,
			inResp:   badJSONTest,
			outUsers: nil,
			outErr:   "error",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultUsers []dbapp.User
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			resultUsers, resultErr = service.PetitionGetAllUsers(mock, tt.inURL)

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr.Error(), tt.outErr)
			}

			assert.Equal(t, tt.outUsers, resultUsers)
		})
	}
}

func TestPetitionGetIDByUsername(t *testing.T) {
	t.Parallel()

	goodResponseTest := dbapp.GetIDByUsernameResponse{
		ID: idTest,
	}

	badResponseTest := dbapp.GetIDByUsernameResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name              string
		inURL, inUsername string
		outErr            string
		inResp            []byte
		outID             int
	}{
		{
			name:       "NoError",
			inURL:      urlTest,
			inUsername: usernameTest,
			inResp:     goodJSONTest,
			outID:      idTest,
			outErr:     "",
		},
		{
			name:       "ErrorBadURL",
			inURL:      "%%",
			inUsername: usernameTest,
			inResp:     []byte("{}"),
			outID:      0,
			outErr:     `parse "%%": invalid URL escape "%%"`,
		},
		{
			name:       "ErrorNoJSON",
			inURL:      urlTest,
			inUsername: usernameTest,
			inResp:     []byte(""),
			outID:      0,
			outErr:     "unexpected end of JSON input",
		},
		{
			name:       "ErrorBadJSON",
			inURL:      urlTest,
			inUsername: usernameTest,
			inResp:     badJSONTest,
			outID:      0,
			outErr:     "error",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultID int
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			resultID, resultErr = service.PetitionGetIDByUsername(
				mock,
				tt.inURL,
				dbapp.GetIDByUsernameRequest{
					Username: tt.inUsername,
				},
			)

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr.Error(), tt.outErr)
			}

			assert.Equal(t, tt.outID, resultID)
		})
	}
}

func TestPetitionGetUserByID(t *testing.T) {
	t.Parallel()

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
		assert.Error(t, err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name    string
		inURL   string
		outErr  string
		outUser dbapp.User
		inResp  []byte
		inID    int
	}{
		{
			name:    "NoError",
			inURL:   urlTest,
			inID:    idTest,
			inResp:  goodJSONTest,
			outUser: goodResponseTest.User,
			outErr:  "",
		},
		{
			name:    "ErrorBadURL",
			inURL:   "%%",
			inID:    idTest,
			inResp:  []byte("{}"),
			outUser: dbapp.User{},
			outErr:  `parse "%%": invalid URL escape "%%"`,
		},
		{
			name:    "ErrorNoJSON",
			inURL:   urlTest,
			inID:    idTest,
			inResp:  []byte(""),
			outUser: dbapp.User{},
			outErr:  "unexpected end of JSON input",
		},
		{
			name:    "ErrorBadJSON",
			inURL:   urlTest,
			inID:    idTest,
			inResp:  badJSONTest,
			outUser: dbapp.User{},
			outErr:  "error",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultUser dbapp.User
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			resultUser, resultErr = service.PetitionGetUserByID(
				mock,
				tt.inURL,
				dbapp.GetUserByIDRequest{
					ID: tt.inID,
				},
			)

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr.Error(), tt.outErr)
			}

			assert.Equal(t, tt.outUser, resultUser)
		})
	}
}

func TestPetitionGetUserByUsernameAndPassword(t *testing.T) {
	t.Parallel()

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
		assert.Error(t, err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name       string
		inURL      string
		inUsername string
		inPassword string
		outErr     string
		outUser    dbapp.User
		inResp     []byte
	}{
		{
			name:       "NoError",
			inURL:      urlTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inResp:     goodJSONTest,
			outUser:    goodResponseTest.User,
			outErr:     "",
		},
		{
			name:       "ErrorBadURL",
			inURL:      "%%",
			inUsername: usernameTest,
			inPassword: passwordTest,
			inResp:     []byte("{}"),
			outUser:    dbapp.User{},
			outErr:     `parse "%%": invalid URL escape "%%"`,
		},
		{
			name:       "ErrorNoJSON",
			inURL:      urlTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inResp:     []byte(""),
			outUser:    dbapp.User{},
			outErr:     "unexpected end of JSON input",
		},
		{
			name:       "ErrorBadJSON",
			inURL:      urlTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inResp:     badJSONTest,
			outUser:    dbapp.User{},
			outErr:     "error",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultUser dbapp.User
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			resultUser, resultErr = service.PetitionGetUserByUsernameAndPassword(
				mock,
				tt.inURL,
				dbapp.GetUserByUsernameAndPasswordRequest{
					Username: tt.inUsername,
					Password: tt.inPassword,
				},
			)

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr.Error(), tt.outErr)
			}

			assert.Equal(t, tt.outUser, resultUser)
		})
	}
}

func TestPetitionInsertUser(t *testing.T) {
	t.Parallel()

	goodResponseTest := dbapp.InsertUserResponse{}

	badResponseTest := dbapp.InsertUserResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name       string
		inURL      string
		inUsername string
		inPassword string
		inEmail    string
		outErr     string
		inResp     []byte
	}{
		{
			name:       "NoError",
			inURL:      urlTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inResp:     goodJSONTest,
			outErr:     "",
		},
		{
			name:       "ErrorBadURL",
			inURL:      "%%",
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inResp:     []byte("{}"),
			outErr:     `parse "%%": invalid URL escape "%%"`,
		},
		{
			name:       "ErrorNoJSON",
			inURL:      urlTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inResp:     []byte(""),
			outErr:     "unexpected end of JSON input",
		},
		{
			name:       "ErrorBadJSON",
			inURL:      urlTest,
			inUsername: usernameTest,
			inPassword: passwordTest,
			inEmail:    emailTest,
			inResp:     badJSONTest,
			outErr:     "error",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			resultErr = service.PetitionInsertUser(
				mock,
				tt.inURL,
				dbapp.InsertUserRequest{
					Username: tt.inUsername,
					Password: tt.inPassword,
					Email:    tt.inEmail,
				},
			)

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr.Error(), tt.outErr)
			}
		})
	}
}

func TestPetitionDeleteUser(t *testing.T) {
	t.Parallel()

	goodResponseTest := dbapp.DeleteUserResponse{}

	badResponseTest := dbapp.DeleteUserResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name   string
		inURL  string
		outErr string
		inResp []byte
		inID   int
	}{
		{
			name:   "NoError",
			inURL:  urlTest,
			inID:   idTest,
			inResp: goodJSONTest,
			outErr: "",
		},
		{
			name:   "ErrorBadURL",
			inURL:  "%%",
			inID:   idTest,
			inResp: []byte("{}"),
			outErr: `parse "%%": invalid URL escape "%%"`,
		},
		{
			name:   "ErrorNoJSON",
			inURL:  urlTest,
			inID:   idTest,
			inResp: []byte(""),
			outErr: "unexpected end of JSON input",
		},
		{
			name:   "ErrorBadJSON",
			inURL:  urlTest,
			inID:   idTest,
			inResp: badJSONTest,
			outErr: "error",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			resultErr = service.PetitionDeleteUser(
				mock,
				tt.inURL,
				dbapp.DeleteUserRequest{
					ID: tt.inID,
				},
			)

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr.Error(), tt.outErr)
			}
		})
	}
}

func TestPetitionGenerateToken(t *testing.T) {
	t.Parallel()

	goodResponseTest := tokenapp.GenerateTokenResponse{
		Token: tokenTest,
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name       string
		inURL      string
		inUsername string
		inEmail    string
		inSecret   string
		outToken   string
		outErr     string
		inResp     []byte
		inID       int
	}{
		{
			name:       "NoError",
			inURL:      urlTest,
			inID:       idTest,
			inUsername: usernameTest,
			inEmail:    emailTest,
			inSecret:   secretTest,
			inResp:     goodJSONTest,
			outToken:   tokenTest,
			outErr:     "",
		},
		{
			name:       "ErrorBadURL",
			inURL:      "%%",
			inID:       idTest,
			inUsername: usernameTest,
			inEmail:    emailTest,
			inSecret:   secretTest,
			inResp:     []byte("{}"),
			outToken:   "",
			outErr:     `parse "%%": invalid URL escape "%%"`,
		},
		{
			name:       "ErrorNoJSON",
			inURL:      urlTest,
			inID:       idTest,
			inUsername: usernameTest,
			inEmail:    emailTest,
			inSecret:   secretTest,
			inResp:     []byte(""),
			outToken:   "",
			outErr:     "unexpected end of JSON input",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultToken string
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			resultToken, resultErr = service.PetitionGenerateToken(
				mock,
				tt.inURL,
				tokenapp.GenerateTokenRequest{
					ID:       tt.inID,
					Username: tt.inUsername,
					Email:    tt.inEmail,
					Secret:   tt.inSecret,
				},
			)

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr.Error(), tt.outErr)
			}

			assert.Equal(t, tt.outToken, resultToken)
		})
	}
}

func TestPetitionExtractToken(t *testing.T) {
	t.Parallel()

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
		assert.Error(t, err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name        string
		inURL       string
		inToken     string
		inSecret    string
		outUsername string
		outEmail    string
		outErr      string
		inResp      []byte
		outID       int
	}{
		{
			name:        "NoError",
			inURL:       urlTest,
			inToken:     tokenTest,
			inSecret:    secretTest,
			inResp:      goodJSONTest,
			outID:       idTest,
			outUsername: usernameTest,
			outEmail:    emailTest,
			outErr:      "",
		},
		{
			name:        "ErrorBadURL",
			inURL:       "%%",
			inToken:     tokenTest,
			inSecret:    secretTest,
			inResp:      []byte("{}"),
			outID:       0,
			outUsername: "",
			outEmail:    "",
			outErr:      `parse "%%": invalid URL escape "%%"`,
		},
		{
			name:        "ErrorNoJSON",
			inURL:       urlTest,
			inToken:     tokenTest,
			inSecret:    secretTest,
			inResp:      []byte(""),
			outID:       0,
			outUsername: "",
			outEmail:    "",
			outErr:      "unexpected end of JSON input",
		},
		{
			name:        "ErrorBadJSON",
			inURL:       urlTest,
			inToken:     tokenTest,
			inSecret:    secretTest,
			inResp:      badJSONTest,
			outID:       0,
			outUsername: "",
			outEmail:    "",
			outErr:      "error",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultID int
			var resultUsername, resultEmail string
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			resultID, resultUsername, resultEmail, resultErr = service.PetitionExtractToken(
				mock,
				tt.inURL,
				tokenapp.ExtractTokenRequest{
					Token:  tt.inToken,
					Secret: tt.inSecret,
				},
			)

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr.Error(), tt.outErr)
			}

			assert.Equal(t, tt.outID, resultID)
			assert.Equal(t, tt.outUsername, resultUsername)
			assert.Equal(t, tt.outEmail, resultEmail)
		})
	}
}

func TestPetitionSetToken(t *testing.T) {
	t.Parallel()

	goodResponseTest := tokenapp.SetTokenResponse{}

	badResponseTest := tokenapp.SetTokenResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name    string
		inURL   string
		inToken string
		outErr  string
		inResp  []byte
	}{
		{
			name:    "NoError",
			inURL:   urlTest,
			inToken: tokenTest,
			inResp:  goodJSONTest,
			outErr:  "",
		},
		{
			name:    "ErrorBadURL",
			inURL:   "%%",
			inToken: tokenTest,
			inResp:  []byte("{}"),
			outErr:  `parse "%%": invalid URL escape "%%"`,
		},
		{
			name:    "ErrorNoJSON",
			inURL:   urlTest,
			inToken: tokenTest,
			inResp:  []byte(""),
			outErr:  "unexpected end of JSON input",
		},
		{
			name:    "ErrorBadJSON",
			inURL:   urlTest,
			inToken: tokenTest,
			inResp:  badJSONTest,
			outErr:  "error",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			resultErr = service.PetitionSetToken(
				mock,
				tt.inURL,
				tokenapp.SetTokenRequest{
					Token: tt.inToken,
				},
			)

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr.Error(), tt.outErr)
			}
		})
	}
}

func TestPetitionDeleteToken(t *testing.T) {
	t.Parallel()

	goodResponseTest := tokenapp.DeleteTokenResponse{}

	badResponseTest := tokenapp.DeleteTokenResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name    string
		inURL   string
		inToken string
		outErr  string
		inResp  []byte
	}{
		{
			name:    "NoError",
			inURL:   urlTest,
			inToken: tokenTest,
			inResp:  goodJSONTest,
			outErr:  "",
		},
		{
			name:    "ErrorBadURL",
			inURL:   "%%",
			inToken: tokenTest,
			inResp:  []byte("{}"),
			outErr:  `parse "%%": invalid URL escape "%%"`,
		},
		{
			name:    "ErrorNoJSON",
			inURL:   urlTest,
			inToken: tokenTest,
			inResp:  []byte(""),
			outErr:  "unexpected end of JSON input",
		},
		{
			name:    "ErrorBadJSON",
			inURL:   urlTest,
			inToken: tokenTest,
			inResp:  badJSONTest,
			outErr:  "error",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			resultErr = service.PetitionDeleteToken(
				mock,
				tt.inURL,
				tokenapp.DeleteTokenRequest{
					Token: tt.inToken,
				},
			)

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr.Error(), tt.outErr)
			}
		})
	}
}

func TestPetitionCheckToken(t *testing.T) {
	t.Parallel()

	goodResponseTest := tokenapp.CheckTokenResponse{
		Check: true,
	}

	badResponseTest := tokenapp.DeleteTokenResponse{
		Err: "error",
	}

	goodJSONTest, err := json.Marshal(goodResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	badJSONTest, err := json.Marshal(badResponseTest)
	if err != nil {
		assert.Error(t, err)
	}

	for _, tt := range []struct {
		name     string
		inURL    string
		inToken  string
		outErr   string
		inResp   []byte
		outCheck bool
	}{
		{
			name:     "NoError",
			inURL:    urlTest,
			inToken:  tokenTest,
			inResp:   goodJSONTest,
			outCheck: true,
			outErr:   "",
		},
		{
			name:     "ErrorBadURL",
			inURL:    "%%",
			inToken:  tokenTest,
			inResp:   []byte("{}"),
			outCheck: false,
			outErr:   `parse "%%": invalid URL escape "%%"`,
		},
		{
			name:     "ErrorNoJSON",
			inURL:    urlTest,
			inToken:  tokenTest,
			inResp:   []byte(""),
			outCheck: false,
			outErr:   "unexpected end of JSON input",
		},
		{
			name:     "ErrorBadJSON",
			inURL:    urlTest,
			inToken:  tokenTest,
			inResp:   badJSONTest,
			outCheck: false,
			outErr:   "error",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var resultCheck bool
			var resultErr error

			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(tt.inResp)),
				}, nil
			})

			resultCheck, resultErr = service.PetitionCheckToken(
				mock,
				tt.inURL,
				tokenapp.CheckTokenRequest{
					Token: tt.inToken,
				},
			)

			if tt.name == nameNoError {
				assert.Empty(t, resultErr)
			} else {
				assert.Contains(t, resultErr.Error(), tt.outErr)
			}

			assert.Equal(t, tt.outCheck, resultCheck)
		})
	}
} */
