package models

import "github.com/google/uuid"

type Folder struct {
	Base
	Name           string
	AccessTypeId   uint8
	AccessType     AccessType
	ParentFolderID *uuid.UUID
	ParentFolder   *Folder
	Permissions    []Permission `gorm:"polymorphic:Storage;polymorphicValue:folder"`
}
