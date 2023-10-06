package models

import "gorm.io/gorm"

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

func FetchDeckById(db *gorm.DB, id string) (*Deck, error) {
	var deck Deck

	err := db.First(&deck, id).Error
	if err != nil {
		return nil, err
	}

	return &deck, nil
}

func FetchDecksByParentId(db *gorm.DB, parentFolderId string) (*[]Deck, error) {
	var decks []Deck

	err := db.Find(&decks, "parent_folder_id = ?", parentFolderId).Error
	if err != nil {
		return nil, err
	}

	return &decks, nil
}
