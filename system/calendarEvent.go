package system

import (
	"calendar_service/global"
)

type CalendarEvent struct {
	global.CalendarEventModel // 包含ID、创建时间、更新时间、删除时间
	ApiCalendarEvent
}
type ApiCalendarEvent struct {
	EventID        uint   `json:"eventId" gorm:"primarykey"` // 主键
	EventName      string `json:"eventName"`                 // 名称
	Start          uint64 `json:"start"`                     // 开始时间
	End            uint64 `json:"end"`                       // 结束时间
	EventColor     string `json:"eventColor"`                // 颜色
	AllDay         bool   `json:"allDay"`                    // 全天
	UserName       string `json:"userName"`                  // 创建者
	EventLocation  string `json:"eventLocation"`             // 地点
	EventPersonnel string `json:"eventPersonnel"`            // 人员
	EventTimed     bool   `json:"eventTimed"`                // 显示该事件
	GroupId        uint64 `json:"groupId"`                   // 属于对应的日历组
}
