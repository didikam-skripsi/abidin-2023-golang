package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Kelas struct {
	gorm.Model
	UUID uuid.UUID `gorm:"size:36;not null;unique" json:"uuid"`
	Name string    `gorm:"size:255;not null;" json:"name"`
}

func (kelas *Kelas) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate UUID v4 and assign it to ID field before creating the record
	kelas.UUID = uuid.New()
	return nil
}
