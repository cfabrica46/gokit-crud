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

var errPetition = errors.New("error to petition get all users")

// MakePetition ...
func MakePetition(client httpClient, url, httpMethod string, bodyStruct interface{},
) (dataResp []byte, err error) {
	var dataReq []byte

	if bodyStruct != nil {
		dataReq, err = json.Marshal(bodyStruct)
		if err != nil {
			return nil, fmt.Errorf("error to make petition: %w", err)
		}
	}

	ctx, ctxCancel := context.WithTimeout(context.TODO(), time.Minute)
	defer ctxCancel()

	req, err := http.NewRequestWithContext(ctx, httpMethod, url, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, fmt.Errorf("error to make petition: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error to make petition: %w", err)
	}
	defer resp.Body.Close()

	dataResp, _ = io.ReadAll(resp.Body)

	return dataResp, nil
}

// PetitionGetAllUsers ...
func PetitionGetAllUsers(client httpClient, url string) (users []dbapp.User, err error) {
	var response dbapp.GetAllUsersResponse

	dataResp, err := MakePetition(client, url, http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("error to petition get all users: %w", err)
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return nil, fmt.Errorf("error to petition get all users: %w", err)
	}

	if response.Err != "" {
		return nil, fmt.Errorf("%w: %s", errPetition, response.Err)
	}

	users = response.Users

	return users, nil
}

// PetitionGetUserByID ...
func PetitionGetUserByID(client httpClient, url string, body dbapp.GetUserByIDRequest,
) (user dbapp.User, err error) {
	var response dbapp.GetUserByIDResponse

	// convert password to Sha256
	// in ENDPOINT
	// password = fmt.Sprintf("%x", sha256.Sum256([]byte(password))).

	dataResp, err := MakePetition(client, url, http.MethodGet, body)
	if err != nil {
		return dbapp.User{}, fmt.Errorf("error to petition get user by ID: %w", err)
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return dbapp.User{}, fmt.Errorf("error to petition get user by ID: %w", err)
	}

	if response.Err != "" {
		return dbapp.User{}, fmt.Errorf("%w: %s", errPetition, response.Err)
	}

	user = response.User

	return user, nil
}

// PetitionGetUserByUsernameAndPassword ...
func PetitionGetUserByUsernameAndPassword(client httpClient, url string,
	body dbapp.GetUserByUsernameAndPasswordRequest,
) (user dbapp.User, err error) {
	var response dbapp.GetUserByUsernameAndPasswordResponse

	// convert password to Sha256
	// password = fmt.Sprintf("%x", sha256.Sum256([]byte(password))).
	dataResp, err := MakePetition(client, url, http.MethodGet, body)
	if err != nil {
		return dbapp.User{}, fmt.Errorf("error to petition get user by username and password: %w", err)
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return dbapp.User{}, fmt.Errorf("error to petition get user by username and password: %w", err)
	}

	if response.Err != "" {
		return dbapp.User{}, fmt.Errorf("%w: %s", errPetition, response.Err)
	}

	user = response.User

	return user, nil
}

// PetitionGetIDByUsername ...
func PetitionGetIDByUsername(client httpClient, url string, body dbapp.GetIDByUsernameRequest,
) (id int, err error) {
	var response dbapp.GetIDByUsernameResponse

	dataResp, err := MakePetition(client, url, http.MethodGet, body)
	if err != nil {
		return 0, fmt.Errorf("error to petition get id by username: %w", err)
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return 0, fmt.Errorf("error to petition get id by username: %w", err)
	}

	if response.Err != "" {
		return 0, fmt.Errorf("%w: %s", errPetition, response.Err)
	}

	id = response.ID

	return id, nil
}

// PetitionInsertUser ...
func PetitionInsertUser(client httpClient, url string, body dbapp.InsertUserRequest) (err error) {
	var response dbapp.InsertUserResponse

	// convert password to Sha256
	// in ENDPOINT
	// password = fmt.Sprintf("%x", sha256.Sum256([]byte(password))).

	dataResp, err := MakePetition(client, url, http.MethodPost, body)
	if err != nil {
		return fmt.Errorf("error to petition insert user: %w", err)
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return fmt.Errorf("error to petition insert user: %w", err)
	}

	if response.Err != "" {
		return fmt.Errorf("%w: %s", errPetition, response.Err)
	}

	return nil
}

