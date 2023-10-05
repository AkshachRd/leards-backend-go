package models

type Tag struct {
	Base
	Name string `gorm:"size:255; unique; not null"`
}
