package system

import "calendar_service/global"

/*
*
	todo
	日历用户
*/

type CalendarUser struct {
	global.CalendarEventModel
	ApiCalendarUser
}
type ApiCalendarUser struct {
	UserID    uint64 `json:"userId" gorm:"primarykey"` // 主键
	UserName  string `json:"nickName"`                 // 用户名
	HeaderUrl string `json:"headerUrl"`                // 头像地址
}
