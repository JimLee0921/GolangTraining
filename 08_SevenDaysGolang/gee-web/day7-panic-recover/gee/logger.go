package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(context *Context) {
		// 开始时间
		t := time.Now()
		// 执行逻辑
		context.Next()
		// 计算总耗费时间并记录
		log.Printf("[%d] %s in %v", context.StatusCode, context.Req.RequestURI, time.Since(t))
	}
}
