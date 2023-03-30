package system

import (
	"calendar_service/service/system"
	. "calendar_service/system"
	"context"
	"gorm.io/gorm"
)

type initCalendarEvent struct {
}

var calendarEventOrder = calendarGroupOrder + 1

func init() {
	system.RegisterInit(calendarEventOrder, &initCalendarEvent{})
}

// 创建表

func (e initCalendarEvent) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&CalendarEvent{})
}

// 判断是否已经创建了表格

func (e initCalendarEvent) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&CalendarEvent{})
}
