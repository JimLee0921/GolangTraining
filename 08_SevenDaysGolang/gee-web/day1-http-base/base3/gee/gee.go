package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义自己的框架接口层
/*
可以在将来提供：中间件机制（next() 调用链）
上下文对象（Context 包含路径参数、JSON 解析、状态码设置）
路由分组（Group）
等
*/
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Engine 结构体，Gee 框架的核心引擎，类似 http.ServeMux，存储路由映射
type Engine struct {
	router map[string]HandlerFunc
}

// New 对外构造函数，初始化 Engine 对象（带空路由表）方便用户一行构建框架实例
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute 添加路由规则，内部方法
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
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
/*
这是整个框架的请求分发逻辑
当有请求进来时：

1. 组合 key，例如 "GET-/hello"
2. 去 router 里查找对应的 handler
3. 如果找到 -> 执行 handler
4. 找不到 -> 返回 404
*/
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
