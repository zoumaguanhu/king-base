package pg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"king.com/king/base/common/gorm_pkg"
	"time"
)

type Db struct {
	DSN     string
	DbCPool *gorm_pkg.GormConnPoolConf
	Mode    string `json:"mode,optional,default=dev"`
}

func New(dsn string, mode string, c *gorm_pkg.GormConnPoolConf) *Db {
	return &Db{DSN: dsn, Mode: mode, DbCPool: c}
}
func (d *Db) C(mode string) *gorm.DB {
	conn, err := gorm.Open(
		postgres.Open(d.DSN),
		&gorm.Config{Logger: gorm_pkg.NewGormLogger(mode)},
	)
	if err != nil {
		fmt.Println("error:", err)
	}

	if sqlDB, err := conn.DB(); err == nil {
		pool := d.DbCPool
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(pool.MaxIdleConns)
		sqlDB.SetMaxOpenConns(pool.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Duration(pool.MaxLifeTime) * time.Second)
	}

	return conn
}
