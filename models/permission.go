package models

import "gorm.io/gorm"

const (
	PermissionTypeOwner = iota
)

type Permission struct {
	Model
	StorageID      string `gorm:"size:36; not null"`
	StorageType    string `gorm:"size:255; not null"`
	UserID         string `gorm:"size:36; not null"`
	User           User
	PermissionType uint8 `gorm:"not null"`
}

func NewPermission(db *gorm.DB, storageId string, storageType string, userId string, permissionType uint8) (*Permission, error) {
	permission := Permission{StorageID: storageId, StorageType: storageType, UserID: userId, PermissionType: permissionType}

	err := db.Create(&permission).Error
	if err != nil {
		return &Permission{}, err
	}

	return &permission, nil
}
