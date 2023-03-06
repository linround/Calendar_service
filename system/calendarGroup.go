package system

import "calendar_service/global"

type CalendarGroup struct {
	global.CalendarEventModel
	ApiCalendarGroup
}
type ApiCalendarGroup struct {
	GroupID    uint   `json:"groupId" gorm:"primarykey"` // 主键
	UserID     uint64 `json:"userId"`                    // 所属用户
	GroupType  uint   `json:"groupType"`                 // 0 自己的日历 1 订阅日历
	GroupColor string `json:"groupColor"`                //颜色
	GroupName  string `json:"groupName"`                 // 名称
}
