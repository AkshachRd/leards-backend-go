package models

import "github.com/google/uuid"

type Deck struct {
	Base
	Name           string
	ParentFolderId uuid.UUID
	ParentFolder   Folder
	AccessTypeID   uint8
	AccessType     AccessType
	Cards          []Card
	Permissions    []Permission    `gorm:"polymorphic:Storage;polymorphicValue:deck"`
	StorageHasTags []StorageHasTag `gorm:"polymorphic:Storage;polymorphicValue:deck"`
}
