package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(context *gee.Context) {
		context.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(context *gee.Context) {
		// expect /hello?name=jimlee
		context.String(http.StatusOK, "hello %s, you're at %s\n", context.Query("name"), context.Path)
	})
	r.GET("/hello/:name", func(context *gee.Context) {
		// expect /hello/jimlee
		context.String(http.StatusOK, "hello %s, you are at %s", context.Param("name"), context.Path)
	})
	r.GET("/assets/*filepath", func(context *gee.Context) {
		context.JSON(http.StatusOK, gee.H{"filename": context.Param("filepath")})
	})

	r.Run(":9999")
}

/*
使用 Trie 树使得 gee 框架支持动态路由
:name 表示当前片段是一个参数，例如 /p/:lang/doc 可匹配 /p/go/doc。
*filepath 通配符模式，表示匹配后续所有片段，例如 /static/*filepath 匹配 /static/css/main.css。
*/
