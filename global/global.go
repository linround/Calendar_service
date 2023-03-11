package global

import (
	"calendar_service/config"
	"gorm.io/gorm"
)

// 定义全局变量

var (
	CalendarDB     *gorm.DB
	CalendarConfig config.Server
)
