package controller

import (
	modelBursaKerja "github.com/Fatwaff/be_bursa-kerja/model"
	moduleBursaKerja "github.com/Fatwaff/be_bursa-kerja/module"
	"github.com/Fatwaff/ws-fatwa/config"
	"github.com/gofiber/fiber/v2"
)

func GetAllLowongan(c *fiber.Ctx) error {
	var data []modelBursaKerja.Lowongan
	hasil := moduleBursaKerja.GetAllDocs(config.BursaKerjamongoconn, "lowongan", data)
	return c.JSON(hasil)
}