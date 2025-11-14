package main

import (
	"fmt"
	"gee"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.New()
	// 全局中间件
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")

	r.Static("/assets", "./static")
	stu1 := &student{
		Name: "JimLee",
		Age:  20,
	}
	stu2 := &student{
		Name: "JamesBond",
		Age:  17,
	}
	r.GET("/", func(context *gee.Context) {
		context.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(context *gee.Context) {
		context.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.GET("/data", func(context *gee.Context) {
		context.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2025, 10, 14, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}

/*
使用 HTML 模板
实现静态资源服务(Static Resource)
支持HTML模板渲染
*/
