package models

import (
	"github.com/AkshachRd/leards-backend-go/settings"
	"gorm.io/gorm"
)

type Setting string

const (
	Locale Setting = "locale"
	Theme  Setting = "theme"
)

type UserSetting struct {
	Model
	UserID       string `gorm:"size:36"`
	User         User
	SettingName  Setting `gorm:"size:255"`
	SettingValue string  `gorm:"size:255"`
}

func getUserSettingsMap() map[Setting]string {
	return map[Setting]string{
		Locale: settings.AppSettings.EnvVars.DefaultLocale,
		Theme:  "light",
	}
}

func NewUserSetting(db *gorm.DB, userId string, settingName Setting, settingValue string) (*UserSetting, error) {
	userSetting := UserSetting{UserID: userId, SettingName: settingName, SettingValue: settingValue}
	err := db.Create(&userSetting).Error
	if err != nil {
		return &UserSetting{}, err
	}

	return &userSetting, nil
}

func NewUserSettings(db *gorm.DB, userId string) (*[]UserSetting, error) {
	var userSettings []UserSetting

	for settingName, settingValue := range getUserSettingsMap() {
		userSetting, err := NewUserSetting(db, userId, settingName, settingValue)
		if err != nil {
			return nil, err
		}

		userSettings = append(userSettings, *userSetting)
	}

	return &userSettings, nil
}

func UpdateUserSettings(userId string, settings map[Setting]string) (*[]UserSetting, error) {
	userSettings, err := FetchUserSettingsByUserId(userId)
	if err != nil {
		return nil, err
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for i, userSetting := range *userSettings {
		if settingValue, ok := settings[userSetting.SettingName]; ok {
			err = userSetting.Update(db, settingValue)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
			(*userSettings)[i] = userSetting
		}
	}

	err = tx.Commit().Error
	return userSettings, err
}

func (us *UserSetting) Update(db *gorm.DB, settingValue string) error {
	us.SettingValue = settingValue

	err := db.Save(us).Error
	if err != nil {
		return err
	}

	return nil
}

func FetchUserSettingsByUserId(userId string) (*[]UserSetting, error) {
	var userSettings []UserSetting

	err := db.Find(&userSettings, "user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}

	return &userSettings, nil
}
