package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/index", func(context *gee.Context) {
		context.HTML(http.StatusOK, "<h1>Hi! this is index</h1>")
	})
	// 使用路由分组
	v1 := r.Group("/v1")
	// 加个视觉结构块
	{
		v1.GET("/", func(context *gee.Context) {
			context.HTML(http.StatusOK, "<h1>Hello Gee v1</h1>")
		})
		v1.GET("/hello", func(context *gee.Context) {
			// expect /hello?name=JimLee
			context.String(http.StatusOK, "hello %s, you're at %s\n", context.Query("name"), context.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/JimLee
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	// 启动服务
	r.Run(":9999")
}

/*
对路由进行分组控制，比如：
以/post开头的路由匿名可访问
以/admin开头的路由需要鉴权
以/api开头的路由是 RESTful 接口，可以对接第三方平台，需要三方平台鉴权
*/
