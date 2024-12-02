package ginx

import (
	"github.com/gin-gonic/gin"
)

type middleware struct{}

// Middleware 中间件
var Middleware = &middleware{}

// Logger 日志中间件
func (m *middleware) Logger(formatter func(params LogFormatterParams) string) HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(params LogFormatterParams) string {
			return formatter(params)
		},
	})
}

// Recovery 恢复中间件
func (m *middleware) Recovery() any {
	return recover()
}
