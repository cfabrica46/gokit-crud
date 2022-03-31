package service_test

import (
	"errors"
)

const (
	idTest       int    = 1
	usernameTest string = "username"
	passwordTest string = "password"
	emailTest    string = "email@email.com"
	secretTest   string = "secret"

	urlTest       string = "localhost:8080"
	dbHostTest    string = "db"
	tokenHostTest string = "token"
	portTest      string = "8080"
	tokenTest     string = "token"

	schemaNameTest string = "%v"
)

var (
	errWebServer        = errors.New("error from web server")
	errNotTypeIndicated = errors.New("response is not of the type indicated")
)

// func TestSignUp(t *testing.T) {
// 	for index, table := range []struct {
// 		inUsername, inPassword, inEmail string
// 		isError                         bool
// 		url                             string
// 		method                          string
// 	}{
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			inEmail:    emailTest,
// 			isError:    false,
// 			url:        "http://token:8080/generate",
// 			method:     http.MethodPost,
// 		},
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			inEmail:    emailTest,
// 			isError:    true,
// 			url:        "http://db:8080/user",
// 			method:     http.MethodPost,
// 		},
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			inEmail:    emailTest,
// 			isError:    true,
// 			url:        "http://db:8080/id/username",
// 			method:     http.MethodGet,
// 		},
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			inEmail:    emailTest,
// 			isError:    true,
// 			url:        "http://token:8080/generate",
// 			method:     http.MethodPost,
// 		},
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			inEmail:    emailTest,
// 			isError:    true,
// 			url:        "http://token:8080/token",
// 			method:     http.MethodPost,
// 		},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultToken string
// 			var resultErr error
// 			var tokenResponse, errorResponse string

// 			if table.isError {
// 				errorResponse = errWebServer.Error()
// 			} else {
// 				tokenResponse = tokenTest
// 			}

// 			testResp := struct {
// 				ID    int    `json:"id"`
// 				Token string `json:"token"`
// 				Err   string `json:"err"`
// 			}{
// 				ID:    idTest,
// 				Token: tokenResponse,
// 				Err:   errorResponse,
// 			}

// 			jsonData, err := json.Marshal(testResp)
// 			if err != nil {
// 				t.Error(err)
// 			}

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				response := &http.Response{Body: io.NopCloser(bytes.NewReader([]byte("{}")))}

// 				if req.URL.String() == table.url {
// 					if req.Method == table.method {
// 						response = &http.Response{
// 							StatusCode: http.StatusOK,
// 							Body:       io.NopCloser(bytes.NewReader(jsonData)),
// 						}
// 					}
// 				}

// 				return response, nil
// 			})

// 			svc := service.NewService(
// 				mock,
// 				dbHostTest,
// 				portTest,
// 				tokenHostTest,
// 				portTest,
// 				secretTest,
// 			)

// 			resultToken, resultErr = svc.SignUp(table.inUsername, table.inPassword, table.inEmail)

// 			if !table.isError {
// 				assert.Nil(t, resultErr)
// 			} else {
// 				assert.ErrorContains(t, resultErr, errorResponse)
// 			}
// 			assert.Equal(t, tokenResponse, resultToken)
// 		})
// 	}
// }

// func TestSignIn(t *testing.T) {
// 	for index, table := range []struct {
// 		inUsername, inPassword string
// 		isError                bool
// 		url                    string
// 		method                 string
// 	}{
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			isError:    false,
// 			url:        "http://token:8080/generate",
// 			method:     http.MethodPost,
// 		},
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			isError:    true,
// 			url:        "http://db:8080/user/username_password",
// 			method:     http.MethodGet,
// 		},
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			isError:    true,
// 			url:        "http://token:8080/generate",
// 			method:     http.MethodPost,
// 		},
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			isError:    true,
// 			url:        "http://token:8080/token",
// 			method:     http.MethodPost,
// 		},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultToken string
// 			var resultErr error
// 			var tokenResponse, errorResponse string

// 			if table.isError {
// 				errorResponse = errWebServer.Error()
// 			} else {
// 				tokenResponse = tokenTest
// 			}

