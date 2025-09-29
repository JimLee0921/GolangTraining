package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		通常用于io.Reader（文件、网络、字符串 reader 等）读取
		fmt.Fscan
			func Fscan(r io.Reader, a ...any) (n int, err error)
			从 r 里连续读取数据，用空格和换行作为分隔符按顺序填充到 a 里
		fmt.Fscanln
			func Fscanln(r io.Reader, a ...any) (n int, err error)
			和 Fscan 类似，但在遇到换行符时停止，不会继续跨行读取，如果行里字段不足，会报错
		fmt.Fscanf
			func Fscanf(r io.Reader, format string, a ...any) (n int, err error)
			类似 Fscan，但可以指定格式化字符串来解析输入，非常适合解析固定格式的文本
	*/
	// fmt.Fscan
	data1 := "Jim 20"
	r1 := strings.NewReader(data1)

	var name1 string
	var age1 int
	n, err := fmt.Fscan(r1, &name1, &age1)
	if err != nil {
		fmt.Println("读取错误:", err)
		return
	}
	fmt.Printf("成功读取 %d 个字段: name=%s age=%d\n", n, name1, age1)

	// fmt.Fscanln
	data2 := "Alice 30\nBob 25\n"
	r2 := strings.NewReader(data2)

	var name2 string
	var age2 int
	_, _ = fmt.Fscanln(r2, &name2, &age2)
	fmt.Println(name2, age2) // Alice 30

	data3 := "Name:John Age:28"
	r := strings.NewReader(data3)

	var name3 string
	var age3 int
	_, _ = fmt.Fscanf(r, "Name:%s Age:%d", &name3, &age3)
	fmt.Println(name3, age3) // John 28

}
