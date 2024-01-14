package models

import (
	"math/rand"
	"time"
)

type Card struct {
	Model
	FrontSide string `gorm:"type:text; not null"`
	BackSide  string `gorm:"type:text; not null"`
	DeckID    string `gorm:"size:36; not null"`
	Deck      Deck   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func UpdateCards(cards []Card) error {
	if len(cards) == 0 {
		return nil
	}

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

func CreateCards(rowCards *[]Card) (*[]Card, error) {
	cards := *rowCards
	if len(cards) == 0 {
		return rowCards, nil
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for i, _ := range cards {
		if err := tx.Create(&cards[i]).Error; err != nil {
			tx.Rollback()
			return rowCards, err
		}
	}

	return &cards, tx.Commit().Error
}

func DeleteCards(cards []Card) error {
	if len(cards) == 0 {
		return nil
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Delete(&cards).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func FetchCardsByDeckId(deckId string) (*[]Card, error) {
	var cards []Card

	err := db.Find(&cards, "deck_id = ?", deckId).Error
	if err != nil {
		return nil, err
	}

	return &cards, nil
}

func FetchCardsByFolderId(id string) (*[]Card, error) {
	var cards []Card

	err := db.Table("folder").
		Select("card.front_side, card.back_side, card.id_card").
		Joins("left join deck on deck.parent_folder_id = folder.id_folder").
		Joins("left join card on card.deck_id = deck.id_deck").
		Where("card.id_card is not null and folder.id_folder = ?", id).
		Find(&cards).Error
	if err != nil {
		return &[]Card{}, err
	}

	var childFolders []Folder
	err = db.Table("folder").
		Select("id_folder, parent_folder_id").
		Where("parent_folder_id = ?", id).
		Find(&childFolders).Error
	if err != nil {
		return &[]Card{}, err
	}

	for _, childFolder := range childFolders {
		childCards, err := FetchCardsByFolderId(childFolder.ID)
		if err != nil {
			return &[]Card{}, err
		}
		cards = append(cards, *childCards...)
	}

	return &cards, nil
}

func FetchRandomCardByStorageIdAndStorageType(storageType, storageId string) (*Card, error) {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	// TODO: write queries for fetching random card
	if storageType == StorageTypeFolder {
		rowCards, err := FetchCardsByFolderId(storageId)
		cards := *rowCards
		if err != nil {
			return &Card{}, err
		}

		elementIndex := r.Intn(len(cards))
		return &cards[elementIndex], nil
	} else if storageType == StorageTypeDeck {
		rowCards, err := FetchCardsByDeckId(storageId)
		cards := *rowCards
		if err != nil {
			return &Card{}, err
		}

		elementIndex := r.Intn(len(cards))
		return &cards[elementIndex], nil
	} else {
		return &Card{}, nil
	}
}
