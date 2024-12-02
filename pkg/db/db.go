package db

import (
	"github.com/wyuu874/zcore/internal/gormx"
	"github.com/wyuu874/zcore/pkg/config"
)

// DB 获取数据库实例
func DB() *gormx.GormEngine {
	return gormx.GetDB()
}

// Init 初始化数据库连接
func Init() error {
	conf := config.Database{}
	config.GetConfig("database", &conf)
	gormxConfig := gormx.Config{
		Debug:        conf.Debug,
		Host:         conf.Host,
		Port:         conf.Port,
		User:         conf.Username,
		Pass:         conf.Password,
		Name:         conf.Database,
		Char:         conf.Charset,
		MaxIdleConns: conf.MaxIdleConns,
		MaxOpenConns: conf.MaxOpenConns,
		MaxLifetime:  conf.MaxLifetime,
	}
	return gormx.Initialize(gormxConfig)
}

// Close 关闭数据库连接
func Close() error {
	return gormx.Close()
}
