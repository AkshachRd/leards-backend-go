package models

import "gorm.io/gorm"

type Card struct {
	Base
	FrontSide string `gorm:"type:text; not null"`
	BackSide  string `gorm:"type:text; not null"`
	DeckID    string `gorm:"size:36; not null"`
	Deck      Deck
}

func FetchCardsByDeckId(db *gorm.DB, deckId string) (*[]Card, error) {
	var cards []Card

	err := db.Find(&cards, "deck_id = ?", deckId).Error
	if err != nil {
		return nil, err
	}

	return &cards, nil
}
