package service

// GetAllUsers
type GetAllUsersRequest struct {
}

type GetAllUsersResponse struct {
	Users []User `json:"users"`
	Err   string `json:"err,omitempty"`
}

// GetUserByID
type GetUserByIDRequest struct {
	ID int `json:"id"`
}

type GetUserByIDResponse struct {
	User User   `json:"user"`
	Err  string `json:"err,omitempty"`
}

// GetUserByUsernameAndPassword
type GetUserByUsernameAndPasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUserByUsernameAndPasswordResponse struct {
	User User   `json:"user"`
	Err  string `json:"err,omitempty"`
}

// GetIDByUsername
type GetIDByUsernameRequest struct {
	Username string `json:"username"`
}

type GetIDByUsernameResponse struct {
	ID  int    `json:"id"`
	Err string `json:"err,omitempty"`
}

// InsertUser
type InsertUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type InsertUserResponse struct {
	Err string `json:"err,omitempty"`
}

// DeleteUser
type DeleteUserRequest struct {
	ID int `json:"id"`
}

type DeleteUserResponse struct {
	RowsAffected int    `json:"rowsAffected"`
	Err          string `json:"err,omitempty"`
}
