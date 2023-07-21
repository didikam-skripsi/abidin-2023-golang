package response

import "gostarter-backend/models"

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func (res UserResponse) Collections(datas []models.User) interface{} {
	collection := make([]UserResponse, 0)

	for index := range datas {
		collection = append(collection, UserResponse{
			ID:       datas[index].ID,
			Username: datas[index].Username,
		})
	}
	return collection
}

func (res UserResponse) Response(data models.User) interface{} {
	return UserResponse{
		ID:       data.ID,
		Username: data.Username,
	}
}
