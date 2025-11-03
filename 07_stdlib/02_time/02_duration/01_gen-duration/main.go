package main

import (
	"fmt"
	"time"
)

func DemoOne() {
	// 数字 × 时间单位常量生成 Duration 对象
	d1 := 10 * time.Second       // 10秒
	d2 := 500 * time.Millisecond // 500 毫秒（0.5秒）
	d3 := 2 * time.Hour          // 2小时
	fmt.Println(d1)              // 10s
	fmt.Println(d2)              // 500ms
	fmt.Println(d3)              // 2h0m0s
}

func DemoTwo() {
	// 时间差计算返回 Duration 对象
	t1 := time.Now()
	t2 := t1.Add(75 * time.Minute)

	d := t2.Sub(t1)
	fmt.Println(d) // 1h15m0s
}

func DemoThree() {
	/*
		time.ParseDuration() 进行字符串解析
			配置文件 （config.yaml / JSON）
			环境变量
			命令行参数
			用户输入
	*/
	d, err := time.ParseDuration("1h30m")
	if err != nil {
		panic(err)
	}
	fmt.Println(d) // 1h30m0s
}

func main() {
	DemoOne()
	DemoTwo()
	DemoThree()
}
