package service

import (
	"context"
	"encoding/json"
	"net/http"
)

func DecodeGetAllUsersRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var request GetAllUsersRequest
	return request, nil
}

func DecodeGetUserByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetUserByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetUserByUsernameAndPasswordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetUserByUsernameAndPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetIDByUsernameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetIDByUsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeInsertUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request InsertUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DeleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
