package system

import (
	"calendar_service/global"
	"calendar_service/model/system/request"
	"calendar_service/system"
	"errors"

	"gorm.io/gorm"
)

type CalendarUserService struct {
}

func (calendarService *CalendarUserService) RegisterUser(user system.CalendarUser) (err error) {
	//查重
	var preUser request.Register
	if !errors.Is(global.CalendarDB.Model(&system.CalendarUser{}).Where("user_account = ?", user.UserAccount).First(&preUser).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户已存在")
	}
	// create 会自动反射到表名上
	return global.CalendarDB.Create(&user).Error

}
func (calendarService *CalendarUserService) GetUser(params request.SearchCalendarUserParams) (user system.CalendarUser, err error) {
	err = global.CalendarDB.Model(&system.CalendarUser{}).Find(&user).Error
	return
}
func (calendarService *CalendarUserService) UpdateUser(user system.CalendarUser) (err error) {
	err = global.CalendarDB.Model(&system.CalendarUser{}).Where("user_id = ?", user.UserID).Updates(map[string]interface{}{
		"user_name":  user.UserName,
		"user_email": user.UserEmail,
		"avatar_url": user.AvatarUrl,
	}).Error
	return
}
func (calendarService *CalendarUserService) DeleteUser(user system.CalendarUser) (err error) {
	err = global.CalendarDB.Delete(&system.CalendarUser{}, user.UserID).Error
	return
}
