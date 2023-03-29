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
	EventName      string `json:"eventName" gorm:"comment:事件名称"`
	Start          uint64 `json:"start" gorm:"comment:开始时间"`
	End            uint64 `json:"end" gorm:"comment:结束时间"`
	EventColor     string `json:"eventColor" gorm:"comment:事件颜色"`
	AllDay         bool   `json:"allDay" gorm:"comment:全天设置"`
	EventLocation  string `json:"eventLocation" gorm:"comment:事件地点"`
	EventPersonnel string `json:"eventPersonnel" gorm:"comment:事件人员"`
	EventTimed     bool   `json:"eventTimed" gorm:"comment:显示该事件"`
	GroupID        uint64 `json:"groupId" gorm:"comment:事件对应的日历组id"`
	UserAccount    string `json:"userAccount" gorm:"comment:事件创建者"`
	UserID         uint64 `json:"userID" gorm:"comment:事件所属的用户"`
}

func (e CalendarEvent) TableName() string {
	return "calendar_events"
}
