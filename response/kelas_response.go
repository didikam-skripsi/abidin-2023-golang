package response

import (
	"gostarter-backend/models"

	"github.com/google/uuid"
)

type KelasResponse struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

func (res KelasResponse) Collections(datas []models.Kelas) interface{} {
	collection := make([]KelasResponse, 0)
	for _, data := range datas {
		collection = append(collection, res.Response(data))
	}
	return collection
}

func (this KelasResponse) Response(data models.Kelas) KelasResponse {
	this.UUID = data.UUID
	this.Name = data.Name
	return this
}
