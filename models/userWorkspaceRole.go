package models

type role string

const (
	Owner       role = "owner"
	Participant      = "participant"
)

type UserWorkspaceRole struct {
	ID   uint8 `gorm:"primaryKey"`
	Role role
}
