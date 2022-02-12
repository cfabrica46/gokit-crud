package service

/* func MakeGetAllUsersEndpoint(svc serviceDBInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		users, err := svc.GetAllUsers()
		if err != nil {
			return getAllUsersResponse{users, err.Error()}, nil
		}
		return getAllUsersResponse{users, ""}, nil
	}
} */

/* func MakeGetUserByIDEndpoint(svc serviceDBInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserByIDRequest)
		user, err := svc.GetUserByID(req.ID)
		if err != nil {
			return getUserByIDResponse{user, err.Error()}, nil
		}
		return getUserByIDResponse{user, ""}, nil
	}
} */

/* func MakeGetUserByUsernameAndPasswordEndpoint(svc serviceDBInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserByUsernameAndPasswordRequest)
		user, err := svc.GetUserByUsernameAndPassword(req.Username, req.Password)
		if err != nil {
			return getUserByUsernameAndPasswordResponse{user, err.Error()}, nil
		}
		return getUserByUsernameAndPasswordResponse{user, ""}, nil
	}
} */

/* func MakeGetIDByUsernameEndpoint(svc serviceDBInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getIDByUsernameRequest)
		id, err := svc.GetIDByUsername(req.Username)
		if err != nil {
			return getIDByUsernameResponse{id, err.Error()}, nil
		}
		return getIDByUsernameResponse{id, ""}, nil
	}
} */

/* func MakeInsertUserEndpoint(svc serviceDBInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(insertUserRequest)
		err := svc.InsertUser(req.Username, req.Password, req.Email)
		if err != nil {
			return insertUserResponse{err.Error()}, nil
		}
		return insertUserResponse{""}, nil
	}
} */

/* func MakeDeleteUserEndpoint(svc serviceDBInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteUserRequest)
		rowsAffected, err := svc.DeleteUser(req.Username, req.Password, req.Email)
		if err != nil {
			return deleteUserResponse{rowsAffected, err.Error()}, nil
		}
		return deleteUserResponse{rowsAffected, ""}, nil
	}
} */
