package models

type StorageHasTag struct {
	StorageID   string `gorm:"primaryKey; size:36;"`
	StorageType string `gorm:"primaryKey; size:255"`
	TagID       uint64 `gorm:"primaryKey"`
	Tag         Tag
}

func FetchOrCreateStorageHasTag(storageId string, storageType string, tagId uint64) (*StorageHasTag, error) {
	var storageHasTag StorageHasTag

	err := db.FirstOrCreate(
		&storageHasTag,
		StorageHasTag{StorageID: storageId, StorageType: storageType, TagID: tagId},
	).Error
	if err != nil {
		return &StorageHasTag{}, err
	}

	return &storageHasTag, nil
}

func DeleteStorageHasTagByStorageIdAndTagId(storageId string, storageType string, tagId uint64) error {
	return db.Delete(StorageHasTag{StorageID: storageId, StorageType: storageType, TagID: tagId}).Error
}
