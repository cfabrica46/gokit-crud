package service

import (
	"context"
	"encoding/json"
	"net/http"
)

// DecodeSignUpRequest ...
func DecodeSignUpRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeSignInRequest ...
func DecodeSignInRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeLogOutRequest ...
func DecodeLogOutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request LogOutRequest
	request.Token = r.Header.Get("Authorization")
	return request, nil
}

// DecodeGetAllUsersRequest ...
func DecodeGetAllUsersRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var request GetAllUsersRequest
	return request, nil
}

// DecodeProfileRequest ...
func DecodeProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ProfileRequest
	request.Token = r.Header.Get("Authorization")
	return request, nil
}

// DecodeDeleteAccountRequest ...
func DecodeDeleteAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DeleteAccountRequest
	request.Token = r.Header.Get("Authorization")
	return request, nil
}

// EncodeResponse ...
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
