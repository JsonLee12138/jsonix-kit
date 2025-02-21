package middleware

import (
	"github.com/JsonLee12138/jsonix/pkg/configs"
	"github.com/JsonLee12138/jsonix/pkg/core"
	"github.com/gofiber/fiber/v2"
)

func I18n(cnf configs.I18nConfig) fiber.Handler {
	return core.NewI18n(cnf)
}
