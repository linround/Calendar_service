package system

import (
	"calendar_service/model/system/request"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInitHandler struct {
}

func NewMysqlInitHandler() *MysqlInitHandler {
	return &MysqlInitHandler{}
}

type SubInitializer interface {
	InitializerName() string
	DataInserted(ctx context.Context) bool
	TableCreated(ctx context.Context) bool
	InitializeData(ctx context.Context) (next context.Context, err error)
}

func (h MysqlInitHandler) EnsureDB(ctx context.Context, config *request.InitDB) (next context.Context, err error) {
	c := config.ToMysqlConfig()
	// 在上下文传递config
	next = context.WithValue(ctx, "config", c)
	dsn := config.MysqlDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS '%s'", c.DbName)
	err = createDatabase(dsn, "mysql", createSql)
	if err != nil {
		return nil, err
	}
	var db *gorm.DB
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: c.Dsn(),
	}))
	next = context.WithValue(next, "db", db)
	return next, err

}

func (h MysqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}
func (h MysqlInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) {
		c()
	}(cancel)
	for _, init := range inits {
		if init.DataInserted(next) {
			continue
		}
		n, err := init.InitializeData(next)
		if err != nil {
			return err
		}
		next = n
	}
	return nil
}

type orderedInitializer struct {
	order int
	SubInitializer
}
type initSlice []*orderedInitializer

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
		if init.TableCreated(next) {
			continue
		}
		return errors.New("创建表失败")
	}
	return nil
}
