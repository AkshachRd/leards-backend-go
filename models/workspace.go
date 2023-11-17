package models

type Workspace struct {
	Model
	Name         string `gorm:"size:255"`
	RootFolderID string `gorm:"size:36"`
	RootFolder   Folder
}
