package main

import (
	"github.com/TechMaster/golang/08Fiber/Repository/database"
	"github.com/TechMaster/golang/08Fiber/Repository/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})
	database.Connect()
	app.Static("/", "./public")
	//ngăn truy cập / request từ cổng khác 3000
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)
	app.Listen(":3000")
}
