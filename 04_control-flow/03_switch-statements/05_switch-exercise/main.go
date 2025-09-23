package main

import "fmt"

func classifyStatusRange(code int) string {
	switch {
	case code >= 100 && code < 200:
		return "信息(1xx)"
	case code >= 200 && code < 300:
		return "成功(2xx)"
	case code >= 300 && code < 400:
		return "跳转(3xx)"
	case code >= 400 && code < 500:
		return "客户端错误(4xx)"
	case code >= 500 && code < 600:
		return "服务端错误(5xx)"
	default:
		return "未知/无效状态码"
	}
}

// main HTTP 状态码判断返回信息
func main() {
	fmt.Println(classifyStatusRange(200))
	fmt.Println(classifyStatusRange(302))
	fmt.Println(classifyStatusRange(500))

}
