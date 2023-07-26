package services

import (
	"fmt"
	"gostarter-backend/models"
	"gostarter-backend/request"

	"github.com/google/uuid"
)

type TransformasiService struct{}

func (s TransformasiService) Store(input request.NilaiRequest) (models.Nilai, error) {
	nilai := models.Nilai{
		SiswaUuid: input.SiswaUuid,
		Uts:       input.Uts,
		Uas:       input.Uas,
		Tugas:     input.Tugas,
		Absen:     input.Absen,
		Sikap:     input.Sikap,
	}
	err := models.DB.Create(&nilai).Error
	if err != nil {
		return nilai, fmt.Errorf("gagal simpan data")
	}
	return nilai, nil
}

func (s TransformasiService) Update(input request.NilaiRequest) (models.Nilai, error) {
	nilai := models.Nilai{
		SiswaUuid: input.SiswaUuid,
		Uts:       input.Uts,
		Uas:       input.Uas,
		Tugas:     input.Tugas,
		Absen:     input.Absen,
		Sikap:     input.Sikap,
	}

	if models.DB.Model(&nilai).Where("siswa_uuid = ?", input.SiswaUuid).Updates(&nilai).RowsAffected == 0 {
		return nilai, fmt.Errorf("gagal update data dengan SiswaUuid %d", input.SiswaUuid)
	}
	return nilai, nil
}

func (s TransformasiService) Show(siswaUuid uuid.UUID) (models.Nilai, error) {
	var nilai models.Nilai
	if err := models.DB.Where("siswa_uuid = ?", siswaUuid).First(&nilai).Error; err != nil {
		return nilai, err
	}

	return nilai, nil
}

func (s TransformasiService) ShowSiswaNilai(siswaUuid uuid.UUID) (models.Siswa, error) {
	var siswa models.Siswa
	if err := models.DB.Preload("Nilai").Preload("Kelas").Where("uuid = ?", siswaUuid).First(&siswa).Error; err != nil {
		return siswa, err
	}
	return siswa, nil
}

func (s TransformasiService) CountBayes(nilai models.Nilai) error {
	err := s.UpdateTransformasi(&nilai)
	if err != nil {
		return err
	}
	var nilaiTransformasi models.Transformasi
	query := models.DB
	if err := query.Where("siswa_uuid = ?", nilai.SiswaUuid).First(&nilaiTransformasi).Error; err != nil {
		return fmt.Errorf("Gagal ambil data siswa")
	}
	bayesYa, err := getNilaiYa(nilaiTransformasi)
	bayesTidak, err := getNilaiTidak(nilaiTransformasi)
	if bayesYa > bayesTidak {
		nilai.Class = "Ya"
	} else {
		nilai.Class = "Tidak"
	}
	if err := models.DB.Save(&nilai).Error; err != nil {
		return fmt.Errorf("Gagal save nilai baru")
	}
	return nil
}

func getNilaiYa(nilaiTransformasi models.Transformasi) (float64, error) {
	var yaCount int64
	queryYa := models.DB.Model(&models.Transformasi{}).Where("class_name = ?", "Ya")
	if err := queryYa.Count(&yaCount).Error; err != nil {
		return 0, fmt.Errorf("Gagal count ya")
	}
	var yaCountUts int64
	queryYaUts := queryYa.Where("uts = ?", nilaiTransformasi.Uts)
	if err := queryYaUts.Count(&yaCountUts).Error; err != nil {
		return 0, fmt.Errorf("Gagal count uts ya")
	}
	var yaCountUas int64
	queryYaUas := queryYa.Where("uas = ?", nilaiTransformasi.Uas)
	if err := queryYaUas.Count(&yaCountUas).Error; err != nil {
		return 0, fmt.Errorf("Gagal count uas ya")
	}
	var yaCountTugas int64
	queryYaTugas := queryYa.Where("tugas = ?", nilaiTransformasi.Tugas)
	if err := queryYaTugas.Count(&yaCountTugas).Error; err != nil {
		return 0, fmt.Errorf("Gagal count tugas ya")
	}
	var yaCountAbsen int64
	queryYaAbsen := queryYa.Where("tugas = ?", nilaiTransformasi.Tugas)
	if err := queryYaAbsen.Count(&yaCountAbsen).Error; err != nil {
		return 0, fmt.Errorf("Gagal count absen ya")
	}
	var yaCountSikap int64
	queryYaSikap := queryYa.Where("sikap = ?", nilaiTransformasi.Sikap)
	if err := queryYaSikap.Count(&yaCountSikap).Error; err != nil {
		return 0, fmt.Errorf("Gagal count sikap ya")
	}
	floatYaCount := float64(yaCount)
	bayesYa := (float64(yaCountUts) / floatYaCount) * (float64(yaCountUas) / floatYaCount) * (float64(yaCountTugas) / floatYaCount) * (float64(yaCountAbsen) / floatYaCount) * (float64(yaCountSikap) / floatYaCount)
	return bayesYa, nil
}

