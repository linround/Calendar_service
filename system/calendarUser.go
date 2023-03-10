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
	UserID      uint64 `json:"userId" gorm:"primarykey"` // 主键
	Password    string `json:"-"`                        //  - 符号代表注释掉 注册时 可更改
	UserAccount string `json:"userAccount"`              // 用户账户  注册时  不可更改
	UserName    string `json:"userName"`                 // 用户名    可设置 可更改
	AvatarUrl   string `json:"avatarUrl"`                // 头像地址  可设置 可更改
	UserEmail   string `json:"userEmail"`                // 用户邮箱  可设置 可更改
}
