package models

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
