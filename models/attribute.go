package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Attribute struct {
	gorm.Model
	UUID       uuid.UUID `gorm:"size:36;not null;unique" json:"uuid"`
	Type       string    `gorm:"size:255;not null;" json:"type"`
	ScopeName  string    `gorm:"size:255;not null;" json:"scope_name"`
	Scope      string    `gorm:"size:255;not null;" json:"scope"`
	Name       string    `gorm:"size:255;not null;" json:"name"`
	Value      string    `gorm:"size:255;" json:"value"`
	RangeStart int       `gorm:"size:11;" json:"range_start"`
	RangeEnd   int       `gorm:"size:11;" json:"range_end"`
}

func (attribute *Attribute) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate UUID v4 and assign it to ID field before creating the record
	attribute.UUID = uuid.New()
	return nil
}
