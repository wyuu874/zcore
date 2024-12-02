package route

import "github.com/gin-gonic/gin"

// Router 路由注册器
type Router struct {
	engine *gin.Engine
}

// NewRouter 创建路由注册器
func NewRouter(engine *gin.Engine) *Router {
	return &Router{engine: engine}
}

// Register 注册路由组
func (r *Router) Register(groups ...Group) {
	for _, group := range groups {
		r.registerGroup(r.engine, group)
	}
}

// registerGroup 注册单个路由组
func (r *Router) registerGroup(engine gin.IRouter, group Group) {
	// 创建路由组
	g := engine.Group(group.Prefix, group.Middlewares...)

	// 注册路由
	for _, route := range group.Routes {
		handlers := append(route.Middlewares, route.Handler)
		g.Handle(route.Method, route.Path, handlers...)
	}

	// 递归注册子组
	for _, subGroup := range group.Groups {
		r.registerGroup(g, subGroup)
	}
}
