package httputils

type BasicResponse struct {
	Message string `json:"message" example:"Successfully"`
} // @name BasicResponse

type TokenResponse struct {
	BasicResponse
	Token string `json:"token" example:"<token>"`
} // @name TokenResponse

type UserResponse struct {
	TokenResponse
	UserId string `json:"userId" example:"53f4cf69-9da6-49e4-8651-450b74abdf9e"`
} // @name UserResponse

type FolderResponse struct {
	BasicResponse
	Folder Folder `json:"folder"`
} // @name FolderResponse

type DeckResponse struct {
	BasicResponse
	Deck Deck
} // @name DeckResponse
