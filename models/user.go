package models

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

const TokenLength = 32
const TokenExpiration = time.Hour

type User struct {
	Model
	Name                string         `gorm:"size:255; not null"`
	Email               string         `gorm:"size:255; not null; unique"`
	PasswordHashed      string         `gorm:"size:255; not null"`
	AuthToken           sql.NullString `gorm:"size:255; index"`
	AuthTokenExpiration sql.NullTime
	ProfileIconPath     sql.NullString `gorm:"size:255"`
	RootFolderID        string         `gorm:"size:36; not null"`
	RootFolder          Folder
	Settings            []UserSetting
}

func (u *User) Update(db *gorm.DB, column string, value interface{}) error {
	return db.Model(u).Update(column, value).Error
}

func (u *User) SetProfileIconPath(profileIconPath string) error {
	err := u.Update(db, "profile_icon_path", profileIconPath)
	if err != nil {
		return err
	}

	u.ProfileIconPath = sql.NullString{String: profileIconPath, Valid: true}
	return nil
}

func getUserPreloadQuery(index int) string {
	return []string{"RootFolder", "Settings"}[index]
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewUser(db *gorm.DB, name string, email string, password string, rootFolderId string) (*User, error) {
	passwordHashed, err := hashPassword(password)
	if err != nil {
		return &User{}, err
	}

	user := User{Name: name, Email: email, PasswordHashed: passwordHashed, RootFolderID: rootFolderId}

	err = db.Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

func CreateUser(name string, email string, password string) (*User, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	rootFolder, err := NewFolder(tx, "rootFolder", AccessTypePrivate, nil)
	if err != nil {
		tx.Rollback()
		return &User{}, nil
	}

	user, err := NewUser(tx, name, email, password, rootFolder.ID)
	if err != nil {
		tx.Rollback()
		return &User{}, nil
	}

	_, err = NewUserSettings(tx, user.ID)
	if err != nil {
		tx.Rollback()
		return &User{}, nil
	}

	_, err = NewPermission(tx, rootFolder.ID, "folder", user.ID, PermissionTypeOwner)
	if err != nil {
		tx.Rollback()
		return &User{}, nil
	}

	err = tx.Commit().Error
	if err != nil {
		return &User{}, nil
	}

	return user, nil
}

func (u *User) SetRootFolderId(rootFolderId string) error {
	err := u.Update(db, "root_folder_id", rootFolderId)
	if err != nil {
		return err
	}

	u.RootFolderID = rootFolderId

	return nil
}

func (u *User) SetPassword(password string) error {
	passwordHashed, err := hashPassword(password)
	if err != nil {
		return err
	}

	err = u.Update(db, "password_hashed", passwordHashed)
	if err != nil {
		return err
	}

	u.PasswordHashed = passwordHashed

	return nil
}

// FetchUserByEmail
//
// Preload args: "RootFolder", "Settings"
func FetchUserByEmail(email string, preloadArgs ...bool) (*User, error) {
	var user User

	query := db
	for i, arg := range preloadArgs {
		if arg {
			query = query.Preload(getUserPreloadQuery(i))
		}
	}

	err := query.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func FetchUserById(id string) (*User, error) {
	var user User

	err := db.First(&user, "id_user = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
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

func (u *User) RefreshAuthToken() error {
	return u.Update(db, "auth_token_expiration", sql.NullTime{
		Time: time.Now().UTC().Add(TokenExpiration), Valid: true,
	})
}

func (u *User) GenerateAuthToken() error {
	randomToken, err := generateRandomToken(TokenLength)
	if err != nil {
		return err
	}

	authToken := sql.NullString{String: randomToken, Valid: true}
	err = u.Update(db, "auth_token", authToken)
	if err != nil {
		return err
	}
	u.AuthToken = authToken

	authTokenExpiration := sql.NullTime{
		Time: time.Now().UTC().Add(TokenExpiration), Valid: true,
	}
	err = u.Update(db, "auth_token_expiration", authTokenExpiration)
	if err != nil {
		return err
	}
	u.AuthTokenExpiration = authTokenExpiration

	return nil
}

func (u *User) RevokeAuthToken() error {
	u.AuthTokenExpiration = sql.NullTime{Time: time.Now().UTC(), Valid: true}

	err := db.Save(u).Error
	return err
}

func FetchUserByToken(authToken string) (*User, error) {
	var user User

	err := db.First(&user, "auth_token = ?", authToken).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) IsTokenValid() bool {
	return u.AuthTokenExpiration.Valid && u.AuthTokenExpiration.Time.After(time.Now().UTC())
}

func UpdateUserById(id string, name string, email string, password string) (*User, error) {
	user, err := FetchUserById(id)
	if err != nil {
		return &User{}, err
	}

	if name != "" && user.Name != name {
		err = user.Update(db, "name", name)
		if err != nil {
			return &User{}, err
		}
		user.Name = name
	}

	if email != "" && user.Email != email {
		err = user.Update(db, "email", email)
		if err != nil {
			return &User{}, err
		}
		user.Email = email
	}

	if password != "" && !user.IsPasswordCorrect(password) {
		// TODO: Инкапсулировать обновление каждого поля как у пароля
		err = user.SetPassword(password)
		if err != nil {
			return &User{}, err
		}
	}

	return user, nil
}
