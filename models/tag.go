package models

type Tag struct {
	Model
	Name string `gorm:"size:255; unique; not null"`
}
