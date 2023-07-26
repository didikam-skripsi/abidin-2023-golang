package response

import (
	"gostarter-backend/models"
)

type NilaiResponse struct {
	Uts   int    `json:"uts"`
	Uas   int    `json:"uas"`
	Tugas int    `json:"tugas"`
	Absen int    `json:"absen"`
	Sikap string `json:"sikap"`
	Class string `json:"class"`
}

func NewNilaiResponse() NilaiResponse {
	return NilaiResponse{}
}

func (res NilaiResponse) Collections(datas []*models.Nilai) interface{} {
	collection := make([]NilaiResponse, 0)
	for _, data := range datas {
		resNilai := res.Response(data)
		nilRes := NilaiResponse{
			Uts:   resNilai.Uts,
			Uas:   resNilai.Uas,
			Tugas: resNilai.Tugas,
			Absen: resNilai.Absen,
			Sikap: resNilai.Sikap,
			Class: resNilai.Class,
		}
		collection = append(collection, nilRes)
	}
	return collection
}

func (this NilaiResponse) Response(data *models.Nilai) *models.Nilai {
	return data
}
