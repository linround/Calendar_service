package system

import "calendar_service/global"

/*
*
	todo
	日历用户
*/

type CalendarUser struct {
	global.CalendarEventModel
	UserID    uint64 `json:"userId" gorm:"primarykey"`
	NickName  string `json:"nickName"`
	HeaderUrl string `json:"headerUrl"`
}
