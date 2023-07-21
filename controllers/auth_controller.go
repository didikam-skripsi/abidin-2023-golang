package controllers

import (
	helper "gostarter-backend/helpers"
	"gostarter-backend/helpers/token"
	"gostarter-backend/request"
	"gostarter-backend/response"
	"gostarter-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService  services.UserService
	userResponse response.UserResponse
}

func (this AuthController) Register(c *gin.Context) {
	var input request.RegisterRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		response.APIResponse(c, http.StatusBadRequest, "Input field invalid", errors)
		return
	}

	user, err := this.userService.Register(input)

	if err != nil {
		response.APIResponse(c, http.StatusInternalServerError, "Register failed", err.Error())
		return
	}

	response.APIResponse(c, http.StatusOK, "Registration successfully", this.userResponse.Response(user))

}

func (this AuthController) Login(c *gin.Context) {
	var input request.LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		response.APIResponse(c, http.StatusBadRequest, "Input field invalid", errors)
		return
	}

	token, err := this.userService.LoginCheck(input)

	if err != nil {
		response.APIResponse(c, http.StatusBadRequest, "username or password is incorrect", nil)
		return
	}

	response.APIResponse(c, http.StatusOK, "Login successfully", gin.H{"token": token})
	return

}

func (this AuthController) CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		response.APIResponse(c, http.StatusBadRequest, "Token failed", err.Error())
		return
	}
	user, err := this.userService.GetUserByID(user_id)
	if err != nil {
		response.APIResponse(c, http.StatusBadRequest, "Get user failed", nil)
		return
	}

	response.APIResponse(c, http.StatusOK, "Show data successfully", this.userResponse.Response(user))
	return
}
