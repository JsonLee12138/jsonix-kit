package configs

import (
	"github.com/JsonLee12138/json-server/pkg/configs"
)

type Config struct {
	CaptchaConfig CaptchaConfig        `mapstructure:"captcha" json:"captcha" yaml:"captcha" toml:"captcha"`
	System        configs.SystemConfig `mapstructure:"system" json:"system" yaml:"system" toml:"system"`
	Mysql         configs.MysqlConfig  `mapstructure:"mysql" json:"mysql" yaml:"mysql" toml:"mysql"`
	Redis         configs.RedisConfig  `mapstructure:"redis" json:"redis" yaml:"redis" toml:"redis"`
	I18n          configs.I18nConfig   `mapstructure:"i18n" json:"i18n" yaml:"i18n" toml:"i18n"`
	Logger        configs.LogConfig    `mapstructure:"logger" json:"logger" yaml:"logger" toml:"logger"`
	Cors          configs.CorsConfig   `mapstructure:"cors" json:"cors" yaml:"cors" toml:"cors"`
}
