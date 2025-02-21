package main

import (
	"fmt"
	"jsonix-kit/apps/example"
	"jsonix-kit/auto_migrate"
	"jsonix-kit/middleware"
	"net/http"

	"jsonix-kit/configs"

	"github.com/JsonLee12138/jsonix/pkg/core"
	"github.com/JsonLee12138/jsonix/pkg/utils"
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
	uparser := utils.Raise(uaparser.New("./config/regexes.yaml"))
	logger := core.NewLogger(cnf.Logger)
	app.Use(middleware.Cors(&cnf.Cors))
	redisClient := utils.Raise(core.NewRedis(cnf.Redis))
	app.Use(middleware.Logger(func(vo middleware.LogVO) {
		v, _ := core.MarshalForFiber(vo)
		if vo.Code == http.StatusOK {
			logger.Info(string(v))
		} else {
			logger.Error(string(v))
		}
	}))
	mysql := utils.Raise(core.NewGormMysql(cnf.Mysql))
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
	utils.RaiseVoid(container.Provide(func() *gorm.DB {
		return mysql
	}))
	utils.RaiseVoid(container.Provide(func() *redis.Client {
		return redisClient
	}))
	utils.RaiseVoid(example.ExampleModuleSetup(container))
	core.StartApp(app)
}
