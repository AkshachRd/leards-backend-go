package services

import (
	"fmt"
	"time"

	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/open-spaced-repetition/go-fsrs"
)

const MAX_REPETITION_TIME_OFFSET = time.Minute * 15

type RepetitionStats map[fsrs.State]int

type RepetitionService struct {
	fsrsPrameters fsrs.Parameters
}

func NewRepetitionService() *RepetitionService {
	return &RepetitionService{fsrsPrameters: fsrs.DefaultParam()}
}

func (r *RepetitionService) ReviewCard(userId string, cardId string, reviewAnswer string) error {
	var card fsrs.Card

	repetition, err := models.FetchRepetitionByUserIDAndCardID(userId, cardId)
	if err != nil {
		return err
	}

	card = r.repetitionToFSRSCard(repetition)
	now := time.Now()

	rating, err := r.reviewAnswerToFSRSRating(reviewAnswer)
	if err != nil {
		return err
	}

	schedulingInfo := r.fsrsPrameters.Repeat(card, now)[rating]
	cardAfterRepetition := schedulingInfo.Card
	err = models.UpdateRepetition(r.fsrsCardToRepetition(cardAfterRepetition, userId, cardId))
	if err != nil {
		return err
	}

	return nil
}

func (r *RepetitionService) CreateRepetition(userId, cardId string) error {
	card := fsrs.NewCard()
	err := models.UpdateRepetition(r.fsrsCardToRepetition(card, userId, cardId))
	if err != nil {
		return err
	}

	return nil
}

func (r *RepetitionService) FetchNextRepetitionCard(userId, storageId, storageType string) (*models.Card, error) {
	repetition, err := models.FetchNextRepetitionByUserIdAndStorageIdAndStorageType(userId, storageId, storageType)
	if err != nil {
		return &models.Card{}, err
	}

	if !repetition.Due.Before(time.Now().Add(MAX_REPETITION_TIME_OFFSET)) {
		return &models.Card{}, nil
	}

	return &repetition.Card, nil
}

func (r *RepetitionService) GetStorageStats(userId, storageId, storageType string) (*RepetitionStats, error) {
	rowRepetitionStats, err := models.FetchRepetitionStats(storageId, storageType)
	if err != nil {
		return &RepetitionStats{}, err
	}

	repetitionStats := make(RepetitionStats)
	for _, stat := range *rowRepetitionStats {
		repetitionStats[fsrs.State(stat.State)] = stat.StateCount
	}

	return &repetitionStats, nil
}

func (r *RepetitionService) repetitionToFSRSCard(repetition *models.Repetition) fsrs.Card {
	return fsrs.Card{
		Due:           repetition.Due,
		Stability:     repetition.Stability,
		Difficulty:    repetition.Difficulty,
		ElapsedDays:   repetition.ElapsedDays,
		ScheduledDays: repetition.ScheduledDays,
		Reps:          repetition.Reps,
		Lapses:        repetition.Lapses,
		State:         fsrs.State(repetition.State),
		LastReview:    repetition.LastReview,
	}
}

func (r *RepetitionService) fsrsCardToRepetition(card fsrs.Card, userId, cardId string) *models.Repetition {
	return &models.Repetition{
		UserID:        userId,
		CardID:        cardId,
		Due:           card.Due,
		Stability:     card.Stability,
		Difficulty:    card.Difficulty,
		ElapsedDays:   card.ElapsedDays,
		ScheduledDays: card.ScheduledDays,
		Reps:          card.Reps,
		Lapses:        card.Lapses,
		State:         models.RepetitionState(card.State),
		LastReview:    card.LastReview,
	}
}

func (r *RepetitionService) reviewAnswerToFSRSRating(reviewAnswer string) (fsrs.Rating, error) {
	switch reviewAnswer {
	case "repeat":
		return fsrs.Again, nil
	case "hard":
		return fsrs.Hard, nil
	case "good":
		return fsrs.Good, nil
	case "easy":
		return fsrs.Easy, nil
	default:
		return fsrs.Good, fmt.Errorf("unknown review answer: %s", reviewAnswer)
	}
}
