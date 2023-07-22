package response

import (
	"gostarter-backend/models"

	"github.com/google/uuid"
)

type UserResponse struct {
	UUID     uuid.UUID       `json:"uuid"`
	Username string          `json:"username"`
	Role     models.RoleType `json:"role"`
}

func (res UserResponse) Collections(datas []models.User) interface{} {
	collection := make([]UserResponse, 0)

	for index := range datas {
		collection = append(collection, UserResponse{
			UUID:     datas[index].UUID,
			Username: datas[index].Username,
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
		Username: data.Username,
	}
}

func (res UserResponse) ResponseWithAccess(data *models.User) interface{} {
	if data == nil {
		return nil
	}
	return UserResponse{
		UUID:     data.UUID,
		Username: data.Username,
		Role:     data.Role,
	}
}
