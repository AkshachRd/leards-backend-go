package httputils

import (
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/AkshachRd/leards-backend-go/services"
)

func ConvertDecksToContent(decks *[]models.Deck) *[]Content {
	content := make([]Content, 0)

	for _, deck := range *decks {
		content = append(content, Content{Id: deck.ID, Name: deck.Name, Type: "deck"})
	}

	return &content
}

func ConvertFoldersToContent(folders *[]models.Folder) *[]Content {
	content := make([]Content, 0)

	for _, folder := range *folders {
		content = append(content, Content{Id: folder.ID, Name: folder.Name, Type: "folder"})
	}

	return &content
}

func ConvertFolder(folder *models.Folder) *Folder {
	path := []Path{{Name: folder.Name, Id: folder.ID}}
	for parentFolder := folder.ParentFolder; parentFolder != nil; {
		path = append(
			[]Path{{Name: parentFolder.Name, Id: parentFolder.ID}},
			path...,
		)
		parentFolder = parentFolder.ParentFolder
	}

	content := make([]Content, 0)
	content = append(content, *ConvertFoldersToContent(&folder.Folders)...)
	content = append(content, *ConvertDecksToContent(&folder.Decks)...)
	accessType, _ := folder.AccessType.String()

	return &Folder{
		FolderId:   folder.ID,
		Name:       folder.Name,
		Path:       path,
		Content:    content,
		Tags:       *ConvertStorageHasTagsToTags(&folder.StorageHasTags),
		AccessType: accessType,
	}
}

func ConvertDeck(deck *models.Deck) *Deck {
	accessType, _ := deck.AccessType.String()

	return &Deck{
		DeckId:         deck.ID,
		Name:           deck.Name,
		Content:        *ConvertCards(&deck.Cards),
		ParentFolderId: deck.ParentFolderID,
		Tags:           *ConvertStorageHasTagsToTags(&deck.StorageHasTags),
		AccessType:     accessType,
	}
}

func ConvertCard(card *models.Card) *Card {
	return &Card{card.ID, card.FrontSide, card.BackSide}
}

func ConvertCards(cards *[]models.Card) *[]Card {
	convertedCards := make([]Card, 0)

	for _, card := range *cards {
		convertedCards = append(convertedCards, *ConvertCard(&card))
	}

	return &convertedCards
}

func ConvertUserSettings(userSettings *[]models.UserSetting) *Settings {
	settings := make(Settings)

	for _, userSetting := range *userSettings {
		settings[userSetting.SettingName] = userSetting.SettingValue
	}

	return &settings
}

func ConvertProfileIcon(filename string) string {
	// TODO: заменить путь до аватарки на хосте на переменную
	return "/api/v1/accounts/avatars/" + filename
}

func ConvertUser(user *models.User) *User {
	convertedUser := User{
		UserId:       user.ID,
		Name:         user.Name,
		Email:        user.Email,
		AuthToken:    user.AuthToken.String,
		RootFolderId: user.RootFolderID,
		Settings:     *ConvertUserSettings(&user.Settings),
	}

	if user.ProfileIconPath.Valid {
		convertedUser.ProfileIcon = ConvertProfileIcon(user.ProfileIconPath.String)
	}

	return &convertedUser
}

func ConvertFavoriteStoragesContentToContent(favoriteStoragesContent *[]models.FavoriteStoragesContent) *[]Content {
	content := make([]Content, 0)

	for _, favoriteStorageContent := range *favoriteStoragesContent {
		content = append(content, Content{
			Id:   favoriteStorageContent.StorageId,
			Type: favoriteStorageContent.StorageType,
			Name: favoriteStorageContent.Name,
		})
	}

	return &content
}

func ConvertStorageHasTagsToTags(storageHasTags *[]models.StorageHasTag) *[]string {
	tags := make([]string, 0)

	for _, storageHasTag := range *storageHasTags {
		tags = append(tags, storageHasTag.Tag.Name)
	}

	return &tags
}

func ConvertDeckToStorageSettings(deck *models.Deck) *StorageSettings {
	accessType, _ := deck.AccessType.String()

	return &StorageSettings{
		Tags:       *ConvertStorageHasTagsToTags(&deck.StorageHasTags),
		Name:       deck.Name,
		AccessType: accessType,
	}
}

func ConvertFolderToStorageSettings(folder *models.Folder) *StorageSettings {
	accessType, _ := folder.AccessType.String()

	return &StorageSettings{
		Tags:       *ConvertStorageHasTagsToTags(&folder.StorageHasTags),
		Name:       folder.Name,
		AccessType: accessType,
	}
}

func ConvertRepetitionStats(repetitionStats *services.RepetitionStats) *RepetitionStats {
	return &RepetitionStats{
		New:        (*repetitionStats)[0],
		Learning:   (*repetitionStats)[1],
		Review:     (*repetitionStats)[2],
		Relearning: (*repetitionStats)[3],
	}
}

func ConvertSearchResults(searchResults *[]services.SearchResult) []SearchResult {
	remappedSearchResults := make([]SearchResult, 0)

	for _, searchResult := range *searchResults {
		result := SearchResult{
			ID:         searchResult.ID,
			Name:       searchResult.Name,
			Rating:     searchResult.Rating,
			Type:       searchResult.Type,
			AuthorName: searchResult.AuthorName,
			Tags:       searchResult.Tags,
		}

		if len(searchResult.ProfileIconPath) > 0 {
			result.ProfileIcon = ConvertProfileIcon(searchResult.ProfileIconPath)
		}

		remappedSearchResults = append(remappedSearchResults, result)
	}

	return remappedSearchResults
}
