package routes

import (
	"github.com/bit-wiz/data-store-a/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api")

	route.Get("/", controllers.Root)
}
