package models

import "github.com/google/uuid"

type StorageHasTag struct {
	StorageID   uuid.UUID `gorm:"primaryKey"`
	StorageType string    `gorm:"primaryKey"`
	TagID       uuid.UUID `gorm:"primaryKey"`
	Tag         Tag
}
