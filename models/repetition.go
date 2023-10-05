package models

import (
	"time"
)

type Repetition struct {
	UserID        string `gorm:"primaryKey; size:36"`
	User          User
	CardID        string `gorm:"primaryKey; size:36"`
	Card          Card
	Due           time.Time `json:"due" gorm:"not null"`
	Stability     float64   `json:"stability" gorm:"not null"`
	Difficulty    float64   `json:"difficulty" gorm:"not null"`
	ElapsedDays   uint64    `json:"elapsedDays" gorm:"not null"`
	ScheduledDays uint64    `json:"scheduledDays" gorm:"not null"`
	Reps          uint64    `json:"reps" gorm:"not null"`
	Lapses        uint64    `json:"lapses" gorm:"not null"`
	StateID       uint8     `json:"stateID" gorm:"not null"`
	State         RepetitionState
	LastReview    time.Time `json:"lastReview" gorm:"not null"`
}
