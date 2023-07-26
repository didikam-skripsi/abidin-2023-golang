package controllers

import (
	helper "gostarter-backend/helpers"
	"gostarter-backend/helpers/token"
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

type SiswaController struct {
	siswaResponse response.SiswaResponse
	siswaService  services.SiswaService
}

func (this SiswaController) GetPaginate(c *fiber.Ctx) error {
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
	if err := query.Preload("Kelas").Order("id ASC").Find(&siswas).Error; err != nil {
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

func (this SiswaController) Store(c *fiber.Ctx) error {
	var input request.SiswaRequest

	c.BodyParser(&input)
	validator := helper.NewValidator()
	if errs := validator.Validate(input); len(errs) > 0 && errs[0].Error {
		return response.APIResponse(c, http.StatusUnprocessableEntity, "Input tidak valid", errs)
	}
	authUser, err := token.ExtractTokenUser(c)
	if err != nil {
		response.APIResponse(c, http.StatusUnauthorized, "Token tidak valid", err.Error())
		return nil
	}
	siswa, err := this.siswaService.Store(authUser.UUID, input)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Gagal tambah data", err)
		return nil
	}
	response.APIResponse(c, http.StatusOK, "Berhasil tambah data", this.siswaResponse.Response(siswa))
	return nil
}

func (this SiswaController) Show(c *fiber.Ctx) error {
	input_uuid := c.Params("uuid")
	parse_uuid, err := uuid.Parse(input_uuid)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "UUID tidak valid", err)
	}
	siswa, err := this.siswaService.Show(parse_uuid)

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

func (this SiswaController) Update(c *fiber.Ctx) error {
	var input request.SiswaRequest
	input_uuid := c.Params("uuid")
	parse_uuid, err := uuid.Parse(input_uuid)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "UUID tidak valid", err)
	}
	c.BodyParser(&input)
	validator := helper.NewValidator()
	if errs := validator.Validate(input); len(errs) > 0 && errs[0].Error {
		return response.APIResponse(c, http.StatusUnprocessableEntity, "Input tidak valid", errs)
	}
	siswa, err := this.siswaService.Show(parse_uuid)
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

	siswa, err = this.siswaService.Update(parse_uuid, input)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Gagal update data", err)
		return nil
	}

	response.APIResponse(c, http.StatusOK, "Berhasil update data", this.siswaResponse.Response(siswa))
	return nil
}

func (this SiswaController) Delete(c *fiber.Ctx) error {
	input_uuid := c.Params("uuid")
	parse_uuid, err := uuid.Parse(input_uuid)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "UUID tidak valid", err)
	}
	siswa, err := this.siswaService.Show(parse_uuid)
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

	err = this.siswaService.Delete(parse_uuid)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Gagal hapus data", err)
		return nil
	}

	response.APIResponse(c, http.StatusOK, "Berhasil hapus data", this.siswaResponse.Response(siswa))
	return nil
}
