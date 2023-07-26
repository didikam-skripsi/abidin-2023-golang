package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Siswa struct {
	gorm.Model
	UUID         uuid.UUID    `gorm:"size:36;not null;unique" json:"uuid"`
	Name         string       `gorm:"size:255;not null;" json:"name"`
	Nisn         string       `gorm:"size:255;not null;" json:"nisn"`
	KelasUuid    uuid.UUID    `gorm:"size:36;not null" json:"kelas_uuid"`
	Kelas        *Kelas       `gorm:"foreignKey:KelasUuid"`
	Nilai        *Nilai       `gorm:"foreignKey:SiswaUuid"` // Hubungan HasOne dengan model Nilai
	Transformasi Transformasi `gorm:"foreignKey:SiswaUuid"` // Hubungan HasOne dengan model Transformasi
}

func (siswa *Siswa) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate UUID v4 and assign it to ID field before creating the record
	siswa.UUID = uuid.New()
	return nil
}
