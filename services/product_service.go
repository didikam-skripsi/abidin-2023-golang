package services

import (
	"fmt"
	"gostarter-backend/models"
	"gostarter-backend/request"
)

type ProductService struct{}

func (s ProductService) Store(input request.ProductRequest) (models.Product, error) {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
	}

	err := models.DB.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s ProductService) Update(ID uint, input request.ProductRequest) (models.Product, error) {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
	}

	if models.DB.Model(&product).Where("id = ?", ID).Updates(&product).RowsAffected == 0 {
		return product, fmt.Errorf("failed to update product with ID %d", ID)
	}
	product.ID = ID
	return product, nil
}

func (s ProductService) Delete(ID uint) error {
	product := models.Product{}
	if models.DB.Delete(&product, ID).RowsAffected == 0 {
		return fmt.Errorf("failed to delete product with ID %d", ID)
	}
	return nil
}

func (s ProductService) Show(ID uint) (models.Product, error) {
	var product models.Product
	if err := models.DB.First(&product, ID).Error; err != nil {
		return product, err
	}

	return product, nil
}
