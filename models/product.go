package models

import "github.com/jinzhu/gorm"

type Product struct {
	Name        string `gorm:"size:255;not null;" json:"name"`
	Description string `gorm:"size:255;not null;" json:"description"`
	gorm.Model
}
