package gee

import (
	"log"
	"net/http"
)

// HandlerFunc 定义自己的框架接口层
type HandlerFunc func(*Context)

// Engine 结构体，Gee 框架的核心引擎，类似 http.ServeMux，存储路由映射
type Engine struct {
	*RouterGroup // 字段嵌入
	router       *router
	groups       []*RouterGroup // 存储所有的分组
}

/*
RouterGroup 路由分组
每一个 RouterGroup 表示一组路由（有相同前缀）
比如 /v1、/v2、/admin
它也可以嵌套，比如 /v1/admin
*/
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // 支持中间件（下个章节）
	parent      *RouterGroup  // 支持父子关系
	engine      *Engine       // 所有分组共享一个 Engine
}

// New 对外构造函数，初始化 Engine 对象
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 方法用来创建新的路由分组，所有分组共享同一个 Engine 实例
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: engine.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// addRoute 添加路由规则
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET 对外的注册 GET 请求路由接口
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 对外的注册 POST 请求路由接口
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run 启动服务器
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 实现 http.Handler 接口（核心）
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
