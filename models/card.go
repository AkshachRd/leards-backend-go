package models

import (
	"gorm.io/gorm"
)

type Card struct {
	Base
	FrontSide string `gorm:"type:text; not null"`
	BackSide  string `gorm:"type:text; not null"`
	DeckID    string `gorm:"size:36; not null"`
	Deck      Deck
}

func UpdateCards(db *gorm.DB, cards []Card) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, card := range cards {
		if err := tx.Model(&Card{}).Where("id_card = ?", card.ID).Updates(map[string]interface{}{
			"front_side": card.FrontSide,
			"back_side":  card.BackSide,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func CreateCards(db *gorm.DB, cards []Card) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, card := range cards {
		if err := tx.Create(&card).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func DeleteCards(db *gorm.DB, cards []Card) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, card := range cards {
		if err := tx.Delete(&card).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func FetchCardsByDeckId(db *gorm.DB, deckId string) (*[]Card, error) {
	var cards []Card

	err := db.Find(&cards, "deck_id = ?", deckId).Error
	if err != nil {
		return nil, err
	}

	return &cards, nil
}
