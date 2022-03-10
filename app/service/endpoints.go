package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// MakeSignUpEndpoint ...
func MakeSignUpEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(SignUpRequest)
		token, err := svc.SignUp(req.Username, req.Password, req.Email)
		if err != nil {
			return SignUpResponse{token, err.Error()}, nil
		}
		return SignUpResponse{token, ""}, nil
	}
}

// MakeSignInEndpoint ...
func MakeSignInEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(SignInRequest)
		token, err := svc.SignIn(req.Username, req.Password)
		if err != nil {
			return SignInResponse{token, err.Error()}, nil
		}
		return SignInResponse{token, ""}, nil
	}
}

// MakeLogOutEndpoint ...
func MakeLogOutEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(LogOutRequest)
		err := svc.LogOut(req.Token)
		if err != nil {
			return LogOutResponse{err.Error()}, nil
		}
		return LogOutResponse{""}, nil
	}
}

// MakeGetAllUsersEndpoint ...
func MakeGetAllUsersEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		// req := request.(GetAllUsersRequest)
		users, err := svc.GetAllUsers()
		if err != nil {
			return GetAllUsersResponse{users, err.Error()}, nil
		}
		return GetAllUsersResponse{users, ""}, nil
	}
}

// MakeProfileEndpoint ...
func MakeProfileEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(ProfileRequest)
		user, err := svc.Profile(req.Token)
		if err != nil {
			return ProfileResponse{user, err.Error()}, nil
		}
		return ProfileResponse{user, ""}, nil
	}
}

// MakeDeleteAccountEndpoint ...
func MakeDeleteAccountEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAccountRequest)
		err := svc.DeleteAccount(req.Token)
		if err != nil {
			return DeleteAccountResponse{err.Error()}, nil
		}
		return DeleteAccountResponse{""}, nil
	}
}
