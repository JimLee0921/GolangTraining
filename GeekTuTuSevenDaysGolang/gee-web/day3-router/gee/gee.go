package gee

import (
	"net/http"
)

// HandlerFunc 定义自己的框架接口层
type HandlerFunc func(*Context)

// Engine 结构体，Gee 框架的核心引擎，类似 http.ServeMux，存储路由映射
type Engine struct {
	router *router
}

// New 对外构造函数，初始化 Engine 对象
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 添加路由规则，内部方法
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET 对外的注册 GET 请求路由接口
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 对外的注册 POST 请求路由接口
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
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
