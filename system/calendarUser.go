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
	AvatarUrl string `json:"avatarUrl"`                // 头像地址
	UserEmail string `json:"userEmail"`                // 用户邮箱
}
