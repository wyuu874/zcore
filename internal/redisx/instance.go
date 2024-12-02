package redisx

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/wyuu874/zcore/pkg/logger"
	"go.uber.org/zap"
	"sync"
)

// RedisEngine Redis 引擎
type RedisEngine = redis.Client

// instance Redis 实例
type instance struct {
	engine *RedisEngine
}

// Config Redis 配置
type Config struct {
	Host string // 服务地址
	Port int    // 服务端口
	Pass string // 密码
	DB   int    // 数据库
}

var (
	defaultInstance *instance
	once            sync.Once
	mu              sync.RWMutex
)

// Initialize 初始化Redis连接
func Initialize(conf Config) error {
	var err error
	once.Do(func() {
		err = initRDB(conf)
	})
	return err
}

// initRDB 初始化Redis连接
func initRDB(conf Config) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Pass,
		DB:       conf.DB,
	})

	logger.Info("测试redis连接", zap.String("addr", fmt.Sprintf("%s:%d", conf.Host, conf.Port)))

	// 测试连接
	result, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("redis连接失败: %v", err)
	}

	logger.Info("redis连接成功", zap.String("result", result))

	defaultInstance = &instance{engine: rdb}
	return nil
}

// GetRDB 获取Redis实例
func GetRDB() *RedisEngine {
	mu.RLock()
	defer mu.RUnlock()

	if defaultInstance == nil {
		panic("Redis未初始化")
	}
	return defaultInstance.engine
}

// Close 关闭Redis连接
func Close() error {
	mu.Lock()
	defer mu.Unlock()

	if defaultInstance != nil {
		err := defaultInstance.engine.Close()
		if err != nil {
			return err
		}

		defaultInstance = nil
	}
	return nil
}
