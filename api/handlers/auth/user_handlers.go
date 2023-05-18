package api

import (
	"nom/services"

	"github.com/gofiber/fiber/v2"
)

func UserLogin(c *fiber.Ctx, userService *services.AuthService) error {
	return c.SendString("Welcome to an Awesome API")
}
