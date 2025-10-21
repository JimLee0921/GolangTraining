package main

import (
	"fmt"
)

func main() {
	/*
		非阻塞操作（default 分支）
		如果 case 都不满足且有 default 关键字则走 default 逻辑
	*/
	ch := make(chan int, 1)

	// 初始写入一个值，占满缓冲
	ch <- 19
	fmt.Println("initial data: 19")

	// 使用 select 尝试再写入
	select {
	case ch <- 10:
		fmt.Println("write data 10")
	default:
		fmt.Println("channel is full")
	}

	// 再读出一个数据
	val := <-ch
	fmt.Println("read: ", val)

	// 再次尝试写入，此时有空间
	select {
	case ch <- 20:
		fmt.Println("write data 20")
	default:
		fmt.Println("channel is full")
	}

	// 读出最后的值
	val2 := <-ch
	fmt.Println("finial data: ", val2)
}
