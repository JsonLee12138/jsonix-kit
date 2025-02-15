package main

import (
	"fmt"
	"json-server-kit/apps/auth"
	"json-server-kit/apps/example"
	auto_migrate "json-server-kit/auto_migrate_local"
	"json-server-kit/middleware"
	"net/http"

	"json-server-kit/configs"

	utils2 "json-server-kit/utils"

	"github.com/JsonLee12138/json-server/pkg/core"
	"github.com/JsonLee12138/json-server/pkg/utils"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/ua-parser/uap-go/uaparser"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	app := core.NewApp()
	configInstance := utils.Raise(core.NewConfig())
	var cnf configs.Config
	configInstance.Bind(&cnf)
	core.Validator.RegisterValidation("phone", utils2.ValidatePhoneNumber)
	uparser := utils.Raise(uaparser.New("./config/regexes.yaml"))
	logger := core.NewLogger(cnf.Logger)
	app.Use(middleware.BotDetection(uparser))
	app.Use(middleware.I18n(cnf.I18n))
	app.Use(middleware.Response())
	app.Use(middleware.Logger(func(vo middleware.LogVO) {
		v, _ := core.MarshalForFiber(vo)
		if vo.Code == http.StatusOK {
			logger.Info(string(v))
		} else {
			logger.Error(string(v))
		}
	}))
	mysql := core.NewGormMysql(cnf.Mysql)
	utils.RaiseVoid(auto_migrate.AutoMigrate(mysql))
	fmt.Println("数据库自动迁移完成")
	container := dig.New()
	utils.RaiseVoid(container.Provide(func() *zap.Logger {
		return logger
	}))
	utils.RaiseVoid(container.Provide(func() *configs.Config {
		return &cnf
	}))
	utils.RaiseVoid(container.Provide(func() *uaparser.Parser {
		return uparser
	}))
	utils.RaiseVoid(container.Provide(func() *fiber.App {
		return app
	}))
	utils.RaiseVoid(container.Provide(func(config *configs.Config) *gorm.DB {
		return mysql
	}))
	utils.RaiseVoid(container.Provide(func(config *configs.Config) *redis.Client {
		return core.NewRedis(config.Redis)
	}))
	utils.RaiseVoid(example.ExampleModuleSetup(container))
	utils.RaiseVoid(auth.AuthModuleSetup(container))
	core.StartApp(app)
}
