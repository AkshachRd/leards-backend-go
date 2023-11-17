package models

type Access string

const (
	Public  Access = "public"
	Private Access = "private"
)

type AccessType struct {
	ID   uint8  `gorm:"primaryKey"`
	Type string `gorm:"size:255; not null; unique;"`
}

func FillAccessTypes() error {
	acc := AccessType{Type: string(Public)}
	err := db.Create(&acc).Error
	if err != nil {
		return err
	}

	acc = AccessType{Type: string(Private)}
	err = db.Create(&acc).Error
	if err != nil {
		return err
	}

	return nil
}
