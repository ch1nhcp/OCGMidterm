package controller

import (
	"github.com/TechMaster/golang/08Fiber/Repository/database"
	"github.com/TechMaster/golang/08Fiber/Repository/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllFilm(c *fiber.Ctx) error {
	var film []models.Film
	database.DB.Find(&film)
	return c.JSON(film)
}
