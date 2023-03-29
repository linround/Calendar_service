package system

import (
	"calendar_service/service/system"
	. "calendar_service/system"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var calendarUserOrder = system.InitOrderSystem + 1

func init() {
	system.RegisterInit(calendarUserOrder, &initCalendarUser{})
}

type initCalendarUser struct {
}

func (u *initCalendarUser) InitializerName() string {
	return CalendarUser{}.TableName()
}
func (u *initCalendarUser) DataInserted(ctx context.Context) bool {
	return true
}
func (u *initCalendarUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	// 检查对应的表是否存在
	return db.Migrator().HasTable(&CalendarUser{})
}

func (u *initCalendarUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []CalendarUser{
		{
			Password:    "123456",
			UserAccount: "admin",
			UserName:    "admin",
			AvatarUrl:   "https://avatars.githubusercontent.com/u/44738166?v=4",
			UserEmail:   "yuanlincuc@gmail.com",
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return nil, errors.Wrap(err, CalendarUser{}.TableName()+"表初始化失败")
	}
	next = context.WithValue(ctx, u.InitializerName(), entities)
	return
}
