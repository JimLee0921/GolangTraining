package main

/*
writer 接口也是一种类型定义
它定义的是一组行为（方法集），而非数据结构
任何实现这些方法的类型都自动满足该接口
*/
type writer interface {
	writer(p []byte) (n int, err error)
}

func main() {

}
