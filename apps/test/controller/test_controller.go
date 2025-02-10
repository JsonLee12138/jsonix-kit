package controller

import (
	"github.com/gofiber/fiber/v2"
	"json-server-kit/apps/test/service"
)

type TestController struct {
	service *service.TestService
}

func NewTestController(service *service.TestService) *TestController {
	return &TestController{
		service,
	}
}

func (c *TestController) HelloWord(ctx *fiber.Ctx) error {
	return ctx.SendString(c.service.HelloWord())
}
