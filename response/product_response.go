package response

import "gostarter-backend/models"

type ProductResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (res ProductResponse) Collections(datas []models.Product) interface{} {
	collection := make([]ProductResponse, 0)

	for index := range datas {
		collection = append(collection, ProductResponse{
			ID:          datas[index].ID,
			Name:        datas[index].Name,
			Description: datas[index].Description,
		})
	}
	return collection
}

func (res ProductResponse) Response(data models.Product) interface{} {
	return ProductResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
	}
}
