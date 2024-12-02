package rdb

import (
	"github.com/wyuu874/zcore/internal/redisx"
	"github.com/wyuu874/zcore/pkg/config"
)

// RDB 获取Redis实例
func RDB() *redisx.RedisEngine {
	return redisx.GetRDB()
}

// Init 初始化Redis连接
func Init() error {
	conf := config.Redis{}
	config.GetConfig("redis", &conf)
	redisConfig := redisx.Config{
		Host: conf.Host,
		Port: conf.Port,
		Pass: conf.Pass,
		DB:   conf.DB,
	}
	return redisx.Initialize(redisConfig)
}

// Close 关闭Redis连接
func Close() error {
	return redisx.Close()
}
