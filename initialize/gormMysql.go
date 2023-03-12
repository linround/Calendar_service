package initialize

import (
	"calendar_service/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	//定义数据胡信息
	mySql := global.CalendarConfig.Mysql
	fmt.Println(mySql.Dsn())
	mysqlConfig := mysql.Config{
		DSN:                       mySql.Dsn(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	//	建立数据库连接
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接错误")
		return nil
	}
	// 获取通用的数据库对象
	sqlDB, _ := db.DB()

	// 设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	return db
}
