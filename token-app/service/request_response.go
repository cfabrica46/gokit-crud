package service

/* GenerateToken(int, string, string) (string, error)
ExtractData(string) (int, string, string, error)
SetToken(string) error
DeleteToken(string) error
CheckToken(string) (bool, error) */

// GenerateToken
type generateTokenRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Secret   string `json:"secret"`
}

type generateTokenResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

// ExtractData
type extractDataRequest struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

type extractDataResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Err      string `json:"err,omitempty"`
}

// SetToken
type setTokenRequest struct {
	Token string `json:"token"`
}

type setTokenResponse struct {
	Err string `json:"err,omitempty"`
}

// DeleteToken
type deleteTokenRequest struct {
	Token string `json:"token"`
}

type deleteTokenResponse struct {
	Err string `json:"err,omitempty"`
}

// CheckToken
type checkTokenRequest struct {
	Token string `json:"token"`
}

type checkTokenResponse struct {
	Check bool   `json:"check"`
	Err   string `json:"err,omitempty"`
}
