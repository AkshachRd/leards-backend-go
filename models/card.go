package models

import (
	"database/sql"
	"github.com/google/uuid"
)

type Card struct {
	Base
	FrontSide                string
	BackSide                 string
	NextPracticeTime         sql.NullTime
	ConsecutiveCorrectAnswer uint16
	LastTimeEasy             sql.NullTime
	DeckID                   uuid.UUID
	Deck                     Deck
}
