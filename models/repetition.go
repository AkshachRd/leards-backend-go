package models

import (
	"github.com/google/uuid"
	"time"
)

type Repetition struct {
	UserID        uuid.UUID `gorm:"primaryKey"`
	User          User
	CardID        uuid.UUID `gorm:"primaryKey"`
	Card          Card
	Due           time.Time       `json:"due"`
	Stability     float64         `json:"stability"`
	Difficulty    float64         `json:"difficulty"`
	ElapsedDays   uint64          `json:"elapsedDays"`
	ScheduledDays uint64          `json:"scheduledDays"`
	Reps          uint64          `json:"reps"`
	Lapses        uint64          `json:"lapses"`
	StateID       uint8           `json:"stateID"`
	State         RepetitionState `json:"state"`
	LastReview    time.Time       `json:"lastReview"`
}
