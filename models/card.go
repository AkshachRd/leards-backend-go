package models

import (
	"github.com/google/uuid"
)

type Card struct {
	Base
	FrontSide string
	BackSide  string
	DeckID    uuid.UUID
	Deck      Deck
}
