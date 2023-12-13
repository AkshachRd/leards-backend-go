package models

import (
	"fmt"
)

type SearchResult struct {
	ID              string `gorm:"column:id"`
	Name            string `gorm:"column:name"`
	Rating          uint   `gorm:"column:rating"`
	Type            string `gorm:"column:type"`
	ProfileIconPath string `gorm:"column:profile_icon_path"`
}

func SearchByNameWithPagination(name string, sortType string, orderBy string, page int, pageSize int) (*[]SearchResult, error) {
	var results []SearchResult

	folderQuery := db.Table("folder").
		Select("folder.id_folder AS id, folder.name AS name, COUNT(favorite_storage.storage_id) AS rating, ? AS type, user.profile_icon_path", StorageTypeFolder).
		Joins("LEFT JOIN favorite_storage ON favorite_storage.storage_id = folder.id_folder AND favorite_storage.storage_type = type").
		Joins("LEFT JOIN permission ON permission.storage_id = folder.id_folder AND permission.storage_type = type").
		Joins("LEFT JOIN user ON permission.user_id = user.id_user").
		Where("folder.name LIKE ? AND folder.access_type = ?", fmt.Sprintf("%%%s%%", name), AccessTypePrivate).
		Group("folder.id_folder")
	// TODO: Change to AccessTypePublic
	deckQuery := db.Table("deck").
		Select("deck.id_deck AS id, deck.name AS name, COUNT(favorite_storage.storage_id) AS rating, ? AS type, user.profile_icon_path", StorageTypeDeck).
		Joins("LEFT JOIN favorite_storage ON favorite_storage.storage_id = deck.id_deck AND favorite_storage.storage_type = type").
		Joins("LEFT JOIN permission ON permission.storage_id = deck.id_deck AND permission.storage_type = type").
		Joins("LEFT JOIN user ON permission.user_id = user.id_user").
		Where("deck.name LIKE ? AND deck.access_type = ?", fmt.Sprintf("%%%s%%", name), AccessTypePrivate).
		Group("deck.id_deck")

	offset := (page - 1) * pageSize

	err := db.
		Raw(fmt.Sprintf("SELECT * FROM (?) UNION ALL SELECT * FROM (?) ORDER BY %s %s LIMIT %d OFFSET %d",
			sortType, orderBy, pageSize, offset), folderQuery, deckQuery).
		Scan(&results).Error
	if err != nil {
		return &[]SearchResult{}, err
	}

	return &results, nil
}
