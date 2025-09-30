package main

import "fmt"

func main() {
	/*
		channel 操作符
		1. 发送(send)：ch <- value
			把 value 发送到 channel ch
			如果是 无缓冲 channel，发送会阻塞，直到有接收方准备好取走
			如果是 有缓冲 channel，当缓冲区未满时，发送不会阻塞
		2. 接收(receive)：value := <-ch
			从 channel ch 接收一个值，并赋给 value，接收的次序是按照发送的次序，先进先出
			如果没有数据：
				无缓冲 -> 阻塞，直到有人发送
				有缓冲 -> 阻塞，直到缓冲里有数据
			接收还可以带 逗号 ok 形式：value, ok := <-ch
				ok == true -> 接收到正常数据
				ok == false -> channel 已关闭，且读到的是该类型的零值
		3. 关闭(close)：close(ch)
			表示 不会再向 channel 发送数据
			已经在 channel 里的数据仍可继续接收
			再向已关闭的 channel 发送 -> panic
			从已关闭的 channel 接收 -> 得到零值，并且 ok=false
	*/
	ch := make(chan int, 2)

	// 发送
	ch <- 20
	ch <- 50

	// 接收
	fmt.Println(<-ch) // 20
	fmt.Println(<-ch) // 50

	// 再发送
	ch <- 59
	close(ch) // 关闭 channel

	// 接收 + ok
	val, ok := <-ch
	fmt.Println(val, ok) // 59 true

	val, ok = <-ch
	fmt.Println(val, ok) // 0 false （零值 + 关闭状态）
}
