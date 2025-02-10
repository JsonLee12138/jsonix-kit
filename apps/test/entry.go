package test

import (
	"json-server-kit/apps/test/controller"
	"json-server-kit/apps/test/repository"
	"json-server-kit/apps/test/service"

	"github.com/JsonLee12138/json-server/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

func inject(scope *dig.Scope) {
	utils.RaiseVoid(scope.Provide(controller.NewTestController))
	utils.RaiseVoid(scope.Provide(service.NewTestService))
	utils.RaiseVoid(scope.Provide(repository.NewTestRepository))
}
func TestModuleSetup(container *dig.Container) error {
	return utils.TryCatchVoid(func() {
		scope := container.Scope("test")
		utils.RaiseVoid(scope.Provide(controller.NewTestController))
		utils.RaiseVoid(scope.Provide(service.NewTestService))
		utils.RaiseVoid(scope.Provide(repository.NewTestRepository))
		inject(scope)
		utils.RaiseVoid(scope.Invoke(func(app *fiber.App, testController *controller.TestController) {
			group := app.Group("test")
			group.Get("/", testController.HelloWord)
		}))
	}, utils.DefaultErrorHandler)
}
