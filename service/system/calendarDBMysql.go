package system

import (
	"calendar_service/model/system/request"
	"context"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInitHandler struct {
}

func NewMysqlInitHandler() *MysqlInitHandler {
	return &MysqlInitHandler{}
}

func (h MysqlInitHandler) EnsureDB(ctx context.Context, config *request.InitDB) (next context.Context, err error) {
	c := config.ToMysqlConfig()
	// 在上下文传递config
	next = context.WithValue(ctx, "config", c)
	dsn := config.MysqlDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8 COLLATE utf8_General_ci ", c.DbName)
	err = createDatabase(dsn, "mysql", createSql)
	if err != nil {
		return nil, err
	}
	var db *gorm.DB
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: c.Dsn(),
	}), &gorm.Config{
		// 初始化过程禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	next = context.WithValue(next, "db", db)
	return next, err

}

func (h MysqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}
func (h MysqlInitHandler) InitData(ctx context.Context) error {
	_, cancel := context.WithCancel(ctx)
	defer func(c func()) {
		c()
	}(cancel)
	return nil
}

/*
todo
这里定义了 初始化时候的对象
order 用于初始化时的顺序处理
SubInitializer
*/
type orderedInitializer struct {
	order int
	SubInitializer
}
type initSlice []*orderedInitializer

func (a initSlice) Len() int {
	return len(a)
}

func (a initSlice) Less(i, j int) bool {
	return a[i].order < a[j].order
}
func (a initSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func() {
		db.Close()
	}()

	_, err = db.Exec(createSql)
	return err
}
func createTables(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) {
		c()
	}(cancel)
	for _, init := range inits {
		// 检查表是否存在
		// 存在就跳过
		if init.TableCreated(next) {
			continue
		}
		n, err := init.MigrateTable(ctx)
		if err != nil {
			return err
		}
		next = n
	}
	return nil
}
