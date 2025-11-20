package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 快捷类型别名（在后续 JSON 输出时方便使用）
type H map[string]any

/*
Context 上下文对象，封装了请求和响应信息
Writer：封装响应的写入端
Req：封装请求
Path / Method：从请求中提取的关键信息
StatusCode：记录响应状态（方便日志等）
*/
type Context struct {
	// 原始字段
	Writer http.ResponseWriter
	Req    *http.Request
	// 请求信息
	Path   string
	Method string
	Params map[string]string
	// 响应信息
	StatusCode int
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// newContext 构造函数，每次有请求到来时，框架就会创建一个新的 Context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm 请求相关辅助方法，从请求中获取 POST 表单数据
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 请求相关辅助方法，从请求 URL 查询参中获取 GET 参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 响应相关辅助方法，设置响应状态码（例如 200, 404, 500）
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 响应相关辅助方法，设置响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 快捷设置响应内容
func (c *Context) String(code int, format string, value ...any) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, value...)))
}

// JSON 快捷设置 JSON 响应内容，设置 Header 为 application/json
func (c *Context) JSON(code int, obj any) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 快捷设置二进制响应内容
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 快捷设置 HTML 响应内容
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
