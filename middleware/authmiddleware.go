package middleware

import (

	// "encoding/base64"

	"github.com/gofiber/fiber/v2"

	// "github.com/twilio/twilio-go/rest/api/v2010"

	// "strconv"
	"xactscore/utils"
)

func Isauthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := utils.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()
}
