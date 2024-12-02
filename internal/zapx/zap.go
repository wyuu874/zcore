package zapx

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

// Field 是 zap 的日志字段
type Field = zap.Field

// Logger 是 zap 的日志记录器
type Logger = zap.Logger

// InitLogger 初始化日志记录器
func InitLogger(level string, channel string, logPath string, maxsize int, maxAge int, maxBackup int) *Logger {
	// 创建日志核心
	cores := []zapcore.Core{}

	// 创建日志级别
	atomicLevel := newLevel(level)

	// 如果channel为console或logPath为空，则输出到控制台，否则输出到文件
	if channel == "console" || logPath == "" {
		cores = append(cores, newConsoleCore(atomicLevel))
	} else if channel == "daily" {
		cores = append(cores, newDailyFileCore(atomicLevel, logPath, maxsize, maxAge, maxBackup))
	} else {
		cores = append(cores, newSingleFileCore(atomicLevel, logPath))
	}
	core := zapcore.NewTee(cores...)

	// 创建logger
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

// newLevel 创建日志级别
func newLevel(level string) zap.AtomicLevel {
	// 如果level为空，则设置为info
	if level == "" {
		level = "info"
	}

	// 日志级别映射
	levelMap := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}

	// 创建atomicLevel
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(levelMap[level])

	return atomicLevel
}

// newEncoderConfig 创建编码器配置
func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// newConsoleCore 创建控制台核心
func newConsoleCore(atomicLevel zap.AtomicLevel) zapcore.Core {
	// 创建编码器配置
	encoderConfig := newEncoderConfig()

	// 创建编码
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		atomicLevel,
	)
}

// newSingleFileCore 创建单个文件核心
func newSingleFileCore(atomicLevel zap.AtomicLevel, logPath string) zapcore.Core {
	// 确保日志目录存在
	_ = newLogDir(logPath)

	// 创建日志文件
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("创建日志文件失败: %v", err))
	}

	// 创建编码器
	encoderConfig := newEncoderConfig()

	// 创建日志核心
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(logFile),
		atomicLevel,
	)
}

// newDailyFileCore 创建每日文件核心
func newDailyFileCore(atomicLevel zap.AtomicLevel, logPath string, maxsize int, maxAge int, maxBackup int) zapcore.Core {
	// 确保日志目录存在
	logDir := newLogDir(logPath)

	// 获取当前日期
	now := time.Now()
	date := now.Format("2006-01-02")

	// 创建日志文件
	filename := filepath.Join(logDir, fmt.Sprintf("%s.log", date))

	// 创建lumberjack日志记录器
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  // 文件位置
		MaxSize:    maxsize,   // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxAge:     maxAge,    // 保留旧文件的最大天数
		MaxBackups: maxBackup, // 保留旧文件的最大个数
		Compress:   false,     // 是否压缩/归档旧文件
	}

	// 创建编码器
	encoderConfig := newEncoderConfig()

	// 创建日志核心
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(lumberJackLogger),
		atomicLevel,
	)
}

// newLogDir 创建日志目录
func newLogDir(logPath string) string {
	wd, _ := os.Getwd()
	logDir := filepath.Join(wd, filepath.Dir(logPath))
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(fmt.Sprintf("创建日志目录失败: %s", err.Error()))
	}
	return logDir
}
