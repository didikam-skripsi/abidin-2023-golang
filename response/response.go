package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func APIResponse(context *gin.Context, code int, message string, data interface{}) {
	jsonResponse := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	context.JSON(code, jsonResponse)
}
