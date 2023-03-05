package system

import "calendar_service/global"

type CalendarGroup struct {
	global.CalendarEventModel
	ID     uint   `json:"id" gorm:"primarykey"` // 主键
	userID uint64 `json:"userId"`
	Color  string `json:"color"`
	Name   string `json:"name"`
}
