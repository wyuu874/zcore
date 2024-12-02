package httpx

import (
	"github.com/wyuu874/zcore/internal/ginx"
	"github.com/wyuu874/zcore/internal/ginx/route"
	"net/http"
)

// NewRouter 创建路由注册器
func NewRouter(engine *ginx.GinEngine) *route.Router {
	return route.NewRouter(engine)
}

// WithMiddlewares 添加中间件
func WithMiddlewares(middlewares ...ginx.HandlerFunc) route.RouteOption {
	return func(r *route.Route) {
		r.Middlewares = append(r.Middlewares, middlewares...)
	}
}

// NewRoute 创建路由
func NewRoute(method, path string, handler ginx.HandlerFunc, opts ...route.RouteOption) route.Route {
	route := route.Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}

	for _, opt := range opts {
		opt(&route)
	}

	return route
}

// GET 创建 GET 路由
func GET(path string, handler ginx.HandlerFunc, opts ...route.RouteOption) route.Route {
	return NewRoute(http.MethodGet, path, handler, opts...)
}

// POST 创建 POST 路由
func POST(path string, handler ginx.HandlerFunc, opts ...route.RouteOption) route.Route {
	return NewRoute(http.MethodPost, path, handler, opts...)
}

// NewGroup 创建路由组
func NewGroup(prefix string, opts ...route.GroupOption) route.Group {
	group := route.Group{
		Prefix: prefix,
	}

	for _, opt := range opts {
		opt(&group)
	}

	return group
}

// WithGroupMiddlewares 添加组中间件
func WithGroupMiddlewares(middlewares ...ginx.HandlerFunc) route.GroupOption {
	return func(g *route.Group) {
		g.Middlewares = append(g.Middlewares, middlewares...)
	}
}

// WithRoutes 添加路由
func WithRoutes(routes ...route.Route) route.GroupOption {
	return func(g *route.Group) {
		g.Routes = append(g.Routes, routes...)
	}
}

// WithGroups 添加子组
func WithGroups(groups ...route.Group) route.GroupOption {
	return func(g *route.Group) {
		g.Groups = append(g.Groups, groups...)
	}
}
