package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	// mysql的用户 root
	// mysql的登录密码为 linyuan333
	publicMysqlPassword = "linyuan333" // 数据库密码
	publicHost          = "121.199.1.247"
	publicPassword      = "Linyuan333" // XShell XFtp 连接密码
	nginxPath           = "usr/local/nginx/html"
	servicePath         = "/calendar_service"
	host                = "localhost"
	database            = "calendar"
	user                = "root"
	password            = "123456"
)

func GormMysql() *gorm.DB {
	//定义数据胡信息

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local", user, password, host, database) // 本地dsn
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local", user, publicMysqlPassword, publicHost, database) // 云服务器dsn
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
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
