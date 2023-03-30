package system

import (
	"calendar_service/global"
	"calendar_service/model/system/request"
	"context"
	"errors"
	"gorm.io/gorm"
	"sort"
)

var (
	InitOrderSystem = 10
)

var ErrMissingDBContext = errors.New("缺少上下文")

var initializers initSlice

type SubInitializer interface {
	TableCreated(ctx context.Context) bool                     // 判断这个表是否已经被创建
	MigrateTable(ctx context.Context) (context.Context, error) // 创建表
}

type DBInitHandler interface {
	EnsureDB(ctx context.Context, config *request.InitDB) (context.Context, error)
	InitTables(ctx context.Context, inits initSlice) error
	InitData(ctx context.Context) error
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
	if len(initializers) == 0 {
		return errors.New("无初始化对象")
	}
	sort.Sort(&initializers)

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

	return nil
}
