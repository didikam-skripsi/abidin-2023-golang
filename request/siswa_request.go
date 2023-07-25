package request

import "github.com/google/uuid"

type SiswaRequest struct {
	Name      string    `validate:"required" json:"name"`
	Nisn      string    `validate:"required" json:"nisn"`
	KelasUuid uuid.UUID `validate:"required" json:"kelas_uuid"`
}
