package system

import (
	"calendar_service/global"
)

type CalendarEvent struct {
	global.CalendarEventModel // 包含ID、创建时间、更新时间、删除时间
	ApiCalendarEvent
}
type ApiCalendarEvent struct {
	EventID        uint   `json:"eventId" gorm:"primarykey;comment:事件id;"`
	EventName      string `json:"eventName" gorm:"事件名称"`
	Start          uint64 `json:"start" gorm:"comment:开始时间"`
	End            uint64 `json:"end" gorm:"结束时间"`
	EventColor     string `json:"eventColor" gorm:"事件颜色"`
	AllDay         bool   `json:"allDay" gorm:"全天设置"`
	UserName       string `json:"userName" gorm:"事件创建者"`
	EventLocation  string `json:"eventLocation" gorm:"事件地点"`
	EventPersonnel string `json:"eventPersonnel" gorm:"事件人员"`
	EventTimed     bool   `json:"eventTimed" gorm:"显示该事件"`
	GroupId        uint64 `json:"groupId" gorm:"事件对应的日历组id"`
}
