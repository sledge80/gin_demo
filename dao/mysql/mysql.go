package mysql

import (
	"fmt"
	"gin_demo/settings"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		settings.Conf.MySQLConfig.User,
		settings.Conf.MySQLConfig.Password,
		settings.Conf.MySQLConfig.Host,
		settings.Conf.MySQLConfig.Port,
		settings.Conf.MySQLConfig.Dbname,
	)
	// 也可以使用MustConnect连接不成功就panic
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed, err:%v\n", zap.Error(err))
		return
	}
	Db.SetMaxOpenConns(settings.Conf.MySQLConfig.MaxOpenConns)
	Db.SetMaxIdleConns(settings.Conf.MySQLConfig.MaxIdleConns)
	return
}

func Close() {
	_ = Db.Close()
}
