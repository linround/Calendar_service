package core

import (
	"calendar_service/core/internal"
	"calendar_service/global"
	"fmt"
	"github.com/spf13/viper"
)

// viper 是一个go程序配置解决方案

func Viper() *viper.Viper {
	var config string
	config = internal.ConfigDefaultFile
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	err = v.Unmarshal(&global.CalendarConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
	return v
}
