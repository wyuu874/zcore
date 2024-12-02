package gormx

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

// GormEngine GORM 数据库引擎
type GormEngine = gorm.DB

// instance 定义一个数据库实例结构体
type instance struct {
	engine *GormEngine
}

// Config 数据库配置
type Config struct {
	Debug        bool   // 是否开启调试模式
	Host         string // 服务地址
	Port         string // 服务端口
	User         string // 用户名
	Pass         string // 密码
	Name         string // 数据库名称
	Char         string // 字符集
	MaxIdleConns int    // 最大空闲连接数
	MaxOpenConns int    // 最大打开连接数
	MaxLifetime  int    // 连接最大生命周期
}

var (
	defaultInstance *instance
	once            sync.Once
	mu              sync.RWMutex
)

// Initialize 初始化数据库连接
func Initialize(conf Config) error {
	var err error
	once.Do(func() {
		err = initDB(conf)
	})
	return err
}

// initDB 初始化数据库连接
func initDB(conf Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.Name,
		conf.Char,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 newLogger(conf.Debug),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return fmt.Errorf("连接数据库失败: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("获取数据库实例失败: %v", err)
	}

	// 设置连接池
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(conf.MaxLifetime) * time.Second)

	defaultInstance = &instance{engine: db}
	return nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	mu.RLock()
	defer mu.RUnlock()

	if defaultInstance == nil {
		panic("数据库未初始化")
	}
	return defaultInstance.engine
}

// Close 关闭数据库连接
func Close() error {
	mu.Lock()
	defer mu.Unlock()

	if defaultInstance != nil {
		sqlDB, err := defaultInstance.engine.DB()
		if err != nil {
			return err
		}
		err = sqlDB.Close()
		if err != nil {
			return err
		}
		defaultInstance = nil
	}
	return nil
}
