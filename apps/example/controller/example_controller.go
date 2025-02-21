package controller

import (
	"github.com/gofiber/fiber/v2"
	"json-server-kit/apps/example/service"
)

type ExampleController struct {
    service *service.ExampleService
}

func NewExampleController(service *service.ExampleService) *ExampleController {
	return &ExampleController{
	    service,
	}
}

func (c *ExampleController) HelloWorld(ctx *fiber.Ctx) error {
	return ctx.SendString(c.service.HelloWorld())
}
