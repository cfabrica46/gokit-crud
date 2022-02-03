package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetAllUsersEndpoint(svc serviceDBInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		// req := request.(getAllUsersRequest)
		users, err := svc.GetAllUsers()
		if err != nil {
			return getAllUsersResponse{users, err.Error()}, nil
		}
		return getAllUsersResponse{users, ""}, nil
	}
}

func makeGetUserByIDEndpoint(svc serviceDBInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserByIDRequest)
		user, err := svc.GetUserByID(req.ID)
		if err != nil {
			return getUserByIDResponse{user, err.Error()}, nil
		}
		return getUserByIDResponse{user, ""}, nil
	}
}

func makeGetUserByUsernameAndPasswordEndpoint(svc serviceDBInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserByUsernameAndPasswordRequest)
		user, err := svc.GetUserByUsernameAndPassword(req.Username, req.Password)
		if err != nil {
			return getUserByUsernameAndPasswordResponse{user, err.Error()}, nil
		}
		return getUserByUsernameAndPasswordResponse{user, ""}, nil
	}
}

func makeGetIDByUsernameEndpoint(svc serviceDBInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getIDByUsernameRequest)
		id, err := svc.GetIDByUsername(req.Username)
		if err != nil {
			return getIDByUsernameResponse{id, err.Error()}, nil
		}
		return getIDByUsernameResponse{id, ""}, nil
	}
}

func makeInsertUserEndpoint(svc serviceDBInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(insertUserRequest)
		err := svc.InsertUser(req.Username, req.Password, req.Email)
		if err != nil {
			return insertUserResponse{err.Error()}, nil
		}
		return insertUserResponse{""}, nil
	}
}

func makeDeleteUserByUsernameEndpoint(svc serviceDBInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteUserByUsernameRequest)
		rowsAffected, err := svc.DeleteUserByUsername(req.Username)
		if err != nil {
			return deleteUserByUsernameResponse{rowsAffected, err.Error()}, nil
		}
		return deleteUserByUsernameResponse{rowsAffected, ""}, nil
	}
}
