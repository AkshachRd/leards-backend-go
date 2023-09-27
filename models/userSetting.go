package models

import "github.com/google/uuid"

type Setting string

const (
	Language Setting = "language"
	Theme    Setting = "theme"
)

type UserSetting struct {
	Base
	UserID       uuid.UUID
	User         User
	SettingName  Setting
	SettingValue string
}
