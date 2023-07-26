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

type AttributeController struct {
	attributeResponse response.AttributeResponse
	attributeService  services.AttributeService
}

func NewAttributeController() AttributeController {
	return AttributeController{}
}

func (this AttributeController) GetPaginate(c *fiber.Ctx) error {
	var attributes []models.Attribute
	var totalRecords int64

	page := c.Query("page", "1")
	perPage := c.Query("per_page", "10")
	searchQuery := c.Query("search")
	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)
	query := models.DB
	if searchQuery != "" {
		query = query.Where("scope_name LIKE ?", "%"+searchQuery+"%")
	}
	query.Model(&models.Attribute{}).Count(&totalRecords)
	offset := (pageInt - 1) * perPageInt
	query = query.Limit(perPageInt).Offset(offset)
	if err := query.Order("id ASC").Find(&attributes).Error; err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Gagal ambil data", err.Error())
		return nil
	}
	response := fiber.Map{
		"current_page": pageInt,
		"per_page":     perPageInt,
		"total":        totalRecords,
		"data":         this.attributeResponse.Collections(attributes),
	}

	return c.JSON(response)
}

func (this AttributeController) Show(c *fiber.Ctx) error {
	input_uuid := c.Params("uuid")
	parse_uuid, err := uuid.Parse(input_uuid)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "UUID tidak valid", err)
	}
	attribute, err := this.attributeService.Show(parse_uuid)
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
	response.APIResponse(c, http.StatusOK, "Berhasil ambil data", this.attributeResponse.Response(attribute))
	return nil
}

func (this AttributeController) Update(c *fiber.Ctx) error {
	var input request.AttributeRequest
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
	attribute, err := this.attributeService.Show(parse_uuid)
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

	attribute, err = this.attributeService.Update(parse_uuid, input)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Gagal update data", err)
		return nil
	}

	response.APIResponse(c, http.StatusOK, "Berhasil update data", this.attributeResponse.Response(attribute))
	return nil
}
