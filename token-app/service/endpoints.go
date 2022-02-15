package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

/* GenerateToken(int, string, string) (string, error)
ExtractData(string) (int, string, string, error)
SetToken(string) error
DeleteToken(string) error
CheckToken(string) (bool, error) */

func MakeGenerateTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(generateTokenRequest)
		token := svc.GenerateToken(req.ID, req.Username, req.Email, []byte(req.Secret))
		return generateTokenResponse{token}, nil
	}
}

func MakeExtractTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(extractTokenRequest)
		id, username, email, err := svc.ExtractToken(req.Token, []byte(req.Secret))
		if err != nil {
			return extractTokenResponse{id, username, email, err.Error()}, nil
		}
		return extractTokenResponse{id, username, email, ""}, nil
	}
}

func MakeSetTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(setTokenRequest)
		err := svc.SetToken(req.Token)
		if err != nil {
			return setTokenResponse{err.Error()}, nil
		}
		return setTokenResponse{""}, nil
	}
}

func MakeDeleteTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteTokenRequest)
		err := svc.DeleteToken(req.Token)
		if err != nil {
			return deleteTokenResponse{err.Error()}, nil
		}
		return deleteTokenResponse{""}, nil
	}
}

func MakeCheckTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(checkTokenRequest)
		check, err := svc.CheckToken(req.Token)
		if err != nil {
			return checkTokenResponse{check, err.Error()}, nil
		}
		return checkTokenResponse{check, ""}, nil
	}
}
