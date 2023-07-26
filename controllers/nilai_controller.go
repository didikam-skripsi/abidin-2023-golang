package controllers

import (
	helper "gostarter-backend/helpers"
	"gostarter-backend/models"
	"gostarter-backend/request"
	"gostarter-backend/response"
	"gostarter-backend/services"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type NilaiController struct {
	transformasiService services.TransformasiService
	siswaService        services.SiswaService
	siswaResponse       response.SiswaResponse
	nilaiResponse       response.NilaiResponse
	nilaiService        services.NilaiService
}

func NewNilaiController() NilaiController {
	return NilaiController{}
}

func (this NilaiController) GetPaginate(c *fiber.Ctx) error {
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
	if err := query.Preload("Kelas").Preload("Nilai").Order("id ASC").Find(&siswas).Error; err != nil {
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

func (this NilaiController) Store(c *fiber.Ctx) error {
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

	siswa, err := this.siswaService.Show(input.SiswaUuid)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response.APIResponse(c, http.StatusNotFound, "Siswa tidak ditemukan", err)
			return nil
		default:
			response.APIResponse(c, http.StatusInternalServerError, "Gagal ambil data siswa", err)
			return nil
		}
	}

	err = this.transformasiService.UpdateTransformasi(siswa.Nilai)
	if err != nil {
		return response.APIResponse(c, http.StatusNotFound, "Gagal update transformasi", err)
	}

	response.APIResponse(c, http.StatusOK, "Berhasil tambah data", this.nilaiResponse.Response(&nilai))
	return nil
}

func (this NilaiController) Show(c *fiber.Ctx) error {
	input_uuid := c.Params("uuid")
	parse_uuid, err := uuid.Parse(input_uuid)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "UUID tidak valid", err)
	}
	siswa, err := this.nilaiService.ShowSiswaNilai(parse_uuid)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response.APIResponse(c, http.StatusNotFound, "Data tidak ditemukan", err)
			return nil
		default:
			response.APIResponse(c, http.StatusInternalServerError, "Gagal ambil data", err)
			return nil
		}
	}
	response.APIResponse(c, http.StatusOK, "Berhasil ambil data", this.siswaResponse.Response(siswa))
	return nil
}
