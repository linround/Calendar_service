package system

import (
	"calendar_service/global"
)

type CalendarGroup struct {
	global.CalendarEventModel `yaml:",inline"`
	ApiCalendarGroup          `yaml:",inline"`
	CalendarEvents            []CalendarEvent `gorm:"foreignKey:group_id"`
}
type ApiCalendarGroup struct {
	GroupType  uint   `json:"groupType" gorm:"comment:日历组说明0自己的日历；1订阅的日历"`
	GroupColor string `json:"groupColor" gorm:"comment:日历颜色配置"`
	GroupName  string `json:"groupName" gorm:"comment:日历名称"`
	UserID     uint64 `json:"userId" gorm:"comment:所属用户"`
	GroupID    uint64 ` json:"groupId" gorm:"primarykey" mapstructure:"groupId"`
}

/***

 	gorm:"-"  ====== 表示字段只是属性，不在数据库表中创建列
 	gorm:" comment:- 符号代表注释掉 注册时 可更改"   ====  对数据库中键值的描述，可查看键的属性说明
	gorm:"index" 使用索引，防止重复


	UUU     string `mapstructure:"state"`
	mapstructure 的含义在于当将数据某个数据填入时，uuu对应的时数据的state的值
***/

// mapstructure
//UUU     string `mapstructure:",remain"`
// ,remain 中，未被映射的值会放在该字段中
// ,squash 中，嵌套的结构被认为是拥有该结构体名字的另一个字段 提升字段
// ,omitempty 在对象解码成字符串是 在只有默认值的情况下不会解码该字段

// yaml
//UUU     string `yaml:",omitempty"`
// ,omitempty 在对象解码成字符串是 在只有默认值的情况下不会解码该字段 父结构体无法访问
// ,flow 父结构体无法访问
// ,inline 父结构体可以访问
