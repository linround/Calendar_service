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
	UserID         uint64          `json:"userId" gorm:"not null;unique;primary_key;comment:用户ID;size:90"`
	Password       string          `json:"-" gorm:" comment:- 符号代表注释掉 注册时 可更改"`
	UserAccount    string          `json:"userAccount" gorm:"index;comment:用户账户"`
	UserName       string          `json:"userName" gorm:"comment: 用户名"`
	AvatarUrl      string          `json:"avatarUrl" gorm:"not null;comment:头像地址"`
	UserEmail      string          `json:"userEmail" gorm:"comment:用户邮箱;"`
	CalendarGroups []CalendarGroup `gorm:"foreignKey:user_id"`
	CalendarEvents []CalendarEvent `gorm:"foreignKey:user_id"`
}

func (u CalendarUser) TableName() string {
	return "calendar_users"
}
