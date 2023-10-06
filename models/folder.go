package models

import "gorm.io/gorm"

type Folder struct {
	Base
	Name           string `gorm:"size:255; not null"`
	AccessTypeId   uint8  `gorm:"not null"`
	AccessType     AccessType
	ParentFolderID *string `gorm:"size:36"`
	ParentFolder   *Folder
	Permissions    []Permission    `gorm:"polymorphic:Storage;polymorphicValue:folder"`
	StorageHasTags []StorageHasTag `gorm:"polymorphic:Storage;polymorphicValue:folder"`
}

func NewFolder(db *gorm.DB, name string, accessTypeId uint8) *Folder {
	folder := Folder{Name: name, AccessTypeId: accessTypeId}
	db.Create(&folder)
	return &folder
}

func FetchFolderById(db *gorm.DB, id string) (*Folder, error) {
	var folder Folder

	err := db.First(&folder, "id_folder = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &folder, nil
}

func FetchFoldersByParentFolderId(db *gorm.DB, parentFolderId string) (*[]Folder, error) {
	var folders []Folder

	err := db.Find(&folders, "parent_folder_id = ?", parentFolderId).Error
	if err != nil {
		return nil, err
	}

	return &folders, nil
}
