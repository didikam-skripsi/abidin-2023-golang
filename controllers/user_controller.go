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

type UserController struct {
	UserService  services.UserService
	userResponse response.UserResponse
}

func (this UserController) GetPostPaginate(c *fiber.Ctx) error {
	var users []models.User
	var totalRecords int64

	page := c.Query("page", "1")
	perPage := c.Query("per_page", "10")
	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)
	models.DB.Model(&models.User{}).Count(&totalRecords)
	offset := (pageInt - 1) * perPageInt
	query := models.DB.Limit(perPageInt).Offset(offset)
	if err := query.Order("id DESC").Find(&users).Error; err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Data find failed", err.Error())
		return nil
	}
	response := fiber.Map{
		"current_page": pageInt,
		"per_page":     perPageInt,
		"total":        totalRecords,
		"data":         this.userResponse.Collections(users),
	}

	return c.JSON(response)
}

func (this UserController) Show(c *fiber.Ctx) error {
	input_uuid := c.Params("uuid")
	parse_uuid, err := uuid.Parse(input_uuid)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "Invalid ID", err)
	}
	user, err := this.UserService.Show(parse_uuid)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response.APIResponse(c, http.StatusNotFound, "Data not found", err)
			return nil
		default:
			response.APIResponse(c, http.StatusInternalServerError, "Failed to show data", err)
			return nil
		}
	}
	response.APIResponse(c, http.StatusOK, "Data show successfully", this.userResponse.Response(&user))
	return nil
}

func (this UserController) Store(c *fiber.Ctx) error {
	var input request.UserCreateRequest
	c.BodyParser(&input)
	validator := helper.NewValidator()
	if errs := validator.Validate(input); len(errs) > 0 && errs[0].Error {
		return response.APIResponse(c, http.StatusUnprocessableEntity, "Input field invalid", errs)
	}
	exist, err := this.UserService.IsExists(input.Username)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Failed to get data", err.Error())
		return nil
	}
	if exist {
		response.APIResponse(c, http.StatusBadRequest, "Failed, username is exist", nil)
		return nil
	}
	data, err := this.UserService.Store(input)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Failed to create data", err.Error())
		return nil
	}
	response.APIResponse(c, http.StatusOK, "Data created successfully", this.userResponse.Response(&data))
	return nil
}

func (this UserController) Update(c *fiber.Ctx) error {
	var input request.UserUpdateRequest
	input_uuid := c.Params("uuid")
	parse_uuid, err := uuid.Parse(input_uuid)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "Invalid ID", err)
	}
	c.BodyParser(&input)
	validator := helper.NewValidator()
	if errs := validator.Validate(input); len(errs) > 0 && errs[0].Error {
		return response.APIResponse(c, http.StatusUnprocessableEntity, "Input field invalid", errs)
	}
	data, err := this.UserService.Show(parse_uuid)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response.APIResponse(c, http.StatusNotFound, "Data not found", err)
			return nil
		default:
			response.APIResponse(c, http.StatusInternalServerError, "Failed to find data", err)
			return nil
		}
	}
	data, err = this.UserService.Update(parse_uuid, input)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Failed to update data", err)
		return nil
	}
	response.APIResponse(c, http.StatusOK, "Data updated successfully", this.userResponse.Response(&data))
	return nil
}

func (this UserController) Delete(c *fiber.Ctx) error {
	input_uuid := c.Params("uuid")
	parse_uuid, err := uuid.Parse(input_uuid)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "Invalid UUID", err)
	}
	data, err := this.UserService.Show(parse_uuid)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response.APIResponse(c, http.StatusNotFound, "Data not found", err)
			return nil
		default:
			response.APIResponse(c, http.StatusInternalServerError, "Failed to find data", err)
			return nil
		}
	}

	err = this.UserService.Delete(parse_uuid)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Failed to delete data", err)
		return nil
	}

	response.APIResponse(c, http.StatusOK, "Data delete successfully", this.userResponse.Response(&data))
	return nil
}
