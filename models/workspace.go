package models

type Workspace struct {
	Base
	Name         string `gorm:"size:255"`
	RootFolderID string `gorm:"size:36"`
	RootFolder   Folder
}
