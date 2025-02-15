package middleware

import (
	"json-server-kit/apps/auth/model/enum"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ua-parser/uap-go/uaparser"
)

func BotDetection(uparser *uaparser.Parser) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userAgent := c.Get("User-Agent")
		client := uparser.Parse(userAgent)
		if enum.IsBot(client.UserAgent.Family) {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"message": "Forbidden"})
		}
		return c.Next()
	}
}
