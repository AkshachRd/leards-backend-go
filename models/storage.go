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

	combinedQuery := db.Raw("? UNION ALL ?",
		db.Table("folder").
			Select("folder.id_folder AS id, folder.name, folder.access_type, ? AS type", StorageTypeFolder),
		db.Table("deck").
			Select("deck.id_deck AS id, deck.name, deck.access_type, ? AS type", StorageTypeDeck),
	)

	err := db.Table("(?) AS combined", combinedQuery).
		Select("combined.id, combined.name, combined.type, COUNT(favorite_storage.storage_id) AS rating, user.profile_icon_path").
		Joins("LEFT JOIN favorite_storage ON favorite_storage.storage_id = combined.id AND favorite_storage.storage_type = combined.type").
		Joins("LEFT JOIN permission ON permission.storage_id = combined.id AND permission.storage_type = combined.type").
		Joins("LEFT JOIN user ON permission.user_id = user.id_user").
		Where("combined.name LIKE ? AND combined.access_type = ?", fmt.Sprintf("%%%s%%", name), AccessTypePublic).
		Group("combined.id").
		Order(fmt.Sprintf("%s %s", sortType, orderBy)).
		Scopes(Paginate(page, pageSize)).
		Scan(&results).
		Error
	if err != nil {
		return &[]SearchResult{}, err
	}

	return &results, nil
}
