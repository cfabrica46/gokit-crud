package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// MakeGenerateTokenEndpoint ...
func MakeGenerateTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, _ := request.(GenerateTokenRequest)
		token := svc.GenerateToken(req.ID, req.Username, req.Email, []byte(req.Secret))

		return GenerateTokenResponse{token}, nil
	}
}

// MakeExtractTokenEndpoint ...
func MakeExtractTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(ExtractTokenRequest)

		id, username, email, err := svc.ExtractToken(req.Token, []byte(req.Secret))
		if err != nil {
			errMessage = err.Error()
		}

		return ExtractTokenResponse{id, username, email, errMessage}, nil
	}
}

// MakeSetTokenEndpoint ...
func MakeSetTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(SetTokenRequest)

		err := svc.SetToken(req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return SetTokenResponse{errMessage}, nil
	}
}

// MakeDeleteTokenEndpoint ...
func MakeDeleteTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(DeleteTokenRequest)

		err := svc.DeleteToken(req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return DeleteTokenResponse{errMessage}, nil
	}
}

// MakeCheckTokenEndpoint ...
func MakeCheckTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(CheckTokenRequest)

		check, err := svc.CheckToken(req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return CheckTokenResponse{check, errMessage}, nil
	}
}
