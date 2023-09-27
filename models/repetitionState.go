package models

type state string

const (
	New        state = "new"
	Learning         = "learning"
	Review           = "review"
	Relearning       = "relearning"
)

type RepetitionState struct {
	ID    uint8 `gorm:"primaryKey"`
	State state
}
