package httputils

type TokenResponse struct {
	Message   string `json:"message" example:"Successfully"`
	Token     string `json:"token" example:"<token>"`
	TokenType string `json:"token_type" example:"bearer"`
}

type BasicResponse struct {
	Message string `json:"message" example:"Successfully"`
}
