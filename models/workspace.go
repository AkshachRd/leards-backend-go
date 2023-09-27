package models

import "github.com/google/uuid"

type Workspace struct {
	Base
	Name         string
	RootFolderID uuid.UUID
	RootFolder   Folder
}
