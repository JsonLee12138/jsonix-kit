package controller

import (
	exampleService "json-server-kit/apps/example/service"

	"github.com/gofiber/fiber/v2"
)

type ExampleController struct {
	service *exampleService.ExampleService
}

func NewExampleController(service *exampleService.ExampleService) *ExampleController {
	return &ExampleController{
		service,
	}
}

func (c *ExampleController) HelloWord(ctx *fiber.Ctx) error {
	return ctx.SendString(c.service.HelloWord())
}
