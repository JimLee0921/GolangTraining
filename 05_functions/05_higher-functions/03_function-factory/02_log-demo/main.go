package main

import "fmt"

func makeLogger(tag string) func(string) {
	return func(msg string) {
		fmt.Printf("[%s] %s\n", tag, msg)
	}
}
func main() {
	/*
		工厂函数返回一个闭包
		每个闭包“记住”了自己的 tag
		实际开发中常用于日志、HTTP、调度系统
	*/
	infoLog := makeLogger("INFO")
	errorLog := makeLogger("ERROR")

	infoLog("System started")
	errorLog("Disk full")
}
