package configs

type CaptchaConfig struct {
	Width           int  `mapstructure:"width" json:"width" yaml:"width" toml:"width"`
	Height          int  `mapstructure:"height" json:"height" yaml:"height" toml:"height"`
	MaxAge          int  `mapstructure:"max_age" json:"max_age" yaml:"max_age" toml:"max_age"`
	NoiseCount      int  `mapstructure:"noise_count" json:"noise_count" yaml:"noise_count" toml:"noise_count"`
	ShowLineOptions int  `mapstructure:"show_line_options" json:"show_line_options" yaml:"show_line_options" toml:"show_line_options"`
	Enable          bool `mapstructure:"enable" json:"enable" yaml:"enable" toml:"enable"`
}
