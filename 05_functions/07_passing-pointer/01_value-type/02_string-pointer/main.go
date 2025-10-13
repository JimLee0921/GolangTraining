package main

import "fmt"

// main string 指针传递
func main() {
	/*
		Go 的 string 是不可变的
		不能修改它内部的字节
		只能用新的字符串替换整个值
	*/
	name := "JimLee"
	fmt.Println(&name) // 0xc000022070
	changeStr(name)
	fmt.Println(name) // JimLee
	changeStrByPointer(&name)
	fmt.Println(name) // JimLee????
}

func changeStr(str string) {
	fmt.Println(str) // JimLee
	str += "????"
	fmt.Println(str) // JimLee????
}

func changeStrByPointer(str *string) {
	/*
		同样合法，用来整体替换
	*/
	fmt.Println(str) // 0xc000022070
	*str += "????"   // 相当于整体重新绑定了
	fmt.Println(str) // 0xc000022070
}
func replaceHead(ps *string) {
	/*
		合法，用来构造新字符串实现修改效果
		直接改 (*ps)[0] = 'X' —— 编译错误，因为 string 不可修改
	*/
	if len(*ps) == 0 {
		return
	}
	b := []byte(*ps) // 拷贝一份可变副本
	b[0] = 'X'
	*ps = string(b) // 生成新字符串再写回
}
