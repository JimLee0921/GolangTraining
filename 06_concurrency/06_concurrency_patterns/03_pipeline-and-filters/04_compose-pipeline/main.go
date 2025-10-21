package main

import (
	"fmt"
	"strings"
)

// Filter 类型：通用函数签名
type Filter func(<-chan string) <-chan string

func Compose(filters ...Filter) Filter {
	return func(in <-chan string) <-chan string {
		out := in
		for _, f := range filters {
			out = f(out)
		}
		return out
	}
}

// 定义几个过滤器
func trimSpace(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			out <- strings.TrimSpace(s)
		}
	}()
	return out
}

func toUpper(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			out <- strings.ToUpper(s)
		}
	}()
	return out
}

func addSuffix(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			out <- s + "_DONE"
		}
	}()
	return out
}

func generator(strs ...string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, s := range strs {
			out <- s
		}
	}()
	return out
}

func main() {
	/*
		高度模块化与复用
		用函数组合（Compose）构建完整 pipeline
		每个 filter 都是可复用组件
		非常接近函数式管道流的风格
	*/
	data := generator(" go  ", " pipeline ", "filters ")
	pipeline := Compose(trimSpace, toUpper, addSuffix)
	for result := range pipeline(data) {
		fmt.Println(result)
	}
}
