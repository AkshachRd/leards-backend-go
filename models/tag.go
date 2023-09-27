package models

type Tag struct {
	Base
	Name string `gorm:"unique;not null"`
}
