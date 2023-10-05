package models

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        string `gorm:"primary_key; unique; type:char(36)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.NewString()
	return
}
