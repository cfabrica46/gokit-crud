package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type myDoFunc func(req *http.Request) (*http.Response, error)

type mockClient struct {
	doFunc myDoFunc
}

func getMockClient(d myDoFunc) *mockClient {
	return &mockClient{d}
}

func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	return m.doFunc(req)
}

func TestSignUp(t *testing.T) {
	for i, tt := range []struct {
		inUsername, inPassword, inEmail, inSecret string
		outErr                                    string
	}{
		{"cesar", "01234", "cesar@email.com", "secret", ""},
		{"cesar", "01234", "cesar@email.com", "secret", "Error from web server"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var resultErr string

			testResp := struct {
				ID    int    `json:"id"`
				Token string `json:"token"`
				Err   string `json:"err"`
			}{
				ID:    1,
				Token: "token",
				Err:   tt.outErr,
			}

			jsonData, err := json.Marshal(testResp)
			if err != nil {
				t.Error(err)
			}

			mock := getMockClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(jsonData)))}, nil
			})

			svc := GetService("localhost", "8080", "localhost", "8080", mock)

			_, err = svc.SignUp(tt.inUsername, tt.inPassword, tt.inEmail, tt.inSecret)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, tt.outErr, resultErr)
		})
	}
}
