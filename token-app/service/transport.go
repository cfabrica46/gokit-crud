package service

import (
	"context"
	"encoding/json"
	"net/http"
)

//DecodeGenerateTokenRequest ...
func DecodeGenerateTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GenerateTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//DecodeExtractTokenRequest ...
func DecodeExtractTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ExtractTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//DecodeSetTokenRequest ...
func DecodeSetTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request SetTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//DecodeDeleteTokenRequest ...
func DecodeDeleteTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DeleteTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//DecodeCheckTokenRequest ...
func DecodeCheckTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CheckTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//EncodeResponse ...
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
