package models

import (
	"gorm.io/gorm"
)

type Deck struct {
	Model
	Name             string            `gorm:"size:255; not null"`
	ParentFolderID   string            `gorm:"size:36; not null"`
	ParentFolder     Folder            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AccessType       uint8             `gorm:"not null"`
	Cards            []Card            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Permissions      []Permission      `gorm:"polymorphic:Storage;polymorphicValue:deck"`
	FavoriteStorages []FavoriteStorage `gorm:"polymorphic:Storage;polymorphicValue:deck"`
	StorageHasTags   []StorageHasTag   `gorm:"polymorphic:Storage;polymorphicValue:deck"`
}

func getDeckPreloadQuery(index int) string {
	return []string{"Cards", "AccessType"}[index]
}

func NewDeck(name string, accessType uint8, parentFolderId string) (*Deck, error) {
	deck := Deck{Name: name, AccessType: accessType, ParentFolderID: parentFolderId}
	err := db.Create(&deck).Error
	if err != nil {
		return &Deck{}, err
	}

	return &deck, nil
}

func UpdateDeckById(id string, name string) (*Deck, error) {
	deck, err := FetchDeckById(id, false, true)
	if err != nil {
		return &Deck{}, err
	}

	deck.Name = name

	err = db.Save(deck).Error
	if err != nil {
		return &Deck{}, err
	}

	deck, err = FetchDeckById(id, true, true)
	if err != nil {
		return &Deck{}, err
	}

	return deck, nil
}

func DeleteDeckById(id string) error {
	deck, err := FetchDeckById(id, true)
	if err != nil {
		return err
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = deck.Delete(db); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (d *Deck) Delete(db *gorm.DB) error {
	cards := d.Cards
	if cards == nil {
		if err := db.Find(&cards, "deck_id = ?", d.ID).Error; err != nil {
			return err
		}
	}

	if len(cards) != 0 {
		err := db.Delete(&cards).Error
		if err != nil {
			return err
		}
	}

	err := db.Delete(d).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchDeckById
//
// Preload args: Cards, AccessType
func FetchDeckById(id string, preloadArgs ...bool) (*Deck, error) {
	var deck Deck

	query := db
	for i, arg := range preloadArgs {
		if arg {
			query = query.Preload(getDeckPreloadQuery(i))
		}
	}

	err := query.First(&deck, "id_deck = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &deck, nil
}

func FetchDecksByParentId(parentFolderId string) (*[]Deck, error) {
	var decks []Deck

	err := db.Find(&decks, "parent_folder_id = ?", parentFolderId).Error
	if err != nil {
		return nil, err
	}

	return &decks, nil
}

func FetchPublicDecksWithPagination(page int, pageSize int) (*[]Deck, error) {
	var decks []Deck

	err := db.Scopes(Paginate(page, pageSize)).Find(&decks, "access_type = ?", AccessTypePublic).Error
	if err != nil {
		return nil, err
	}

	return &decks, nil
}
