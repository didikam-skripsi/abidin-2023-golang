package services

import (
	"fmt"
	"gostarter-backend/models"
	"gostarter-backend/request"

	"github.com/google/uuid"
)

type SiswaService struct{}

func (s SiswaService) Store(UserUuid uuid.UUID, input request.SiswaRequest) (models.Siswa, error) {
	siswa := models.Siswa{
		Name:      input.Name,
		Nisn:      input.Nisn,
		KelasUuid: input.KelasUuid,
	}
	err := models.DB.Create(&siswa).Error
	if err != nil {
		return siswa, fmt.Errorf("gagal simpan data")
	}
	return siswa, nil
}

func (s SiswaService) Update(UUID uuid.UUID, input request.SiswaRequest) (models.Siswa, error) {
	siswa := models.Siswa{
		Name:      input.Name,
		Nisn:      input.Nisn,
		KelasUuid: input.KelasUuid,
	}

	if models.DB.Model(&siswa).Where("uuid = ?", UUID).Updates(&siswa).RowsAffected == 0 {
		return siswa, fmt.Errorf("gagal update data dengan UUID %d", UUID)
	}
	siswa.UUID = UUID
	return siswa, nil
}

func (s SiswaService) Delete(UUID uuid.UUID) error {
	siswa := models.Siswa{}
	if models.DB.Unscoped().Where("uuid = ?", UUID).Delete(&siswa).RowsAffected == 0 {
		return fmt.Errorf("Gagal hapus data dengan UUID %d", UUID)
	}
	return nil
}

func (s SiswaService) Show(UUID uuid.UUID) (models.Siswa, error) {
	var siswa models.Siswa
	if err := models.DB.Preload("Kelas").Preload("Nilai").Preload("Transformasi").Where("uuid = ?", UUID).First(&siswa).Error; err != nil {
		return siswa, err
	}

	return siswa, nil
}
