package main

import (
	"os"

	configs "github.com/bit-wiz/data-store-a/pkg/config"
	"github.com/bit-wiz/data-store-a/pkg/middleware"
	"github.com/bit-wiz/data-store-a/pkg/routes"
	"github.com/bit-wiz/data-store-a/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)

	if os.Getenv("MODE") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
