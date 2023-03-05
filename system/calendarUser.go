package system

import "calendar_service/global"

/*
*
	todo
	日历用户
*/

type CalendarUser struct {
	global.CalendarEventModel
	ID        uint   `json:"id" gorm:"primarykey"` // 主键
	NickName  string `json:"nickName"`
	HeaderUrl string `json:"headerUrl"`
}
