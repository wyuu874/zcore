package logger

import (
	"github.com/wyuu874/zcore/internal/zapx"
	"github.com/wyuu874/zcore/pkg/config"
)

// Field 是 zap 的日志字段
type Field = zapx.Field

// Logger 是 zap 的日志记录器
var Logger *zapx.Logger

// Init 初始化日志记录器
func Init() {
	// 获取日志配置
	loggerConfig := config.Logger{}
	config.GetConfig("logger", &loggerConfig)

	// 初始化日志记录器
	Logger = zapx.InitLogger(
		loggerConfig.Level,
		loggerConfig.Channel,
		loggerConfig.Path,
		loggerConfig.MaxSize,
		loggerConfig.MaxAge,
		loggerConfig.MaxBackups,
	)
}

// Info 记录info级别日志
func Info(msg string, fields ...Field) {
	Logger.Info(msg, fields...)
}

// Error 记录error级别日志
func Error(msg string, fields ...Field) {
	Logger.Error(msg, fields...)
}

// Debug 记录debug级别日志
func Debug(msg string, fields ...Field) {
	Logger.Debug(msg, fields...)
}

// Warn 记录warn级别日志
func Warn(msg string, fields ...Field) {
	Logger.Warn(msg, fields...)
}

// Fatal 记录fatal级别日志
func Fatal(msg string, fields ...Field) {
	Logger.Fatal(msg, fields...)
}
