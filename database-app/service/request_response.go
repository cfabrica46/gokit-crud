package service

// GetAllUsersRequest ...
type GetAllUsersRequest struct{}

// GetAllUsersResponse ...
type GetAllUsersResponse struct {
	Users []User `json:"users"`
	Err   string `json:"err,omitempty"`
}

// GetUserByIDRequest ...
type GetUserByIDRequest struct {
	ID int `json:"id"`
}

// GetUserByIDResponse ...
type GetUserByIDResponse struct {
	User User   `json:"user"`
	Err  string `json:"err,omitempty"`
}

// GetUserByUsernameAndPasswordRequest ...
type GetUserByUsernameAndPasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetUserByUsernameAndPasswordResponse ...
type GetUserByUsernameAndPasswordResponse struct {
	User User   `json:"user"`
	Err  string `json:"err,omitempty"`
}

// GetIDByUsernameRequest ...
type GetIDByUsernameRequest struct {
	Username string `json:"username"`
}

// GetIDByUsernameResponse ...
type GetIDByUsernameResponse struct {
	ID  int    `json:"id"`
	Err string `json:"err,omitempty"`
}

// InsertUserRequest ...
type InsertUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// InsertUserResponse ...
type InsertUserResponse struct {
	Err string `json:"err,omitempty"`
}

// DeleteUserRequest ...
type DeleteUserRequest struct {
	ID int `json:"id"`
}

// DeleteUserResponse ...
type DeleteUserResponse struct {
	RowsAffected int    `json:"rowsAffected"`
	Err          string `json:"err,omitempty"`
}
