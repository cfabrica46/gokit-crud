package service

import (
	"bytes"
	"encoding/json"
	"errors"
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

func petitionGetAllUsers(client httpClient, url string) (user []dbapp.User, err error) {
	var response dbapp.GetAllUsersResponse

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
	user = response.Users
	return
}

func petitionGetUserByUsernameAndPassword(client httpClient, url string, body dbapp.GetUserByUsernameAndPasswordRequest) (user dbapp.User, err error) {
	var response dbapp.GetUserByUsernameAndPasswordResponse

	//convert password to Sha256
	// password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	dataResp, err := makePetition(client, url, http.MethodGet, body)
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

func petitionInsertUser(client httpClient, url string, body dbapp.InsertUserRequest) (err error) {
	var response dbapp.InsertUserResponse

	//convert password to Sha256
	//in ENDPOINT
	// password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	dataResp, err := makePetition(client, url, http.MethodPost, body)
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

func petitionDeleteUser(client httpClient, url string, body dbapp.DeleteUserRequest) (err error) {
	var response dbapp.DeleteUserResponse

	//convert password to Sha256
	// password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	dataResp, err := makePetition(client, url, http.MethodDelete, body)
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

func petitionGenerateToken(client httpClient, url string, body tokenapp.GenerateTokenRequest) (token string, err error) {
	var response tokenapp.GenerateTokenResponse

	dataResp, err := makePetition(client, url, http.MethodPost, body)
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

func petitionExtractToken(client httpClient, url string, body tokenapp.ExtractTokenRequest) (id int, username, email string, err error) {
	var response tokenapp.ExtractTokenResponse

	dataResp, err := makePetition(client, url, http.MethodPost, body)
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
	username = response.Username
	email = response.Email
	return
}

func petitionSetToken(client httpClient, url string, body tokenapp.SetTokenRequest) (err error) {
	var response tokenapp.SetTokenResponse

	dataResp, err := makePetition(client, url, http.MethodPost, body)
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

func petitionDeleteToken(client httpClient, url string, body tokenapp.DeleteTokenRequest) (err error) {
	var response tokenapp.DeleteTokenResponse

	dataResp, err := makePetition(client, url, http.MethodDelete, body)
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

func petitionCheckToken(client httpClient, url string, body tokenapp.CheckTokenRequest) (err error) {
	var response tokenapp.CheckTokenRequest

	dataResp, err := makePetition(client, url, http.MethodPost, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}
	return
}
