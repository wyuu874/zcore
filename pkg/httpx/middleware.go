package httpx

import (
	"fmt"
	"github.com/wyuu874/zcore/internal/ginx"
	"github.com/wyuu874/zcore/pkg/config"
	"github.com/wyuu874/zcore/pkg/logger"
	"go.uber.org/zap"
	"runtime/debug"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware() ginx.HandlerFunc {
	return ginx.Middleware.Logger(func(params ginx.LogFormatterParams) string {
		// 更详细的日志格式
		return fmt.Sprintf("[GIN] %s |%d| %s | %s | %s | %d | %s | %s | %s\n",
			params.TimeStamp.Format("2006/01/02 - 15:04:05"),
			params.StatusCode,
			params.Latency,
			params.ClientIP,
			params.Method,
			params.BodySize,
			params.Path,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	})
}

// RecoveryMiddleware 恢复中间件
func RecoveryMiddleware() ginx.HandlerFunc {
	return func(c *ginx.GinContext) {
		defer func() {
			if err := ginx.Middleware.Recovery(); err != nil {
				// 获取堆栈信息
				stack := debug.Stack()

				// 记录错误日志
				logger.Error("[Recovery] panic recovered:", zap.Error(err.(error)), zap.String("stack", string(stack)))

				// 返回 500 错误响应
				ApiInternal(c, "服务器内部错误", nil)
			}
		}()

		c.Next()
	}
}

// I18nMiddleware 国际化中间件
func I18nMiddleware() ginx.HandlerFunc {
	return func(c *ginx.GinContext) {
		// 获取语言
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			// 如果语言为空，则从配置中获取
			lang = config.GetString("locale.default_lang")
		}
		// 设置语言
		c.Set("lang", lang)

		c.Next()
	}
}

// CustomMiddleware 自定义中间件
func CustomMiddleware(f func(c *ginx.GinContext)) ginx.HandlerFunc {
	return func(c *ginx.GinContext) {
		f(c)
		c.Next()
	}
}
