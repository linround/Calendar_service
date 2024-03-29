package initialize

import (
	"calendar_service/system"
	"fmt"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

// 注册数据表

func RegisterTables(db *gorm.DB) {
	// db.AutoMigrate
	// 会创建表、缺失的外键、约束、列和索引
	// 出于保护数据的目的，其不会删除未使用的列

	// 这里处理中文编码的问题
	err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(

		// 记录日历事件表
		system.CalendarEvent{},
		//	日历组表
		system.CalendarGroup{},
		// 用户表
		system.CalendarUser{},
		// 网站埋点表
		system.SiteMonitoring{},
		// 图片字段存储
		system.PictureStorage{},
	)
	if err != nil {
		fmt.Println("表注册失败")
		os.Exit(0)
	}
	fmt.Println("表注册成功")
}
