package controller

import (
	"errors"
	"fmt"
	"net/http"

	modelBursaKerja "github.com/Fatwaff/be_bursa-kerja/model"
	moduleBursaKerja "github.com/Fatwaff/be_bursa-kerja/module"
	"github.com/Fatwaff/ws-fatwa/config"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllLowongan(c *fiber.Ctx) error {
	var data []modelBursaKerja.Lowongan
	hasil := moduleBursaKerja.GetAllDocs(config.BursaKerjamongoconn, "lowongan", data)
	return c.JSON(hasil)
}

func GetLowonganFromID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	data, err := moduleBursaKerja.GetLowonganFromID(objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(data)
}