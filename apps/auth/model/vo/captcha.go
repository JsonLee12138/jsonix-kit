package vo

//go:generate easyjson -all ./captcha.go

type CaptchaVO struct {
	ID      string `json:"id"`
	Captcha string `json:"captcha"`
	Enabled bool   `json:"enabled"`
}
