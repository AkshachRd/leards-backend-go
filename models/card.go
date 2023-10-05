package models

type Card struct {
	Base
	FrontSide string `gorm:"type:text; not null"`
	BackSide  string `gorm:"type:text; not null"`
	DeckID    string `gorm:"size:36; not null"`
	Deck      Deck
}
