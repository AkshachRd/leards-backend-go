package models

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

func CreateCards(cards []Card) error {
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
		if err := tx.Create(&card).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
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
