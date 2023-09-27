package httputils

import "github.com/AkshachRd/leards-backend-go/libs"

type BasicResponse struct {
	Message string `json:"message" example:"Successfully"`
}

type TokenResponse struct {
	BasicResponse
	Token string `json:"token" example:"<token>"`
}

type UserResponse struct {
	TokenResponse
	UserId string `json:"userId" example:"53f4cf69-9da6-49e4-8651-450b74abdf9e"`
}

type LoginResponse struct {
	BasicResponse
	user  libs.User
	Token string `json:"token" example:"<token>"`
}
