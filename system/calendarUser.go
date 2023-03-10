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
	UserID      uint64 `json:"userId" gorm:"not null;unique;primary_key;comment:用户ID;size:90"`
	Password    string `json:"-" gorm:" comment:- 符号代表注释掉 注册时 可更改"`
	UserAccount string `json:"userAccount" gorm:"comment:用户账户"`
	UserName    string `json:"userName" gorm:"comment: 用户名;default:用户名"`
	AvatarUrl   string `json:"avatarUrl" gorm:"not null;comment:头像地址"`
	UserEmail   string `json:"userEmail" gorm:"comment:用户邮箱;default:邮箱"`
}

/***

 gorm:"-"  ====== 表示字段只是属性，不在数据库表中创建列
 gorm:" comment:- 符号代表注释掉 注册时 可更改"   ====  对数据库中键值的描述，可查看键的属性说明
***/
