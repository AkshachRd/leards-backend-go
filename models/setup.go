package models

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Setup() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&AccessType{}, &Card{}, &Deck{}, &Folder{}, &Permission{}, &PermissionType{}, &User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
