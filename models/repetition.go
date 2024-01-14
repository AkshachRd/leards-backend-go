package models

import (
	"fmt"
	"time"
)

type RepetitionState int8

const (
	New RepetitionState = iota
	Learning
	Review
	Relearning
)

type Repetition struct {
	UserID        string `gorm:"primaryKey; size:36"`
	User          User
	CardID        string `gorm:"primaryKey; size:36"`
	Card          Card
	Due           time.Time       `json:"due" gorm:"not null"`
	Stability     float64         `json:"stability" gorm:"not null"`
	Difficulty    float64         `json:"difficulty" gorm:"not null"`
	ElapsedDays   uint64          `json:"elapsedDays" gorm:"not null"`
	ScheduledDays uint64          `json:"scheduledDays" gorm:"not null"`
	Reps          uint64          `json:"reps" gorm:"not null"`
	Lapses        uint64          `json:"lapses" gorm:"not null"`
	State         RepetitionState `json:"state" gorm:"not null"`
	LastReview    time.Time       `json:"lastReview" gorm:"not null"`
}

func FetchRepetitionByUserIDAndCardID(userID string, cardID string) (*Repetition, error) {
	var repetition Repetition
	err := db.Where("user_id = ? AND card_id = ?", userID, cardID).First(&repetition).Error
	if err != nil {
		return &Repetition{}, err
	}

	return &repetition, nil
}

func UpdateRepetition(repetition *Repetition) error {
	return db.Save(repetition).Error
}

func FetchNextRepetitionCardByUserId(userID string) (*Card, error) {
	var repetition Repetition
	err := db.Joins("Card").Where("user_id = ?", userID).Order("due ASC").First(&repetition).Error
	if err != nil {
		return &Card{}, err
	}
	return &repetition.Card, nil
}

func FetchNextRepetitionByUserIdAndStorageIdAndStorageType(userID string, storageId string, storageType string) (*Repetition, error) {
	var repetition Repetition

	var cards *[]Card
	var err error

	if storageType == StorageTypeFolder {
		cards, err = FetchCardsByFolderId(storageId)
	} else if storageType == StorageTypeDeck {
		cards, err = FetchCardsByDeckId(storageId)
	} else {
		return &Repetition{}, fmt.Errorf("unknown storage type: %s", storageType)
	}

	cardIds := make([]string, 0)
	for _, card := range *cards {
		cardIds = append(cardIds, card.ID)
	}

	err = db.
		Joins("Card").
		Where("repetition.card_id IN (?) AND repetition.user_id = ?", cardIds, userID).
		Order("due ASC").
		First(&repetition).
		Error
	if err != nil {
		return &Repetition{}, err
	}

	return &repetition, nil
}

type RepetitionStats []struct {
	State      int
	StateCount int
}

func FetchRepetitionStats(storageId, storageType string) (*RepetitionStats, error) {
	repetitionStats := make(RepetitionStats, 0)
	var cards *[]Card
	var err error

	if storageType == StorageTypeFolder {
		cards, err = FetchCardsByFolderId(storageId)
	} else if storageType == StorageTypeDeck {
		cards, err = FetchCardsByDeckId(storageId)
	} else {
		return &RepetitionStats{}, fmt.Errorf("unknown storage type: %s", storageType)
	}

	if err != nil {
		return &RepetitionStats{}, err
	}

	cardIds := make([]string, 0)
	for _, card := range *cards {
		cardIds = append(cardIds, card.ID)
	}

	err = db.
		Table("repetition").
		Select("repetition.state, COUNT(repetition.state) as state_count").
		Where("repetition.card_id IN (?)", cardIds).
		Group("repetition.state").
		Scan(&repetitionStats).
		Error
	if err != nil {
		return &RepetitionStats{}, err
	}

	return &repetitionStats, nil
}
