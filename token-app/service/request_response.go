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

/*
// GenerateTokenRequest ...
type GenerateTokenRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Secret   string `json:"secret"`
	ID       int    `json:"id"`
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
	Username string `json:"username"`
	Email    string `json:"email"`
	Err      string `json:"err,omitempty"`
	ID       int    `json:"id"`
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
	Err   string `json:"err,omitempty"`
	Check bool   `json:"check"`
}
*/
