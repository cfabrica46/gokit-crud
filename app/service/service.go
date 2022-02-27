package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	tokenapp "github.com/cfabrica46/gokit-crud/token-app/service"
	_ "github.com/lib/pq"
)

type serviceInterface interface {
	SignUp(string, string, string) (string, error)
	SignIn(string, string) (string, error)
}

type service struct {
	dbHost, dbPort, tokenHost, tokenPort string
}

func GetService(dbHost, dbPort, tokenHost, tokenPort string) *service {
	return &service{dbHost, dbPort, tokenHost, tokenPort}
}

func (s service) SignUp(username, password, email string) (token string, err error) {
	var insertResponse dbapp.InsertUserResponse
	var getIDResponse dbapp.GetIDByUsernameResponse
	var generateTokenResponse tokenapp.GenerateTokenResponse

	//convert password to Sha256
	password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	// insert user
	url := fmt.Sprintf("%s:%s/user", s.dbHost, s.dbPort)

	dataResp, err := makePetition(url, http.MethodPost, dbapp.InsertUserRequest{Username: username, Password: password, Email: email})
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &insertResponse)
	if err != nil {
		return
	}
	if insertResponse.Err != "" {
		err = errors.New(insertResponse.Err)
		return
	}

	// get id by username
	url = fmt.Sprintf("%s:%s/id/%s", s.dbHost, s.dbPort, username)

	dataResp, err = makePetition(url, http.MethodGet)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &getIDResponse)
	if err != nil {
		return
	}
	if getIDResponse.Err != "" {
		err = errors.New(insertResponse.Err)
		return
	}
	id := getIDResponse.ID

	// generate token
	url = fmt.Sprintf("%s:%s/generate", s.tokenHost, s.tokenPort)

	dataResp, err = makePetition(url, http.MethodPost, tokenapp.GenerateTokenRequest{ID: id, Username: username, Email: email, Secret: "secret"})
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &generateTokenResponse)
	if err != nil {
		return
	}
	token = generateTokenResponse.Token
	return
}

func makePetition(url, httpMethod string, bodyStruct ...interface{}) (dataResp []byte, err error) {
	dataReq, err := json.Marshal(bodyStruct[0])
	if err != nil {
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(dataReq))
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dataResp, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

/* func (serviceInterface) SignIn(username, password string) (token string, err error) {
	url := fmt.Sprintf("%s:%s/user/username_password", s.dbHost, s.dbPort)
	return
} */
