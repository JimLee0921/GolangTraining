package main

import "fmt"

type Task struct {
	Name  string
	Count int
}

func (t *Task) Run(n int) {
	t.Count += n
	fmt.Println("Running:", t.Name, "Count:", t.Count)
}

func main() {
	t := Task{Name: "Download"}
	// 绑定方法值
	runLater := t.Run // 绑定实例

	// 可以随时调用（t 被捕获在闭包里）
	runLater(3) // Running: Download Count: 3
	runLater(2) // Running: Download Count: 5

	// 延迟执行
	defer t.Run(1) //Running: Download Count: 6（退出时）
}
