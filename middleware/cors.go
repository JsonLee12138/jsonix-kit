package middleware

import (
	"github.com/JsonLee12138/jsonix/pkg/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors(cnf *configs.CorsConfig) fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: cnf.GetAllowOriginsString(),
		AllowMethods: cnf.GetAllowMethodsString(),
		AllowHeaders: cnf.GetAllowHeadersString(),
		MaxAge:       cnf.GetMaxAgeSeconds(),
	})
}
