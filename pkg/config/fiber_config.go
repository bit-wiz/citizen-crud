package configs

import (
	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {

	return fiber.Config{
		ServerHeader: "API",
		AppName:      "App V1",
	}
}
