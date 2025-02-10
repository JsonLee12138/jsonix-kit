package main

import (
	"json-server-kit/apps/auth"
	"json-server-kit/apps/example"
	"json-server-kit/middleware"

	"json-server-kit/configs"

	selfCore "json-server-kit/core"

	"github.com/JsonLee12138/json-server/pkg/core"
	"github.com/JsonLee12138/json-server/pkg/utils"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	app := core.NewApp()
	configInstance := utils.Raise(core.NewConfig())
	app.Use(fiberi18n.New(&fiberi18n.Config{
		RootPath: "./locales",
	}))
	app.Use(middleware.Response())
	core.Validator.RegisterValidation("phone", selfCore.ValidatePhoneNumber)
	var cnf configs.Config
	configInstance.Bind(&cnf)
	container := dig.New()
	utils.RaiseVoid(container.Provide(func() *configs.Config {
		return &cnf
	}))
	utils.RaiseVoid(container.Provide(func() *fiber.App {
		return app
	}))
	utils.RaiseVoid(container.Provide(func(config *configs.Config) *gorm.DB {
		return core.NewGormMysql(config.Mysql)
	}))
	utils.RaiseVoid(container.Provide(func(config *configs.Config) *redis.Client {
		return core.NewRedis(config.Redis)
	}))
	utils.RaiseVoid(example.ExampleModuleSetup(container))
	utils.RaiseVoid(auth.AuthModuleSetup(container))
	core.StartApp(app)
}
