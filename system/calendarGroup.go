package system

import (
	"calendar_service/global"
)

type CalendarGroup struct {
	global.CalendarEventModel
	ApiCalendarGroup
}
type ApiCalendarGroup struct {
	GroupID    uint   `mapstructure:"" json:"groupId" gorm:"primarykey"`
	UserID     uint64 `json:"userId" gorm:"comment:所属用户"`
	GroupType  uint   `json:"groupType" gorm:"comment:日历组说明0自己的日历；1订阅的日历"`
	GroupColor string `json:"groupColor" gorm:"comment:日历颜色配置"`
	GroupName  string `json:"groupName" gorm:"comment:日历名称"`
}

/***

 	gorm:"-"  ====== 表示字段只是属性，不在数据库表中创建列
 	gorm:" comment:- 符号代表注释掉 注册时 可更改"   ====  对数据库中键值的描述，可查看键的属性说明
	gorm:"index" 使用索引，防止重复


	UUU     string `mapstructure:"state"`
	mapstructure的含义在于当将数据某个数据填入时，uuu对应的时数据的state的值
***/
