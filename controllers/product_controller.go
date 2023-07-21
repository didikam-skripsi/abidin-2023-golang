package controllers

import (
	helper "gostarter-backend/helpers"
	"gostarter-backend/models"
	"gostarter-backend/request"
	"gostarter-backend/response"
	"gostarter-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ProductController struct {
	productService  services.ProductService
	productResponse response.ProductResponse
}

func (this ProductController) GetPostPaginate(c *gin.Context) {
	var products []models.Product
	var totalRecords int64

	page := c.DefaultQuery("page", "1")
	perPage := c.DefaultQuery("per_page", "10")
	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)
	models.DB.Model(&models.Product{}).Count(&totalRecords)
	offset := (pageInt - 1) * perPageInt
	query := models.DB.Limit(perPageInt).Offset(offset)
	if err := query.Find(&products).Error; err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Product find failed", nil)
		return
	}
	response := gin.H{
		"current_page": pageInt,
		"per_page":     perPageInt,
		"total":        totalRecords,
		"data":         this.productResponse.Collections(products),
	}

	c.JSON(http.StatusOK, response)
}

func (this ProductController) Index(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": this.productResponse.Collections(products)})
}

func (this ProductController) Store(c *gin.Context) {
	var input request.ProductRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		response.APIResponse(c, http.StatusUnprocessableEntity, "Input field invalid", errors)
		return
	}

	product, err := this.productService.Store(input)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Failed to create product", err)
		return
	}
	response.APIResponse(c, http.StatusOK, "Product created successfully", this.productResponse.Response(product))
	return
}

func (this ProductController) Show(c *gin.Context) {
	var input request.ProductRequestID

	if err := c.ShouldBindUri(&input); err != nil {
		errors := helper.FormatValidationError(err)
		response.APIResponse(c, http.StatusUnprocessableEntity, "Input field invalid", errors)
		return
	}

	product, err := this.productService.Show(input.ID)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response.APIResponse(c, http.StatusNotFound, "Data not found", err)
			return
		default:
			response.APIResponse(c, http.StatusInternalServerError, "Failed to show product", err)
			return
		}
	}
	response.APIResponse(c, http.StatusOK, "Product show successfully", this.productResponse.Response(product))
	return
}

func (this ProductController) Update(c *gin.Context) {
	var input request.ProductRequest
	var inputID request.ProductRequestID

	if err := c.ShouldBindUri(&inputID); err != nil {
		errors := helper.FormatValidationError(err)
		response.APIResponse(c, http.StatusUnprocessableEntity, "ID field invalid", errors)
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		response.APIResponse(c, http.StatusUnprocessableEntity, "Input field invalid", errors)
		return
	}

	product, err := this.productService.Show(inputID.ID)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response.APIResponse(c, http.StatusNotFound, "Data not found", err)
			return
		default:
			response.APIResponse(c, http.StatusInternalServerError, "Failed to find product", err)
			return
		}
	}

	product, err = this.productService.Update(inputID.ID, input)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Failed to update product", err)
		return
	}

	response.APIResponse(c, http.StatusOK, "Product updated successfully", this.productResponse.Response(product))
	return
}

func (this ProductController) Delete(c *gin.Context) {
	var inputID request.ProductRequestID

	if err := c.ShouldBindUri(&inputID); err != nil {
		errors := helper.FormatValidationError(err)
		response.APIResponse(c, http.StatusUnprocessableEntity, "ID field invalid", errors)
		return
	}

	product, err := this.productService.Show(inputID.ID)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response.APIResponse(c, http.StatusNotFound, "Data not found", err)
			return
		default:
			response.APIResponse(c, http.StatusInternalServerError, "Failed to find product", err)
			return
		}
	}

	err = this.productService.Delete(inputID.ID)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Failed to delete product", err)
		return
	}

	response.APIResponse(c, http.StatusOK, "Product delete successfully", this.productResponse.Response(product))
	return
}
