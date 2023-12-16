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
	accessTypeInt, ok := ss.convertAccessTypeToInt(accessType)
	if !ok {
		return fmt.Errorf("unkown access type: %s", accessType)
	}

	var err error
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

func (ss *SharingService) setFolderAccess(folderId string, accessType uint8) error {
	folder := models.Folder{Model: models.Model{ID: folderId}}
	return folder.SetAccessType(accessType)
}

func (ss *SharingService) setDeckAccess(deckId string, accessType uint8) error {
	deck := models.Deck{Model: models.Model{ID: deckId}}
	return deck.SetAccessType(accessType)
}

func (ss *SharingService) convertAccessTypeToInt(accessType string) (uint8, bool) {
	switch accessType {
	case "public":
		return models.AccessTypePublic, true
	case "private":
		return models.AccessTypePrivate, true
	case "shared":
		return models.AccessTypeShared, true
	default:
		return 0, false
	}
}
