package main

import (
	"fmt"
	"time"
)

// main 回调函数示例
func main() {
	// 定义一个回调函数
	callbackFun := func(taskName string, err error) {
		if err != nil {
			fmt.Println(taskName, "任务执行出现错误", err)
		} else {
			fmt.Println(taskName, "任务执行成功")
		}
	}
	// 模拟执行任务
	runTask("task1", callbackFun)
	runTask("task2", callbackFun)
	runTask("task3", callbackFun)
	runTask("task4", callbackFun)
}

// 定义一个执行任务的函数，接收任务和一个处理错误的回调函数
func runTask(taskName string, callback func(string, error)) {
	fmt.Println("开始执行任务:", taskName)
	// 模拟任务耗时
	time.Sleep(2 * time.Second)

	// 模拟结果
	if taskName == "task2" {
		// 模拟错误
		callback(taskName, fmt.Errorf("任务执行失败: %s", taskName))
		return
	}
	// 执行成功
	callback(taskName, nil)
}
