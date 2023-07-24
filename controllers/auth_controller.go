package controllers

import (
	helper "gostarter-backend/helpers"
	"gostarter-backend/helpers/token"
	"gostarter-backend/request"
	"gostarter-backend/response"
	"gostarter-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type AuthController struct {
	userService  services.UserService
	userResponse response.UserResponse
}

func (this AuthController) Register(c *fiber.Ctx) error {
	var input request.RegisterRequest

	c.BodyParser(&input)
	validator := helper.NewValidator()
	if errs := validator.Validate(input); len(errs) > 0 && errs[0].Error {
		return response.APIResponse(c, http.StatusUnprocessableEntity, "Input field invalid", errs)
	}

	user, err := this.userService.Register(input)

	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Register failed", err.Error())
		return nil
	}

	response.APIResponse(c, http.StatusOK, "Registration successfully", this.userResponse.Response(user))
	return nil
}

func (this AuthController) Login(c *fiber.Ctx) error {
	var input request.LoginRequest
	c.BodyParser(&input)
	validator := helper.NewValidator()
	if errs := validator.Validate(input); len(errs) > 0 && errs[0].Error {
		return response.APIResponse(c, http.StatusUnprocessableEntity, "Input field invalid", errs)
	}

	token, err := this.userService.LoginCheck(input)

	if err != nil {
		response.APIResponse(c, http.StatusBadRequest, "username or password is incorrect", nil)
		return nil
	}

	response.APIResponse(c, http.StatusOK, "Login successfully", gin.H{"token": token})
	return nil
}

func (this AuthController) CurrentUser(c *fiber.Ctx) error {
	authUser, err := token.ExtractTokenUser(c)
	if err != nil {
		response.APIResponse(c, http.StatusUnauthorized, "Token failed", err.Error())
		return nil
	}
	user, err := this.userService.GetUserByUUID(authUser.UUID)
	if err != nil {
		response.APIResponse(c, http.StatusBadRequest, "Get user failed", nil)
		return nil
	}

	response.APIResponse(c, http.StatusOK, "Show data successfully", this.userResponse.Response(user))
	return nil
}

func (this AuthController) Profile(c *fiber.Ctx) error {
	var input request.ProfileUpdateRequest
	authUser, err := token.ExtractTokenUser(c)
	if err != nil {
		response.APIResponse(c, http.StatusUnauthorized, "Profil tidak ditemukan", err.Error())
		return nil
	}
	c.BodyParser(&input)
	validator := helper.NewValidator()
	if errs := validator.Validate(input); len(errs) > 0 && errs[0].Error {
		return response.APIResponse(c, http.StatusUnprocessableEntity, "Input tidak valid", errs)
	}
	data, err := this.userService.Show(authUser.UUID)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response.APIResponse(c, http.StatusNotFound, "Data tidak ditemukan", err)
			return nil
		default:
			response.APIResponse(c, http.StatusInternalServerError, "Gagal menemukan data", err)
			return nil
		}
	}
	data, err = this.userService.UpdateProfile(authUser.UUID, input)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Gagal update data", err)
		return nil
	}

	tokenData, err := token.GenerateToken(data)
	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Gagal update token", err)
		return nil
	}
	response.APIResponse(c, http.StatusOK, "Berhasil update data", gin.H{"token": tokenData})
	return nil
}
