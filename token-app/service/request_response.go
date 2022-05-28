package service

// GenerateTokenRequest ...
type GenerateTokenRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Secret   string `json:"secret"`
	ID       int    `json:"id"`
}

// ExtractTokenRequest ...
type ExtractTokenRequest struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

// Token ...
type Token struct {
	Token string `json:"token"`
}

// ExtractTokenResponse ...
type ExtractTokenResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Err      string `json:"err,omitempty"`
	ID       int    `json:"id"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Err string `json:"err,omitempty"`
}

// CheckTokenResponse ...
type CheckTokenResponse struct {
	Err   string `json:"err,omitempty"`
	Check bool   `json:"check"`
}
