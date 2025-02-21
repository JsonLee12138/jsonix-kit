package example

import (
	"json-server-kit/apps/example/controller"
	"json-server-kit/apps/example/service"
	"json-server-kit/apps/example/repository"

	"github.com/JsonLee12138/jsonix/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

func ProvideController(scope *dig.Scope) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(scope.Provide(controller.NewExampleController))
	}, utils.DefaultErrorHandler)
}

func ProvideService(scope *dig.Scope) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(scope.Provide(service.NewExampleService))
	}, utils.DefaultErrorHandler)
}

func ProvideRepository(scope *dig.Scope) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(scope.Provide(repository.NewExampleRepository))
	}, utils.DefaultErrorHandler)
}

func RouterSetup(app *fiber.App, exampleController *controller.ExampleController) {
	group := app.Group("example")
	group.Get("/", exampleController.HelloWorld)
}

func ExampleModuleSetup(container *dig.Container) error {
    return utils.TryCatchVoid(func() {
      scope := container.Scope("example")
      utils.RaiseVoid(ProvideController(scope))
			utils.RaiseVoid(ProvideService(scope))
			utils.RaiseVoid(ProvideRepository(scope))
      utils.RaiseVoid(scope.Invoke(RouterSetup))
    }, utils.DefaultErrorHandler)
}
