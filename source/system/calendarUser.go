package system

import (
	"calendar_service/service/system"
	. "calendar_service/system"
	"context"
	"gorm.io/gorm"
)

var calendarUserOrder = system.InitOrderSystem + 1

type initCalendarUser struct {
}

func init() {
	system.RegisterInit(calendarUserOrder, &initCalendarUser{})
}

func (u *initCalendarUser) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&CalendarUser{})

}
func (u *initCalendarUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	// 检查对应的表是否存在
	return db.Migrator().HasTable(&CalendarUser{})
}
