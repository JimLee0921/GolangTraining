package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}
type File struct{}

func (File) Read()  { fmt.Println("Reading file") }
func (File) Write() { fmt.Println("Writing file") }
func main() {
	/*
		类型断言还可以用于从一个接口转为另一个接口
		只要底层类型实现了目标接口
	*/
	var rw ReadWriter = File{}

	r := rw.(Reader) // 断言成功
	w := rw.(Writer) // 断言成功
	r.Read()
	w.Write()

	/*
		reader 虽然最初被声明为 Reader 接口类型
		但内部装的是一个 File 实例
		而 File 类型本身实现了：Read() 和 Write()
		所以确实也满足 ReadWriter 接口
		因此类型断言 reader.(ReadWriter) 成功
		得到的 rrr 现在就是一个 ReadWriter 接口变量
	*/
	var reader Reader = File{}

	rrr := reader.(ReadWriter)
	rrr.Write()
	rrr.Read()
}
