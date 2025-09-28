package main

import "fmt"

type animal struct {
	name string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func introduce(a interface{}) {
	fmt.Println(a)
}

// Go 1.18+ 新版本写法
func introduces(a any) {
	fmt.Println(a)
}

func main() {
	/*
		函数参数使用空接口表示可以接收任何类型
	*/
	husky := dog{animal{"哈士奇"}, true}
	persian := cat{animal{"波斯猫"}, true}
	introduces(husky)
	introduce(persian)
}
