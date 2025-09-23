package main

import "fmt"

// main fallthrough 关键字
func main() {
	/*
		在 Go 里，switch 的 默认行为和 C/Java 不一样：
		Go 默认不会“贯穿”执行下一个 case。
		也就是说，匹配到一个 case 后，执行完就退出，不会自动跑到下一个。
		如果想强制执行下一个 case，就要用 fallthrough
	*/
	// 这里匹配到 Marcus 后由于 fallthrough关键字会向下一直穿透直到 没有 fallthrough 或跑完整个语句
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}

}
