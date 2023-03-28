package system

import (
	"calendar_service/global"
	"calendar_service/model/system/request"
	"context"
	"gorm.io/gorm"
)

var initializers initSlice

type DBInitHandler interface {
	EnsureDB(ctx context.Context, config *request.InitDB) (context.Context, error)
	InitTables(ctx context.Context, inits initSlice) error
	InitData(ctx context.Context, inits initSlice) error
}

type CalendarDBService struct {
}

func (d *CalendarDBService) InitDB(config request.InitDB) (err error) {
	// context 的作用
	ctx := context.TODO()
	var initHandler DBInitHandler
	initHandler = NewMysqlInitHandler()
	ctx = context.WithValue(ctx, "dbType", "mysql")
	ctx, err = initHandler.EnsureDB(ctx, &config)
	if err != nil {
		return err
	}
	db := ctx.Value("db").(*gorm.DB)
	global.CalendarDB = db

	if err = initHandler.InitTables(ctx, initializers); err != nil {
		return err
	}

	if err = initHandler.InitData(ctx, initializers); err != nil {
		return err
	}
	return nil
}
