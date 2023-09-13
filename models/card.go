package models

import (
	"database/sql"
)

type Card struct {
	Base
	FrontSide                string
	BackSide                 string
	NextPracticeTime         sql.NullTime
	ConsecutiveCorrectAnswer uint16
	LastTimeEasy             sql.NullTime
	Deck                     Deck
}
