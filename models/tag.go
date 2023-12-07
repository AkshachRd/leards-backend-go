package models

type Tag struct {
	Model
	ID   uint64 `gorm:"primary_key"`
	Name string `gorm:"size:255; unique; not null"`
}

func FetchOrCreateTagByName(name string) (*Tag, error) {
	var tag Tag

	err := db.FirstOrCreate(&tag, Tag{Name: name}).Error
	if err != nil {
		return &Tag{}, err
	}

	return &tag, nil
}

func FetchTagsByStorageIdByStorageId(storageId string, storageType string) (*[]Tag, error) {
	var tags []Tag

	err := db.Table("tag").
		Joins("LEFT JOIN storage_has_tag ON tag.id_tag = storage_has_tag.tag_id").
		Where("storage_has_tag.storage_id = ? AND storage_has_tag.storage_type = ?", storageId, storageType).
		Find(&tags).Error
	if err != nil {
		return &[]Tag{}, err
	}

	return &tags, nil
}
