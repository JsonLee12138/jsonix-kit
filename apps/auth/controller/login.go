package controller

import (
	"errors"
	"json-server-kit/apps/auth/model/dto"
	"json-server-kit/apps/auth/service"
	"json-server-kit/middleware"
	"json-server-kit/utils"

	utils2 "github.com/JsonLee12138/jsonix/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type LoginController struct {
	LoginService   *service.LoginService
	CaptchaService *service.CaptchaService
}

func NewLoginController(LoginService *service.LoginService, CaptchaService *service.CaptchaService) *LoginController {
	return &LoginController{
		LoginService,
		CaptchaService,
	}
}

func (c *LoginController) HelloWord(ctx *fiber.Ctx) error {
	return ctx.SendString(c.LoginService.HelloWord())
}

// @Summary 用户名密码登录
// @Description 用户名密码登录(需要验证码)
// @Tags auth
// @Accept json
// @Produce json
// @Param loginDTO body dto.UsernameLoginWithCaptchaDTO true "登录信息"
// @Success 200 {object} utils.BaseResponseVO{data=entity.User}
// @Router /auth/login [post]
func (c *LoginController) UsernameLoginWithCaptcha(ctx *fiber.Ctx) error {
	return utils2.TryCatchVoid(func() {
		var loginDTO dto.UsernameLoginWithCaptchaDTO
		utils2.RaiseVoidByErrorHandler(ctx.BodyParser(&loginDTO), func(err error) error {
			return utils.NewBadParameters(err)
		})
		check := c.CaptchaService.Verify(loginDTO.CaptchaId, loginDTO.Captcha)
		if !check {
			panic(errors.New("verify_captcha_failed"))
		}
		userInfo := utils2.Raise(c.LoginService.GetUserInfoByUsername(loginDTO.Username))
		utils2.RaiseVoid(c.LoginService.UserEnabled(userInfo))

		utils2.RaiseVoid(c.LoginService.CheckPassword(userInfo, loginDTO.Password))

		// 生成token
		// 存token到redis
		// 判断登录平台是网页还是app
		// 网页 设置token到cookie
		// app 设置token到header
		// 返回用户信息
		ctx.Locals(middleware.ResponseDataKey, userInfo)
	}, utils2.DefaultErrorHandler)
}
