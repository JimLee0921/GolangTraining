package main

import "fmt"

// main 匿名结构体
func main() {
	/*
		匿名结构体就是：没有名字的 struct 类型
		它在定义的同时直接使用，不需要先写 type MyStruct struct { … }
		临时使用一次的结构体，不需要提前 type 定义
		可以只声明不初始化
		如果匿名结构体的字段名、顺序、类型、tag 都完全一致，那么它们就是同一个类型，可以互相赋值
		但是即使字段都相同但是顺序不一样也不能赋值
		使用场景
			一次性使用的数据组合
				例如某个函数里临时需要把几个字段打包，不想额外定义新类型
			快速 mock 数据
				写测试、演示代码时，用匿名结构体省事
			返回多个字段
				如果不想写命名结构体，可以用匿名结构体作为返回值
	*/
	// 声明并初始化
	config := struct {
		Host string
		Port int
	}{"localhost", 8080}
	fmt.Println(config.Host, config.Port)

	// 声明不初始化
	newConfig := struct {
		Host string
		Port int
	}{}
	newConfig.Host, newConfig.Port = "127.0.0.1", 8080
	fmt.Println(newConfig.Host, newConfig.Port)

	config = newConfig // 顺序相同字段相同可以相互赋值
	fmt.Println(config.Host, config.Port)

	newNewConfig := struct {
		Port int
		Host string
	}{Port: 8080, Host: "127.0.0.1"}
	fmt.Println(newNewConfig.Host, newNewConfig.Port)
	// config = newNewConfig	字段顺序不同不能直接赋值
}
