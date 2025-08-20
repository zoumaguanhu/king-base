package pg

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"king.com/king/base/common/constants"
	"king.com/king/base/common/gorm_pkg"
	"king.com/king/base/common/secret"
	"time"
)

type Db struct {
	DSN        string
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
	DbCPool    *gorm_pkg.GormConnPoolConf
	Mode       string `json:"mode,optional,default=dev"`
}

func New(dsn string, mode string, c *gorm_pkg.GormConnPoolConf) *Db {
	return &Db{DSN: dsn, Mode: mode, DbCPool: c}
}
func (d *Db) C(mode string) *gorm.DB {
	pass, b := secret.Parse(d.Mode, constants.DB_PASSWORD_FILE)
	if b {
		d.DbPassword = pass
		logx.Infof("db pass:%v", pass)
	}
	dsn := fmt.Sprintf(d.DSN, d.DbHost, d.DbUser, d.DbPassword, d.DbName, d.DbPort)
	conn, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{Logger: gorm_pkg.NewGormLogger(mode)},
	)
	if err != nil {
		logx.Errorf("db error:%v", err)
		panic(err)
	}

	if sqlDB, err := conn.DB(); err == nil {
		pool := d.DbCPool
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(pool.MaxIdleConns)
		sqlDB.SetMaxOpenConns(pool.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Duration(pool.MaxLifeTime) * time.Second)
	}
	logx.Infof("db conn success")
	return conn
}
