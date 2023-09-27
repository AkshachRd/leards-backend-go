package models

import "github.com/google/uuid"

type Permission struct {
	Base
	StorageID        uuid.UUID
	StorageType      string
	UserID           uuid.UUID
	User             User
	PermissionTypeID uint8
	PermissionType   PermissionType
}
