package services

import (
	"fmt"
	"gostarter-backend/models"
	"gostarter-backend/request"

	"github.com/google/uuid"
)

type AttributeService struct{}

func (s AttributeService) Update(UUID uuid.UUID, input request.AttributeRequest) (models.Attribute, error) {
	attribute := models.Attribute{
		Name:       input.Name,
		Value:      input.Value,
		RangeStart: input.RangeStart,
		RangeEnd:   input.RangeEnd,
	}
	updateData := map[string]interface{}{
		"name":        attribute.Name,
		"value":       attribute.Value,
		"range_start": attribute.RangeStart,
		"range_end":   attribute.RangeEnd,
	}

	if models.DB.Model(&attribute).Where("uuid = ?", UUID).Updates(updateData).RowsAffected == 0 {
		return attribute, fmt.Errorf("gagal update data dengan UUID %d", UUID)
	}
	attribute.UUID = UUID
	return attribute, nil
}

func (s AttributeService) Show(UUID uuid.UUID) (models.Attribute, error) {
	var attribute models.Attribute
	if err := models.DB.Where("uuid = ?", UUID).First(&attribute).Error; err != nil {
		return attribute, err
	}

	return attribute, nil
}
