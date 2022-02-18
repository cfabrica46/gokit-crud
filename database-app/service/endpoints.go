package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetAllUsersEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		users, err := svc.GetAllUsers()
		if err != nil {
			return GetAllUsersResponse{users, err.Error()}, nil
		}
		return GetAllUsersResponse{users, ""}, nil
	}
}

func MakeGetUserByIDEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserByIDRequest)
		user, err := svc.GetUserByID(req.ID)
		if err != nil {
			return GetUserByIDResponse{user, err.Error()}, nil
		}
		return GetUserByIDResponse{user, ""}, nil
	}
}

func MakeGetUserByUsernameAndPasswordEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserByUsernameAndPasswordRequest)
		user, err := svc.GetUserByUsernameAndPassword(req.Username, req.Password)
		if err != nil {
			return GetUserByUsernameAndPasswordResponse{user, err.Error()}, nil
		}
		return GetUserByUsernameAndPasswordResponse{user, ""}, nil
	}
}

func MakeGetIDByUsernameEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetIDByUsernameRequest)
		id, err := svc.GetIDByUsername(req.Username)
		if err != nil {
			return GetIDByUsernameResponse{id, err.Error()}, nil
		}
		return GetIDByUsernameResponse{id, ""}, nil
	}
}

func MakeInsertUserEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(InsertUserRequest)
		err := svc.InsertUser(req.Username, req.Password, req.Email)
		if err != nil {
			return InsertUserResponse{err.Error()}, nil
		}
		return InsertUserResponse{""}, nil
	}
}

func MakeDeleteUserEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		rowsAffected, err := svc.DeleteUser(req.Username, req.Password, req.Email)
		if err != nil {
			return DeleteUserResponse{rowsAffected, err.Error()}, nil
		}
		return DeleteUserResponse{rowsAffected, ""}, nil
	}
}
