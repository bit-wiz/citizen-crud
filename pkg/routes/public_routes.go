package routes

import (
	"github.com/bit-wiz/data-store-a/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api")

	// citizen routes
	route.Get("/", controllers.Root)
	route.Get("/fields", controllers.GetFields)
	route.Get("/allcitizen", controllers.GetAllCitizens)
	route.Get("/citizen/:id", controllers.GetCitizen)
	route.Post("/citizen", controllers.CreateCitizen)
	route.Patch("/citizen/:id", controllers.UpdateCitizen)
	route.Delete("/citizen/:id", controllers.DeleteCitizen)
}
