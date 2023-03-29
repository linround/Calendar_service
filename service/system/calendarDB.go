package system

import (
	"calendar_service/global"
	"calendar_service/model/system/request"
	"context"
	"errors"
	"gorm.io/gorm"
)

var (
	InitOrderSystem = 10
)

var ErrMissingDBContext = errors.New("missing db in context")

var initializers initSlice

type SubInitializer interface {
	InitializerName() string
	DataInserted(ctx context.Context) bool
	TableCreated(ctx context.Context) bool
	InitializeData(ctx context.Context) (next context.Context, err error)
}

type DBInitHandler interface {
	EnsureDB(ctx context.Context, config *request.InitDB) (context.Context, error)
	InitTables(ctx context.Context, inits initSlice) error
	InitData(ctx context.Context, inits initSlice) error
}

type CalendarDBService struct {
}

func RegisterInit(order int, initalizer SubInitializer) {
	if initializers == nil {
		initializers = initSlice{}
	}
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
