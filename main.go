package main

import (
	"calendar_service/core"
	"calendar_service/global"
	"calendar_service/initialize"
)

func main() {
	core.Viper() // 读取yaml的配置信息
	// 连接数据库
	// 将连接的数据库保存在全局上
	global.CalendarDB = initialize.Gorm()
	if global.CalendarDB != nil {
		//	初始化表
		initialize.RegisterTables(global.CalendarDB)
		db, _ := global.CalendarDB.DB()
		//	关闭数据库连接
		defer func() {
			db.Close()
		}()
	}
	core.RunServer()
}
