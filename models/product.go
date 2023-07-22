package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"size:36;not null;unique" json:"uuid"`
	UserUuid    uuid.UUID `gorm:"size:36;not null" json:"user_uuid"`
	Name        string    `gorm:"size:255;not null;" json:"name"`
	Description string    `gorm:"size:255;" json:"description"`
	User        *User     `gorm:"foreignKey:UserUuid"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate UUID v4 and assign it to ID field before creating the record
	product.UUID = uuid.New()
	return nil
}
