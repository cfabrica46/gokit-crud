package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var errFailedGetHeader = errors.New("failed to get header")

// DecodeSignUpRequest ...
func DecodeSignUpRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request SignUpRequest

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request: %w", err)
	}

	return request, nil
}

// DecodeSignInRequest ...
func DecodeSignInRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request SignInRequest

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request: %w", err)
	}

	return request, nil
}

// DecodeLogOutRequest ...
func DecodeLogOutRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request LogOutRequest

	if r.Header.Get("Authorization") == "" {
		return nil, errFailedGetHeader
	}

	request.Token = r.Header.Get("Authorization")

	return request, nil
}

// DecodeGetAllUsersRequest ...
func DecodeGetAllUsersRequest(_ context.Context, _ *http.Request) (req interface{}, err error) {
	var request GetAllUsersRequest

	return request, nil
}

// DecodeProfileRequest ...
func DecodeProfileRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request ProfileRequest

	if r.Header.Get("Authorization") == "" {
		return nil, errFailedGetHeader
	}

	request.Token = r.Header.Get("Authorization")

	return request, nil
}

// DecodeDeleteAccountRequest ...
func DecodeDeleteAccountRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request DeleteAccountRequest

	if r.Header.Get("Authorization") == "" {
		return nil, errFailedGetHeader
	}

	request.Token = r.Header.Get("Authorization")

	return request, nil
}

// EncodeResponse ...
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if err = json.NewEncoder(w).Encode(response); err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}

	return nil
}
