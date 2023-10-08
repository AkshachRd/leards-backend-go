package httputils

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
