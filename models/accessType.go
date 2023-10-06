package models

type Access string

const (
	Public  Access = "public"
	Private Access = "private"
)

type AccessType struct {
	ID   uint8  `gorm:"primaryKey"`
	Type string `gorm:"size:255; not null; unique;"`
}
