package models

type Folder struct {
	Base
	Name           string `gorm:"size:255; not null"`
	AccessTypeId   uint8  `gorm:"not null"`
	AccessType     AccessType
	ParentFolderID *string `gorm:"size:36"`
	ParentFolder   *Folder
	Permissions    []Permission    `gorm:"polymorphic:Storage;polymorphicValue:folder"`
	StorageHasTags []StorageHasTag `gorm:"polymorphic:Storage;polymorphicValue:folder"`
}
