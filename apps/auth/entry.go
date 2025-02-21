package auth

import (
	"json-server-kit/apps/auth/controller"
	"json-server-kit/apps/auth/repository"
	"json-server-kit/apps/auth/service"
	"json-server-kit/core"

	"github.com/JsonLee12138/jsonix/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

// TODO: 登录应该有的功能
// 用户名密码登录
// IP黑名单
// 用户黑名单
// 登录失败次数限制
// 记录登录日志
// 验证码
// 登录成功后生成token
// token有效期
// 双token
// token刷新

// 微信登录
// 手机号登录
// 邮箱登录

// 登出
// 登出后token失效

// 双重验证(谷歌验证、验证码验证)

func ProvideController(scope *dig.Scope) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(scope.Provide(controller.NewCaptchaController))
		utils.RaiseVoid(scope.Provide(controller.NewLoginController))
	}, utils.DefaultErrorHandler)
}

func ProvideService(scope *dig.Scope) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(scope.Provide(service.NewCaptchaService))
		utils.RaiseVoid(scope.Provide(service.NewLoginService))
	}, utils.DefaultErrorHandler)
}

func ProvideRepository(scope *dig.Scope) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(scope.Provide(repository.NewUserRepository))
	}, utils.DefaultErrorHandler)
}

func RouterSetup(app *fiber.App, exampleController *controller.LoginController, captchaController *controller.CaptchaController) {
	group := app.Group("auth")
	group.Get("/", exampleController.HelloWord)
	group.Get("/captcha", captchaController.Get)
	group.Post("/login", exampleController.UsernameLoginWithCaptcha)
}

func AuthModuleSetup(container *dig.Container) error {
	return utils.TryCatchVoid(func() {
		scope := container.Scope("auth")
		utils.RaiseVoid(scope.Provide(core.NewCaptchaStore))
		utils.RaiseVoid(ProvideController(scope))
		utils.RaiseVoid(ProvideService(scope))
		utils.RaiseVoid(ProvideRepository(scope))
		utils.RaiseVoid(scope.Invoke(RouterSetup))
	}, utils.DefaultErrorHandler)
}
