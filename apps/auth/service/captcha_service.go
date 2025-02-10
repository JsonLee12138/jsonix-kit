package service

import (
	"json-server-kit/apps/auth/model/vo"
	"json-server-kit/configs"
	"json-server-kit/core"

	"github.com/mojocn/base64Captcha"
)

type CaptchaService struct {
	captchaStore *core.CaptchaStore
	cnf          *configs.Config
}

func NewCaptchaService(captchaStore *core.CaptchaStore, cnf *configs.Config) *CaptchaService {
	return &CaptchaService{
		captchaStore,
		cnf,
	}
}

func (s *CaptchaService) Get() (vo.CaptchaVO, error) {
	var result vo.CaptchaVO
	var driver base64Captcha.DriverMath
	driver.Width = s.cnf.CaptchaConfig.Width
	driver.Height = s.cnf.CaptchaConfig.Height
	driver.NoiseCount = s.cnf.CaptchaConfig.NoiseCount
	driver.ShowLineOptions = s.cnf.CaptchaConfig.ShowLineOptions
	result.Enabled = s.cnf.CaptchaConfig.Enable
	cp := base64Captcha.NewCaptcha(driver.ConvertFonts(), s.captchaStore)
	if result.Enabled {
		id, b64s, _, err := cp.Generate()
		result.ID = id
		result.Captcha = b64s
		if err != nil {
			return result, err
		}
	}
	return result, nil
}
