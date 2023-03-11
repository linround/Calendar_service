package config

type GeneralDB struct {
	Host         string `json:"host" yaml:"host" mapstructure:"host"`
	Port         string `json:"port" yaml:"port" mapstructure:"port"`
	UserName     string `json:"userName" yaml:"userName" mapstructure:"userName"`
	Password     string `json:"password" yaml:"password" mapstructure:"password"`
	Config       string `json:"config" yaml:"config" mapstructure:"mapstructure"`
	DbName       string `json:"db-name" yaml:"db-name" mapstructure:"db-name"`
	MaxIdleConns int    `json:"max-idle-conns" yaml:"max-idle-conns" mapstructure:"max-idle-conns"`
	MaxOpenConns int    `json:"max-open-conns" yaml:"max-open-conns" mapstructure:"max-open-conns"`
}
