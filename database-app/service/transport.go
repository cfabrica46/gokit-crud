package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// DecodeGetAllUsersRequest ...
func DecodeGetAllUsersRequest(_ context.Context, _ *http.Request) (any, error) {
	var request GetAllUsersRequest

	return request, nil
}

// DecodeGetUserByIDRequest ...
func DecodeGetUserByIDRequest(_ context.Context, r *http.Request) (any, error) {
	var request GetUserByIDRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request: %w", err)
	}

	return request, nil
}

// DecodeGetUserByUsernameAndPasswordRequest ...
func DecodeGetUserByUsernameAndPasswordRequest(
	_ context.Context,
	r *http.Request,
) (
	any,
	error,
) {
	var request GetUserByUsernameAndPasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request: %w", err)
	}

	return request, nil
}

// DecodeGetIDByUsernameRequest ...
func DecodeGetIDByUsernameRequest(_ context.Context, r *http.Request) (any, error) {
	var request GetIDByUsernameRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request: %w", err)
	}

	return request, nil
}

// DecodeInsertUserRequest ...
func DecodeInsertUserRequest(_ context.Context, r *http.Request) (any, error) {
	var request InsertUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request: %w", err)
	}

	return request, nil
}

// DecodeDeleteUserRequest ...
func DecodeDeleteUserRequest(_ context.Context, r *http.Request) (any, error) {
	var request DeleteUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request: %w", err)
	}

	return request, nil
}

// EncodeResponse ...
func EncodeResponse(_ context.Context, w http.ResponseWriter, response any) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}

	return nil
}
