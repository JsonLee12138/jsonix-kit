package configs

import (
	"github.com/JsonLee12138/json-server/pkg/configs"
)

type Config struct {
	CaptchaConfig CaptchaConfig        `mapstructure:"captcha" json:"captcha" yaml:"captcha" toml:"captcha"`
	System        configs.SystemConfig `mapstructure:"system" json:"system" yaml:"system" toml:"system"`
	Mysql         configs.MysqlConfig  `mapstructure:"mysql" json:"mysql" yaml:"mysql" toml:"mysql"`
	Redis         configs.RedisConfig  `mapstructure:"redis" json:"redis" yaml:"redis" toml:"redis"`
}

type CaptchaConfig struct {
	Width           int  `mapstructure:"width" json:"width" yaml:"width" toml:"width"`
	Height          int  `mapstructure:"height" json:"height" yaml:"height" toml:"height"`
	MaxAge          int  `mapstructure:"max_age" json:"max_age" yaml:"max_age" toml:"max_age"`
	NoiseCount      int  `mapstructure:"noise_count" json:"noise_count" yaml:"noise_count" toml:"noise_count"`
	ShowLineOptions int  `mapstructure:"show_line_options" json:"show_line_options" yaml:"show_line_options" toml:"show_line_options"`
	Enable          bool `mapstructure:"enable" json:"enable" yaml:"enable" toml:"enable"`
}
