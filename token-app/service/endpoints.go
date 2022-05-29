package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

var ErrRequest = errors.New("error to request")

// MakeGenerateTokenEndpoint ...
func MakeGenerateTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		req, ok := request.(IDUsernameEmailSecretRequest)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type GenerateTokenRequest", ErrRequest)
		}

		token := svc.GenerateToken(req.ID, req.Username, req.Email, []byte(req.Secret))

		return Token{Token: token}, nil
	}
}

// MakeExtractTokenEndpoint ...
func MakeExtractTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, ok := request.(TokenSecretRequest)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type GenerateTokenRequest", ErrRequest)
		}

		id, username, email, err := svc.ExtractToken(req.Token, []byte(req.Secret))
		if err != nil {
			errMessage = err.Error()
		}

		return IDUsernameEmailErrResponse{ID: id, Username: username, Email: email, Err: errMessage}, nil
	}
}

// MakeManageTokenEndpoint ...
func MakeManageTokenEndpoint(svc serviceInterface, st State) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, ok := request.(Token)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type Token", ErrRequest)
		}

		err := svc.ManageToken(st, req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return ErrorResponse{Err: errMessage}, nil
	}
}

// MakeCheckTokenEndpoint ...
func MakeCheckTokenEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, ok := request.(Token)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type Token", ErrRequest)
		}

		check, err := svc.CheckToken(req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return CheckErrResponse{Check: check, Err: errMessage}, nil
	}
}
