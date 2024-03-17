package main

import (
	"embed"
	"net/http"
	"os"

	"github.com/bit-wiz/data-store-a/app/queries"
	configs "github.com/bit-wiz/data-store-a/pkg/config"
	"github.com/bit-wiz/data-store-a/pkg/middleware"
	"github.com/bit-wiz/data-store-a/pkg/routes"
	"github.com/bit-wiz/data-store-a/pkg/utils"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joho/godotenv"
)

//go:embed frontend/dist
var Frontend embed.FS

func main() {
	godotenv.Load()

	err := queries.NewMongo("citizen")
	if err != nil {
		panic(err)
	}

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(Frontend),
		PathPrefix: "frontend/dist",
		Browse:     true,
	}))

	routes.PublicRoutes(app)

	if os.Getenv("MODE") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
