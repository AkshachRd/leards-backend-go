package models

type Folder struct {
	Base
	Name           string
	AccessType     AccessType
	ParentFolderID *string
	ParentFolder   *Folder
	Permissions    []Permission `gorm:"polymorphic:Storage;polymorphicValue:folder"`
}
