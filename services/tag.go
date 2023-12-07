package services

import "github.com/AkshachRd/leards-backend-go/models"

type TagService struct {
	StorageId   string
	StorageType string
	Tags        map[string]uint64
}

func NewTagService(storageId string, storageType string) (*TagService, error) {
	fetchedTags, err := models.FetchTagsByStorageIdByStorageId(storageId, storageType)
	if err != nil {
		return &TagService{}, nil
	}

	tags := make(map[string]uint64)
	for _, tag := range *fetchedTags {
		tags[tag.Name] = tag.ID
	}

	return &TagService{storageId, storageType, tags}, nil
}

func (t *TagService) AddTags(tags []string) error {
	for _, tagName := range tags {
		tag, err := models.FetchOrCreateTagByName(tagName)
		if err != nil {
			return err
		}

		_, err = models.FetchOrCreateStorageHasTag(t.StorageId, t.StorageType, tag.ID)
		if err != nil {
			return err
		}

		t.Tags[tagName] = tag.ID
	}

	return nil
}

func (t *TagService) RemoveTags(tags []string) error {
	for _, tagName := range tags {
		if tagId, ok := t.Tags[tagName]; ok {
			if err := models.DeleteStorageHasTagByStorageIdAndTagId(t.StorageId, t.StorageType, tagId); err != nil {
				return err
			}

			delete(t.Tags, tagName)
		}
	}

	return nil
}

func (t *TagService) GetTags() []string {
	tags := make([]string, 0)

	for tag := range t.Tags {
		tags = append(tags, tag)
	}

	return tags
}
