package service

// GenerateTokenRequest ...
type GenerateTokenRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Secret   string `json:"secret"`
}

// GenerateTokenResponse ...
type GenerateTokenResponse struct {
	Token string `json:"token"`
}

// ExtractTokenRequest ...
type ExtractTokenRequest struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

// ExtractTokenResponse ...
type ExtractTokenResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Err      string `json:"err,omitempty"`
}

// SetTokenRequest ...
type SetTokenRequest struct {
	Token string `json:"token"`
}

// SetTokenResponse ...
type SetTokenResponse struct {
	Err string `json:"err,omitempty"`
}

// DeleteTokenRequest ...
type DeleteTokenRequest struct {
	Token string `json:"token"`
}

// DeleteTokenResponse ...
type DeleteTokenResponse struct {
	Err string `json:"err,omitempty"`
}

// CheckTokenRequest ...
type CheckTokenRequest struct {
	Token string `json:"token"`
}

// CheckTokenResponse ...
type CheckTokenResponse struct {
	Check bool   `json:"check"`
	Err   string `json:"err,omitempty"`
}
