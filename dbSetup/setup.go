package dbSetup

import (
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func Setup() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{NamingStrategy: CustomNamingStrategy{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}, Logger: newLogger})
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
