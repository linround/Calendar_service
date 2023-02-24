package system

import (
	"calendar_service/global"
)

type CalendarEvent struct {
	global.CalendarEventModel // 包含ID、创建时间、更新时间、删除时间
	ApiCalendarEvent
}
type ApiCalendarEvent struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`      // 名称
	Start     uint64 `json:"start"`     // 开始时间
	End       uint64 `json:"end"`       // 结束时间
	Color     string `json:"color"`     // 颜色
	AllDay    bool   `json:"all_day"`   // 全天
	Author    string `json:"author"`    // 创建者
	Location  string `json:"location"`  // 地点
	Personnel string `json:"personnel"` // 人员
	Timed     bool   `json:"timed"`     // 显示该事件
}
