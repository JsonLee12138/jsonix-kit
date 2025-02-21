package common

import (
	"json-server-kit/apps/common/controller"
	"json-server-kit/apps/common/repository"
	"json-server-kit/apps/common/service"

	"github.com/JsonLee12138/jsonix/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

func ProvideController(scope *dig.Scope) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(scope.Provide(controller.NewDictController))
	}, utils.DefaultErrorHandler)
}

func ProvideService(scope *dig.Scope) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(scope.Provide(service.NewDictService))
	}, utils.DefaultErrorHandler)
}

func ProvideRepository(scope *dig.Scope) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(scope.Provide(repository.NewDictRepository))
	}, utils.DefaultErrorHandler)
}

func RouterSetup(app *fiber.App, dictController *controller.DictController) {
	group := app.Group("common")
	group.Get("/dict/types", dictController.GetDictTypes)
}

func CommonModuleSetup(container *dig.Container) error {
	return utils.TryCatchVoid(func() {
		scope := container.Scope("common")
		utils.RaiseVoid(ProvideController(scope))
		utils.RaiseVoid(ProvideService(scope))
		utils.RaiseVoid(ProvideRepository(scope))
		utils.RaiseVoid(scope.Invoke(RouterSetup))
	}, utils.DefaultErrorHandler)
}
