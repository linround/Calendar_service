package system

import "calendar_service/global"

type CalendarGroup struct {
	global.CalendarEventModel
	ApiCalendarGroup
}
type ApiCalendarGroup struct {
	GroupID    uint   `json:"groupId" gorm:"primarykey"`
	UserID     uint64 `json:"userId" gorm:"comment:所属用户"`
	GroupType  uint   `json:"groupType" gorm:"comment:日历组说明0自己的日历；1订阅的日历"`
	GroupColor string `json:"groupColor" gorm:"comment:日历颜色配置"`
	GroupName  string `json:"groupName" gorm:"comment:日历名称"`
}
