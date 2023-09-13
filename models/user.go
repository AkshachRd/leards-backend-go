package models

import (
	"database/sql"
)

type User struct {
	Base
	Name                string
	Email               string
	PasswordHashed      string
	AuthToken           sql.NullString `gorm:"index"`
	AuthTokenExpiration sql.NullTime
	ProfileIconPath     sql.NullString
	RootFolderID        sql.NullString
	RootFolder          Folder `gorm:"foreignKey:FolderID;references:RootFolderID"`
}