// 			testResp := struct {
// 				User  dbapp.User `json:"user"`
// 				Token string     `json:"token"`
// 				Err   string     `json:"err"`
// 			}{
// 				User: dbapp.User{
// 					ID:       idTest,
// 					Username: usernameTest,
// 					Password: passwordTest,
// 					Email:    emailTest,
// 				},
// 				Token: tokenResponse,
// 				Err:   errorResponse,
// 			}

// 			jsonData, err := json.Marshal(testResp)
// 			if err != nil {
// 				t.Error(err)
// 			}

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				response := &http.Response{Body: io.NopCloser(bytes.NewReader([]byte("{}")))}

// 				if req.URL.String() == table.url {
// 					if req.Method == table.method {
// 						response = &http.Response{
// 							StatusCode: http.StatusOK,
// 							Body:       io.NopCloser(bytes.NewReader(jsonData)),
// 						}
// 					}
// 				}

// 				return response, nil
// 			})

// 			svc := service.NewService(
// 				mock,
// 				dbHostTest,
// 				portTest,
// 				tokenHostTest,
// 				portTest,
// 				secretTest,
// 			)

// 			resultToken, resultErr = svc.SignIn(table.inUsername, table.inPassword)

// 			if !table.isError {
// 				assert.Nil(t, resultErr)
// 			} else {
// 				assert.ErrorContains(t, resultErr, errorResponse)
// 			}
// 			assert.Equal(t, tokenResponse, resultToken)
// 		})
// 	}
// }

// func TestLogOut(t *testing.T) {
// 	for index, table := range []struct {
// 		inToken  string
// 		outCheck bool

// 		isError bool
// 		url     string
// 		method  string
// 	}{
// 		{
// 			inToken:  tokenTest,
// 			outCheck: true,
// 			isError:  false,
// 			url:      "http://token:8080/check",
// 			method:   http.MethodPost,
// 		},
// 		{
// 			inToken:  tokenTest,
// 			outCheck: true,
// 			isError:  true,
// 			url:      "http://token:8080/check",
// 			method:   http.MethodPost,
// 		},
// 		{
// 			inToken:  tokenTest,
// 			outCheck: false,
// 			isError:  true,
// 			url:      "http://token:8080/check",
// 			method:   http.MethodPost,
// 		},
// 		{
// 			tokenTest,
// 			true,
// 			true,
// 			"http://token:8080/token",
// 			http.MethodDelete,
// 		},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultErr error
// 			var errorResponse string

// 			if table.isError {
// 				errorResponse = errWebServer.Error()
// 			}

// 			testResp := struct {
// 				Check bool   `json:"check"`
// 				Err   string `json:"err"`
// 			}{
// 				Check: table.outCheck,
// 				Err:   errorResponse,
// 			}

// 			if !table.outCheck {
// 				testResp.Err = ""
// 				errorResponse = service.ErrTokenNotValid.Error()
// 			}

