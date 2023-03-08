package system

import (
	"calendar_service/global"
	"calendar_service/model/system/request"
	"calendar_service/system"
)

type CalendarUserService struct {
}

func (calendarService *CalendarUserService) CreateUser(user system.CalendarUser) (err error) {
	return global.CalendarDB.Create(&user).Error

}
func (calendarService *CalendarUserService) GetUser(params request.SearchCalendarUserParams) (user system.ApiCalendarUser, err error) {
	err = global.CalendarDB.Model(&system.CalendarUser{}).Find(&user).Error
	return
}
func (calendarService *CalendarUserService) UpdateUser(user system.ApiCalendarUser) (err error) {
	err = global.CalendarDB.Model(&system.CalendarUser{}).Where("user_id = ?", user.UserID).Updates(map[string]interface{}{
		"user_name":  user.UserName,
		"user_email": user.UserEmail,
		"avatar_url": user.AvatarUrl,
	}).Error
	return
}
func (calendarService *CalendarUserService) DeleteUser(user system.ApiCalendarUser) (err error) {
	err = global.CalendarDB.Delete(&system.CalendarUser{}, user.UserID).Error
	return
}
