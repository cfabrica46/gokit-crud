package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

var ErrRequest = errors.New("error to request")

// MakeGetAllUsersEndpoint ...
func MakeGetAllUsersEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, _ any) (any, error) {
		var errMessage string

		users, err := svc.GetAllUsers()
		if err != nil {
			errMessage = err.Error()
		}

		return UsersErrorResponse{Users: users, Err: errMessage}, nil
	}
}

// MakeGetUserByIDEndpoint ...
func MakeGetUserByIDEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, ok := request.(IDRequest)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type GenerateTokenRequest", ErrRequest)
		}

		user, err := svc.GetUserByID(req.ID)
		if err != nil {
			errMessage = err.Error()
		}

		return UserErrorResponse{User: user, Err: errMessage}, nil
	}
}

// MakeGetUserByUsernameAndPasswordEndpoint ...
func MakeGetUserByUsernameAndPasswordEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, ok := request.(UsernamePasswordRequest)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type GenerateTokenRequest", ErrRequest)
		}

		user, err := svc.GetUserByUsernameAndPassword(req.Username, req.Password)
		if err != nil {
			errMessage = err.Error()
		}

		return UserErrorResponse{User: user, Err: errMessage}, nil
	}
}

// MakeGetIDByUsernameEndpoint ...
func MakeGetIDByUsernameEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, ok := request.(UsernameRequest)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type GenerateTokenRequest", ErrRequest)
		}

		id, err := svc.GetIDByUsername(req.Username)
		if err != nil {
			errMessage = err.Error()
		}

		return IDErrorResponse{ID: id, Err: errMessage}, nil
	}
}

// MakeInsertUserEndpoint ...
func MakeInsertUserEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, ok := request.(UsernamePasswordEmailRequest)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type GenerateTokenRequest", ErrRequest)
		}

		err := svc.InsertUser(req.Username, req.Password, req.Email)
		if err != nil {
			errMessage = err.Error()
		}

		return ErrorResponse{errMessage}, nil
	}
}

// MakeDeleteUserEndpoint ...
func MakeDeleteUserEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, ok := request.(IDRequest)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type GenerateTokenRequest", ErrRequest)
		}

		rowsAffected, err := svc.DeleteUser(req.ID)
		if err != nil {
			errMessage = err.Error()
		}

		return RowsErrorResponse{RowsAffected: rowsAffected, Err: errMessage}, nil
	}
}
