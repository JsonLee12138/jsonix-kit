package dto

//go:generate easyjson -all ./login.go

type UsernameLoginDTO struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type UsernameLoginWithCaptchaDTO struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type PhoneSMSLoginDTO struct {
	Phone string `json:"phone"` // 手机号
	Code  string `json:"code"`  // 验证码
}

type EmailSMSLoginDTO struct {
	Email string `json:"email"` // 邮箱
	Code  string `json:"code"`  // 验证码
}

type AccountLoginDTO struct {
	Account  string `json:"account"`  // 账号
	Password string `json:"password"` // 密码
}

type AccountLoginWithCaptchaDTO struct {
	Account   string `json:"account"`   // 账号
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}
