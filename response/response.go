package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func APIResponse(c *fiber.Ctx, code int, message string, data interface{}) error {
	jsonResponse := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	return c.Status(code).JSON(jsonResponse)
}
