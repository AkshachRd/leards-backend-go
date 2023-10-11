package httputils

import "github.com/AkshachRd/leards-backend-go/models"

type Card struct {
	CardId    string `json:"cardId"`
	FrontSide string `json:"frontSide"`
	BackSide  string `json:"backSide"`
} // @name Card

type Deck struct {
	DeckId  string `json:"deckId"`
	Name    string `json:"name"`
	Content []Card `json:"content"`
} // @name Deck

type Path struct {
	Name string `json:"name"`
	Id   string `json:"id"`
} // @name Path

type Content struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Name string `json:"name"`
} // @name Content

type Folder struct {
	FolderId string    `json:"folderId"`
	Name     string    `json:"name"`
	Path     []Path    `json:"path"`
	Content  []Content `json:"content"`
} // @name Folder

type Settings map[models.Setting]string // @name Settings

type User struct {
	UserId       string   `json:"userId"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	AuthToken    string   `json:"authToken"`
	ProfileIcon  string   `json:"profileIcon,omitempty"`
	RootFolderId string   `json:"rootFolderId"`
	Settings     Settings `json:"settings"`
} // @name User