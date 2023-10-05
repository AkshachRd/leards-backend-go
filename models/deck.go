package models

type Deck struct {
	Base
	Name           string `gorm:"size:255; not null"`
	ParentFolderId string `gorm:"size:36; not null"`
	ParentFolder   Folder
	AccessTypeID   uint8 `gorm:"not null"`
	AccessType     AccessType
	Cards          []Card
	Permissions    []Permission    `gorm:"polymorphic:Storage;polymorphicValue:deck"`
	StorageHasTags []StorageHasTag `gorm:"polymorphic:Storage;polymorphicValue:deck"`
}
