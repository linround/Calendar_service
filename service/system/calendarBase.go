package system

import (
	"calendar_service/global"
	"calendar_service/system"
	"calendar_service/utils"
	"errors"

	"gorm.io/gorm"
)

type CalendarBaseService struct {
}

func (calendarService *CalendarBaseService) Login(u *system.CalendarUser) (user *system.CalendarUser, err error) {
	var preUser system.CalendarUser
	err = global.CalendarDB.Model(&system.CalendarUser{}).Where("user_account = ?", u.UserAccount).First(&preUser).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, preUser.Password); !ok {
			return nil, errors.New("密码错误")
		}
		return &preUser, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("账户不存在")
	}
	return
}
