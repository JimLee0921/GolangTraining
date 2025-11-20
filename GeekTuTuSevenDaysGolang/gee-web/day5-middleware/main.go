package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(context *gee.Context) {
		// 开始时间
		t := time.Now()
		// 模拟发生 500 错误
		context.Fail(http.StatusInternalServerError, "Internal Server error")
		// 计算用时
		log.Printf("[%d] %s in %v for group v2", context.StatusCode, context.Req.URL, time.Since(t))
	}
}

func main() {
	r := gee.New()
	// 全局中间件
	r.Use(gee.Logger())
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
	v2.Use(onlyForV2()) // v2 路由组中间件
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
中间件设计，非业务的技术类组件，允许用户自己定义功能，对 handler 进行加工，主要用于
日志记录（记录请求起止时间、路径、状态码）
身份验证 / 鉴权
恶意请求过滤
统一处理响应头 / 错误处理

本质就是定义 HandlerFunc 扩展、在请求链中按顺序执行
*/
