package controllers

import (
	"gostarter-backend/models"
	"gostarter-backend/response"

	"github.com/gofiber/fiber/v2"
)

type KelasController struct {
	siswaResponse response.SiswaResponse
	kelasResponse response.KelasResponse
}

func (this KelasController) Index(c *fiber.Ctx) error {
	var kelas []models.Kelas
	models.DB.Find(&kelas)
	c.JSON(fiber.Map{"datas": this.kelasResponse.Collections(kelas)})
	return nil
}
