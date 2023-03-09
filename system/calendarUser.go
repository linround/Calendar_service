package system

import "calendar_service/global"

/*
*
	todo
	日历用户
*/

// 注册用户时所需要的完整字段

type CalendarUser struct {
	global.CalendarEventModel
	Password string `json:"password"` // 注册时 可更改
	ApiCalendarUser
}

// 服务端传递给前端

type ApiCalendarUser struct {
	UserID      uint64 `json:"userId" gorm:"primarykey"` // 主键
	UserAccount string `json:"userAccount"`              // 用户账户  注册时  不可更改
	UserName    string `json:"userName"`                 // 用户名    可设置 可更改
	AvatarUrl   string `json:"avatarUrl"`                // 头像地址  可设置 可更改
	UserEmail   string `json:"userEmail"`                // 用户邮箱  可设置 可更改
}
