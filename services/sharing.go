package services

import (
	"fmt"
	"github.com/AkshachRd/leards-backend-go/models"
)

type SharingService struct {
}

func NewSharingService() *SharingService {
	return &SharingService{}
}

func (ss *SharingService) SetStorageAccess(storageId string, storageType string, accessType string) error {
	accessTypeInt, err := models.AccessTypeToString(accessType)
	if err != nil {
		return fmt.Errorf("unkown access type: %s", accessType)
	}

	switch storageType {
	case models.StorageTypeFolder:
		err = ss.setFolderAccess(storageId, accessTypeInt)
	case models.StorageTypeDeck:
		err = ss.setDeckAccess(storageId, accessTypeInt)
	default:
		err = fmt.Errorf("unkown storage type: %s", storageType)
	}
	if err != nil {
		return err
	}

	return nil
}

func (ss *SharingService) setFolderAccess(folderId string, accessType models.AccessType) error {
	folder := models.Folder{Model: models.Model{ID: folderId}}
	return folder.SetAccessType(accessType)
}

func (ss *SharingService) setDeckAccess(deckId string, accessType models.AccessType) error {
	deck := models.Deck{Model: models.Model{ID: deckId}}
	return deck.SetAccessType(accessType)
}
