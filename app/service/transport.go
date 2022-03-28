package service

import (
	"context"
	"encoding/json"
	"net/http"
)

// DecodeSignUpRequest ...
func DecodeSignUpRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request SignUpRequest

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	req = request

	return
}

// DecodeSignInRequest ...
func DecodeSignInRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request SignInRequest

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	req = request

	return
}

// DecodeLogOutRequest ...
func DecodeLogOutRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request LogOutRequest

	request.Token = r.Header.Get("Authorization")

	req = request

	return
}

// DecodeGetAllUsersRequest ...
func DecodeGetAllUsersRequest(_ context.Context, _ *http.Request) (req interface{}, err error) {
	var request GetAllUsersRequest

	req = request

	return
}

// DecodeProfileRequest ...
func DecodeProfileRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request ProfileRequest

	request.Token = r.Header.Get("Authorization")

	req = request

	return
}

// DecodeDeleteAccountRequest ...
func DecodeDeleteAccountRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request DeleteAccountRequest

	request.Token = r.Header.Get("Authorization")

	req = request

	return
}

// EncodeResponse ...
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if err = json.NewEncoder(w).Encode(response); err != nil {
		return
	}

	return
}
