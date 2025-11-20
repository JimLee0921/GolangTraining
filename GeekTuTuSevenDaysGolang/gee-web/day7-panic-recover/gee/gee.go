package gee

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

// HandlerFunc 定义自己的框架接口层
type HandlerFunc func(*Context)

// Engine 结构体，Gee 框架的核心引擎，类似 http.ServeMux，存储路由映射
type Engine struct {
	*RouterGroup // 字段嵌入
	router       *router
	groups       []*RouterGroup // 存储所有的分组
	// HTML 模板渲染
	htmlTemplates *template.Template
	funcMap       template.FuncMap
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

// Default 默认使用 Logger() 和 Recovery 中间件
func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
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

// ServeHTTP 收集所有中间件
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 请求进来时决定处理哪些中间件
	var middlewares []HandlerFunc
	// 遍历所有分组，判断哪个分组的 prefix 与当前 URL 匹配
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			// 把匹配到的 middlewares 放入 Context
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = engine
	// 让它继续去匹配路由、执行 handler
	engine.router.handle(c)
}

// Use 中间件应用到某个 Group
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// createStaticHandler 创建一个静态文件处理器
// relativePath: 相对于路由组的路径前缀
// fs: 静态文件系统
// 返回值: 处理静态文件请求的HandlerFunc
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(context *Context) {
		file := context.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			context.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(context.Writer, context.Req)
	}
}

// Static 注册静态文件服务路由
// relativePath: URL路径前缀，用于访问静态文件
// root: 本地文件系统目录路径，作为静态文件的根目录
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// 注册 GET handlers
	group.GET(urlPattern, handler)
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}
