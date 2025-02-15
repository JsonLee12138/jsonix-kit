package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidatePhoneNumber(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	if phone == "" {
		return true
	}
	// E.164 格式正则表达式：+86 后面跟着 11 位数字的中国手机号
	phoneRegex := `^\\+?[1-9]\\d{1,14}$`

	// 使用正则表达式进行匹配
	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(phone)
}
