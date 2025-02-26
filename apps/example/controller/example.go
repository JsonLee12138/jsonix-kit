package controller

import (
	"github.com/gofiber/fiber/v2"
	"jsonix-kit/apps/example/service"
)

type ExampleController struct {
	service *service.ExampleService
}

func NewExampleController(service *service.ExampleService) *ExampleController {
	return &ExampleController{
		service,
	}
}

// @Summary HelloWorld
// @Description HelloWorld
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello World"
// @Router /example [get]
func (c *ExampleController) HelloWorld(ctx *fiber.Ctx) error {
	return ctx.SendString(c.service.HelloWorld())
}
