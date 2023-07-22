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

type ProductController struct {
	productService  services.ProductService
	productResponse response.ProductResponse
}

func (this ProductController) GetPostPaginate(c *fiber.Ctx) error {
	var products []models.Product
	var totalRecords int64

	page := c.Query("page", "1")
	perPage := c.Query("per_page", "10")
	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)
	models.DB.Model(&models.Product{}).Count(&totalRecords)
	offset := (pageInt - 1) * perPageInt
	query := models.DB.Limit(perPageInt).Offset(offset)
	if err := query.Preload("User").Find(&products).Error; err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Product find failed", err.Error())
		return nil
	}
	response := fiber.Map{
		"current_page": pageInt,
		"per_page":     perPageInt,
		"total":        totalRecords,
		"data":         this.productResponse.Collections(products),
	}

	return c.JSON(response)
}

func (this ProductController) Index(c *fiber.Ctx) error {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(fiber.Map{"products": this.productResponse.Collections(products)})
	return nil
}

func (this ProductController) Store(c *fiber.Ctx) error {
	var input request.ProductRequest

	c.BodyParser(&input)
	if errs := helper.Validate(input); len(errs) > 0 && errs[0].Error {
		return response.APIResponse(c, http.StatusBadRequest, "Input field invalid", errs)
	}
	authUser, err := token.ExtractTokenUser(c)
	if err != nil {
		response.APIResponse(c, http.StatusUnauthorized, "Token failed", err.Error())
		return nil
	}
	product, err := this.productService.Store(authUser.UUID, input)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Failed to create product", err)
		return nil
	}
	response.APIResponse(c, http.StatusOK, "Product created successfully", this.productResponse.Response(product))
	return nil
}

func (this ProductController) Show(c *fiber.Ctx) error {
	input_uuid := c.Params("uuid")
	parse_uuid, err := uuid.Parse(input_uuid)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "Invalid ID", err)
	}
	product, err := this.productService.Show(parse_uuid)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response.APIResponse(c, http.StatusNotFound, "Data not found", err)
			return nil
		default:
			response.APIResponse(c, http.StatusInternalServerError, "Failed to show product", err)
			return nil
		}
	}
	response.APIResponse(c, http.StatusOK, "Product show successfully", this.productResponse.Response(product))
	return nil
}

func (this ProductController) Update(c *fiber.Ctx) error {
	var input request.ProductRequest
	input_uuid := c.Params("id")
	parse_uuid, err := uuid.Parse(input_uuid)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "Invalid ID", err)
	}

	c.BodyParser(&input)
	if errs := helper.Validate(input); len(errs) > 0 && errs[0].Error {
		return response.APIResponse(c, http.StatusBadRequest, "Input field invalid", errs)
	}
	product, err := this.productService.Show(parse_uuid)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response.APIResponse(c, http.StatusNotFound, "Data not found", err)
			return nil
		default:
			response.APIResponse(c, http.StatusInternalServerError, "Failed to find product", err)
			return nil
		}
	}

	product, err = this.productService.Update(parse_uuid, input)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Failed to update product", err)
		return nil
	}

	response.APIResponse(c, http.StatusOK, "Product updated successfully", this.productResponse.Response(product))
	return nil
}

func (this ProductController) Delete(c *fiber.Ctx) error {
	input_uuid := c.Params("uuid")
	parse_uuid, err := uuid.Parse(input_uuid)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "Invalid UUID", err)
	}
	product, err := this.productService.Show(parse_uuid)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response.APIResponse(c, http.StatusNotFound, "Data not found", err)
			return nil
		default:
			response.APIResponse(c, http.StatusInternalServerError, "Failed to find product", err)
			return nil
		}
	}

	err = this.productService.Delete(parse_uuid)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Failed to delete product", err)
		return nil
	}

	response.APIResponse(c, http.StatusOK, "Product delete successfully", this.productResponse.Response(product))
	return nil
}
