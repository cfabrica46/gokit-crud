package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// MakeSignUpEndpoint ...
func MakeSignUpEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(UEP)

		token, err := svc.SignUp(req.Username, req.Password, req.Email)
		if err != nil {
			errMessage = err.Error()
		}

		return TokenErr{token, errMessage}, nil
	}
}

// MakeSignInEndpoint ...
func MakeSignInEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(UP)

		token, err := svc.SignIn(req.Username, req.Password)
		if err != nil {
			errMessage = err.Error()
		}

		return TokenErr{token, errMessage}, nil
	}
}

// MakeLogOutEndpoint ...
func MakeLogOutEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(Token)

		err := svc.LogOut(req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return Err{errMessage}, nil
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

		return UsersErr{users, errMessage}, nil
	}
}

// MakeProfileEndpoint ...
func MakeProfileEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(Token)

		user, err := svc.Profile(req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return UserErr{user, errMessage}, nil
	}
}

// MakeDeleteAccountEndpoint ...
func MakeDeleteAccountEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(Token)

		err := svc.DeleteAccount(req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return Err{errMessage}, nil
	}
}
