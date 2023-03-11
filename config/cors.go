package config

type CORS struct {
	Whitelist []CORSWhitelist `mapstructure:"whitelist" json:"whitelist" yaml:"whitelist"`
}

type CORSWhitelist struct {
	AllowOrigin string `mapstructure:"allow-origin" json:"allow-origin" yaml:"allow-origin"`
}