// PetitionDeleteUser ...
func PetitionDeleteUser(client httpClient, url string, body dbapp.DeleteUserRequest) (err error) {
	var response dbapp.DeleteUserResponse

	// convert password to Sha256
	// password = fmt.Sprintf("%x", sha256.Sum256([]byte(password))).

	dataResp, err := MakePetition(client, url, http.MethodDelete, body)
	if err != nil {
		return fmt.Errorf("error to petition delete user: %w", err)
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return fmt.Errorf("error to petition delete user: %w", err)
	}

	if response.Err != "" {
		return fmt.Errorf("%w: %s", errPetition, response.Err)
	}

	return nil
}

// PetitionGenerateToken ...
func PetitionGenerateToken(client httpClient, url string, body tokenapp.GenerateTokenRequest,
) (token string, err error) {
	// var response tokenapp.GenerateTokenResponse.
	var response struct {
		Err string
		tokenapp.GenerateTokenResponse
	}

	dataResp, err := MakePetition(client, url, http.MethodPost, body)
	if err != nil {
		return "", fmt.Errorf("error to petition generate token: %w", err)
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return "", fmt.Errorf("error to petition generate token: %w", err)
	}

	if response.Err != "" {
		return "", fmt.Errorf("%w: %s", errPetition, response.Err)
	}

	token = response.Token

	return token, nil
}

// PetitionExtractToken ...
func PetitionExtractToken(client httpClient, url string, body tokenapp.ExtractTokenRequest,
) (id int, username, email string, err error) {
	var response tokenapp.ExtractTokenResponse

	dataResp, err := MakePetition(client, url, http.MethodPost, body)
	if err != nil {
		return 0, "", "", fmt.Errorf("error to petition extract token: %w", err)
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return 0, "", "", fmt.Errorf("error to petition extract token: %w", err)
	}

	if response.Err != "" {
		return 0, "", "", fmt.Errorf("%w: %s", errPetition, response.Err)
	}

	id = response.ID
	username = response.Username
	email = response.Email

	return id, username, email, nil
}

// PetitionSetToken ...
func PetitionSetToken(client httpClient, url string, body tokenapp.SetTokenRequest) (err error) {
	var response tokenapp.SetTokenResponse

	dataResp, err := MakePetition(client, url, http.MethodPost, body)
	if err != nil {
		return fmt.Errorf("error to petition set token: %w", err)
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return fmt.Errorf("error to petition set token: %w", err)
	}

	if response.Err != "" {
		return fmt.Errorf("%w: %s", errPetition, response.Err)
	}

	return nil
}

// PetitionDeleteToken ...
func PetitionDeleteToken(client httpClient, url string, body tokenapp.DeleteTokenRequest,
) (err error) {
	var response tokenapp.DeleteTokenResponse

	dataResp, err := MakePetition(client, url, http.MethodDelete, body)
	if err != nil {
		return fmt.Errorf("error to petition delete token: %w", err)
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return fmt.Errorf("error to petition delete token: %w", err)
	}

	if response.Err != "" {
		return fmt.Errorf("%w: %s", errPetition, response.Err)
	}

	return nil
}

// PetitionCheckToken ...
func PetitionCheckToken(client httpClient, url string, body tokenapp.CheckTokenRequest,
) (check bool, err error) {
	var response tokenapp.CheckTokenResponse

	dataResp, err := MakePetition(client, url, http.MethodPost, body)
	if err != nil {
		return false, fmt.Errorf("error to petition check token: %w", err)
	}

	err = json.Unmarshal(dataResp, &response)
	if err != nil {
		return false, fmt.Errorf("error to petition check token: %w", err)
	}

	if response.Err != "" {
		return false, fmt.Errorf("%w: %s", errPetition, response.Err)
	}

	check = response.Check

	return check, nil
}