func getNilaiTidak(nilaiTransformasi models.Transformasi) (float64, error) {
	var tidakCount int64
	queryTidak := models.DB.Model(&models.Transformasi{}).Where("class_name = ?", "Tidak")
	if err := queryTidak.Count(&tidakCount).Error; err != nil {
		return 0, fmt.Errorf("Gagal count tidak")
	}
	var tidakCountUts int64
	queryTidakUts := queryTidak.Where("uts = ?", nilaiTransformasi.Uts)
	if err := queryTidakUts.Count(&tidakCountUts).Error; err != nil {
		return 0, fmt.Errorf("Gagal count uts tidak")
	}
	var tidakCountUas int64
	queryTidakUas := queryTidak.Where("uas = ?", nilaiTransformasi.Uas)
	if err := queryTidakUas.Count(&tidakCountUas).Error; err != nil {
		return 0, fmt.Errorf("Gagal count uas tidak")
	}
	var tidakCountTugas int64
	queryTidakTugas := queryTidak.Where("tugas = ?", nilaiTransformasi.Tugas)
	if err := queryTidakTugas.Count(&tidakCountTugas).Error; err != nil {
		return 0, fmt.Errorf("Gagal count tugas tidak")
	}
	var tidakCountAbsen int64
	queryTidakAbsen := queryTidak.Where("tugas = ?", nilaiTransformasi.Tugas)
	if err := queryTidakAbsen.Count(&tidakCountAbsen).Error; err != nil {
		return 0, fmt.Errorf("Gagal count absen tidak")
	}
	var tidakCountSikap int64
	queryTidakSikap := queryTidak.Where("sikap = ?", nilaiTransformasi.Sikap)
	if err := queryTidakSikap.Count(&tidakCountSikap).Error; err != nil {
		return 0, fmt.Errorf("Gagal count sikap tidak")
	}
	floattidakCount := float64(tidakCount)
	bayesTidak := (float64(tidakCountUts) / floattidakCount) * (float64(tidakCountUas) / floattidakCount) * (float64(tidakCountTugas) / floattidakCount) * (float64(tidakCountAbsen) / floattidakCount) * (float64(tidakCountSikap) / floattidakCount)
	return bayesTidak, nil
}

