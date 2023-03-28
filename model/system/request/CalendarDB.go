package request

import (
	"calendar_service/config"
	"fmt"
)

type InitDB struct {
	DBType   string `json:"dbType"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password"`
	DBName   string `json:"dbName" binding:"required"`
}

func (i *InitDB) ToMysqlConfig() config.Mysql {
	return config.Mysql{
		GeneralDB: config.GeneralDB{
			Host:         i.Host,
			Port:         i.Port,
			UserName:     i.UserName,
			Password:     i.Password,
			Config:       "charset=utf8mb4&parseTime=True&loc=Local",
			DbName:       i.DBName,
			MaxIdleConns: 10,
			MaxOpenConns: 100,
		},
	}

}
func (i *InitDB) MysqlDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.Port == "" {
		i.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", i.UserName, i.Password, i.Host, i.Port)
}
