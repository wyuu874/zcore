package gormx

import (
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

// newLogger 创建 GORM 日志配置
func newLogger(debug bool) logger.Interface {
	var logLevel logger.LogLevel
	var writer io.Writer = os.Stdout

	if debug {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
		writer = io.Discard // 生产环境不输出 SQL 日志
	}

	return logger.New(
		log.New(writer, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:                  logLevel,               // 日志级别
			IgnoreRecordNotFoundError: true,                   // 忽略记录未找到的错误
			Colorful:                  debug,                  // 彩色输出
		},
	)
}
