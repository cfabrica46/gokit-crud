package service

// GenerateToken
type GenerateTokenRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Secret   string `json:"secret"`
}

type GenerateTokenResponse struct {
	Token string `json:"token"`
}

// ExtractData
type ExtractTokenRequest struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

type ExtractTokenResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Err      string `json:"err,omitempty"`
}

// SetToken
type SetTokenRequest struct {
	Token string `json:"token"`
}

type SetTokenResponse struct {
	Err string `json:"err,omitempty"`
}

// DeleteToken
type DeleteTokenRequest struct {
	Token string `json:"token"`
}

type DeleteTokenResponse struct {
	Err string `json:"err,omitempty"`
}

// CheckToken
type CheckTokenRequest struct {
	Token string `json:"token"`
}

type CheckTokenResponse struct {
	Check bool   `json:"check"`
	Err   string `json:"err,omitempty"`
}
