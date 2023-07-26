package request

import "github.com/google/uuid"

type NilaiRequest struct {
	SiswaUuid uuid.UUID `validate:"required" json:"siswa_uuid"`
	Uts       int       `validate:"required" json:"uts"`
	Uas       int       `validate:"required" json:"uas"`
	Tugas     int       `validate:"required" json:"tugas"`
	Absen     int       `validate:"required" json:"absen"`
	Sikap     string    `validate:"required" json:"sikap"`
	Class     string    `validate:"" json:"class"`
}
