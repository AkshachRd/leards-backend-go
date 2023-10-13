package httputils

type BasicResponse struct {
	Message string `json:"message" example:"Successfully"`
} // @name BasicResponse

type TokenResponse struct {
	BasicResponse
	Token string `json:"token" example:"<token>"`
} // @name TokenResponse

type UserResponse struct {
	BasicResponse
	User User `json:"user"`
} // @name UserResponse

type FolderResponse struct {
	BasicResponse
	Folder Folder `json:"folder"`
} // @name FolderResponse

type DeckResponse struct {
	BasicResponse
	Deck Deck `json:"deck"`
} // @name DeckResponse

type CardsResponse struct {
	BasicResponse
	Cards []Card `json:"cards"`
} // @name CardsResponse
