package models

type FavoriteStorage struct {
	UserID      string `gorm:"primaryKey; size:36;"`
	User        User
	StorageID   string `gorm:"primaryKey; size:36; not null"`
	StorageType string `gorm:"primaryKey; size:255; not null"`
}

type FavoriteStoragesContent struct {
	StorageId   string
	StorageType string
	Name        string
}

func FetchFavoriteStoragesContentByUserId(userId string) (*[]FavoriteStoragesContent, error) {
	var favoriteStoragesContent []FavoriteStoragesContent

	err := db.Table("favorite_storage").
		Select("favorite_storage.storage_id, favorite_storage.storage_type, CASE favorite_storage.storage_type WHEN 'folder' THEN folder.name ELSE deck.name END as name").
		Joins("LEFT JOIN permission ON favorite_storage.user_id = permission.user_id").
		Joins("LEFT JOIN folder ON favorite_storage.storage_type = ? AND favorite_storage.storage_id = folder.id_folder", StorageTypeFolder).
		Joins("LEFT JOIN deck ON favorite_storage.storage_type = ? AND favorite_storage.storage_id = deck.id_deck", StorageTypeDeck).
		Where("favorite_storage.user_id = ?", userId).
		Group("favorite_storage.storage_id").Scan(&favoriteStoragesContent).Error
	if err != nil {
		return &[]FavoriteStoragesContent{}, err
	}

	return &favoriteStoragesContent, nil
}

func NewFavoriteStorage(userId string, storageId string, storageType string) (*FavoriteStorage, error) {
	favoriteStorage := FavoriteStorage{UserID: userId, StorageID: storageId, StorageType: storageType}

	err := db.Create(&favoriteStorage).Error
	if err != nil {
		return &FavoriteStorage{}, err
	}

	return &favoriteStorage, nil
}

func DeleteFavoriteStorage(userId string, storageId string, storageType string) error {
	favoriteStorage := FavoriteStorage{UserID: userId, StorageID: storageId, StorageType: storageType}
	return db.Delete(&favoriteStorage).Error
}
