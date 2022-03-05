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
)

func makePetition(client httpClient, url, httpMethod string, bodyStruct ...interface{}) (dataResp []byte, err error) {
	var dataReq []byte

	if len(bodyStruct) > 0 {
		dataReq, err = json.Marshal(bodyStruct[0])
		if err != nil {
			return
		}
	}

	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(dataReq))
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dataResp, _ = io.ReadAll(resp.Body)
	return
}

func petitionGetUserByUsernameAndPassword(client httpClient, url string, username, password string) (user dbapp.User, err error) {
	var response dbapp.GetUserByUsernameAndPasswordResponse

	dataResp, err := makePetition(client, url, http.MethodGet, dbapp.GetUserByUsernameAndPasswordRequest{Username: username, Password: password})
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	if response.Err != "" {
		err = errors.New(response.Err)
		return
	}
	user = response.User
	return
}

func petitionGetIDByUsername(client httpClient, url string) (id int, err error) {
	var response dbapp.GetIDByUsernameResponse

	dataResp, err := makePetition(client, url, http.MethodGet)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	if response.Err != "" {
		err = errors.New(response.Err)
		return
	}
	id = response.ID
	return
}

func petitionInsertUser(client httpClient, url string, username, password, email string) (err error) {
	var response dbapp.InsertUserResponse

	//convert password to Sha256
	password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	dataResp, err := makePetition(client, url, http.MethodPost, dbapp.InsertUserRequest{Username: username, Password: password, Email: email})
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}
	if response.Err != "" {
		err = errors.New(response.Err)
		return
	}
	return
}

func petitionGenerateToken(client httpClient, url string, id int, username, email, secret string) (token string, err error) {
	var response tokenapp.GenerateTokenResponse

	dataResp, err := makePetition(client, url, http.MethodGet, tokenapp.GenerateTokenRequest{ID: id, Username: username, Email: email, Secret: secret})
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	token = response.Token
	return
}

func petitionSetToken(client httpClient, url string, token string) (err error) {
	var response tokenapp.SetTokenResponse

	dataResp, err := makePetition(client, url, http.MethodPost, tokenapp.SetTokenRequest{Token: token})
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}
	return
}

func petitionDeleteToken(client httpClient, url string, token string) (err error) {
	var response tokenapp.DeleteTokenResponse

	dataResp, err := makePetition(client, url, http.MethodDelete, tokenapp.DeleteTokenRequest{Token: token})
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}
	return
}

func petitionCheckToken(client httpClient, url string, token string) (err error) {
	var response tokenapp.CheckTokenRequest

	dataResp, err := makePetition(client, url, http.MethodPost, tokenapp.CheckTokenRequest{Token: token})
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}
	return
}
