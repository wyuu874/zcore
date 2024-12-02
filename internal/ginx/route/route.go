package route

import (
	"github.com/wyuu874/zcore/internal/ginx"
)

// Route 定义单个路由
type Route struct {
	Method      string
	Path        string
	Handler     ginx.HandlerFunc
	Middlewares []ginx.HandlerFunc
}

// Group 定义路由组
type Group struct {
	Prefix      string
	Routes      []Route
	Groups      []Group
	Middlewares []ginx.HandlerFunc
}

// RouteOption 路由选项函数
type RouteOption func(*Route)

// GroupOption 路由组选项函数
type GroupOption func(*Group)
