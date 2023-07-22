package response

import (
	"gostarter-backend/models"

	"github.com/google/uuid"
)

type ProductResponse struct {
	UUID        uuid.UUID   `json:"uuid"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	User        interface{} `json:"user"`
}

func (res ProductResponse) Collections(datas []models.Product) interface{} {
	collection := make([]ProductResponse, 0)
	for _, data := range datas {
		collection = append(collection, res.Response(data))
	}
	return collection
}

func (this ProductResponse) Response(data models.Product) ProductResponse {
	this.UUID = data.UUID
	this.Name = data.Name
	this.Description = data.Description

	if data.User != nil {
		this.User = (&UserResponse{}).Response(data.User)
	} else {
		this.User = nil
	}

	return this
}
