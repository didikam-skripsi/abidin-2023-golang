package response

import (
	"gostarter-backend/models"

	"github.com/google/uuid"
)

type UserResponse struct {
	UUID     uuid.UUID       `json:"uuid"`
	Name     string          `json:"name"`
	Username string          `json:"username"`
	Role     models.RoleType `json:"role"`
}

func (res UserResponse) Collections(datas []models.User) interface{} {
	collection := make([]UserResponse, 0)

	for index := range datas {
		collection = append(collection, UserResponse{
			UUID:     datas[index].UUID,
			Name:     datas[index].Name,
			Username: datas[index].Username,
			Role:     datas[index].Role,
		})
	}
	return collection
}

func (res UserResponse) Response(data *models.User) interface{} {
	if data == nil {
		return nil
	}
	return UserResponse{
		UUID:     data.UUID,
		Name:     data.Name,
		Username: data.Username,
		Role:     data.Role,
	}
}

func (res UserResponse) ResponseWithAccess(data *models.User) interface{} {
	if data == nil {
		return nil
	}
	return UserResponse{
		UUID:     data.UUID,
		Name:     data.Name,
		Username: data.Username,
		Role:     data.Role,
	}
}
