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
	if err := query.Find(&products).Error; err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Product find failed", nil)
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

	product, err := this.productService.Store(input)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Failed to create product", err)
		return nil
	}
	response.APIResponse(c, http.StatusOK, "Product created successfully", this.productResponse.Response(product))
	return nil
}

func (this ProductController) Show(c *fiber.Ctx) error {
	input_id := c.Params("id")
	parse_id, err := strconv.ParseUint(input_id, 10, 64)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "Invalid ID", err)
	}
	ID := uint(parse_id)
	product, err := this.productService.Show(ID)

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
	input_id := c.Params("id")
	parse_id, err := strconv.Atoi(input_id)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "Invalid ID", err)
	}
	inputID := uint(parse_id)

	c.BodyParser(&input)
	if errs := helper.Validate(input); len(errs) > 0 && errs[0].Error {
		return response.APIResponse(c, http.StatusBadRequest, "Input field invalid", errs)
	}
	product, err := this.productService.Show(inputID)
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

	product, err = this.productService.Update(inputID, input)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Failed to update product", err)
		return nil
	}

	response.APIResponse(c, http.StatusOK, "Product updated successfully", this.productResponse.Response(product))
	return nil
}

func (this ProductController) Delete(c *fiber.Ctx) error {
	input_id := c.Params("id")
	parse_id, err := strconv.Atoi(input_id)
	if err != nil {
		return response.APIResponse(c, http.StatusBadRequest, "Invalid ID", err)
	}
	inputID := uint(parse_id)
	product, err := this.productService.Show(inputID)
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

	err = this.productService.Delete(inputID)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Failed to delete product", err)
		return nil
	}

	response.APIResponse(c, http.StatusOK, "Product delete successfully", this.productResponse.Response(product))
	return nil
}
