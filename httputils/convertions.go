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
	for _, contentDeck := range folder.Decks {
		content = append(content, Content{Id: contentDeck.ID, Name: contentDeck.Name, Type: "deck"})
	}
	for _, contentFolder := range folder.Folders {
		content = append(content, Content{Id: contentFolder.ID, Name: contentFolder.Name, Type: "folder"})
	}

	return &Folder{FolderId: folder.ID, Name: folder.Name, Path: path, Content: content}
}
