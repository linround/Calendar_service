package config

type Server struct {
	Mysql Mysql `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	Cors  CORS  `json:"cors" yaml:"cors" mapstructure:"cors"`
	JWT   JWT   `json:"jwt" yaml:"jwt" mapstructure:"jwt"`
}
