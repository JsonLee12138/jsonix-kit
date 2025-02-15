package vo

//go:generate easyjson -all ./captcha.go

type CaptchaVO struct {
	ID      string `json:"id"`      // 验证码ID
	Captcha string `json:"captcha"` // 验证码
	Enabled bool   `json:"enabled"` // 是否启用
}
