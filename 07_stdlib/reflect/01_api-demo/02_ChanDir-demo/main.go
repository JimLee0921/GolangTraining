package main

import (
	"fmt"
	"reflect"
)

func main() {
	tElem := reflect.TypeOf(0)
	// 使用 ChanOf 动态创建 channel 类型，第一个参数就是对应的 channel 方向
	tChan := reflect.ChanOf(reflect.SendDir, tElem)
	fmt.Println(tChan) // chan<- int
	// 底层调用 String 方法打印方向
	fmt.Println(reflect.SendDir, reflect.BothDir, reflect.RecvDir)
}
