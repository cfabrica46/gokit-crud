package service

/* import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// MakeSignUpEndpoint ...
func MakeSignUpEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, _ := request.(UsernamePasswordEmailRequest)

		token, err := svc.SignUp(req.Username, req.Password, req.Email)
		if err != nil {
			errMessage = err.Error()
		}

		return TokenErrorResponse{Token: token, Err: errMessage}, nil
	}
}

// MakeSignInEndpoint ...
func MakeSignInEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, _ := request.(UsernamePasswordRequest)

		token, err := svc.SignIn(req.Username, req.Password)
		if err != nil {
			errMessage = err.Error()
		}

		return TokenErrorResponse{Token: token, Err: errMessage}, nil
	}
}

// MakeLogOutEndpoint ...
func MakeLogOutEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, _ := request.(TokenRequest)

		err := svc.LogOut(req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return ErrorResponse{Err: errMessage}, nil
	}
}

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

// MakeProfileEndpoint ...
func MakeProfileEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, _ := request.(TokenRequest)

		user, err := svc.Profile(req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return UserErrorResponse{User: user, Err: errMessage}, nil
	}
}

// MakeDeleteAccountEndpoint ...
func MakeDeleteAccountEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, _ := request.(TokenRequest)

		err := svc.DeleteAccount(req.Token)
		if err != nil {
			errMessage = err.Error()
		}

		return ErrorResponse{Err: errMessage}, nil
	}
} */
