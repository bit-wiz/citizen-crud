package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Root(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"msg": "hello ^0^",
	})
}
