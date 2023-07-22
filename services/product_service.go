package services

import (
	"fmt"
	"gostarter-backend/models"
	"gostarter-backend/request"

	"github.com/google/uuid"
)

type ProductService struct{}

func (s ProductService) Store(UserUuid uuid.UUID, input request.ProductRequest) (models.Product, error) {
	product := models.Product{
		UserUuid:    UserUuid,
		Name:        input.Name,
		Description: input.Description,
	}

	err := models.DB.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s ProductService) Update(UUID uuid.UUID, input request.ProductRequest) (models.Product, error) {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
	}

	if models.DB.Model(&product).Where("uuid = ?", UUID).Updates(&product).RowsAffected == 0 {
		return product, fmt.Errorf("failed to update product with ID %d", UUID)
	}
	product.UUID = UUID
	return product, nil
}

func (s ProductService) Delete(UUID uuid.UUID) error {
	product := models.Product{}
	// Melakukan hard delete pada data dengan UUID tertentu
	if models.DB.Unscoped().Where("uuid = ?", UUID).Delete(&product).RowsAffected == 0 {
		return fmt.Errorf("failed to delete product with ID %d", UUID)
	}
	return nil
}

func (s ProductService) Show(UUID uuid.UUID) (models.Product, error) {
	var product models.Product
	if err := models.DB.Preload("User").Where("uuid = ?", UUID).First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}
