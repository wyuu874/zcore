package httpx

import (
	"context"
	"fmt"
	"github.com/wyuu874/zcore/internal/ginx"
	"github.com/wyuu874/zcore/internal/i18nx"
	"github.com/wyuu874/zcore/internal/viperx"
	"github.com/wyuu874/zcore/pkg/config"
	"github.com/wyuu874/zcore/pkg/db"
	"github.com/wyuu874/zcore/pkg/logger"
	"github.com/wyuu874/zcore/pkg/rdb"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	shutdownTimeout = 60 * time.Second
	startupTimeout  = 5 * time.Second
)

// Context 上下文
type Context = ginx.GinContext

// Engine 引擎
var Engine *ginx.GinEngine

// 初始化配置和服务
func init() {
	initConfig()
	initEngine()
}

// 初始化配置
func initConfig() {
	// 初始化配置文件
	viperx.InitConfig()

	// 初始化国际化
	var localeConfig config.Locale
	config.GetConfig("locale", &localeConfig)
	i18nx.InitI18n(localeConfig.DefaultLang, localeConfig.Dir)
}

// 初始化引擎
func initEngine() {
	// 初始化 Gin 引擎
	Engine = ginx.GetEngine(config.GetBool("app.debug"))

	// 注册全局中间件
	Engine.Use(
		I18nMiddleware(),     // 国际化
		LoggerMiddleware(),   // 日志
		RecoveryMiddleware(), // 恢复
	)
}

// 初始化服务依赖
func initDependencies() error {
	// 初始化日志记录器
	logger.Init()

	// 初始化数据库连接
	if err := db.Init(); err != nil {
		return fmt.Errorf("数据库初始化失败: %w", err)
	}

	// 初始化Redis连接
	if err := rdb.Init(); err != nil {
		return fmt.Errorf("redis初始化失败: %w", err)
	}

	return nil
}

// 创建 HTTP 服务器
func newServer(conf config.App) *http.Server {
	return &http.Server{
		Addr:           conf.Host + ":" + conf.Port,
		Handler:        Engine,
		ReadTimeout:    time.Duration(conf.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(conf.WriteTimeout) * time.Second,
		IdleTimeout:    time.Duration(conf.IdleTimeout) * time.Second,
		MaxHeaderBytes: conf.MaxHeaderBytes << 20,
	}
}

// Run 启动 HTTP 服务
func Run() {
	// 初始化依赖
	if err := initDependencies(); err != nil {
		logger.Fatal("服务依赖初始化失败", zap.Error(err))
	}
	defer func() {
		logger.Logger.Sync()
		db.Close()
		rdb.Close()
	}()

	// 加载应用配置
	var appConfig config.App
	config.GetConfig("app", &appConfig)

	// 创建服务器
	srv := newServer(appConfig)

	// 启动服务器
	startServer(srv, appConfig)

	// 等待中断信号
	waitForShutdown(srv)
}

// 启动服务器
func startServer(srv *http.Server, conf config.App) {
	go func() {
		logger.Info("正在启动HTTP服务...",
			zap.String("host", conf.Host),
			zap.String("port", conf.Port),
		)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("HTTP服务启动失败",
				zap.String("addr", srv.Addr),
				zap.Error(err),
			)
		}
	}()

	// 等待服务器启动
	time.Sleep(startupTimeout)
	logger.Info("HTTP服务启动成功",
		zap.String("addr", srv.Addr),
		zap.Duration("read_timeout", srv.ReadTimeout),
		zap.Duration("write_timeout", srv.WriteTimeout),
		zap.Duration("idle_timeout", srv.IdleTimeout),
	)
}

// 等待服务器关闭
func waitForShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit

	logger.Info("收到关闭信号，正在关闭HTTP服务...",
		zap.String("signal", sig.String()),
	)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("HTTP服务关闭异常",
			zap.Error(err),
			zap.Duration("timeout", shutdownTimeout),
		)
		return
	}

	logger.Info("HTTP服务已成功关闭")
}
