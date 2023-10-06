package models

import "gorm.io/gorm"

type Deck struct {
	Base
	Name           string `gorm:"size:255; not null"`
	ParentFolderID string `gorm:"size:36; not null"`
	ParentFolder   Folder
	AccessTypeID   uint8 `gorm:"not null"`
	AccessType     AccessType
	Cards          []Card
	Permissions    []Permission    `gorm:"polymorphic:Storage;polymorphicValue:deck"`
	StorageHasTags []StorageHasTag `gorm:"polymorphic:Storage;polymorphicValue:deck"`
}

func NewDeck(db *gorm.DB, name string, accessType Access, parentFolderId string) (*Deck, error) {
	var accType AccessType
	err := db.First(&accType, "type = ?", accessType).Error
	if err != nil {
		return &Deck{}, nil
	}

	deck := Deck{Name: name, AccessTypeID: accType.ID, ParentFolderID: parentFolderId}
	err = db.Create(&deck).Error
	if err != nil {
		return &Deck{}, nil
	}

	return &deck, nil
}

func FetchDeckById(db *gorm.DB, id string, preloadCards bool) (*Deck, error) {
	var deck Deck

	query := db
	if preloadCards {
		query = query.Preload("Cards")
	}

	err := query.First(&deck, "id_deck = ?", id).Error
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
