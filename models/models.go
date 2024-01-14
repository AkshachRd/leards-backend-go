package models

import (
	"database/sql"
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

var db *gorm.DB

type AccessType uint8

const (
	AccessTypePrivate AccessType = iota
	AccessTypePublic
	AccessTypeShared
)

func (at AccessType) String() (string, error) {
	switch at {
	case AccessTypePrivate:
		return "private", nil
	case AccessTypePublic:
		return "public", nil
	case AccessTypeShared:
		return "shared", nil
	}

	return "", fmt.Errorf("unknown access type: %d", at)
}

func AccessTypeToString(accessTypeString string) (AccessType, error) {
	switch accessTypeString {
	case "public":
		return AccessTypePublic, nil
	case "private":
		return AccessTypePrivate, nil
	case "shared":
		return AccessTypeShared, nil
	default:
		return 0, fmt.Errorf("unknown access type: %s", accessTypeString)
	}
}

const (
	StorageTypeFolder = "folder"
	StorageTypeDeck   = "deck"
)

type Model struct {
	ID        string `gorm:"primary_key; unique; type:char(36)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

func (base *Model) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.NewString()
	return
}

type MockUserData struct {
	name     string
	email    string
	password string
}

func GetDatabase() *gorm.DB {
	return db
}

// Setup initializes the database instance
func Setup() {
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

	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{NamingStrategy: CustomNamingStrategy{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}, Logger: newLogger})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	err = db.AutoMigrate(
		&Card{},
		&Deck{},
		&Folder{},
		&Permission{},
		&Repetition{},
		&StorageHasTag{},
		&Tag{},
		&User{},
		&UserHasWorkspace{},
		&UserSetting{},
		&Workspace{},
		&UserWorkspaceRole{},
		&FavoriteStorage{},
	)
	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}
}

type CustomNamingStrategy struct {
	schema.NamingStrategy
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func (c CustomNamingStrategy) ColumnName(table, column string) string {
	if strings.ToLower(column) == "id" {
		column = column + "_" + table
	}

	snake := matchFirstCap.ReplaceAllString(column, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

//
//db.Scopes(Paginate(r)).Find(&users)
//db.Scopes(Paginate(r)).Find(&articles)
