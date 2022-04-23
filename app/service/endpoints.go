package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// MakeSignUpEndpoint ...
func MakeSignUpEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(SignUpRequest)

		token, err := svc.SignUp(req.Username, req.Password, req.Email)
		if err != nil {
			errMessage = err.Error()
		}

		return SignUpResponse{Token: token, Err: errMessage}, nil
	}
}

// MakeSignInEndpoint ...
func MakeSignInEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(SignInRequest)

		token, err := svc.SignIn(req.Username, req.Password)
		if err != nil {
			errMessage = err.Error()
		}

		return SignInResponse{Token: token, Err: errMessage}, nil
	}
}

// MakeLogOutEndpoint ...
func MakeLogOutEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(LogOutRequest)

		err := svc.LogOut(req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return LogOutResponse{Err: errMessage}, nil
	}
}

// MakeGetAllUsersEndpoint ...
func MakeGetAllUsersEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		var errMessage string

		users, err := svc.GetAllUsers()
		if err != nil {
			errMessage = err.Error()
		}

		return GetAllUsersResponse{Users: users, Err: errMessage}, nil
	}
}

// MakeProfileEndpoint ...
func MakeProfileEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(ProfileRequest)

		user, err := svc.Profile(req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return ProfileResponse{User: user, Err: errMessage}, nil
	}
}

// MakeDeleteAccountEndpoint ...
func MakeDeleteAccountEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(DeleteAccountRequest)

		err := svc.DeleteAccount(req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return DeleteAccountResponse{Err: errMessage}, nil
	}
}
