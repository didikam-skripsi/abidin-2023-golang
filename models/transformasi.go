package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Transformasi struct {
	gorm.Model
	Uts       *uuid.UUID `gorm:"size:36;default:null;" json:"uts"`
	UtsName   string     `gorm:"size:255;default:null;" json:"uts_name"`
	Uas       *uuid.UUID `gorm:"size:36;default:null;;" json:"uas"`
	UasName   string     `gorm:"size:255;default:null;" json:"uas_name"`
	Tugas     *uuid.UUID `gorm:"size:36;default:null;;" json:"tugas"`
	TugasName string     `gorm:"size:255;default:null;" json:"tugas_name"`
	Absen     *uuid.UUID `gorm:"size:36;default:null;;" json:"absen"`
	AbsenName string     `gorm:"size:255;default:null;" json:"absen_name"`
	Sikap     *uuid.UUID `gorm:"size:36;default:null;;" json:"sikap"`
	SikapName string     `gorm:"size:255;default:null;" json:"sikap_name"`
	Class     *uuid.UUID `gorm:"size:36;default:null;" json:"class"`
	ClassName string     `gorm:"size:255;default:null;" json:"class_name"`
	SiswaUuid uuid.UUID  `gorm:"size:36;not null" json:"siswa_uuid"`
	Siswa     *Siswa     `gorm:"foreignKey:uuid;references:SiswaUuid"`
	Nilai     *Nilai     `gorm:"foreignkey:SiswaUuid;association_foreignkey:SiswaUuid"`
}

// contoh relasi
// type Nilai struct {
//     gorm.Model
//     Uts          int          `gorm:"size:11;" json:"uts"`
//     SiswaUuid    uuid.UUID    `gorm:"size:36;not null" json:"siswa_uuid"`
//     Siswa        *Siswa       `gorm:"foreignkey:SiswaUuid;association_foreignkey:Uuid"`
//     Transformasi *Transformasi `gorm:"foreignkey:SiswaUuid;association_foreignkey:SiswaUuid"`
// }

// type Transformasi struct {
//     gorm.Model
//     Uts       int       `gorm:"size:11;default:null;" json:"uts"`
//     SiswaUuid uuid.UUID `gorm:"size:36;not null" json:"siswa_uuid"`
//     Siswa     *Siswa    `gorm:"foreignkey:SiswaUuid;association_foreignkey:Uuid"`
//     Nilai     *Nilai    `gorm:"foreignkey:SiswaUuid;association_foreignkey:SiswaUuid"`
// }
