package httputils

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

type Path struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Content struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Folder struct {
	FolderId string    `json:"folderId"`
	Name     string    `json:"name"`
	Path     []Path    `json:"path"`
	Content  []Content `json:"content"`
}

type FolderResponse struct {
	BasicResponse
	Folder Folder `json:"folder"`
}

type Card struct {
	CardId    string `json:"cardId"`
	FrontSide string `json:"frontSide"`
	BackSide  string `json:"backSide"`
}

type Deck struct {
	DeckId  string `json:"deckId"`
	Name    string `json:"name"`
	Content []Card `json:"content"`
}

type DeckResponse struct {
	BasicResponse
	Deck Deck
}
