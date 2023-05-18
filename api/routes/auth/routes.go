package api

import (
	"nom/services"

	userHander "nom/api/handlers/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRouts(api *fiber.App, service *services.AuthService) {
	api.Get("/login", func(c *fiber.Ctx) error {
		return userHander.UserLogin(c, service)
	})
}
