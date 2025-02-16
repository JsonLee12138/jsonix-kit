package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func Compress() fiber.Handler {
	return compress.New(compress.Config{
		Level: compress.LevelBestCompression,
		Next: func(c *fiber.Ctx) bool {
			body := c.Request().Body()
			if body == nil {
				return false
			}
			c.Request().SetBody(body)
			bodySize := len(body)
			return bodySize < 20_1024
		},
	})
}
