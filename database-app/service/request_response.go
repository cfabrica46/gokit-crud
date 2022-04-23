package service

// GetAllUsersRequest ...
type GetAllUsersRequest struct{}

// GetAllUsersResponse ...
type GetAllUsersResponse struct {
	Err   string `json:"err,omitempty"`
	Users []User `json:"users"`
}

// GetUserByIDRequest ...
type GetUserByIDRequest struct {
	ID int `json:"id"`
}

// GetUserByIDResponse ...
type GetUserByIDResponse struct {
	Err  string `json:"err,omitempty"`
	User User   `json:"user"`
}

// GetUserByUsernameAndPasswordRequest ...
type GetUserByUsernameAndPasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetUserByUsernameAndPasswordResponse ...
type GetUserByUsernameAndPasswordResponse struct {
	Err  string `json:"err,omitempty"`
	User User   `json:"user"`
}

// GetIDByUsernameRequest ...
type GetIDByUsernameRequest struct {
	Username string `json:"username"`
}

// GetIDByUsernameResponse ...
type GetIDByUsernameResponse struct {
	Err string `json:"err,omitempty"`
	ID  int    `json:"id"`
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
	Err          string `json:"err,omitempty"`
	RowsAffected int    `json:"rowsAffected"`
}
