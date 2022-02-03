package service

import (
	"context"
	"encoding/json"
	"net/http"
)

func decodeGetAllUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getAllUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetUserByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getUserByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetUserByUsernameAndPassword(_ context.Context, r *http.Request) (interface{}, error) {
	var request getUserByUsernameAndPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetIDByUsername(_ context.Context, r *http.Request) (interface{}, error) {
	var request getIDByUsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeInsertUser(_ context.Context, r *http.Request) (interface{}, error) {
	var request insertUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeDeleteUserByUsername(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteUserByUsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
