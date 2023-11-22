package models

type FavoriteStorage struct {
	UserID      string `gorm:"primaryKey; size:36;"`
	User        User
	StorageID   string `gorm:"primaryKey; size:36; not null"`
	StorageType string `gorm:"primaryKey; size:255; not null"`
}

func FetchFavoriteStoragesByUserId(id string) (*[]FavoriteStorage, error) {
	var favoriteStorage []FavoriteStorage

	err := db.Find(&favoriteStorage, "id_user = ?", id).Error
	if err != nil {
		return &[]FavoriteStorage{}, err
	}

	return &favoriteStorage, nil
}

func FetchFavoriteDecksByUserId(id string) (*[]Deck, error) {
	var decks []Deck

	err := db.Joins("left join permission on deck.id_deck = permission.storage_id and permission.storage_type = 'deck'").
		Joins("left join favorite_storage on favorite_storage.user_id = permission.user_id and favorite_storage.storage_type = 'deck'").
		Where("favorite_storage.user_id = ?", id).
		Find(&decks).
		Error

	if err != nil {
		return &[]Deck{}, err
	}

	return &decks, nil
}

func NewFavoriteStorage(userId string, storageId string, storageType string) (*FavoriteStorage, error) {
	favoriteStorage := FavoriteStorage{UserID: userId, StorageID: storageId, StorageType: storageType}

	err := db.Create(&favoriteStorage).Error
	if err != nil {
		return &FavoriteStorage{}, err
	}

	return &favoriteStorage, nil
}
