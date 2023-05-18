package middleware

import (
	"nom/structs"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var secretKey = []byte("secret")

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{"message": "Missing authorization header"})
		}

		token, err := jwt.ParseWithClaims(authHeader, &structs.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.Status(401).JSON(fiber.Map{"message": "Invalid token signature"})
			}
			return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})

		}

		if claims, ok := token.Claims.(*structs.Claims); ok && token.Valid {
			c.Locals("username", claims.Username)
			return c.Next()
		}

		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}
}
