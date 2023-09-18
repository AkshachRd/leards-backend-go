package models

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

const TokenLength = 32
const TokenExpiration = time.Hour

type User struct {
	Base
	Name                string
	Email               string
	PasswordHashed      string
	AuthToken           sql.NullString `gorm:"index"`
	AuthTokenExpiration sql.NullTime
	ProfileIconPath     sql.NullString
	RootFolderID        uuid.NullUUID
	RootFolder          Folder
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewUser(db *gorm.DB, name string, email string, password string) error {
	passwordHashed, err := hashPassword(password)
	if err != nil {
		return err
	}

	result := db.Create(&User{Name: name, Email: email, PasswordHashed: passwordHashed})
	return result.Error
}

func (u *User) SetPassword(db *gorm.DB, password string) error {
	passwordHashed, err := hashPassword(password)
	if err != nil {
		return err
	}

	u.PasswordHashed = passwordHashed
	result := db.Save(u)
	return result.Error
}

func FetchUserByLogin(db *gorm.DB, login string) (*User, error) {
	var user *User

	result := db.Where("name = ?", login).Or("email = ?", login).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *User) IsPasswordCorrect(password string) bool {
	return checkPasswordHash(password, u.PasswordHashed)
}

func generateRandomToken(length int) (string, error) {
	// Determine the number of bytes needed for the given length
	numBytes := (length * 6) / 8 // 6 bits per base64 character

	// Generate random bytes
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64
	token := base64.URLEncoding.EncodeToString(randomBytes)

	// Trim any padding characters '='
	token = token[:length]

	return token, nil
}

func (u *User) GenerateAuthToken(db *gorm.DB) error {
	randomToken, err := generateRandomToken(TokenLength)
	if err != nil {
		return err
	}

	u.AuthToken = sql.NullString{String: randomToken, Valid: true}
	u.AuthTokenExpiration = sql.NullTime{Time: time.Now().UTC().Add(TokenExpiration), Valid: true}

	result := db.Save(u)
	return result.Error
}

func (u *User) RevokeAuthToken(db *gorm.DB) error {
	u.AuthTokenExpiration = sql.NullTime{Time: time.Now().UTC(), Valid: true}

	result := db.Save(u)
	return result.Error
}

func FetchUserByToken(db *gorm.DB, authToken string) (*User, error) {
	var user *User

	result := db.Where("auth_token = ?", authToken).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