// 			jsonData, err := json.Marshal(testResp)
// 			if err != nil {
// 				t.Error(err)
// 			}

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				testCheck := struct {
// 					Check bool `json:"check"`
// 				}{
// 					Check: table.outCheck,
// 				}

// 				jsonCheck, err := json.Marshal(testCheck)
// 				if err != nil {
// 					t.Error(err)
// 				}

// 				response := &http.Response{Body: io.NopCloser(bytes.NewReader(jsonCheck))}

// 				if req.URL.String() == table.url {
// 					if req.Method == table.method {
// 						response = &http.Response{
// 							StatusCode: http.StatusOK,
// 							Body:       io.NopCloser(bytes.NewReader(jsonData)),
// 						}
// 					}
// 				}

// 				return response, nil
// 			})

// 			svc := service.NewService(
// 				mock,
// 				dbHostTest,
// 				portTest,
// 				tokenHostTest,
// 				portTest,
// 				secretTest,
// 			)

// 			resultErr = svc.LogOut(table.inToken)

// 			if !table.isError {
// 				assert.Nil(t, resultErr)
// 			} else {
// 				assert.ErrorContains(t, resultErr, errorResponse)
// 			}
// 		})
// 	}
// }

// func TestGetAllUsers(t *testing.T) {
// 	log.SetFlags(log.Lshortfile)

// 	for index, table := range []struct {
// 		outUsers []dbapp.User

// 		isError bool
// 		url     string
// 		method  string
// 	}{
// 		{
// 			outUsers: []dbapp.User{
// 				{
// 					ID:       idTest,
// 					Username: usernameTest,
// 					Password: passwordTest,
// 					Email:    emailTest,
// 				},
// 			},
// 			isError: false,
// 			url:     "http://db:8080/users",
// 			method:  http.MethodGet,
// 		},
// 		{
// 			outUsers: nil,
// 			isError:  true,
// 			url:      "http://db:8080/users",
// 			method:   http.MethodGet,
// 		},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultUsers []dbapp.User
// 			var resultErr error
// 			var errorResponse string

// 			if table.isError {
// 				errorResponse = errWebServer.Error()
// 			}

// 			testResp := struct {
// 				Users []dbapp.User `json:"users"`
// 				Err   string       `json:"err"`
// 			}{
// 				Users: []dbapp.User{{
// 					ID:       idTest,
// 					Username: usernameTest,
// 					Password: passwordTest,
// 					Email:    emailTest,
// 				}},
// 				Err: errorResponse,
// 			}

// 			jsonData, err := json.Marshal(testResp)
// 			if err != nil {
// 				t.Error(err)
// 			}

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				return &http.Response{
// 					StatusCode: http.StatusOK,
// 					Body:       ioutil.NopCloser(bytes.NewReader(jsonData)),
// 				}, nil
// 			})

// 			svc := service.NewService(
// 				mock,
// 				dbHostTest,
// 				portTest,
// 				tokenHostTest,
// 				portTest,
// 				secretTest,
// 			)

// 			resultUsers, resultErr = svc.GetAllUsers()

// 			if !table.isError {
// 				assert.Nil(t, resultErr)
// 			} else {
// 				assert.ErrorContains(t, resultErr, errorResponse)
// 			}

// 			assert.Equal(t, table.outUsers, resultUsers)
// 		})
// 	}
// }

// func TestProfile(t *testing.T) {
// 	for index, table := range []struct {
// 		inToken  string
// 		outUser  dbapp.User
// 		outCheck bool

// 		isError bool
// 		url     string
// 		method  string
// 	}{
// 		{
// 			inToken: tokenTest,
// 			outUser: dbapp.User{
// 				ID:       idTest,
// 				Username: usernameTest,
// 				Password: passwordTest,
// 				Email:    emailTest,
// 			},
// 			outCheck: true,
// 			isError:  false,
// 			url:      "http://db:8080/user/id",
// 			method:   http.MethodGet,
// 		},
// 		{
// 			inToken:  tokenTest,
// 			outUser:  dbapp.User{},
// 			outCheck: true,
// 			isError:  true,
// 			url:      "http://token:8080/check",
// 			method:   http.MethodPost,
// 		},
// 		{
// 			inToken:  tokenTest,
// 			outUser:  dbapp.User{},
// 			outCheck: false,
// 			isError:  true,
// 			url:      "http://token:8080/check",
// 			method:   http.MethodPost,
// 		},
// 		{
// 			inToken:  tokenTest,
// 			outUser:  dbapp.User{},
// 			outCheck: true,
// 			isError:  true,
// 			url:      "http://token:8080/extract",
// 			method:   http.MethodPost,
// 		},
// 		{
// 			inToken:  tokenTest,
// 			outUser:  dbapp.User{},
// 			outCheck: true,
// 			isError:  true,
// 			url:      "http://db:8080/user/id",
// 			method:   http.MethodGet,
// 		},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultUser dbapp.User
// 			var resultErr error
// 			var errorResponse string

// 			if table.isError {
// 				errorResponse = errWebServer.Error()
// 			}

// 			testResp := struct {
// 				User     dbapp.User `json:"user"`
// 				ID       int        `json:"id"`
// 				Username string     `json:"username"`
// 				Email    string     `json:"email"`
// 				Check    bool       `json:"check"`
// 				Err      string     `json:"err"`
// 			}{
// 				User:     table.outUser,
// 				ID:       table.outUser.ID,
// 				Username: table.outUser.Username,
// 				Email:    table.outUser.Email,
// 				Check:    table.outCheck,
// 				Err:      errorResponse,
// 			}

// 			if !table.outCheck {
// 				testResp.Err = ""
// 				errorResponse = service.ErrTokenNotValid.Error()
// 			}

// 			jsonData, err := json.Marshal(testResp)
// 			if err != nil {
// 				t.Error(err)
// 			}

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				testCheck := struct {
// 					Check bool `json:"check"`
// 				}{
// 					Check: table.outCheck,
// 				}

// 				jsonCheck, err := json.Marshal(testCheck)
// 				if err != nil {
// 					t.Error(err)
// 				}

// 				response := &http.Response{Body: io.NopCloser(bytes.NewReader(jsonCheck))}

// 				if req.URL.String() == table.url {
// 					if req.Method == table.method {
// 						response = &http.Response{
// 							StatusCode: http.StatusOK,
// 							Body:       io.NopCloser(bytes.NewReader((jsonData))),
// 						}
// 					}
// 				}

// 				return response, nil
// 			})

// 			svc := service.NewService(
// 				mock,
// 				dbHostTest,
// 				portTest,
// 				tokenHostTest,
// 				portTest,
// 				secretTest,
// 			)

// 			resultUser, resultErr = svc.Profile(table.inToken)

// 			if !table.isError {
// 				assert.Nil(t, resultErr)
// 			} else {
// 				assert.ErrorContains(t, resultErr, errorResponse)
// 			}

// 			assert.Equal(t, table.outUser, resultUser)
// 		})
// 	}
// }

// func TestDeleteAccount(t *testing.T) {
// 	for index, table := range []struct {
// 		inToken  string
// 		outCheck bool

// 		isError bool
// 		url     string
// 		method  string
// 	}{
// 		{tokenTest, true, false, "http://token:8080/check", http.MethodPost},
// 		{tokenTest, true, true, "http://token:8080/check", http.MethodPost},
// 		{tokenTest, false, true, "http://token:8080/check", http.MethodPost},
// 		{tokenTest, true, true, "http://token:8080/extract", http.MethodPost},
// 		{tokenTest, true, true, "http://db:8080/user", http.MethodDelete},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultErr error
// 			var errorResponse string

// 			if table.isError {
// 				errorResponse = errWebServer.Error()
// 			}

// 			testResp := struct {
// 				User     dbapp.User `json:"user"`
// 				ID       int        `json:"id"`
// 				Username string     `json:"username"`
// 				Email    string     `json:"email"`
// 				Check    bool       `json:"check"`
// 				Err      string     `json:"err"`
// 			}{
// 				User: dbapp.User{
// 					ID:       idTest,
// 					Username: usernameTest,
// 					Password: passwordTest,
// 					Email:    emailTest,
// 				},
// 				ID:       idTest,
// 				Username: usernameTest,
// 				Email:    emailTest,
// 				Check:    table.outCheck,
// 				Err:      errorResponse,
// 			}

// 			if !table.outCheck {
// 				testResp.Err = ""
// 				errorResponse = service.ErrTokenNotValid.Error()
// 			}

// 			jsonData, err := json.Marshal(testResp)
// 			if err != nil {
// 				t.Error(err)
// 			}

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				testCheck := struct {
// 					Check bool `json:"check"`
// 				}{
// 					Check: table.outCheck,
// 				}

// 				jsonCheck, err := json.Marshal(testCheck)
// 				if err != nil {
// 					t.Error(err)
// 				}

// 				response := &http.Response{Body: io.NopCloser(bytes.NewReader(jsonCheck))}

// 				if req.URL.String() == table.url {
// 					if req.Method == table.method {
// 						response = &http.Response{
// 							StatusCode: http.StatusOK,
// 							Body:       io.NopCloser(bytes.NewReader(jsonData)),
// 						}
// 					}
// 				}

// 				return response, nil
// 			})

// 			svc := service.NewService(
// 				mock,
// 				dbHostTest,
// 				portTest,
// 				tokenHostTest,
// 				portTest,
// 				secretTest,
// 			)

// 			resultErr = svc.DeleteAccount(table.inToken)

// 			if !table.isError {
// 				assert.Nil(t, resultErr)
// 			} else {
// 				assert.ErrorContains(t, resultErr, errorResponse)
// 			}
// 		})
// 	}
// }
