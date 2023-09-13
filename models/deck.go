package models

type Deck struct {
	Base
	Name         string
	ParentFolder Folder
	AccessType   AccessType
	Cards        []Card
	Permissions  []Permission `gorm:"polymorphic:Storage;polymorphicValue:deck"`
}
