package example

import (
	"json-server-kit/apps/example/controller"
	"json-server-kit/apps/example/repository"
	"json-server-kit/apps/example/service"

	"github.com/JsonLee12138/json-server/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

func ExampleModuleSetup(container *dig.Container) error {
	return utils.TryCatchVoid(func() {
		scope := container.Scope("Example")
		utils.RaiseVoid(scope.Provide(controller.NewExampleController))
		utils.RaiseVoid(scope.Provide(service.NewExampleService))
		utils.RaiseVoid(scope.Provide(repository.NewExampleRepository))
		utils.RaiseVoid(scope.Invoke(func(app *fiber.App, exampleController *controller.ExampleController) {
			group := app.Group("example")
			group.Get("/", exampleController.HelloWord)
		}))
	}, utils.DefaultErrorHandler)
}
