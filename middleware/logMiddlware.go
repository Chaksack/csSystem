package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func setupMiddleware(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		// Log incoming request
		log.Printf("Request: %s %s", c.Method(), c.OriginalURL())

		return c.Next()
	})
}
