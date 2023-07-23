package request

import (
	"gostarter-backend/models"
)

type UserCreateRequest struct {
	Name     string          `validate:"required" json:"name"`
	Username string          `validate:"required" json:"username"`
	Password string          `validate:"required" json:"password"`
	Role     models.RoleType `validate:"required,oneof=admin user" json:"role"`
}

type UserUpdateRequest struct {
	Name     string `validate:"required" json:"name"`
	Username string `validate:"required" json:"username"`
	Password string
	Role     models.RoleType `validate:"required,oneof=admin user" json:"role"`
}
