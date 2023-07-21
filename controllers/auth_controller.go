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
)

type AuthController struct {
	userService  services.UserService
	userResponse response.UserResponse
}

func (this AuthController) Register(c *fiber.Ctx) error {
	var input request.RegisterRequest

	c.BodyParser(&input)
	if errs := helper.Validate(input); len(errs) > 0 && errs[0].Error {
		return response.APIResponse(c, http.StatusBadRequest, "Input field invalid", errs)
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
	if errs := helper.Validate(input); len(errs) > 0 && errs[0].Error {
		return response.APIResponse(c, http.StatusBadRequest, "Input field invalid", errs)
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

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		response.APIResponse(c, http.StatusBadRequest, "Token failed", err.Error())
		return nil
	}
	user, err := this.userService.GetUserByID(user_id)
	if err != nil {
		response.APIResponse(c, http.StatusBadRequest, "Get user failed", nil)
		return nil
	}

	response.APIResponse(c, http.StatusOK, "Show data successfully", this.userResponse.Response(user))
	return nil
}
