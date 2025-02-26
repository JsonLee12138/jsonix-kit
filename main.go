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
	//"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/ua-parser/uap-go/uaparser"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 如需正常运行, 请先配置去config.yaml中配置mysql的相关配置或自定义注入其他gorm.DB实例, 否则将无法正常运行
// For normal operation, please configure the config.yaml to configure mysql or customize the injection of other gorm.DB instances, otherwise it will not work properly.
func main() {
	app := core.NewApp()
	configInstance := utils.Raise(core.NewConfig())
	var cnf configs.Config
	configInstance.Bind(&cnf)
	uparser := utils.Raise(uaparser.New("./config/regexes.yaml"))
	logger := core.NewLogger(cnf.Logger)
	app.Use(middleware.Cors(&cnf.Cors))
	// 如需使用redis, 请先配置去config.yaml文件中配置相关配置
	// If you want to use redis, please go to config.yaml first to configure the relevant configurations.
	//redisClient, err := core.NewRedis(cnf.Redis)
	//utils.RaiseVoidByErrorHandler(err, func(err error) error {
	//	return fmt.Errorf("redis connection failure: %w", err)
	//})
	app.Use(middleware.Logger(func(vo middleware.LogVO) {
		v, _ := core.MarshalForFiber(vo)
		if vo.Code == http.StatusOK {
			logger.Info(string(v))
		} else {
			logger.Error(string(v))
		}
	}))
	fmt.Println(cnf.Mysql.DSN())
	// 如需使用mysql, 请先配置去config.yaml文件中配置相关配置
	// If you want to use mysql, please go to config.yaml first to configure the relevant configuration.
	mysql, err := core.NewGormMysql(cnf.Mysql)
	utils.RaiseVoidByErrorHandler(err, func(err error) error {
		return fmt.Errorf("❌ Database connection failure: %w", err)
	})
	// 如需迁移数据库，请先在根目录运行 `jsonix migrate` 再取消注释
	utils.RaiseVoid(auto_migrate.AutoMigrate(mysql))
	fmt.Println("✅ Automated database migration completed")
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
	//utils.RaiseVoid(container.Provide(func() *redis.Client {
	//	return redisClient
	//}))
	utils.RaiseVoid(example.ExampleModuleSetup(container))
	core.StartApp(app)
}
