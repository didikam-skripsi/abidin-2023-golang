package services

import (
	"errors"
	"gostarter-backend/helpers/token"
	"gostarter-backend/models"
	"gostarter-backend/request"

	"github.com/google/uuid"
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
