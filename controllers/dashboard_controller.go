package controllers

import (
	"fmt"
	"gostarter-backend/models"

	"github.com/gofiber/fiber/v2"
)

type DashboardController struct{}

func NewsDashboardController() DashboardController {
	return DashboardController{}
}

func (this DashboardController) Donut(c *fiber.Ctx) error {
	var yaCount int64
	queryYa := models.DB.Model(&models.Transformasi{}).Where("class_name = ?", "Ya")
	if err := queryYa.Count(&yaCount).Error; err != nil {
		return fmt.Errorf("Gagal count ya")
	}
	var tidakCount int64
	queryTidak := models.DB.Model(&models.Transformasi{}).Where("class_name = ?", "Tidak")
	if err := queryTidak.Count(&tidakCount).Error; err != nil {
		return fmt.Errorf("Gagal count tidak")
	}
	response := []int64{yaCount, tidakCount}
	return c.JSON(response)
}

type KelasSummary struct {
	Name  string `gorm:"column:kelas_name"`
	Count int64  `gorm:"column:siswa_count"`
}

func (this DashboardController) Column(c *fiber.Ctx) error {
	var kelasSummaries []KelasSummary
	models.DB.Table("kelas").
		Select("kelas.name as kelas_name, COUNT(siswas.uuid) as siswa_count").
		Joins("LEFT JOIN siswas ON kelas.uuid = siswas.kelas_uuid").
		Group("kelas.uuid").
		Scan(&kelasSummaries)

	categories := make([]string, len(kelasSummaries))
	series := make([]int64, len(kelasSummaries))
	for i, k := range kelasSummaries {
		categories[i] = k.Name
		series[i] = k.Count
	}
	response := fiber.Map{
		"categories": categories,
		"series":     series,
	}
	return c.JSON(response)
}
