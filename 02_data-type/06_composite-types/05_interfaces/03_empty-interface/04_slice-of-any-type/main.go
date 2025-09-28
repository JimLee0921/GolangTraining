package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {
	/*
		使用空接口定义切片类型表示接收任何类型数据
	*/
	husky := dog{animal{"哈士奇"}, true}
	persian := cat{animal{"波斯猫"}, false}
	orange := cat{animal{"橘猫"}, true}
	// animals := []interface{}{husky, persian, orange}
	// Go 1.18+ 新版本写法
	animals := []any{husky, persian, orange}
	fmt.Println(animals)
}
