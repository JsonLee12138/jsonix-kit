package controller

import (
	"errors"
	"json-server-kit/apps/auth/service"
	"json-server-kit/middleware"

	"github.com/gofiber/fiber/v2"
)

type CaptchaController struct {
	captchaService *service.CaptchaService
}

func NewCaptchaController(captchaService *service.CaptchaService) *CaptchaController {
	return &CaptchaController{
		captchaService,
	}
}

// @Summary 获取验证码
// @Description 获取验证码
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} utils.BaseResponseVO{data=vo.CaptchaVO}
// @Router /auth/captcha [get]
func (s *CaptchaController) Get(c *fiber.Ctx) error {
	result, err := s.captchaService.Get()
	if err != nil {
		return errors.New("generate_captcha_failed")
	}
	c.Locals(middleware.ResponseDataKey, result)
	return nil
}
