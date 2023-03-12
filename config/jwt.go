package config

type JWT struct {
	SigningKey string `json:"signing-key" yaml:"signing-key" mapstructure:"signing-key"`
}
