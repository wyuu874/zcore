package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"sync"
)

// GinEngine gin引擎类型别名
type GinEngine = gin.Engine

// GinContext gin上下文类型别名
type GinContext = gin.Context

// HandlerFunc gin处理函数类型别名
type HandlerFunc = gin.HandlerFunc

// LogFormatterParams gin日志格式化参数类型别名
type LogFormatterParams = gin.LogFormatterParams

var (
	// engineInstance 全局引擎实例
	engineInstance *GinEngine
	// engineMutex 确保并发安全的互斥锁
	engineMutex sync.Mutex
)

// GetEngine 获取gin引擎实例
func GetEngine(debug bool) *GinEngine {
	engineMutex.Lock()
	defer engineMutex.Unlock()

	if engineInstance == nil {
		if debug {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
		engineInstance = gin.New()
	}
	return engineInstance
}

// SetValidator 设置请求参数验证器
func SetValidator(validator binding.StructValidator) {
	binding.Validator = validator
}
