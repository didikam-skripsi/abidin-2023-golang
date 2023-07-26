package response

import (
	"gostarter-backend/models"

	"github.com/google/uuid"
)

type AttributeResponse struct {
	UUID       uuid.UUID `json:"uuid"`
	Type       string    `json:"type"`
	ScopeName  string    `json:"scope_name"`
	Scope      string    `json:"scope"`
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	RangeStart int       `json:"range_start"`
	RangeEnd   int       `json:"range_end"`
}

func (res AttributeResponse) Collections(datas []models.Attribute) interface{} {
	collection := make([]AttributeResponse, 0)
	for _, data := range datas {
		collection = append(collection, res.Response(data))
	}
	return collection
}

func (this AttributeResponse) Response(data models.Attribute) AttributeResponse {
	this.UUID = data.UUID
	this.Type = data.Type
	this.ScopeName = data.ScopeName
	this.Scope = data.Scope
	this.Name = data.Name
	this.Value = data.Value
	this.RangeStart = data.RangeStart
	this.RangeEnd = data.RangeEnd
	return this
}
