package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	// Call the next handler in the chain
	err := c.Next()

	// log the request dureation
	duration := time.Since(start)
	fmt.Printf("[%s] %s %s\n", c.Method(), c.Path(), duration)

	return err
}
