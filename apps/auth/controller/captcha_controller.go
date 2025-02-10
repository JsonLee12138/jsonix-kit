package controller

import (
	"json-server-kit/apps/auth/service"
	"json-server-kit/middleware"
	"net/http"

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

func (s *CaptchaController) Get(c *fiber.Ctx) error {
	result, err := s.captchaService.Get()
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Failed to generate captcha")
	}
	c.Locals(middleware.ResponseDataKey, result)
	return nil
}
