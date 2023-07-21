package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type HomeController struct{}

func (con HomeController) Index(c *fiber.Ctx) error {
	data := fiber.Map{
		"title": "hello world gofiber",
	}
	return c.JSON(data)
}
