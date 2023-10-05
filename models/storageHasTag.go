package models

type StorageHasTag struct {
	StorageID   string `gorm:"primaryKey; size:36;"`
	StorageType string `gorm:"primaryKey; size:255"`
	TagID       string `gorm:"primaryKey; size:36;"`
	Tag         Tag
}
