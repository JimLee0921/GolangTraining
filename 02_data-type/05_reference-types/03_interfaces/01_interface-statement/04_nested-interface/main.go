package main

import "fmt"

//type Reader interface {
//	Read(p []byte) (n int, err error)
//}
//
//type Writer interface {
//	Write(p []byte) (n int, err error)
//}
//
//// ReadWriter 组合接口：标准库中的 io 包就这样用
//type ReadWriter interface {
//	Reader
//	Writer
//}

func main() {
	/*
		Go 运行接口嵌套来构造更复杂的行为
		只是把多个接口合并成一个大的接口
	*/
	fmt.Println("haha")

}
