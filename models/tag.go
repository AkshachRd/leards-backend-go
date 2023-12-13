package models

import "strings"

type Tag struct {
	Model
	ID   uint64 `gorm:"primary_key"`
	Name string `gorm:"size:255; unique; not null"`
}

func FetchOrCreateTagByName(name string) (*Tag, error) {
	var tag Tag

	err := db.FirstOrCreate(&tag, Tag{Name: strings.ToLower(name)}).Error
	if err != nil {
		return &Tag{}, err
	}

	return &tag, nil
}

func FetchTagsByStorageIdAndStorageType(storageId string, storageType string) (*[]Tag, error) {
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

func FetchTagNamesByStorageIdAndStorageType(storageId string, storageType string) (*[]string, error) {
	var tagNames []string

	err := db.Table("storage_has_tag").
		Select("tag.name").
		Joins("LEFT JOIN tag ON tag.id_tag = storage_has_tag.tag_id").
		Where("storage_has_tag.storage_id = ? AND storage_has_tag.storage_type = ?", storageId, storageType).
		Pluck("name", &tagNames).
		Error
	if err != nil {
		return &[]string{}, err
	}

	return &tagNames, nil
}
