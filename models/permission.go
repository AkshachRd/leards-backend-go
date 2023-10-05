package models

type Permission struct {
	Base
	StorageID        string `gorm:"size:36; not null"`
	StorageType      string `gorm:"size:255; not null"`
	UserID           string `gorm:"size:36; not null"`
	User             User
	PermissionTypeID uint8 `gorm:"not null"`
	PermissionType   PermissionType
}
