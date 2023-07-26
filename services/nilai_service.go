package services

import (
	"fmt"
	"gostarter-backend/models"
	"gostarter-backend/request"

	"github.com/google/uuid"
)

type NilaiService struct{}

func (s NilaiService) Store(input request.NilaiRequest) (models.Nilai, error) {
	nilai := models.Nilai{
		SiswaUuid: input.SiswaUuid,
		Uts:       input.Uts,
		Uas:       input.Uas,
		Tugas:     input.Tugas,
		Absen:     input.Absen,
		Sikap:     input.Sikap,
		Class:     input.Class,
	}
	err := models.DB.Create(&nilai).Error
	if err != nil {
		return nilai, fmt.Errorf("gagal simpan data")
	}
	return nilai, nil
}

func (s NilaiService) Update(input request.NilaiRequest) (models.Nilai, error) {
	nilai := models.Nilai{
		SiswaUuid: input.SiswaUuid,
		Uts:       input.Uts,
		Uas:       input.Uas,
		Tugas:     input.Tugas,
		Absen:     input.Absen,
		Sikap:     input.Sikap,
		Class:     input.Class,
	}
	if input.Class == "" {
		nilai.Class = "-"
	}
	if models.DB.Model(&nilai).Where("siswa_uuid = ?", input.SiswaUuid).Updates(&nilai).RowsAffected == 0 {
		return nilai, fmt.Errorf("gagal update data dengan SiswaUuid %d", input.SiswaUuid)
	}
	return nilai, nil
}

func (s NilaiService) Show(siswaUuid uuid.UUID) (models.Nilai, error) {
	var nilai models.Nilai
	if err := models.DB.Where("siswa_uuid = ?", siswaUuid).First(&nilai).Error; err != nil {
		return nilai, err
	}

	return nilai, nil
}

func (s NilaiService) ShowSiswaNilai(siswaUuid uuid.UUID) (models.Siswa, error) {
	var siswa models.Siswa
	if err := models.DB.Preload("Nilai").Preload("Kelas").Where("uuid = ?", siswaUuid).First(&siswa).Error; err != nil {
		return siswa, err
	}
	return siswa, nil
}
