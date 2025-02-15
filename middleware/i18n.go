package middleware

import (
	"github.com/JsonLee12138/json-server/pkg/configs"
	"github.com/JsonLee12138/json-server/pkg/core"
	"github.com/gofiber/fiber/v2"
)

func I18n(cnf configs.I18nConfig) fiber.Handler {
	return core.NewI18n(cnf)
}
