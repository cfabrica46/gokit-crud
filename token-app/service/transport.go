package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// DecodeGenerateTokenRequest ...
func DecodeGenerateTokenRequest(_ context.Context, r *http.Request) (any, error) {
	var request GenerateTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request: %w", err)
	}

	return request, nil
}

// DecodeExtractTokenRequest ...
func DecodeExtractTokenRequest(_ context.Context, r *http.Request) (any, error) {
	var request ExtractTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request: %w", err)
	}

	return request, nil
}

// DecodeSetTokenRequest ...
func DecodeSetTokenRequest(_ context.Context, r *http.Request) (any, error) {
	var request SetTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request: %w", err)
	}

	return request, nil
}

// DecodeDeleteTokenRequest ...
func DecodeDeleteTokenRequest(_ context.Context, r *http.Request) (any, error) {
	var request DeleteTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request: %w", err)
	}

	return request, nil
}

// DecodeCheckTokenRequest ...
func DecodeCheckTokenRequest(_ context.Context, r *http.Request) (any, error) {
	var request CheckTokenRequest
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
