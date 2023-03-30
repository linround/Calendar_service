package system

import (
	"calendar_service/service/system"
	. "calendar_service/system"
	"context"
	"gorm.io/gorm"
)

type initCalendarGroup struct {
}

var calendarGroupOrder = calendarUserOrder + 1

func init() {
	system.RegisterInit(calendarUserOrder, &initCalendarGroup{})
}

func (g initCalendarGroup) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&CalendarGroup{})
}

func (g initCalendarGroup) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&CalendarGroup{})
}
