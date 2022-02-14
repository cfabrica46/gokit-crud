package service

import (
	"context"
	"encoding/json"
	"net/http"
)

func DecodeGenerateTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request generateTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeExtractTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request extractTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeSetTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request setTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeDeleteTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeCheckTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request checkTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
