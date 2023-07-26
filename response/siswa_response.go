package response

import (
	"gostarter-backend/models"

	"github.com/google/uuid"
)

type SiswaResponse struct {
	UUID         uuid.UUID     `json:"uuid"`
	Name         string        `json:"name"`
	Nisn         string        `json:"nisn"`
	Transformasi interface{}   `json:"transformasi"`
	Kelas        interface{}   `json:"kelas"`
	Nilai        *models.Nilai `json:"nilai"`
}

func (res SiswaResponse) Collections(datas []models.Siswa) interface{} {
	collection := make([]SiswaResponse, 0)
	for _, data := range datas {
		collection = append(collection, res.Response(data))
	}
	return collection
}

func (this SiswaResponse) Response(data models.Siswa) SiswaResponse {
	this.UUID = data.UUID
	this.Name = data.Name
	this.Nisn = data.Nisn
	this.Kelas = data.Kelas
	this.Transformasi = data.Transformasi
	this.Nilai = NewNilaiResponse().Response(data.Nilai)
	return this
}
