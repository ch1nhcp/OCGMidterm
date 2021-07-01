package routes

import (
	"github.com/TechMaster/golang/08Fiber/Repository/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/film", controller.GetAllFilm)
	// app.Post("/review/add", controller.CreateReview)
	// app.Get("/average/:id", controller.AverageRating)
	// app.Post("/review/delete/:id", controller.DeleteReviewById)
}
