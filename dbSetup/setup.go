package dbSetup

import (
	"github.com/AkshachRd/leards-backend-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Setup() (*gorm.DB, error) {
	dsn := ""
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: CustomNamingStrategy{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.AccessType{},
		&models.Card{},
		&models.Deck{},
		&models.Folder{},
		&models.Permission{},
		&models.PermissionType{},
		&models.Repetition{},
		&models.RepetitionState{},
		&models.StorageHasTag{},
		&models.Tag{},
		&models.User{},
		&models.UserHasWorkspace{},
		&models.UserSetting{},
		&models.Workspace{},
		&models.UserWorkspaceRole{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
