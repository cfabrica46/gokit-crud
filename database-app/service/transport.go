package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DecodeGetAllUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getAllUsersRequest
	return request, nil
}

func DecodeGetUserByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getUserByIDRequest

	idString := mux.Vars(r)["id"]

	// router doesn't allow a non-integer value to be declared
	id, _ := strconv.Atoi(idString)
	request.ID = id

	return request, nil
}

func DecodeGetUserByUsernameAndPasswordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getUserByUsernameAndPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetIDByUsernameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getIDByUsernameRequest

	username := mux.Vars(r)["username"]
	request.Username = username

	return request, nil
}

func DecodeInsertUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request insertUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
