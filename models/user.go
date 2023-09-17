package models

import (
	"database/sql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

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

func NewUser(name string, email string, password string) (*User, error) {
	passwordHashed, err := hashPassword(password)
	return &User{Name: name, Email: email, PasswordHashed: passwordHashed}, err
}

func (u *User) SetPassword(password string) (*User, error) {
	passwordHashed, err := hashPassword(password)
	u.PasswordHashed = passwordHashed
	return u, err
}

func (u *User) IsPasswordCorrect(password string) bool {
	return checkPasswordHash(password, u.PasswordHashed)
}
