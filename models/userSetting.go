package models

import "gorm.io/gorm"

type Setting string

const (
	Language Setting = "language"
	Theme    Setting = "theme"
)

type UserSetting struct {
	Base
	UserID       string `gorm:"size:36"`
	User         User
	SettingName  Setting `gorm:"size:255"`
	SettingValue string  `gorm:"size:255"`
}

func FetchUserSettingsByUserId(db *gorm.DB, userId string) (*[]UserSetting, error) {
	var userSettings []UserSetting

	err := db.Find(&userSettings, "user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}

	return &userSettings, nil
}
