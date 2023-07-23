package services

import (
	"errors"
	"fmt"
	"gostarter-backend/helpers/token"
	"gostarter-backend/models"
	"gostarter-backend/request"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (this UserService) Register(input request.RegisterRequest) (*models.User, error) {
	var err error
	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password
	err = models.DB.Create(&u).Error
	if err != nil {
		return &u, err
	}
	return &u, nil
}

func (this UserService) GetUserByID(uid uint) (models.User, error) {
	var u models.User
	if err := models.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}
	u.PrepareGive()
	return u, nil
}

func (this UserService) GetUserByUUID(uid uuid.UUID) (*models.User, error) {
	var u models.User
	if err := models.DB.Where("uuid = ?", uid).First(&u).Error; err != nil {
		return &u, errors.New("User not found!")
	}
	u.PrepareGive()
	return &u, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (this UserService) LoginCheck(input request.LoginRequest) (string, error) {
	var err error
	u := models.User{}
	err = models.DB.Model(models.User{}).Where("username = ?", input.Username).Take(&u).Error

	if err != nil {
		return "", err
	}
	err = VerifyPassword(input.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := token.GenerateToken(u)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (s UserService) Show(UUID uuid.UUID) (models.User, error) {
	var user models.User
	if err := models.DB.Where("uuid = ?", UUID).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (s UserService) IsExists(username string) (bool, error) {
	var user models.User
	if err := models.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s UserService) Store(input request.UserCreateRequest) (models.User, error) {
	user := models.User{
		Name:     input.Name,
		Username: input.Username,
		Password: input.Password,
		Role:     input.Role,
	}

	err := models.DB.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s UserService) Update(UUID uuid.UUID, input request.UserUpdateRequest) (models.User, error) {
	user := models.User{
		Name:     input.Name,
		Username: input.Username,
		Role:     input.Role,
	}
	if input.Password != "" {
		user.Password = input.Password
	}
	if models.DB.Model(&user).Where("uuid = ?", UUID).Updates(&user).RowsAffected == 0 {
		return user, fmt.Errorf("failed to update data with UUID %d", UUID)
	}
	user.UUID = UUID
	return user, nil
}

func (s UserService) Delete(UUID uuid.UUID) error {
	user := models.User{}
	// Melakukan hard delete pada data dengan UUID tertentu
	if models.DB.Unscoped().Where("uuid = ?", UUID).Delete(&user).RowsAffected == 0 {
		return fmt.Errorf("failed to delete data with UUID %d", UUID)
	}
	return nil
}
