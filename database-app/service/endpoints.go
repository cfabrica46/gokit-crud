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

		return GetAllUsersResponse{Users: users, Err: errMessage}, nil
	}
}

// MakeGetUserByIDEndpoint ...
func MakeGetUserByIDEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, ok := request.(GetUserByIDRequest)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type GenerateTokenRequest", ErrRequest)
		}

		user, err := svc.GetUserByID(req.ID)
		if err != nil {
			errMessage = err.Error()
		}

		return GetUserByIDResponse{User: user, Err: errMessage}, nil
	}
}

// MakeGetUserByUsernameAndPasswordEndpoint ...
func MakeGetUserByUsernameAndPasswordEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, ok := request.(GetUserByUsernameAndPasswordRequest)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type GenerateTokenRequest", ErrRequest)
		}

		user, err := svc.GetUserByUsernameAndPassword(req.Username, req.Password)
		if err != nil {
			errMessage = err.Error()
		}

		return GetUserByUsernameAndPasswordResponse{User: user, Err: errMessage}, nil
	}
}

// MakeGetIDByUsernameEndpoint ...
func MakeGetIDByUsernameEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, ok := request.(GetIDByUsernameRequest)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type GenerateTokenRequest", ErrRequest)
		}

		id, err := svc.GetIDByUsername(req.Username)
		if err != nil {
			errMessage = err.Error()
		}

		return GetIDByUsernameResponse{ID: id, Err: errMessage}, nil
	}
}

// MakeInsertUserEndpoint ...
func MakeInsertUserEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, ok := request.(InsertUserRequest)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type GenerateTokenRequest", ErrRequest)
		}

		err := svc.InsertUser(req.Username, req.Password, req.Email)
		if err != nil {
			errMessage = err.Error()
		}

		return InsertUserResponse{errMessage}, nil
	}
}

// MakeDeleteUserEndpoint ...
func MakeDeleteUserEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, ok := request.(DeleteUserRequest)
		if !ok {
			return nil, fmt.Errorf("%w: isn't of type GenerateTokenRequest", ErrRequest)
		}

		rowsAffected, err := svc.DeleteUser(req.ID)
		if err != nil {
			errMessage = err.Error()
		}

		return DeleteUserResponse{RowsAffected: rowsAffected, Err: errMessage}, nil
	}
}
