package configs

import (
	"github.com/JsonLee12138/jsonix/pkg/configs"
)

type Config struct {
	CaptchaConfig CaptchaConfig        `mapstructure:"captcha" json:"captcha" yaml:"captcha" toml:"captcha"`
	System        configs.SystemConfig `mapstructure:"system" json:"system" yaml:"system" toml:"system"`
	Mysql         configs.MysqlConfig  `mapstructure:"mysql" json:"mysql" yaml:"mysql" toml:"mysql"`
	Redis         configs.RedisConfig  `mapstructure:"redis" json:"redis" yaml:"redis" toml:"redis"`
	I18n          configs.I18nConfig   `mapstructure:"i18n" json:"i18n" yaml:"i18n" toml:"i18n"`
	Logger        configs.LogConfig    `mapstructure:"logger" json:"logger" yaml:"logger" toml:"logger"`
	Cors          configs.CorsConfig   `mapstructure:"cors" json:"cors" yaml:"cors" toml:"cors"`
	RefreshJWT    JWTConfig            `mapstructure:"refresh-jwt" json:"refresh-jwt" yaml:"refresh-jwt" toml:"refresh-jwt"`
	AccessJWT     JWTConfig            `mapstructure:"access-jwt" json:"access-jwt" yaml:"access-jwt" toml:"access-jwt"`
	JWTAuth       JWTAuthConfig        `mapstructure:"jwt-auth" json:"jwt-auth" yaml:"jwt-auth" toml:"jwt-auth"`
}
