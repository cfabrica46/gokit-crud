package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// MakeGenerateTokenEndpoint ...
func MakeGenerateTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GenerateTokenRequest)
		token := svc.GenerateToken(req.ID, req.Username, req.Email, []byte(req.Secret))
		return GenerateTokenResponse{token}, nil
	}
}

// MakeExtractTokenEndpoint ...
func MakeExtractTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(ExtractTokenRequest)
		id, username, email, err := svc.ExtractToken(req.Token, []byte(req.Secret))
		if err != nil {
			return ExtractTokenResponse{id, username, email, err.Error()}, nil
		}
		return ExtractTokenResponse{id, username, email, ""}, nil
	}
}

// MakeSetTokenEndpoint ...
func MakeSetTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(SetTokenRequest)
		err := svc.SetToken(req.Token)
		if err != nil {
			return SetTokenResponse{err.Error()}, nil
		}
		return SetTokenResponse{""}, nil
	}
}

// MakeDeleteTokenEndpoint ...
func MakeDeleteTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteTokenRequest)
		err := svc.DeleteToken(req.Token)
		if err != nil {
			return DeleteTokenResponse{err.Error()}, nil
		}
		return DeleteTokenResponse{""}, nil
	}
}

// MakeCheckTokenEndpoint ...
func MakeCheckTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CheckTokenRequest)
		check, err := svc.CheckToken(req.Token)
		if err != nil {
			return CheckTokenResponse{check, err.Error()}, nil
		}
		return CheckTokenResponse{check, ""}, nil
	}
}
