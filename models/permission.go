package models

type Permission struct {
	Base
	StorageID      string
	StorageType    string
	User           User
	PermissionType PermissionType
}
