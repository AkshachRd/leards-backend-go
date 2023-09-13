package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	base.ID = id
	return
}
