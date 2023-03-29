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
	InitializerName() string                                              // 获取对应要创建的表的名字
	DataInserted(ctx context.Context) bool                                // 在这个表中插入数据的方法
	TableCreated(ctx context.Context) bool                                // 判断这个表是否已经被创建
	InitializeData(ctx context.Context) (next context.Context, err error) // 创建初始数据
	MigrateTable(ctx context.Context) (context.Context, error)            // 创建表
}

type DBInitHandler interface {
	EnsureDB(ctx context.Context, config *request.InitDB) (context.Context, error)
	InitTables(ctx context.Context, inits initSlice) error
	InitData(ctx context.Context, inits initSlice) error
}

type CalendarDBService struct {
}

func RegisterInit(order int, i SubInitializer) {
	if initializers == nil {
		initializers = initSlice{}
	}
	ni := orderedInitializer{
		order, i,
	}
	initializers = append(initializers, &ni)
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
