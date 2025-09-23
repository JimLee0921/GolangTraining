package main

import "fmt"

// main for range 是 go 中的特色写法
func main() {
	/*
		for range 是 Go 用来 遍历集合 的语法糖，可以用来循环：
			数组（array）：对应的就是索引和值
			切片（slice）：对应的就是索引和值
			map：对应的就是键和值，遍历 map 时顺序是不固定的
			字符串：遍历的是字节索引和rune（Unicode 码点，可以直接通过 string(r) 转为对应字符），能正确处理中文、emoji 等字符
			channel：for range 会不断从 channel 里取值，直到 channel 被关闭（没有索引）
		它会返回 两个值：
			1. 索引 / 键
			2. 对应的值
			如果只需要其中一个值，可以用 _ 忽略不需要的值
	*/
	// 遍历数组或切片
	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Println("索引: ", i, "值: ", v)
	}

	// 遍历 map
	books := map[int]string{1: "西游记", 2: "红楼梦", 3: "杀死比尔"}
	for k, v := range books {
		fmt.Println("键: ", k, "值: ", v)
	}

	// 遍历 string（注意一个一个中文占三个字节，所以索引并不是+1递增的）
	hello := "Hello, GO 语言"
	for i, r := range hello {
		fmt.Println(i, r, string(r))         // 这样打印的就是 字符索引 和 rune（Unicode 码点）
		fmt.Printf("索引 %d, rune %c\n", i, r) // 使用 %c 占位符也可以转为 对应字符
	}

	// 遍历 channel
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
