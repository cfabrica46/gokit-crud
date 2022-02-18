package service

// GetAllUsers
type getAllUsersRequest struct {
}

type getAllUsersResponse struct {
	Users []User `json:"users"`
	Err   string `json:"err,omitempty"`
}

// GetUserByID
type getUserByIDRequest struct {
	ID int `json:"id"`
}

type getUserByIDResponse struct {
	User User   `json:"user"`
	Err  string `json:"err,omitempty"`
}

// GetUserByUsernameAndPassword
type getUserByUsernameAndPasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type getUserByUsernameAndPasswordResponse struct {
	User User   `json:"user"`
	Err  string `json:"err,omitempty"`
}

// GetIDByUsername
type getIDByUsernameRequest struct {
	Username string `json:"username"`
}

type getIDByUsernameResponse struct {
	ID  int    `json:"id"`
	Err string `json:"err,omitempty"`
}

// InsertUser
type insertUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type insertUserResponse struct {
	Err string `json:"err,omitempty"`
}

// DeleteUser
type deleteUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type deleteUserResponse struct {
	RowsAffected int    `json:"rowsAffected"`
	Err          string `json:"err,omitempty"`
}
