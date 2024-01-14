package models

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func getStorageQuery() *gorm.DB {
	storageQuery := db.Raw("? UNION ALL ?",
		db.Table("folder").
			Select("folder.id_folder AS id, created_at, updated_at, deleted_at, name, access_type, parent_folder_id, ? AS type", StorageTypeFolder),
		db.Table("deck").
			Select("deck.id_deck AS id, created_at, updated_at, deleted_at, name, access_type, parent_folder_id, ? AS type", StorageTypeDeck),
	)
	return db.Table("(?) AS storage", storageQuery)
}

type SearchResult struct {
	ID              string         `gorm:"column:id"`
	Name            string         `gorm:"column:name"`
	Rating          uint           `gorm:"column:rating"`
	Type            string         `gorm:"column:type"`
	ProfileIconPath sql.NullString `gorm:"column:profile_icon_path"`
	AuthorName      string         `gorm:"column:author_name"`
}

func SearchByNameWithPagination(name string, sortType string, orderBy string, page int, pageSize int) (*[]SearchResult, error) {
	var results []SearchResult

	if sortType == "name" {
		sortType = "storage.name"
	}

	err := getStorageQuery().
		Select("storage.id, storage.name, storage.type, COUNT(favorite_storage.storage_id) AS rating, user.profile_icon_path, user.name AS author_name").
		Joins("LEFT JOIN favorite_storage ON favorite_storage.storage_id = storage.id AND favorite_storage.storage_type = storage.type").
		Joins("LEFT JOIN permission ON permission.storage_id = storage.id AND permission.storage_type = storage.type").
		Joins("LEFT JOIN user ON permission.user_id = user.id_user").
		Where("storage.name LIKE ? AND storage.access_type = ?",
			fmt.Sprintf("%%%s%%", name), AccessTypePublic).
		Group("storage.id").
		Order(fmt.Sprintf("%s %s", sortType, orderBy)).
		Scopes(Paginate(page, pageSize)).
		Scan(&results).
		Error
	if err != nil {
		return &[]SearchResult{}, err
	}

	return &results, nil
}

func SearchByTagsWithPagination(tags []string, sortType string, orderBy string, page int, pageSize int) (*[]SearchResult, error) {
	var results []SearchResult

	tagNames := make([]string, 0)
	for _, tag := range tags {
		tagNames = append(tagNames, strings.ToLower(tag))
	}

	if sortType == "name" {
		sortType = "storage.name"
	}

	err := getStorageQuery().
		Select("storage.id, storage.name, storage.type, COUNT(favorite_storage.storage_id) AS rating, user.profile_icon_path, user.name AS author_name").
		Joins("LEFT JOIN storage_has_tag ON storage_has_tag.storage_id = storage.id AND storage_has_tag.storage_type = storage.type").
		Joins("LEFT JOIN tag ON tag.id_tag = storage_has_tag.tag_id").
		Joins("LEFT JOIN favorite_storage ON favorite_storage.storage_id = storage.id AND favorite_storage.storage_type = storage.type").
		Joins("LEFT JOIN permission ON permission.storage_id = storage.id AND permission.storage_type = storage.type").
		Joins("LEFT JOIN user ON permission.user_id = user.id_user").
		Where("tag.name IN (?) AND storage.access_type = ?", tagNames, AccessTypePublic).
		Group("storage.id").
		Order(fmt.Sprintf("%s %s", sortType, orderBy)).
		Scopes(Paginate(page, pageSize)).
		Scan(&results).
		Error
	if err != nil {
		return &[]SearchResult{}, err
	}

	return &results, nil
}

func SearchByNameOrTagsWithPagination(name string, tags []string, sortType string, orderBy string, page int, pageSize int) (*[]SearchResult, error) {
	var results []SearchResult

	tagNames := make([]string, 0)
	for _, tag := range tags {
		tagNames = append(tagNames, strings.ToLower(tag))
	}

	if sortType == "name" {
		sortType = "storage.name"
	}

	query := getStorageQuery().
		Select("storage.id, storage.name, storage.type, COUNT(favorite_storage.storage_id) AS rating, user.profile_icon_path, user.name AS author_name").
		Joins("LEFT JOIN storage_has_tag ON storage_has_tag.storage_id = storage.id AND storage_has_tag.storage_type = storage.type").
		Joins("LEFT JOIN tag ON tag.id_tag = storage_has_tag.tag_id").
		Joins("LEFT JOIN favorite_storage ON favorite_storage.storage_id = storage.id AND favorite_storage.storage_type = storage.type").
		Joins("LEFT JOIN permission ON permission.storage_id = storage.id AND permission.storage_type = storage.type").
		Joins("LEFT JOIN user ON permission.user_id = user.id_user")

	if name == "" {
		query = query.Where("tag.name IN (?) AND storage.access_type = ?", tagNames, AccessTypePublic)
	} else if len(tags) == 0 {
		query = query.Where("storage.name LIKE ? AND storage.access_type = ?",
			fmt.Sprintf("%%%s%%", name), AccessTypePublic)
	} else {
		query = query.Where("(tag.name IN (?) OR storage.name LIKE ?) AND storage.access_type = ?",
			tagNames, fmt.Sprintf("%%%s%%", name), AccessTypePublic)
	}

	err := query.
		Group("storage.id").
		Order(fmt.Sprintf("%s %s", sortType, orderBy)).
		Scopes(Paginate(page, pageSize)).
		Scan(&results).
		Error
	if err != nil {
		return &[]SearchResult{}, err
	}

	return &results, nil
}
