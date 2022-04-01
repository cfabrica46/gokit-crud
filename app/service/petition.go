package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	dbapp "github.com/cfabrica46/gokit-crud/database-app/service"
	tokenapp "github.com/cfabrica46/gokit-crud/token-app/service"
)

var ErrPrefix = errors.New("error")

func MakePetition(client httpClient, url, httpMethod string, bodyStruct interface{},
) (dataResp []byte, err error) {
	var dataReq []byte

	if bodyStruct != nil {
		dataReq, err = json.Marshal(bodyStruct)
		if err != nil {
			return
		}
	}

	ctx, ctxCancel := context.WithTimeout(context.TODO(), time.Minute)
	defer ctxCancel()

	req, err := http.NewRequestWithContext(ctx, httpMethod, url, bytes.NewBuffer(dataReq))
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

func PetitionGetAllUsers(client httpClient, url string) (users []dbapp.User, err error) {
	var response dbapp.GetAllUsersResponse

	dataResp, err := MakePetition(client, url, http.MethodGet, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	if response.Err != "" {
		err = fmt.Errorf("%w: %s", ErrPrefix, response.Err)

		return
	}

	users = response.Users

	return
}

func PetitionGetUserByID(client httpClient, url string, body dbapp.GetUserByIDRequest,
) (user dbapp.User, err error) {
	var response dbapp.GetUserByIDResponse

	// convert password to Sha256
	// in ENDPOINT
	// password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	dataResp, err := MakePetition(client, url, http.MethodGet, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	if response.Err != "" {
		// err = errors.New(response.Err)
		err = fmt.Errorf("%w: %s", ErrPrefix, response.Err)

		return
	}

	user = response.User

	return
}

func PetitionGetUserByUsernameAndPassword(client httpClient, url string,
	body dbapp.GetUserByUsernameAndPasswordRequest,
) (user dbapp.User, err error) {
	var response dbapp.GetUserByUsernameAndPasswordResponse

	// convert password to Sha256
	// password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	dataResp, err := MakePetition(client, url, http.MethodGet, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	if response.Err != "" {
		err = fmt.Errorf("%w: %s", ErrPrefix, response.Err)
		// err = errors.New(response.Err)

		return
	}

	user = response.User

	return
}

func PetitionGetIDByUsername(client httpClient, url string, body dbapp.GetIDByUsernameRequest,
) (id int, err error) {
	var response dbapp.GetIDByUsernameResponse

	dataResp, err := MakePetition(client, url, http.MethodGet, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	if response.Err != "" {
		err = fmt.Errorf("%w: %s", ErrPrefix, response.Err)

		return
	}

	id = response.ID

	return
}

func PetitionInsertUser(client httpClient, url string, body dbapp.InsertUserRequest) (err error) {
	var response dbapp.InsertUserResponse

	// convert password to Sha256
	// in ENDPOINT
	// password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	dataResp, err := MakePetition(client, url, http.MethodPost, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	if response.Err != "" {
		err = fmt.Errorf("%w: %s", ErrPrefix, response.Err)
		// err = errors.New(response.Err)

		return
	}

	return
}

func PetitionDeleteUser(client httpClient, url string, body dbapp.DeleteUserRequest) (err error) {
	var response dbapp.DeleteUserResponse

	// convert password to Sha256
	// password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	dataResp, err := MakePetition(client, url, http.MethodDelete, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	if response.Err != "" {
		err = fmt.Errorf("%w: %s", ErrPrefix, response.Err)

		return
	}

	return
}

func PetitionGenerateToken(client httpClient, url string, body tokenapp.GenerateTokenRequest,
) (token string, err error) {
	// var response tokenapp.GenerateTokenResponse
	var response struct {
		Err string
		tokenapp.GenerateTokenResponse
	}

	dataResp, err := MakePetition(client, url, http.MethodPost, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	if response.Err != "" {
		err = fmt.Errorf("%w: %s", ErrPrefix, response.Err)

		return
	}

	token = response.Token

	return
}

func PetitionExtractToken(client httpClient, url string, body tokenapp.ExtractTokenRequest,
) (id int, username, email string, err error) {
	var response tokenapp.ExtractTokenResponse

	dataResp, err := MakePetition(client, url, http.MethodPost, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	if response.Err != "" {
		err = fmt.Errorf("%w: %s", ErrPrefix, response.Err)

		return
	}

	id = response.ID
	username = response.Username
	email = response.Email

	return
}

func PetitionSetToken(client httpClient, url string, body tokenapp.SetTokenRequest) (err error) {
	var response tokenapp.SetTokenResponse

	dataResp, err := MakePetition(client, url, http.MethodPost, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	if response.Err != "" {
		err = fmt.Errorf("%w: %s", ErrPrefix, response.Err)

		return
	}

	return
}

func PetitionDeleteToken(client httpClient, url string, body tokenapp.DeleteTokenRequest,
) (err error) {
	var response tokenapp.DeleteTokenResponse

	dataResp, err := MakePetition(client, url, http.MethodDelete, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	if response.Err != "" {
		err = fmt.Errorf("%w: %s", ErrPrefix, response.Err)

		return
	}

	return
}

func PetitionCheckToken(client httpClient, url string, body tokenapp.CheckTokenRequest,
) (check bool, err error) {
	var response tokenapp.CheckTokenResponse

	dataResp, err := MakePetition(client, url, http.MethodPost, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return
	}

	if response.Err != "" {
		err = fmt.Errorf("%w: %s", ErrPrefix, response.Err)

		return
	}

	check = response.Check

	return
}
