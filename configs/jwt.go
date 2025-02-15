package configs

type JWTConfig struct {
	SecretKey     string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key" toml:"secret-key"`
	ExpireTime    string `mapstructure:"expire-time" json:"expire-time" yaml:"expire-time" toml:"expire-time"`
	BufferTime    string `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time" toml:"buffer-time"`
	Issuer        string `mapstructure:"issuer" json:"issuer" yaml:"issuer" toml:"issuer"`
	SigningMethod string `mapstructure:"signing-method" json:"signing-method" yaml:"signing-method" toml:"signing-method"` // 可选值 'HS256', 'HS384', 'HS512', 'RS256', 'RS384', 'RS512', 'ES256', 'ES384', 'ES512', 'PS256', 'PS384', 'PS512'
}
