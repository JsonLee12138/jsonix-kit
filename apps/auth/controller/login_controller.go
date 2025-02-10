package controller

import (
	"json-server-kit/apps/auth/service"

	"github.com/gofiber/fiber/v2"
)

type LoginController struct {
	UsernameService *service.UsernameService
}

func NewLoginController(UsernameService *service.UsernameService) *LoginController {
	return &LoginController{
		UsernameService,
	}
}

func (c *LoginController) HelloWord(ctx *fiber.Ctx) error {
	return ctx.SendString(c.UsernameService.HelloWord())
}
