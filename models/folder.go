package models

import "gorm.io/gorm"

type Folder struct {
	Base
	Name           string `gorm:"size:255; not null"`
	AccessTypeID   uint8  `gorm:"not null"`
	AccessType     AccessType
	ParentFolderID *string `gorm:"size:36"`
	ParentFolder   *Folder
	Folders        []Folder        `gorm:"foreignkey:ParentFolderID"`
	Decks          []Deck          `gorm:"foreignKey:ParentFolderID"`
	Permissions    []Permission    `gorm:"polymorphic:Storage;polymorphicValue:folder"`
	StorageHasTags []StorageHasTag `gorm:"polymorphic:Storage;polymorphicValue:folder"`
}

func getFolderPreloadArgs() []string {
	return []string{"ParentFolder", "Folders", "Decks", "AccessType"}
}

func NewFolder(db *gorm.DB, name string, accessType Access) (*Folder, error) {
	var accType AccessType
	err := db.First(&accType, "type = ?", accessType).Error
	if err != nil {
		return &Folder{}, nil
	}

	folder := Folder{Name: name, AccessTypeID: accType.ID}
	err = db.Create(&folder).Error
	if err != nil {
		return &Folder{}, nil
	}

	return &folder, nil
}

func UpdateFolder(db *gorm.DB, id string, name string, accessType Access) (*Folder, error) {
	folder, err := FetchFolderById(db, id, true, true, true, true)
	if err != nil {
		return &Folder{}, err
	}

	if accessType != Access(folder.AccessType.Type) {
		var accType AccessType
		err = db.First(&accType, "type = ?", accessType).Error
		if err != nil {
			return &Folder{}, err
		}

		folder.AccessTypeID = accType.ID
	}

	folder.Name = name

	err = db.Save(folder).Error
	if err != nil {
		return &Folder{}, err
	}

	return folder, nil
}

// FetchFolderById
//
// Preload args: ParentFolder, AccessType
func FetchFolderById(db *gorm.DB, id string, preloadArgs ...bool) (*Folder, error) {
	var folder Folder

	query := db
	for i, arg := range preloadArgs {
		if arg {
			query = query.Preload(getFolderPreloadArgs()[i])
		}
	}

	err := query.First(&folder, "id_folder = ?", id).Error
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
