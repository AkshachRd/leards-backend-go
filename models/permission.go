package models

import "github.com/google/uuid"

type Permission struct {
	Base
	StorageID        uuid.UUID
	StorageType      string
	UserId           uuid.UUID
	User             User
	PermissionTypeID uint8
	PermissionType   PermissionType
}
