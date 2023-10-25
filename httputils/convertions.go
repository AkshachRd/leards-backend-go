package httputils

import (
	"github.com/AkshachRd/leards-backend-go/models"
)

func ConvertFolder(folder *models.Folder) *Folder {
	var path []Path
	path = append(path, Path{Name: folder.Name, Id: folder.ID})
	for parentFolder := folder.ParentFolder; parentFolder != nil; {
		path = append(
			[]Path{{Name: parentFolder.Name, Id: parentFolder.ID}},
			path...,
		)
		parentFolder = parentFolder.ParentFolder
	}

	var content []Content
	for _, contentFolder := range folder.Folders {
		content = append(content, Content{Id: contentFolder.ID, Name: contentFolder.Name, Type: "folder"})
	}
	for _, contentDeck := range folder.Decks {
		content = append(content, Content{Id: contentDeck.ID, Name: contentDeck.Name, Type: "deck"})
	}

	return &Folder{FolderId: folder.ID, Name: folder.Name, Path: path, Content: content}
}

func ConvertDeck(deck *models.Deck) *Deck {
	var content []Card
	for _, card := range deck.Cards {
		content = append(content, Card{CardId: card.ID, FrontSide: card.FrontSide, BackSide: card.BackSide})
	}

	return &Deck{DeckId: deck.ID, Name: deck.Name, Content: content}
}

func ConvertUserSettings(userSettings *[]models.UserSetting) *Settings {
	settings := make(Settings)

	for _, userSetting := range *userSettings {
		settings[userSetting.SettingName] = userSetting.SettingValue
	}

	return &settings
}
