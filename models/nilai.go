package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Nilai struct {
	gorm.Model
	Uts          int           `gorm:"size:11;" json:"uts"`
	Uas          int           `gorm:"size:11;" json:"uas"`
	Tugas        int           `gorm:"size:11;" json:"tugas"`
	Absen        int           `gorm:"size:11;" json:"absen"`
	Sikap        string        `gorm:"size:11;" json:"sikap"`
	Class        string        `gorm:"size:11;" json:"class"`
	SiswaUuid    uuid.UUID     `gorm:"size:36;not null" json:"siswa_uuid"`
	Siswa        *Siswa        `gorm:"foreignKey:SiswaUuid"` // Hubungan BelongsTo dengan model Siswa
	Transformasi *Transformasi `gorm:"foreignkey:SiswaUuid;association_foreignkey:SiswaUuid"`
}
