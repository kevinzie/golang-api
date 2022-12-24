package routes

import (
	"github.com/gofiber/fiber/v2"
	"project/scratch/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Get("/users", controllers.GetBooks)
}