func (s TransformasiService) GenerateTransformasi() error {
	var siswas []models.Siswa
	query := models.DB
	if err := query.Preload("Nilai").Preload("Nilai.Transformasi").Order("id DESC").Find(&siswas).Error; err != nil {
		return fmt.Errorf("Gagal ambil data siswa")
	}
	for _, siswa := range siswas {
		if siswa.Nilai != nil {
			err := s.UpdateTransformasi(siswa.Nilai)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s TransformasiService) UpdateTransformasi(nilai *models.Nilai) error {
	uts_uuid, uts_name := getUTS(nilai)
	uas_uuid, uas_name := getUAS(nilai)
	tugas_uuid, tugas_name := getTugas(nilai)
	absen_uuid, absen_name := getAbsen(nilai)
	sikap_uuid, sikap_name := getSikap(nilai)
	class_uuid, class_name := getClass(nilai)
	tranformasi := models.Transformasi{
		SiswaUuid: nilai.SiswaUuid,
		Uts:       uts_uuid,
		UtsName:   uts_name,
		Uas:       uas_uuid,
		UasName:   uas_name,
		Tugas:     tugas_uuid,
		TugasName: tugas_name,
		Absen:     absen_uuid,
		AbsenName: absen_name,
		Sikap:     sikap_uuid,
		SikapName: sikap_name,
		Class:     class_uuid,
		ClassName: class_name,
	}
	if nilai.Transformasi != nil {
		if models.DB.Model(&tranformasi).Where("siswa_uuid = ?", nilai.SiswaUuid).Updates(&tranformasi).RowsAffected == 0 {
			return fmt.Errorf("gagal update transformasi dengan SiswaUuid %d", nilai.SiswaUuid)
		}
	} else {
		err := models.DB.Create(&tranformasi).Error
		if err != nil {
			return fmt.Errorf("Gagal tambah nilai transformasi")
		}
	}
	return nil
}

func getUTS(nilai *models.Nilai) (*uuid.UUID, string) {
	var uid uuid.UUID
	var attribute models.Attribute
	query := models.DB.Where("scope = ?", "uts")
	query = query.Where("range_start <= ?", nilai.Uts)
	query = query.Where("range_end >= ?", nilai.Uts)
	if err := query.First(&attribute).Error; err != nil {
		return &uid, ""
	}
	return (*uuid.UUID)(&attribute.UUID), attribute.Name
}

func getUAS(nilai *models.Nilai) (*uuid.UUID, string) {
	var uid uuid.UUID
	var attribute models.Attribute
	query := models.DB.Where("scope = ?", "uas")
	query = query.Where("range_start <= ?", nilai.Uas)
	query = query.Where("range_end >= ?", nilai.Uas)
	if err := query.First(&attribute).Error; err != nil {
		return &uid, ""
	}
	return (*uuid.UUID)(&attribute.UUID), attribute.Name
}

func getTugas(nilai *models.Nilai) (*uuid.UUID, string) {
	var uid uuid.UUID
	var attribute models.Attribute
	query := models.DB.Where("scope = ?", "tugas")
	query = query.Where("range_start <= ?", nilai.Tugas)
	query = query.Where("range_end >= ?", nilai.Tugas)
	if err := query.First(&attribute).Error; err != nil {
		return &uid, ""
	}
	return (*uuid.UUID)(&attribute.UUID), attribute.Name
}

func getAbsen(nilai *models.Nilai) (*uuid.UUID, string) {
	var uid uuid.UUID
	var attribute models.Attribute
	query := models.DB.Where("scope = ?", "absen")
	query = query.Where("range_start <= ?", nilai.Absen)
	query = query.Where("range_end >= ?", nilai.Absen)
	if err := query.First(&attribute).Error; err != nil {
		return &uid, ""
	}
	return (*uuid.UUID)(&attribute.UUID), attribute.Name
}

func getSikap(nilai *models.Nilai) (*uuid.UUID, string) {
	var uid uuid.UUID
	var attribute models.Attribute
	query := models.DB.Where("scope = ?", "sikap")
	query = query.Where("value = ?", nilai.Sikap)
	if err := query.First(&attribute).Error; err != nil {
		return &uid, ""
	}
	return (*uuid.UUID)(&attribute.UUID), attribute.Name
}

func getClass(nilai *models.Nilai) (*uuid.UUID, string) {
	var uid uuid.UUID
	var attribute models.Attribute
	query := models.DB.Where("scope = ?", "class")
	query = query.Where("value = ?", nilai.Class)
	if err := query.First(&attribute).Error; err != nil {
		return &uid, ""
	}
	return (*uuid.UUID)(&attribute.UUID), attribute.Name
}
