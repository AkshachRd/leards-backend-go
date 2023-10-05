package models

type PermissionType struct {
	ID   uint8  `gorm:"primaryKey"`
	Type string `gorm:"size:255; not null; unique"`
}
