package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Folder struct {
	Model
	Name             string            `gorm:"size:255; not null"`
	AccessType       uint8             `gorm:"default:0; not null"`
	ParentFolderID   *string           `gorm:"size:36"`
	ParentFolder     *Folder           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Folders          []Folder          `gorm:"foreignkey:ParentFolderID"`
	Decks            []Deck            `gorm:"foreignKey:ParentFolderID"`
	Permissions      []Permission      `gorm:"polymorphic:Storage;polymorphicValue:folder"`
	FavoriteStorages []FavoriteStorage `gorm:"polymorphic:Storage;polymorphicValue:folder"`
	StorageHasTags   []StorageHasTag   `gorm:"polymorphic:Storage;polymorphicValue:folder"`
}

func getFolderPreloadArgs(query string) []interface{} {
	args := make([]interface{}, 0)
	switch query {
	case "ParentFolder.ParentFolder":
		var preload func(d *gorm.DB) *gorm.DB
		preload = func(d *gorm.DB) *gorm.DB {
			return d.Preload("ParentFolder", preload)
		}
		args = append(args, preload)
	}
	return args
}

func getFolderPreloadQuery(index int) string {
	return []string{
		"ParentFolder",
		"Folders",
		"Decks.Cards",
		"ParentFolder.ParentFolder",
		"StorageHasTags.Tag",
	}[index]
}

func NewFolder(db *gorm.DB, name string, accessType uint8, parentFolderId *string) (*Folder, error) {
	folder := Folder{Name: name, AccessType: accessType}
	if parentFolderId != nil {
		if _, err := uuid.Parse(*parentFolderId); err == nil {
			folder.ParentFolderID = parentFolderId
		}
	}

	err := db.Create(&folder).Error
	if err != nil {
		return &Folder{}, nil
	}

	return &folder, nil
}

func CreateFolder(name string, accessType uint8, parentFolderId *string, userId string) (*Folder, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	folder, err := NewFolder(tx, name, accessType, parentFolderId)
	if err != nil {
		tx.Rollback()
		return &Folder{}, err
	}

	_, err = NewPermission(tx, folder.ID, StorageTypeFolder, userId, PermissionTypeOwner)
	if err != nil {
		tx.Rollback()
		return &Folder{}, nil
	}

	if err = tx.Commit().Error; err != nil {
		return &Folder{}, nil
	}

	return folder, nil
}

func UpdateFolderById(id string, name string) (*Folder, error) {
	folder, err := FetchFolderById(id, false, false, false, false)
	if err != nil {
		return &Folder{}, err
	}

	folder.Name = name

	err = db.Save(folder).Error
	if err != nil {
		return &Folder{}, err
	}

	folder, err = FetchFolderById(id, false, true, true, true, true)
	if err != nil {
		return &Folder{}, err
	}

	return folder, nil
}

func DeleteFolderById(id string) error {
	folder, err := FetchFolderById(id, false, true, true, false, true)
	if err != nil {
		return err
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = folder.Delete(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (f *Folder) Delete(db *gorm.DB) error {
	decks := f.Decks
	if decks == nil {
		if err := db.Find(&decks, "parent_folder_id = ?", f.ID).Error; err != nil {
			return err
		}
	}

	for _, deck := range decks {
		err := deck.Delete(db)
		if err != nil {
			return err
		}
	}

	subfolders := f.Folders
	if subfolders == nil {
		if err := db.Find(&subfolders, "parent_folder_id = ?", f.ID).Error; err != nil {
			return err
		}
	}

	for _, subfolder := range subfolders {
		err := subfolder.Delete(db)
		if err != nil {
			return err
		}
	}

	storageHasTags := f.StorageHasTags
	if storageHasTags == nil {
		if err := db.Find(
			&storageHasTags,
			"storage_id = ? AND storage_type = ?",
			f.ID, StorageTypeFolder).Error; err != nil {
			return err
		}
	}

	if len(storageHasTags) != 0 {
		err := db.Delete(&storageHasTags).Error
		if err != nil {
			return err
		}
	}

	err := db.Delete(f).Error
	if err != nil {
		return err
	}

	return nil
}

// FetchFolderById
//
// Preload args: "ParentFolder", "Folders", "DecksWithCards", "ParentFolderRecursive", "Tags"
func FetchFolderById(id string, preloadArgs ...bool) (*Folder, error) {
	var folder Folder

	query := db
	for i, arg := range preloadArgs {
		if arg {
			queryString := getFolderPreloadQuery(i)
			query = query.Preload(queryString, getFolderPreloadArgs(queryString)...)
		}
	}

	err := query.First(&folder, "id_folder = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &folder, nil
}

func FetchFoldersByParentFolderId(parentFolderId string) (*[]Folder, error) {
	var folders []Folder

	err := db.Find(&folders, "parent_folder_id = ?", parentFolderId).Error
	if err != nil {
		return nil, err
	}

	return &folders, nil
}

func (f *Folder) SetAccessType(accessType uint8) error {
	err := db.Model(f).Update("access_type", accessType).Error
	if err != nil {
		return err
	}

	f.AccessType = accessType

	return nil
}
