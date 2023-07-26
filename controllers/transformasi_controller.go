package controllers

import (
	helper "gostarter-backend/helpers"
	"gostarter-backend/models"
	"gostarter-backend/request"
	"gostarter-backend/response"
	"gostarter-backend/services"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type TransformasiController struct {
	siswaResponse       response.SiswaResponse
	nilaiResponse       response.NilaiResponse
	nilaiService        services.NilaiService
	transformasiService services.TransformasiService
}

type SiswaTransformasi struct {
	models.Siswa
	Uts int
	Uas int
}

func NewTransformasiController() TransformasiController {
	return TransformasiController{}
}

func (this TransformasiController) GetPaginate(c *fiber.Ctx) error {
	if os.Getenv("GENERATE_TF") == "true" {
		err := this.transformasiService.GenerateTransformasi()
		if err != nil {
			response.APIResponse(c, http.StatusInternalServerError, "Gagal generate transformasi", err)
			return nil
		}
	}
	var siswas []models.Siswa
	var totalRecords int64
	page := c.Query("page", "1")
	perPage := c.Query("per_page", "10")
	searchQuery := c.Query("search")
	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)
	query := models.DB
	if searchQuery != "" {
		query = query.Where("name LIKE ?", "%"+searchQuery+"%")
	}
	query.Model(&models.Siswa{}).Count(&totalRecords)
	offset := (pageInt - 1) * perPageInt
	query = query.Limit(perPageInt).Offset(offset)
	query = query.Preload("Transformasi").Preload("Kelas")
	if err := query.Order("id ASC").Find(&siswas).Error; err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Gagal ambil data", err.Error())
		return nil
	}
	response := fiber.Map{
		"current_page": pageInt,
		"per_page":     perPageInt,
		"total":        totalRecords,
		"data":         this.siswaResponse.Collections(siswas),
	}
	return c.JSON(response)
}

func (this TransformasiController) Bayes(c *fiber.Ctx) error {
	var nilais []models.Nilai
	var nilaisCount int64

	query := models.DB.Order("id ASC").Where("class IS NULL").Preload("Transformasi")
	query.Model(&models.Nilai{}).Count(&nilaisCount)
	if err := query.Find(&nilais).Error; err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Gagal ambil data", err.Error())
		return nil
	}
	for _, nilai := range nilais {
		err := this.transformasiService.CountBayes(nilai)
		if err != nil {
			response.APIResponse(c, http.StatusInternalServerError, "Gagal hitung bayes", err.Error())
			return nil
		}
	}
	err := this.transformasiService.GenerateTransformasi()
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Gagal generate transformasi", err)
		return nil
	}
	response.APIResponse(c, http.StatusOK, "Berhasil hitung naive bayes", nilaisCount)
	return nil
}

func (this TransformasiController) Store(c *fiber.Ctx) error {
	var input request.NilaiRequest

	c.BodyParser(&input)
	validator := helper.NewValidator()
	if errs := validator.Validate(input); len(errs) > 0 && errs[0].Error {
		return response.APIResponse(c, http.StatusUnprocessableEntity, "Input tidak valid", errs)
	}

	nilai, err := this.nilaiService.Show(input.SiswaUuid)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			nilai, err = this.nilaiService.Store(input)
			if err != nil {
				response.APIResponse(c, http.StatusInternalServerError, "Gagal tambah data", err)
				return nil
			}
		default:
			response.APIResponse(c, http.StatusInternalServerError, "Gagal ambil data", err)
			return nil
		}
	} else {
		nilai, err = this.nilaiService.Update(input)
		if err != nil {
			response.APIResponse(c, http.StatusInternalServerError, "Gagal update data", err)
			return nil
		}
	}
	response.APIResponse(c, http.StatusOK, "Berhasil tambah data", this.nilaiResponse.Response(&nilai))
	return nil
}
